package test

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

// 定义 XML 结构体
type Dict struct {
	XMLName xml.Name `xml:"dict"`
	Items   []Item   `xml:",innerxml"`
}

type Item struct {
	Key   string `xml:",innerxml"`
	Value string `xml:",innerxml"`
}
func TestJdkXml(t *testing.T) {
	// 假设 XML 数据存储在 xmlData 变量中
	xmlData := `
	<dict>
		<key>JVMArch</key>
		<string>x86_64</string>
		<key>JVMBundleID</key>
		<string>com.oracle.java.jdk</string>
		<key>JVMEnabled</key>
		<true/>
		<key>JVMHomePath</key>
		<string>/Library/Java/JavaVirtualMachines/jdk-17.jdk/Contents/Home</string>
		<key>JVMName</key>
		<string>Java SE 17.0.7</string>
		<key>JVMPlatformVersion</key>
		<string>17.0.7</string>
		<key>JVMVendor</key>
		<string>Oracle Corporation</string>
		<key>JVMVersion</key>
		<string>17.0.7</string>
	</dict>`

	// 解析 XML
	dict := Dict{}
	err := xml.NewDecoder(strings.NewReader(xmlData)).Decode(&dict)
	if err != nil {
		fmt.Println("解析 XML 出错:", err)
		return
	}

	// 将解析后的数据存储到 map 中
	dataMap := make(map[string]string)
	for i := 0; i < len(dict.Items)-1; i += 2 {
		t.Errorf("%v: %v", dict.Items[i].Key, dict.Items[i].Value)
		key := dict.Items[i].Key
		value := dict.Items[i+1].Value
		dataMap[key] = value
	}

	// 打印 map 中的数据
	fmt.Println(dataMap)
	t.Logf("t: %v\n", dataMap)
}