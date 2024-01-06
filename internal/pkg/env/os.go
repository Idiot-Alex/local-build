package env

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// get OS type
// return {OSType} struct such as MacOS, Windows, Linux, Other
func GetOSType() OSType {
	switch runtime.GOOS {
		case "darwin":
			return MacOS
		case "windows":
			return Windows
		case "linux":
			return Linux
		default:
			return Other
	}
}

// get OS arch
// return string such as x86_64 arm ...
func GetOSArch() string {
	osType := GetOSType()
	if MacOS == osType {
		return getMacArch()
	} else if Windows == osType {
		getWindowsArch()
	} else if Linux == osType {

	}
	return "Unknown"
}

func getWindowsArch() {
	panic("unimplemented")
}

// get MacOS arch
func getMacArch() string {
	exCmd := exec.Command("uname", "-m")
	output, err := exCmd.Output()
	if err != nil {
		fmt.Println(err)
		return "Unknown"
	}
	return strings.TrimSpace(string(output))
}

// exec command use witch
func UseWhich(cmd string) string {
	exCmd := exec.Command("which", cmd)
	output, err := exCmd.Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(output))
}
