//go:build linux

package icon

func getIconData(isOnline bool) []byte {
	if isOnline {
		return StumpPng
	}

	return StumpOfflinePng
}
