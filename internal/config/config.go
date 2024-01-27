package config

import (
	"local-build/internal/lblog"
	"local-build/internal/utils"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Path string
	Db   struct {
		FilePath string `toml:"filePath"`
	} `toml:"db"`
}

// load config
func Load(fileName string) Config {
	if fileName == "" {
		fileName = "config.toml"
	}
	homePath := utils.GetHomePath()
	configPath := filepath.Join(homePath, fileName)
	lblog.Info("default configPath: %s", configPath)

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		configPath = filepath.Join(dir, fileName)
		lblog.Info("project configPath: %s", configPath)
	}

	return parseToml(configPath)
}

// parse toml config file
// configPath is file path
func parseToml(configPath string) Config {
	// 读取 TOML 配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	// 将 TOML 解析为 Config 结构体
	config := Config{Path: configPath}
	err = toml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	if !filepath.IsAbs(config.Db.FilePath) {
		dir := filepath.Dir(configPath)
		dbFilePath := filepath.Join(dir, config.Db.FilePath)
		config.Db.FilePath = dbFilePath
	}

	// 输出读取到的配置信息
	lblog.Info("config: %s", config)
	return config
}
