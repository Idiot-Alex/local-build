package test

import (
	"local-build/internal/config"
	"testing"
)

func TestParseToml(t *testing.T) {
	config.Load("config.toml")
}