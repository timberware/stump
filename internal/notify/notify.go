package notify

import (
	_ "embed"
	"github.com/gen2brain/beeep"
	"stump/assets/icon"
	"stump/internal/logger"
)

func Alert(streamerName string) {
	title := "Stump"
	message := streamerName + " is online"

	iconFilePath, err := icon.CreateImagePath()
	if err != nil {
		logger.Error("error ensuring icon file:", err)
		return
	}

	err = beeep.Alert(title, message, iconFilePath)
	if err != nil {
		logger.Error("error creating alert:", err)
		return
	}
}
