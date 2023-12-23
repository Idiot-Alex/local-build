package test

import (
	"local-build/internal/utils"
	"testing"
)

func TestParseToml(t *testing.T) {
	dir := "/Users/zhangxin/Documents/work_space/local-build/"
	file := "config.toml"
	utils.ParseToml(dir, file)
}