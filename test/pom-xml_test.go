package test

import (
	"local-build/internal/utils"
	"testing"
)

func TestPom(t *testing.T) {
	t.Setenv("GO_TEST_TIMEOUT", "60")
	pom := utils.ParsePom("resources/test_pom.xml")
	t.Logf("pom: %#v\n", pom)
}