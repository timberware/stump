//go:build windows

package icon

func getIconLocation(isOnline bool) string {
	if isOnline {
		return "../assets/icon/stump.ico"
	}

	return "../assets/icon/stump_offline.ico"
}
