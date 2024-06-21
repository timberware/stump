package logger

import (
	"log"
	"os"
)

var l *log.Logger

func init() {
	l = log.New(os.Stdout, "", log.Ldate|log.Ltime)
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
