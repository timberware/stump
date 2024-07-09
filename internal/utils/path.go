package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetAppPath() (string, error) {
	var dir string

	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		dir = filepath.Join(appData, "stump")
	case "linux":
		home := os.Getenv("HOME")
		dir = filepath.Join(home, ".local", "share", "stump")
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	return filepath.Join(dir), nil
}
