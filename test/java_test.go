package test

import (
	"testing"
	"local-build/internal/env"
)

func TestJdkXml(t *testing.T) {
	jdkList := env.GetJDKList()
	t.Logf("jdkList: %v\n", jdkList)
}