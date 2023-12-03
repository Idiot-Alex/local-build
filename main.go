package main

import (
	"fmt"
	"runtime"
	_ "os"
	"strings"
	"os/exec"
)


func main() {
	fmt.Println("app started...")
	fmt.Println(GetOS())

	findJava()
}

func findJava() {
	// 查找所有安装的 JDK
	javaHomeCmd := exec.Command("/usr/libexec/java_home", "-v")
	output, err := javaHomeCmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取 JDK 路径
	javaHomeVersions := strings.Split(strings.TrimSpace(string(output)), "\n")
	var javaHomePath string
	for _, version := range javaHomeVersions {
		// 只获取 OpenJDK 的路径
		// if strings.Contains(version, "openjdk") {
			javaHomeCmd := exec.Command("/usr/libexec/java_home", "-v", strings.TrimSpace(version))
			pathOutput, err := javaHomeCmd.Output()
			if err != nil {
				fmt.Println(err)
				continue
			}
			javaHomePath = strings.TrimSpace(string(pathOutput))
			break
		// }
	}

	if javaHomePath == "" {
		fmt.Println("Java runtime not found!")
		return
	}

	fmt.Println("Java home path:", javaHomePath)
}


func GetOS() string {
	os := runtime.GOOS
	if os != "windows" && os !=  "darwin" && os != "linux" {
		return ""
	}
	return os
}