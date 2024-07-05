package icon

import (
	_ "embed"
	"os"
	"stump/internal/logger"
)

//go:embed stump.png
var stumpIconBytes []byte

func CreateTempFile() (string, error) {
	tempFile, err := os.CreateTemp("", "stump-icon-*.png")
	if err != nil {
		return "", err
	}

	_, err = tempFile.Write(stumpIconBytes)
	if err != nil {
		tempFile.Close()
		CleanupTempFile(tempFile.Name())
		return "", err
	}

	err = tempFile.Close()
	if err != nil {
		CleanupTempFile(tempFile.Name())
		return "", err
	}

	return tempFile.Name(), nil
}

func CleanupTempFile(tempFilePath string) {
	err := os.Remove(tempFilePath)
	if err != nil {
		logger.Error("error removing temporary file:", err)
	}
}
