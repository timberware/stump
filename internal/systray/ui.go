package systray

import (
	"github.com/ncruces/zenity"
	"stump/internal/logger"
)

func GetTwitchToken(text, title string) (string, error) {
	result, err := zenity.Entry(text, zenity.Title(title))
	if err != nil {
		logger.Error("Failed to open input dialog", "error", err)
		return "", err
	}

	return result, nil
}

func ConfirmQuit(text, title string) (bool, error) {
	err := zenity.Question(text, zenity.Title(title))
	if err != nil {
		return false, err
	}

	return true, nil
}
