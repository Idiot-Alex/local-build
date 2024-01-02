package test

import (
	"local-build/internal/utils"
	"testing"
)

func TestPackage(t *testing.T) {
	data := utils.ParsePackage("resources/test_package.json")
	t.Logf("data: %#v\n", data)
}