package fileparse

import (
	"encoding/json"
	"encoding/xml"
	"local-build/internal/lblog"
	"local-build/internal/pkg/env"
	"local-build/internal/utils"
	"os"
	"path/filepath"
)

type ParsedInfo struct {
	PomXmlList      []PomXml      `json:"pomXmlList"`
	PackageJsonList []PackageJson `json:"packageJsonList"`
}

// parse directory
func ParseDirectory(dir string) (*ParsedInfo, error) {
	var pomXmlList []PomXml
	var packageJsonList []PackageJson

	// foreach dir tree to parse files
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果出错，直接返回错误
		}
		if !info.IsDir() {
			lblog.Info("current parse file: [%s], name: [%s]", path, info.Name())
			switch info.Name() {
			case env.PACKAGE_JSON:
				data, err := ParsePackageJson(path)
				if err != nil {
					lblog.Error("parse file [%s] error: %s", path, err)
					return err
				}
				packageJsonList = append(packageJsonList, *data)
			case env.POM_XML:
				data, err := ParsePomXml(path)
				if err != nil {
					lblog.Error("parse file [%s] error: %s", path, err)
					return err
				}
				pomXmlList = append(pomXmlList, *data)
			}
		}
		return nil
	}

	// 遍历目录（包括子目录），调用回调函数处理每个文件或目录
	if err := filepath.Walk(dir, walkFunc); err != nil {
		lblog.Error("Walk() failed: %v\n", err)
	}

	return &ParsedInfo{
		PomXmlList:      pomXmlList,
		PackageJsonList: packageJsonList,
	}, nil
}

type PackageJson struct {
	FilePath    string `json:"filePath"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// parse package.json
// return package info
func ParsePackageJson(file string) (*PackageJson, error) {
	// 读取pom.xml文件的内容
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	filePath, _ := filepath.Abs(file)
	packageJson := &PackageJson{FilePath: filePath}
	err = json.Unmarshal(data, &packageJson)
	if err != nil {
		return nil, err
	}

	lblog.Info("package.json: %s", utils.ToJsonString(packageJson))
	return packageJson, nil
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
func ParsePomXml(file string) (*PomXml, error) {
	// 读取pom.xml文件的内容
	xmlData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	filePath, _ := filepath.Abs(file)
	pom := &PomXml{FilePath: filePath}
	err = xml.Unmarshal(xmlData, &pom)
	if err != nil {
		return nil, err
	}

	lblog.Info("pom.xml: %s", utils.ToJsonString(pom))
	return pom, nil
}
