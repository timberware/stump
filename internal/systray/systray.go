package systray

import (
	"stump/internal/auth"
	"stump/internal/icon"
	"stump/internal/logger"
	"stump/internal/user"
	"stump/internal/ws"

	"github.com/getlantern/systray"
)

var twitchUser *systray.MenuItem
var connect *systray.MenuItem
var disconnect *systray.MenuItem
var login *systray.MenuItem
var quit *systray.MenuItem

func OnReady() {
	login = systray.AddMenuItem("Login", "Twitch Login")
	connect = systray.AddMenuItem("Connect", "Connect to Twitch")
	disconnect = systray.AddMenuItem("Disconnect", "Disconnect from Twitch")
	quit = systray.AddMenuItem("Quit", "Quit the app")

	go Handle()
}

func Handle() {
	logo := icon.Data(true)
	logo_offline := icon.Data(false)
	systray.SetTitle("Stump")
	systray.SetIcon(logo_offline)

	var u user.User
	var dc string
	m := make(chan string)

	for {
		select {
		case <-login.ClickedCh:
			dc = auth.GetDeviceCode()
			logger.Info(dc)

		case <-connect.ClickedCh:
			u.Token = auth.GetToken(dc)
			u.GetInfo()
			u.GetAllFollowed()
			logger.Info("Connecting")
			ws.Connect(m, u)
			systray.SetIcon(logo)

		case <-disconnect.ClickedCh:
			logger.Info("Disconnecting")
			ws.Disconnect()
			systray.SetIcon(logo_offline)

		case <-quit.ClickedCh:
			shouldQuit, _ := ConfirmQuit("Do you want to quit?", "Confirm Quit?")

			if shouldQuit {
				systray.Quit()
			}

		case sessionId := <-m:
			for _, f := range u.Followed {
				go ws.SubscribeToEvent(f.Broadcaster_id, sessionId, u.Token)
			}
		}
	}
}

func OnExit() {
	logger.Info("Bye Bye")
}
