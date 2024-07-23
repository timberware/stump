package main

import (
	_ "embed"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"stump/internal/db"
	"stump/internal/logger"
	stump "stump/internal/systray"

	"github.com/getlantern/systray"
)

//go:embed .env
var env string

func main() {
	var err error

	err = logger.Init()
	if err != nil {
		fmt.Printf("failed to initialize logger: %v\n", err)
	}

	envMap, _ := godotenv.Unmarshal(env)
	err = os.Setenv("CLIENT_ID", envMap["CLIENT_ID"])
	if err != nil {
		logger.Error("missing or invalid .env file")
		os.Exit(1)
	}

	conn, err := db.Init()
	if err != nil {
		return
	}
	defer db.Close(conn)

	systray.Run(stump.OnReady, stump.OnExit)
}
