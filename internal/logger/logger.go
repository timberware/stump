package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"stump/internal/utils"
)

var l *log.Logger

const logName = "stump.log"

func Init() error {
	logToFile := utils.IsFlagSet("log")

	var logOutput *os.File

	if logToFile {
		logPath, err := utils.GetAppPath()
		logPath = filepath.Join(logPath, logName)

		if err != nil {
			return fmt.Errorf("failed to get application data path: %v", err)
		}

		logOutput, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open log file: %v", err)
		}
	} else {
		logOutput = os.Stdout
	}

	l = log.New(logOutput, "", log.Ldate|log.Ltime)

	return nil
}

func Info(msg string, args ...any) {
	l.Printf("INFO: "+msg, args...)
}

func Warn(msg string, args ...any) {
	l.Printf("WARN: "+msg, args...)
}

func Error(msg string, args ...any) {
	l.Printf("ERROR: "+msg, args...)
}
