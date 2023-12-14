package env

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type GitInfo struct {
	Name string
	Path string
	Version string
}

// find Git from local machine
// return GitInfo
func GetGitInfo () GitInfo {
	osType := GetOSType()
	if MacOS == osType {
		return getMacGit()
	} else if Windows == osType {
		
	} else if Linux == osType {

	}
	return GitInfo{}
}

// find Git from mac
// use command: which git && git --version
// return {GitInfo} struct or nil when error
func getMacGit() GitInfo {
	gitInfo := new(GitInfo)

	exCmd := exec.Command("which", "git")
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return *gitInfo
	}
	gitPath := strings.TrimSpace(string(output))

	exCmd = exec.Command("git", "--version")
	output, err = exCmd.Output()
	if err != nil {
		log.Println(err)
		return *gitInfo
	}
	gitVersion := strings.TrimSpace(string(output))

	gitInfo.Name = "Git"
	gitInfo.Path = string(gitPath)
	gitInfo.Version = string(gitVersion)

	jsonData, _ := json.MarshalIndent(gitInfo, "", "  ")
	fmt.Println(string(jsonData))

	return *gitInfo
}