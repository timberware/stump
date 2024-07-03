package main

import (
	"github.com/getlantern/systray"
	"stump/internal/db"
	stump "stump/internal/systray"
)

func main() {
	conn, err := db.Init()
	if err != nil {
		return
	}
	defer db.Close(conn)

	systray.Run(stump.OnReady, stump.OnExit)
}
