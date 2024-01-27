package test

import (
	"local-build/internal/fileparse"
	"testing"
)

func TestPomXml(t *testing.T) {
	t.Setenv("GO_TEST_TIMEOUT", "60")
	pom, _ := fileparse.ParsePomXml("resources/test_pom.xml")
	t.Logf("pom: %#v\n", pom)
}
