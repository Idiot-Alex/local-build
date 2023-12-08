package env

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
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

	// Extract column names from the first row
	columns := strings.Fields(rows[0])

	// Find the index of the "Name" column
	nameIndex := -1
	for i, col := range columns {
			if col == "Name" {
					nameIndex = i
					break
			}
	}

	// Extract the "Name" column from each row
	var names []string
	for _, row := range rows[1:] {
			if row != "" {
					cols := strings.Fields(row)
					if nameIndex >= 0 && len(cols) > nameIndex {
							names = append(names, cols[nameIndex])
					}
			}
	}

	// Print the list of program names
	fmt.Println("Program Names:")
	for _, name := range names {
			fmt.Println(name)
	}
}