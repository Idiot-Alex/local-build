package utils

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

// use go-git clone repo
func GitClone() {
	url := "git@github.com:Idiot-Alex/local-build.git"
	directory := "/tmp/test/local-build"
	// Clone the given repository to the given directory
	Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	Error(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	Error(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	Error(err)

	fmt.Println(commit)
}
