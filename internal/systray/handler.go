package systray

import (
	"github.com/getlantern/systray"
	"stump/internal/logger"
)

var connectMenuItem *systray.MenuItem
var disconnectMenuItem *systray.MenuItem
var tokenMenuItem *systray.MenuItem
var quitMenuItem *systray.MenuItem

func SetupMenu() {
	connectMenuItem = systray.AddMenuItem("Connect", "Connect to Twitch")
	disconnectMenuItem = systray.AddMenuItem("Disconnect", "Disconnect from Twitch")
	tokenMenuItem = systray.AddMenuItem("Token", "Twitch Token")
	quitMenuItem = systray.AddMenuItem("Quit", "Quit the app")

	go handleConnect()
	go handleDisconnect()
	go handleToken()
	go handleQuit()
}

func handleConnect() {
	<-connectMenuItem.ClickedCh
	logger.Info("Connecting to Twitch...")
}

func handleDisconnect() {
	<-disconnectMenuItem.ClickedCh
	logger.Info("Disconnecting from Twitch...")
}

func handleToken() {
	<-tokenMenuItem.ClickedCh
	logger.Info("Requesting token")
	GetTwitchToken("Please input your Twitch Token", "Get Token")
	logger.Info("Token received")
}

func handleQuit() {
	<-quitMenuItem.ClickedCh
	logger.Info("Requesting quit")
	confirmQuit, _ := ConfirmQuit("Do you want to continue?", "Confirm Quit?")
	if confirmQuit {
		logger.Info("Quitting confirmed")
		systray.Quit()
	}
}
