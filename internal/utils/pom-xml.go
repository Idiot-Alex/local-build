package utils

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type PomXml struct {
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

	project := PomXml{}
	err = xml.Unmarshal(xmlData, &project)
	if err != nil {
		panic(err)
	}

	jsonData, _ := json.MarshalIndent(project, "", "  ")
	Info("pom: %s", string(jsonData))

	return string(jsonData)
}
