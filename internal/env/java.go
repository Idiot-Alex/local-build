package env

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// type Dict struct {
// 	XMLName xml.Name `xml:"plist"`
// 	Data    []Data   `xml:"array>dict"`
// }

type Data struct {
	Key   string `xml:"key"`
	Value string `xml:"string"`
}

// type Dict struct {
// 	Keys []string `xml:"key"`
// 	Strings []string `xml:"string"`
// 	Bools []bool `xml:"true"`
// }

type Plist struct {
	XMLName xml.Name `xml:"plist"`
	Array   struct {
		Dicts []Dict `xml:"dict"`
	} `xml:"array"`
}

type Dict struct {
	Keys   []string `xml:"key"`
	Values []string `xml:",any"`
}

type JdkInfo struct {
	Name string
	Version string
	Path string
	Vendor string
	Arch string
}

func GetJDKList() {
	osType := GetOSType()
	if MacOS == osType {
		getMacJDKList()
	} else if Windows == osType {
		getWindowsJDKList()
	} else if Linux == osType {

	}
}

// find JDK from mac
// use command: /usr/libexec/java_home -X
// return {JdkInfo} struct or nil when error
func getMacJDKList() []JdkInfo {
	exCmd := exec.Command("/usr/libexec/java_home", "-X")
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(string(output))

	var plist Plist
	err = xml.Unmarshal([]byte(output), &plist)

	var jdkList []JdkInfo
	for _, v := range plist.Array.Dicts {
		jdk := JdkInfo{}
		for i, key := range v.Keys {
			value := v.Values[i]
			switch key {
				case "JVMArch":
					jdk.Arch = value
				case "JVMHomePath":
					jdk.Path = value
				case "JVMVendor":
					jdk.Vendor = value
				case "JVMVersion":
					jdk.Version = value
				case "JVMName":
					jdk.Name = value
			}
		}
		jdkList = append(jdkList, jdk)
	}
	jsonData, _ := json.MarshalIndent(jdkList, "", "  ")
	fmt.Println(string(jsonData))
	
	return jdkList
}

// find JDK from windows
// use Command: wmic product where "Name like 'Java %% Development Kit%%'" get Name, Version, InstallLocation, Vendor
func getWindowsJDKList() {
	exCmd := exec.Command("wmic", "product", "where", `Name like 'Java %% Development Kit%%' get Name, Version, InstallLocation, Vendor`)
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(output))

	// Split output into rows
	rows := strings.Split(string(output), "\n")

	keys := strings.Fields(rows[0])

	var data []map[string]string

	for _, row := range rows[1:] {
		if strings.TrimSpace(row) != "" {
			fields := strings.Fields(row)
			rowMap := make(map[string]string)
			for i, field := range fields {
				if i < len(keys) {
					rowMap[keys[i]] = field
				}
			}
			data = append(data, rowMap)
		}
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))
}
