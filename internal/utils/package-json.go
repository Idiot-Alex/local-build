package utils

import (
	"encoding/json"
	"log"
	"os"
)

type PackageJson struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Description string `json:"description"`
}

// parse package.json
// return package info
func ParsePackage(file string) string {
	// 读取pom.xml文件的内容
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	
	var packageJson PackageJson
	err = json.Unmarshal(data, &packageJson)
	if err != nil {
		panic(err)
	}

	jsonData, _ := json.MarshalIndent(packageJson, "", "  ")
	log.Printf("pom: %v\n", string(jsonData))

	return string(jsonData)
}