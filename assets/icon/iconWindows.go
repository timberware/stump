//go:build windows

package icon

func getIconData(isOnline bool) []byte {
	if isOnline {
		return StumpIco
	}
	
	return StumpOfflineIco
}
