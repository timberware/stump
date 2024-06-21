package systray

import (
	"github.com/getlantern/systray"
	"stump/internal/icon"
	"stump/internal/logger"
)

func OnReady() {
	logger.Info("Application ready")
	systray.SetTitle("Stump")
	systray.SetIcon(icon.Data())

	SetupMenu()
}

func OnExit() {
	logger.Info("Application closed")
}
