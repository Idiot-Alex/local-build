package main

import (
	"local-build/internal/env"
)

func main() {
	env.GetJDKList()

	env.GetGitInfo()

	env.GetMavenInfo()

	env.GetNodeInfo()
}