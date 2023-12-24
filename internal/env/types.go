package env

type OSType string

const (
	MacOS OSType = "macOS"
	Windows OSType = "windows"
	Linux OSType = "Linux"
	Other OSType = "Other"

	GIT string = "GIT"
	JDK string = "JDK"
	MAVEN string = "MAVEN"
	NODE string = "NODE"
)

type Tool struct {
	Name string
	Path string
	Version string
	Vendor string
	Arch string
	Type string
}

type ToolType struct {

}