package env

import (
	"encoding/json"
	"fmt"
	"local-build/internal/store/model"
	"os/exec"
	"strings"
)

// find Git from local machine
// return Tool
func GetGitInfo () model.Tool {
	osType := GetOSType()
	if MacOS == osType {
		return getMacGit()
	} else if Windows == osType {
		
	} else if Linux == osType {

	}
	return model.Tool{}
}

// find Git from mac
// use command: which git && git --version
// return {Tool} struct or nil when error
func getMacGit() model.Tool {
	git := new(model.Tool)

	gitPath := UseWhich("git")

	gitVersion := GitVersion(gitPath)

	git.Name = "Git"
	git.Path = string(gitPath)
	git.Version = string(gitVersion)
	git.Arch = GetOSArch()
	git.Type = GIT

	jsonData, _ := json.MarshalIndent(git, "", "  ")
	fmt.Printf("git: %v\n", string(jsonData))

	return *git
}

// get git version
func GitVersion(gitPath string) string {
	exCmd := exec.Command(gitPath, "--version")
	output, err := exCmd.Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(output))
}