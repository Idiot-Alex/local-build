package test

import (
	"local-build/internal/env"
	"testing"
)

func TestJdkXml(t *testing.T) {
	t.Setenv("GO_TEST_TIMEOUT", "300")
	jdkList := env.GetJDKList()
	t.Logf("jdkList: %v\n", jdkList)
}