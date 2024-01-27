package test

import (
	"local-build/internal/fileparse"
	"testing"
)

func TestPackageJson(t *testing.T) {
	data, _ := fileparse.ParsePackageJson("resources/test_package.json")
	t.Logf("data: %#v\n", data)
}
