package env

import (
	"os/exec"
	"log"
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
}