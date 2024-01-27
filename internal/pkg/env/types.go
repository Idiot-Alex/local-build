package env

type OSType string

const (
	MacOS   OSType = "macOS"
	Windows OSType = "windows"
	Linux   OSType = "Linux"
	Other   OSType = "Other"

	GIT   string = "GIT"
	JDK   string = "JDK"
	MAVEN string = "MAVEN"
	NODE  string = "NODE"

	DIR string = "DIR"
	SVN string = "SVN"

	// 凭据：用户名密码
	CREDENTIALS string = "credentials"
	// ssh 私钥
	SSH_PRIVATE_KEY string = "sshPrivateKey"
	// access token
	ACCESS_TOKEN string = "accessToken"

	PACKAGE_JSON string = "package.json"
	POM_XML      string = "pom.xml"
)
