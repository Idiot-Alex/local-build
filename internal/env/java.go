package env

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"encoding/json"
)

func GetJDKList() {
	getWindowsJDKList()
}

func getWindowsJDKList() {
	// cmd := `wmic product where "Name like 'Java %% Development Kit%%'"`
	exCmd := exec.Command("wmic", "product", "where", `Name like 'Java %% Development Kit%%'`)
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
				if (i < len(keys)) {
					rowMap[keys[i]] = field
				}
			}
			data = append(data, rowMap)
		}
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))

}