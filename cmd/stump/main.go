package main

import (
	"fmt"
	stump "stump/internal/systray"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(stump.OnReady, stump.OnExit)

	fmt.Println("hey sailor")
}
