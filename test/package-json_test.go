package test

import (
	"local-build/internal/fileparse"
	"testing"
)

func TestPackage(t *testing.T) {
	data := fileparse.ParsePackage("resources/test_package.json")
	t.Logf("data: %#v\n", data)
}
