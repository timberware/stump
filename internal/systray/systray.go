package systray

import (
	"fmt"
	"github.com/getlantern/systray"
	"stump/assets/icon"
)

func OnReady() {
	systray.SetTitle("Stump")
	systray.SetIcon(icon.Data())
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func OnExit() {
	fmt.Println("Application closed")
}
