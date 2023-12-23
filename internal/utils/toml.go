package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Db struct {
		FilePath string `toml:"filePath"`
	} `toml:"db"`
}

// parse toml config file
// dir is directory path
// file is toml file name
func ParseToml(dir string, file string) Config {
	configPath := filepath.Join(dir, file)

	// 读取 TOML 配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
			panic(err)
	}

	// 将 TOML 解析为 Config 结构体
	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
			panic(err)
	}

	// 输出读取到的配置信息
	fmt.Printf("config: %+v\n", config)
	return config
}