package fileparse

import (
	"encoding/json"
	"encoding/xml"
	"local-build/internal/lblog"
	"local-build/internal/pkg/env"
	"os"
	"path/filepath"
)

// parse dir
func ParseDirectory(dir string) {
	// foreach dir tree to parse files
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果出错，直接返回错误
		}
		if !info.IsDir() {
			lblog.Info("current file: [%s], name: [%s]", path, info.Name())
			switch info.Name() {
			case env.PACKAGE_JSON:
				ParsePackage(path)
			case env.POM_XML:
				ParsePom(path)
			}
		}
		return nil
	}

	// 遍历目录（包括子目录），调用回调函数处理每个文件或目录
	if err := filepath.Walk(dir, walkFunc); err != nil {
		lblog.Error("Walk() failed: %v\n", err)
	}
}

type PackageJson struct {
	FilePath    string `json:"filePath"`
	Name        string `json:"name"`
	Version     string `json:"version"`
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

	filePath, _ := filepath.Abs(file)
	packageJson := PackageJson{FilePath: filePath}
	err = json.Unmarshal(data, &packageJson)
	if err != nil {
		panic(err)
	}

	jsonData, _ := json.MarshalIndent(packageJson, "", "  ")
	lblog.Info("pom: %s", string(jsonData))

	return string(jsonData)
}

type PomXml struct {
	FilePath   string   `json:"filePath"`
	XMLName    xml.Name `xml:"project"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
	Version    string   `xml:"version"`
	Packaging  string   `xml:"packaging"`
	Profiles   string   `xml:"profiles"`
	Modules    []string `xml:"modules>module"`
}

// parse pom.xml
// return PomXml info
func ParsePom(file string) string {
	// 读取pom.xml文件的内容
	xmlData, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	filePath, _ := filepath.Abs(file)
	project := PomXml{FilePath: filePath}
	err = xml.Unmarshal(xmlData, &project)
	if err != nil {
		panic(err)
	}

	jsonData, _ := json.MarshalIndent(project, "", "  ")
	lblog.Info("pom: %s", string(jsonData))

	return string(jsonData)
}
