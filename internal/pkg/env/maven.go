package env

import (
	"encoding/json"
	"fmt"
	"local-build/internal/store/model"
	"os/exec"
	"strings"
)

// find Maven from local machine
// return Tool
func GetMavenInfo () model.Tool {
	osType := GetOSType()
	if MacOS == osType {
		return getMacMaven()
	} else if Windows == osType {
		
	} else if Linux == osType {

	}
	return model.Tool{}
}

// find Maven from mac
// use command: which mvn && mvn -v
// return Tool struct or nil when error
func getMacMaven() model.Tool {
	maven := new(model.Tool)

	mavenPath := UseWhich("mvn")

	mavenVersion := MavenVersion(mavenPath)

	maven.Name = "Maven"
	maven.Path = string(mavenPath)
	maven.Version = string(mavenVersion)
	maven.Arch = GetOSArch()
	maven.Type = MAVEN

	jsonData, _ := json.MarshalIndent(maven, "", "  ")
	fmt.Printf("maven: %v\n", string(jsonData))

	return *maven
}

// get maven version
func MavenVersion(mavenPath string) string {
	exCmd := exec.Command(mavenPath, "-v")
	output, err := exCmd.Output()
	if err != nil {
		panic(err)
	}
	row := strings.Split(string(output), "\n")[0]
	return strings.TrimSpace(strings.Split(row, "(")[0])
}