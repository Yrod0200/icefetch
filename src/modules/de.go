package modules

import "os"

func getDE() string {

	de := os.Getenv("XDG_CURRENT_DESKTOP")
	if de != "" {
		return de
	}

	if os.Getenv("KDE_FULL_SESSION") == "true" {
		return "KDE"
	}
	if os.Getenv("GNOME_DESKTOP_SESSION_ID") != "" {
		return "GNOME"
	}

	return "Unknown"
}
func DesktopEnvironment() string {

	return getDE()

}
