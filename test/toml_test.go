package test

import (
	"fmt"
	"local-build/internal/utils"
	"path/filepath"
	"testing"
)

func TestParseToml(t *testing.T) {
	path, err := filepath.Abs("resources/test_config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("path: %v\n", path)
	utils.ParseToml(path)
}