package env

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"bufio"
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

// find JDK list from local machine
// return []JdkInfo
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
// use Command: wmic product where "Name like 'Java %% Development Kit%%'" get Name, Version, InstallLocation, Vendor /format:csv
func getWindowsJDKList() []JdkInfo {
	exCmd := exec.Command("wmic", "product", "where", `Name like 'Java %% Development Kit%%'`, "get", "Name,Version,InstallLocation,Vendor", "/format:csv");
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(string(output))

	// output := []byte(`Node,InstallLocation,Name,Vendor,Version
	// DESKTOP-JR7EEIP,C:\Program Files\Java\jdk-1.8\,Java SE Development Kit 8 Update 381 (64-bit),Oracle Corporation,8.0.3810.9
	// DESKTOP-JR7EEIP,C:\Program Files\Java\jdk1.8.0_51\,Java SE Development Kit 8 Update 51 (64-bit),Oracle Corporation,8.0.510.16`)

	var rows []string

	// 将结果按行分割成多行
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r\n\t")
		if strings.TrimSpace(line) != "" {
			rows = append(rows, line)
		}
	}

	fmt.Printf("====%#v", rows)
	keys := strings.Split(rows[0], ",")

	fmt.Printf("====%#v", keys)
	var data []map[string]string

	for _, row := range rows[1:] {
		if strings.TrimSpace(row) != "" {
			fields := strings.Split(row, ",")
			rowMap := make(map[string]string)
			for i, field := range fields {
				if i < len(keys) {
					rowMap[keys[i]] = field
				}
			}
			data = append(data, rowMap)
		}
	}
	fmt.Printf("====%#v", data)

	var jdkList []JdkInfo
	for _, v := range data {
		jdk := JdkInfo{}
		for key, value := range v {
			fmt.Println(key)
			switch key {
				case "InstallLocation":
					jdk.Path = value
				case "Vendor":
					jdk.Vendor = value
				case "Version":
					jdk.Version = value
				case "Name":
					jdk.Name = value
			}
		}
		jdkList = append(jdkList, jdk)
	}
	jsonData, _ := json.MarshalIndent(jdkList, "", "  ")
	fmt.Println(string(jsonData))
	
	return jdkList
}
