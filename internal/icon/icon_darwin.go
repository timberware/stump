//go:build darwin 

package icon

func getIconLocation(isOnline bool) string {
	if isOnline {
		return "assets/icon/stump.png"
	}

	return "assets/icon/stump_offline.png"
}
