package config

import (
	"local-build/internal/utils"
	"log"
	"os"
	"path/filepath"
)

// load config
func Load() {
	fileName := "config.toml"
	homePath := utils.GetHomePath()
	configPath := filepath.Join(homePath, fileName)
	log.Printf("default configPath: %+v\n", configPath)

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		configPath = filepath.Join(dir, fileName)
		log.Printf("project configPath: %+v\n", configPath)
	}

	config := utils.ParseToml(configPath)
	log.Printf("config: %+v\n", config)
}