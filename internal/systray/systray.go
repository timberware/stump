package systray

import (
	"stump/internal/icon"
	"stump/internal/logger"
	"stump/internal/ws"

	"github.com/getlantern/systray"
)

var connect *systray.MenuItem
var disconnect *systray.MenuItem
var token *systray.MenuItem
var quit *systray.MenuItem
var m chan string

func OnReady() {
	systray.SetTitle("Stump")
	systray.SetIcon(icon.Data())

	connect = systray.AddMenuItem("Connect", "Connect to Twitch")
	disconnect = systray.AddMenuItem("Disconnect", "Disconnect from Twitch")
	token = systray.AddMenuItem("Token", "Twitch Token")
	quit = systray.AddMenuItem("Quit", "Quit the app")

	go Handle()
}

func Handle() {
	m = make(chan string)
	for {
		select {
		case <-connect.ClickedCh:
			logger.Info("Connecting")
			ws.Connect(m)
		case <-disconnect.ClickedCh:
			logger.Info("Disconnecting")
			ws.Disconnect()
		case <-token.ClickedCh:
			GetTwitchToken("Please input your Twitch Token", "Get Token")
		case <-quit.ClickedCh:
			shouldQuit, _ := ConfirmQuit("Do you want to continue?", "Confirm Quit?")

			if shouldQuit {
				systray.Quit()
			}
		case message := <-m:
			logger.Info(message)
		}
	}
}

func OnExit() {
	logger.Info("Bye Bye")
}
