package test

import (
	"local-build/internal/pkg/env"
	"log"
	"testing"
)

func TestJdkXml(t *testing.T) {
	t.Setenv("GO_TEST_TIMEOUT", "300")
	jdkList := env.GetJDKList()
	t.Logf("jdkList: %v\n", jdkList)
}

func TestGitVersion(t *testing.T) {
	git := env.GetGitInfo()
	ver := env.GitVersion(git.Path)
	log.Printf("git version: %+v\n", ver)
}

func TestMavenVersion(t *testing.T) {
	mvn := env.GetMavenInfo()
	ver := env.MavenVersion(mvn.Path)
	log.Printf("maven version: %+v\n", ver)
}

func TestNodeVersion(t *testing.T) {
	node := env.GetNodeInfo()
	ver := env.NodeVersion(node.Path)
	log.Printf("node version: %+v\n", ver)
}