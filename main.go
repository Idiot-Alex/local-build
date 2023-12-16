package main

import (
	"local-build/internal/env"
	"local-build/internal/sqlite"
)

func main() {
	env.GetJDKList()

	env.GetGitInfo()

	env.GetMavenInfo()

	env.GetNodeInfo()

	sqlite.SQLiteInMemoryDemo()
}