package user

import (
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"net/url"
	"os"
	"stump/internal/logger"
	"time"
)

type User struct {
	Id        string
	Username  string
	AvatarUrl string
	Token     string
	Followed  []FollowedData
}

type data struct {
	Id                string
	Login             string
	Profile_image_url string
}

type userData struct {
	Data []data
}

type Followed struct {
	Total      int
	Data       []FollowedData
	Pagination Pagination
}

type Pagination struct {
	Cursor string
}

type FollowedData struct {
	Broadcaster_id   string
	Broadcaster_name string
}

var twitchURL = flag.String("twitch_api", "api.twitch.tv", "http service address")

func (u *User) GetInfo() error {
	var ud userData
	clientId := os.Getenv("CLIENT_ID")
	tu := url.URL{Scheme: "https", Host: *twitchURL, Path: "/helix/users"}
	client := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest(http.MethodGet, tu.String(), nil)

	request.Header.Set("Authorization", "Bearer "+u.Token)
	request.Header.Set("Client-ID", clientId)

	response, err := client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)
	if error != nil {
		logger.Error(error.Error())
		return err
	}
	logger.Info(string(body))

	e := json.Unmarshal(body, &ud)
	if e != nil {
		logger.Error(e.Error())
		return err
	}

	u.Id = ud.Data[0].Id
	u.Username = ud.Data[0].Login
	u.AvatarUrl = ud.Data[0].Profile_image_url

	return nil
}

func (u *User) GetAllFollowed() error {
	after := "first"
	u.Followed = []FollowedData{}
	var err error

	for after != "" {
		after, err = u.getFollowed(after)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	return nil
}

func (u *User) getFollowed(after string) (string, error) {
	clientId := os.Getenv("CLIENT_ID")
	FollowedData := Followed{}

	rawQuery := "user_id=" + u.Id
	if after != "first" {
		rawQuery = rawQuery + "&after=" + after
	}

	tu := url.URL{Scheme: "https", Host: *twitchURL, Path: "/helix/channels/followed", ForceQuery: true, RawQuery: rawQuery}
	client := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest(http.MethodGet, tu.String(), nil)

	request.Header.Set("Authorization", "Bearer "+u.Token)
	request.Header.Set("Client-ID", clientId)

	response, err := client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return "client", err
	}
	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)
	if error != nil {
		logger.Error(error.Error())
		return "There was an error extracting the body", err
	}

	logger.Info(string(body))

	e := json.Unmarshal(body, &FollowedData)
	if e != nil {
		logger.Error(e.Error())
		return "There was an error parsing the body", err
	}

	u.Followed = append(u.Followed, FollowedData.Data...)

	return FollowedData.Pagination.Cursor, nil
}
