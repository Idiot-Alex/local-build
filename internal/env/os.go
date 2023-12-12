package env

import (
	"runtime"
)

type OSType string

const (
	MacOS OSType = "macOS"
	Windows OSType = "windows"
	Linux OSType = "Linux"
	Other OSType = "Other"
)

// get OS type
// return {OSType} struct such as MacOS, Windows, Linux, Other
func GetOSType() OSType {
	switch runtime.GOOS {
		case "darwin":
			return MacOS
		case "windows":
			return Windows
		case "linux":
			return Linux
		default:
			return Other
	}
}