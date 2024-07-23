package icon

import (
	"stump/internal/logger"
)

func Data(isOnline bool) []byte {
	iconData := getIconData(isOnline)
	if iconData == nil {
		logger.Error("failed to get systray icon data")
	}

	return iconData
}
