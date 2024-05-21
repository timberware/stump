package main

import (
	"fmt"
	"github.com/getlantern/systray"
)

func onReady() {
	systray.SetTitle("App")
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func onExit() {
	fmt.Println("Application closed")
}
