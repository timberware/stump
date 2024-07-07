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

	tempFilePath, err := icon.CreateTempFile()
	if err != nil {
		logger.Error("error creating temporary file:", err)
		return
	}
	defer icon.CleanupTempFile(tempFilePath)

	err = beeep.Alert(title, message, tempFilePath)
	if err != nil {
		logger.Error("error creating alert:", err)
		return
	}
}
