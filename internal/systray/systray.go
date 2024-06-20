package systray

import (
	"github.com/getlantern/systray"
	"stump/internal/icon"
)

func OnReady() {
	systray.SetTitle("Stump")
	systray.SetIcon(icon.Data())

	SetupMenu()
}

func OnExit() {
	Close()
}
