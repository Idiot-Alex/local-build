package utils

import (
	"fmt"
	"local-build/internal/model"
	"os"

	"github.com/BurntSushi/toml"
)

// parse toml config file
// configPath is file path
func ParseToml(configPath string) model.Config {
	// 读取 TOML 配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
			panic(err)
	}

	// 将 TOML 解析为 Config 结构体
	config := model.Config{Path: configPath}
	err = toml.Unmarshal(data, &config)
	if err != nil {
			panic(err)
	}

	// 输出读取到的配置信息
	fmt.Printf("config: %+v\n", config)
	return config
}