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
func GetNodeInfo () Tool {
	osType := GetOSType()
	if MacOS == osType {
		return getMacNode()
	} else if Windows == osType {
		
	} else if Linux == osType {

	}
	return Tool{}
}

// find Node from mac
// use command: which node && node -v
// return {Tool} struct or nil when error
func getMacNode() Tool {
	node := new(Tool)

	exCmd := exec.Command("which", "node")
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return *node
	}
	nodePath := strings.TrimSpace(string(output))

	exCmd = exec.Command("node", "-v")
	output, err = exCmd.Output()
	if err != nil {
		log.Println(err)
		return *node
	}
	nodeVersion := strings.TrimSpace(string(output))

	node.Name = "Node"
	node.Path = string(nodePath)
	node.Version = string(nodeVersion)
	node.Arch = GetOSArch()

	jsonData, _ := json.MarshalIndent(node, "", "  ")
	fmt.Printf("node: %v\n", string(jsonData))

	return *node
}