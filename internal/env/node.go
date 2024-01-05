package env

import (
	"encoding/json"
	"fmt"
	"local-build/internal/store/model"
	"log"
	"os/exec"
	"strings"
)

// find Git from local machine
// return Tool
func GetNodeInfo () model.Tool {
	osType := GetOSType()
	if MacOS == osType {
		return getMacNode()
	} else if Windows == osType {
		
	} else if Linux == osType {

	}
	return model.Tool{}
}

// find Node from mac
// use command: which node && node -v
// return {Tool} struct or nil when error
func getMacNode() model.Tool {
	node := new(model.Tool)

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
	node.Type = NODE

	jsonData, _ := json.MarshalIndent(node, "", "  ")
	fmt.Printf("node: %v\n", string(jsonData))

	return *node
}