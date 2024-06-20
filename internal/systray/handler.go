package systray

import (
	"fmt"
	"github.com/getlantern/systray"
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
	fmt.Println("Connecting to Twitch...")
}

func handleDisconnect() {
	<-disconnectMenuItem.ClickedCh
	fmt.Println("Disconnecting from Twitch...")
}

func handleToken() {
	<-tokenMenuItem.ClickedCh
	fmt.Println("Requesting token")
	text, _ := GetTwitchToken("Please input your Twitch Token", "Get Token")
	fmt.Println(text)
}

func handleQuit() {
	<-quitMenuItem.ClickedCh
	fmt.Println("Requesting quit")
	confirmQuit, _ := ConfirmQuit("Do you want to continue?", "Confirm Quit?")
	if confirmQuit {
		systray.Quit()
	} else {
		fmt.Println("Continuing")
	}
}

func Close() {
	fmt.Println("Application closed")
}
