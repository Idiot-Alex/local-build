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

	nodePath := UseWhich("node")

	nodeVersion := NodeVersion(nodePath)

	node.Name = "Node"
	node.Path = string(nodePath)
	node.Version = string(nodeVersion)
	node.Arch = GetOSArch()
	node.Type = NODE

	jsonData, _ := json.MarshalIndent(node, "", "  ")
	fmt.Printf("node: %v\n", string(jsonData))

	return *node
}

// get node version
func NodeVersion(nodePath string) string {
	exCmd := exec.Command(nodePath, "-v")
	output, err := exCmd.Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(output))
}