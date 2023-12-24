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

type data struct {
	Key   string `xml:"key"`
	Value string `xml:"string"`
}

type plist struct {
	XMLName xml.Name `xml:"plist"`
	Array   struct {
		Dicts []dict `xml:"dict"`
	} `xml:"array"`
}

type dict struct {
	Keys   []string `xml:"key"`
	Values []string `xml:",any"`
}

// find JDK list from local machine
// return []Tool
func GetJDKList() []Tool {
	osType := GetOSType()
	if MacOS == osType {
		return getMacJDKList()
	} else if Windows == osType {
		return getWindowsJDKList()
	} else if Linux == osType {

	}
	return []Tool{}
}

// find JDK from mac
// use command: /usr/libexec/java_home -X
// return Tool struct or nil when error
func getMacJDKList() []Tool {
	exCmd := exec.Command("/usr/libexec/java_home", "-X")
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return nil
	}

	var plist plist
	err = xml.Unmarshal([]byte(output), &plist)
	if err != nil {
		log.Println(err)
		return nil
	}

	var jdkList []Tool
	for _, v := range plist.Array.Dicts {
		jdk := Tool{ Arch: GetOSArch(), Type: JDK }
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
	fmt.Printf("jdk: %v\n", string(jsonData))
	
	return jdkList
}

// find JDK from windows
// use Command: wmic product where "Name like 'Java %% Development Kit%%'" get Name, Version, InstallLocation, Vendor /format:csv
func getWindowsJDKList() []Tool {
	exCmd := exec.Command("wmic", "product", "where", `Name like 'Java %% Development Kit%%'`, "get", "Name,Version,InstallLocation,Vendor", "/format:csv");
	output, err := exCmd.Output()
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(string(output))

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

	fmt.Printf("%#v", rows)
	keys := strings.Split(rows[0], ",")

	fmt.Printf("%#v", keys)
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
	fmt.Printf("%#v", data)

	var jdkList []Tool
	for _, v := range data {
		jdk := Tool{Arch: GetOSArch()}
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
