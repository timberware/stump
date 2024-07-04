package main

import (
	"os"

	"stump/internal/db"
	"stump/internal/logger"
	stump "stump/internal/systray"

	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("There was a problem loading the .env file")
		os.Exit(1)
	}

	conn, err := db.Init()
	if err != nil {
		return
	}
	defer db.Close(conn)

	systray.Run(stump.OnReady, stump.OnExit)
}
