package env

type OSType string

const (
	MacOS OSType = "macOS"
	Windows OSType = "windows"
	Linux OSType = "Linux"
	Other OSType = "Other"
)

// get OS type
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