package icon

import (
	_ "embed"
)

//go:embed images/stump.ico
var StumpIco []byte

//go:embed images/stumpOffline.ico
var StumpOfflineIco []byte

//go:embed images/stump.png
var StumpPng []byte

//go:embed images/stumpOffline.png
var StumpOfflinePng []byte
