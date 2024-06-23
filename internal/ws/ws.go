package ws

import (
	"flag"
	"net/url"
	"stump/internal/logger"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "eventsub.wss.twitch.tv", "http service address")
var conn *websocket.Conn
var connError error

func Connect(message chan string) {
	flag.Parse()
	interrupt := make(chan error)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws"}

	conn, _, connError = websocket.DefaultDialer.Dial(u.String(), nil)
	if connError != nil {
		logger.Error("dial:", connError)
	}

	logger.Info("connected to %s", u.String())
	go GetMessage(message, interrupt)
}

func GetMessage(m chan string, in chan error) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logger.Info("read:", err)
			in <- err
			return
		}
		m <- string(message)
	}
}

func Disconnect() {
	err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logger.Info("write close:", err)
		return
	}
}
