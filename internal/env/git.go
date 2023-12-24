package env

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// find Git from local machine
// return Tool
func GetGitInfo () Tool {
	osType := GetOSType()
	if MacOS == osType {
		return getMacGit()
	} else if Windows == osType {
		
	} else if Linux == osType {

	}
	return Tool{}
}

// find Git from mac
// use command: which git && git --version
// return {Tool} struct or nil when error
func getMacGit() Tool {
	git := new(Tool)

	exCmd := exec.Command("which", "git")
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return *git
	}
	gitPath := strings.TrimSpace(string(output))

	exCmd = exec.Command("git", "--version")
	output, err = exCmd.Output()
	if err != nil {
		log.Println(err)
		return *git
	}
	gitVersion := strings.TrimSpace(string(output))

	git.Name = "Git"
	git.Path = string(gitPath)
	git.Version = string(gitVersion)
	git.Arch = GetOSArch()
	git.Type = GIT

	jsonData, _ := json.MarshalIndent(git, "", "  ")
	fmt.Printf("git: %v\n", string(jsonData))

	return *git
}