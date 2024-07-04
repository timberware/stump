package auth

import (
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"net/url"
	"os"
	"stump/internal/logger"

	"github.com/pkg/browser"
)

type Device struct {
	Device_code      string
	Expires_in       int
	interval         int
	User_code        string
	Verification_uri string
}

type Token struct {
	Access_token  string
	Expires_in    int
	Refresh_token string
	Token_type    string
}

var twitchURL = flag.String("twitch_addr", "id.twitch.tv", "http service address")
var scopes = "user:read:follows"
var grantTypes = "urn:ietf:params:oauth:grant-type:device_code"

func GetDeviceCode() string {
	var d Device
	clientId := os.Getenv("CLIENT_ID")
	data := url.Values{"client_id": {clientId}, "scopes": {scopes}}
	u := url.URL{Scheme: "https", Host: *twitchURL, Path: "/oauth2/device"}

	resp, err := http.PostForm(u.String(), data)
	if err != nil {
		return "There was an error getting a device code"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "There was an error extracting the body"
	}
	defer resp.Body.Close()

	logger.Info(string(body))
	err = json.Unmarshal(body, &d)
	if err != nil {
		return "There was an error parsing the body"
	}

	err = browser.OpenURL(d.Verification_uri)
	if err != nil {
		return "There was an error opening the login page"
	}

	return d.Device_code
}

func GetToken(dc string) string {
	clientId := os.Getenv("CLIENT_ID")
	var tr Token

	data := url.Values{"client_id": {clientId}, "scopes": {scopes}, "device_code": {dc}, "grant_type": {grantTypes}}
	u := url.URL{Scheme: "https", Host: *twitchURL, Path: "/oauth2/token"}

	resp, err := http.PostForm(u.String(), data)
	if err != nil {
		return "There was an error getting a user token"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "There was an error extracting the body"
	}
	defer resp.Body.Close()

	logger.Info(string(body))
	err = json.Unmarshal(body, &tr)
	if err != nil {
		return "There was an error parsing the body"
	}

	return tr.Access_token
}
