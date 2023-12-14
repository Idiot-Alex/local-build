package env

type OSType string

const (
	MacOS OSType = "macOS"
	Windows OSType = "windows"
	Linux OSType = "Linux"
	Other OSType = "Other"
)

type Tool struct {
	Name string
	Path string
	Version string
	Vendor string
	Arch string
}