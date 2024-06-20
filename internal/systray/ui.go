package systray

import (
	"fmt"
	"github.com/ncruces/zenity"
)

func GetTwitchToken(text, title string) (string, error) {
	result, err := zenity.Entry(text, zenity.Title(title))
	if err != nil {
		return "", fmt.Errorf("failed to open input dialog: %w", err)
	}

	return result, nil
}

func ConfirmQuit(text, title string) (bool, error) {
	err := zenity.Question(text, zenity.Title(title))
	if err != nil {
		return false, fmt.Errorf("failed to open question dialog: %w", err)
	}

	return true, nil
}
