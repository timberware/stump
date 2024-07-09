package main

import (
	"fmt"
	"os"
	"stump/internal/utils"

	"stump/internal/db"
	"stump/internal/logger"
	stump "stump/internal/systray"

	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
)

func main() {
	utils.ParseFlags()

	var err error

	err = godotenv.Load()
	if err != nil {
		logger.Error("There was a problem loading the .env file: %v", err)
		os.Exit(1)
	}

	err = logger.Init()
	if err != nil {
		fmt.Printf("failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	conn, err := db.Init()
	if err != nil {
		return
	}
	defer db.Close(conn)

	systray.Run(stump.OnReady, stump.OnExit)
}
