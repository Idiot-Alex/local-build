package env

import (
	"encoding/json"
	"fmt"
	"local-build/internal/store/model"
	"log"
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

	exCmd := exec.Command("which", "mvn")
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return *maven
	}
	mavenPath := strings.TrimSpace(string(output))

	exCmd = exec.Command("mvn", "-v")
	output, err = exCmd.Output()
	if err != nil {
		log.Println(err)
		return *maven
	}
	row := strings.Split(string(output), "\n")[0]
	mavenVersion := strings.TrimSpace(strings.Split(row, "(")[0])

	maven.Name = "Maven"
	maven.Path = string(mavenPath)
	maven.Version = string(mavenVersion)
	maven.Arch = GetOSArch()
	maven.Type = MAVEN

	jsonData, _ := json.MarshalIndent(maven, "", "  ")
	fmt.Printf("maven: %v\n", string(jsonData))

	return *maven
}