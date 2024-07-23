package icon

import (
	_ "embed"
	"os"
	"path/filepath"
	"stump/internal/utils"
)

func CreateImagePath() (string, error) {
	appPath, err := utils.GetAppPath()
	if err != nil {
		return "", err
	}
	imageFilePath, err := filepath.Join(appPath, "stump.png"), nil

	if _, err := os.Stat(imageFilePath); os.IsNotExist(err) {
		err := os.WriteFile(imageFilePath, StumpPng, 0644)
		if err != nil {
			return "", err
		}
	}

	return imageFilePath, nil
}
