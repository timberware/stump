package icon

import (
	"os"
	"path/filepath"
	"stump/internal/logger"
)

func Data(isOnline bool) []byte {
	iconPath, err := filepath.Abs(filepath.FromSlash(getIconLocation(isOnline)))
	if err != nil {
		logger.Error("Failed to get systray icon path", "path", iconPath, "error", err)
		return nil
	}

	bytes, err := os.ReadFile(iconPath)
	if err != nil {
		return nil
	}

	return bytes
}
