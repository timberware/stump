package systray

import (
	"fmt"
	"github.com/getlantern/systray"
)

func OnReady() {
	systray.SetTitle("App")
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
