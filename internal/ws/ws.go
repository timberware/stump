package ws

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"net/url"
	"os"
	"stump/internal/logger"
	"stump/internal/notify"
	"stump/internal/user"
	"time"

	"github.com/gorilla/websocket"
)

type Condition struct {
	Broadcaster_user_id string `json:"broadcaster_user_id"`
}

type Transport struct {
	Method     string `json:"method"`
	Session_id string `json:"session_id"`
}

type Subscription struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Condition Condition `json:"condition"`
	Transport Transport `json:"transport"`
}

type WSMessage struct {
	Metadata Metadata `json:"metadata"`
	Payload  Payload  `json:"payload"`
}

type Session struct {
	Id string `json:"id"`
}

type Metadata struct {
	Message_type      string `json:"message_type"`
	Subscription_type string `json:"subscription_type"`
}

type Payload struct {
	Session      Session      `json:"session"`
	Subscription Subscription `json:"subscription"`
	Event        Event        `json:"event"`
}

type Event struct {
	Broadcaster_id         string `json:"broadcaster_user_id"`
	Broadcaster_user_login string `json:"broadcaster_user_login"`
	Broadcaster_user_name  string `json:"broadcaster_user_name"`
}

var twitchURL = flag.String("api", "api.twitch.tv", "http service address")
var addr = flag.String("addr", "eventsub.wss.twitch.tv", "http service address")
var conn *websocket.Conn
var connError error

func Connect(m chan string, user user.User) {
	flag.Parse()
	interrupt := make(chan error)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws"}

	conn, _, connError = websocket.DefaultDialer.Dial(u.String(), nil)
	if connError != nil {
		logger.Error("Error:", connError)
	}

	logger.Info("connected to %s", u.String())
	go GetMessage(m, interrupt)
}

func GetMessage(m chan string, in chan error) {
	var ms WSMessage

	for {

		_, message, err := conn.ReadMessage()
		if err != nil {
			logger.Error("Error:", err)
			in <- err
			return
		}

		logger.Info(string(message))

		err = json.Unmarshal(message, &ms)
		if err != nil {
			logger.Error("Error:", err)
			in <- err
			return
		}

		switch {
		case ms.Metadata.Message_type == "session_welcome":
			m <- ms.Payload.Session.Id
		case ms.Metadata.Subscription_type == "stream.online":
			notify.Alert(ms.Payload.Event.Broadcaster_user_name)
		}

		ms.Metadata.Message_type = ""
		ms.Metadata.Subscription_type = ""
	}
}

func SubscribeToEvent(broadcasterId string, sessionId string, token string) error {
	clientId := os.Getenv("CLIENT_ID")
	tu := url.URL{Scheme: "https", Host: *twitchURL, Path: "/helix/eventsub/subscriptions"}

	payload := &Subscription{
		Type:    "stream.online",
		Version: "1",
		Condition: Condition{
			Broadcaster_user_id: broadcasterId,
		},
		Transport: Transport{
			Method:     "websocket",
			Session_id: sessionId,
		},
	}

	marshalled, err := json.Marshal(payload)
	if err != nil {
		logger.Info("write close:", err)
		return err
	}

	client := http.Client{Timeout: 5 * time.Second}
	request, _ := http.NewRequest(http.MethodPost, tu.String(), bytes.NewReader(marshalled))

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Client-Id", clientId)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	body, error := io.ReadAll(response.Body)
	if error != nil {
		logger.Error(error.Error())
		return error
	}
	logger.Info(string(body))

	defer response.Body.Close()

	return nil
}

func Disconnect() {
	err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logger.Info("write close:", err)
		return
	}
}
