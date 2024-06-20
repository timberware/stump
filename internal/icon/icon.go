package icon

import (
	"fmt"
	"os"
	"path/filepath"
)

func Data() []byte {
	iconPath, err := filepath.Abs(filepath.FromSlash(getIconLocation()))
	if err != nil {
		fmt.Println("Failed to get systray icon path:", err)
		return nil
	}

	bytes, err := os.ReadFile(iconPath)
	if err != nil {
		fmt.Println("Failed to find", iconPath, ":", err)
		return nil
	}

	return bytes
}
