package main

import (
	"github.com/getlantern/systray"
	stump "stump/internal/systray"
)

func main() {
	systray.Run(stump.OnReady, stump.OnExit)
}
