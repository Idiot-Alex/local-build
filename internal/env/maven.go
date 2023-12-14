package env

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type MavenInfo struct {
	Name string
	Path string
	Version string
}

// find Maven from local machine
// return MavenInfo
func GetMavenInfo () MavenInfo {
	osType := GetOSType()
	if MacOS == osType {
		return getMacMaven()
	} else if Windows == osType {
		
	} else if Linux == osType {

	}
	return MavenInfo{}
}

// find Maven from mac
// use command: which mvn && mvn -v
// return {MavenInfo} struct or nil when error
func getMacMaven() MavenInfo {
	mavenInfo := new(MavenInfo)

	exCmd := exec.Command("which", "mvn")
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return *mavenInfo
	}
	mavenPath := strings.TrimSpace(string(output))

	exCmd = exec.Command("mvn", "-v")
	output, err = exCmd.Output()
	if err != nil {
		log.Println(err)
		return *mavenInfo
	}
	row := strings.Split(string(output), "\n")[0]
	mavenVersion := strings.TrimSpace(strings.Split(row, "(")[0])

	mavenInfo.Name = "Maven"
	mavenInfo.Path = string(mavenPath)
	mavenInfo.Version = string(mavenVersion)

	jsonData, _ := json.MarshalIndent(mavenInfo, "", "  ")
	fmt.Println(string(jsonData))

	return *mavenInfo
}