package utils

import (
	"github.com/go-git/go-git/v5"
)

type GitConfig struct {
	Url       string
	LocalPath string
	Protocol  string
	UserName  string
	Password  string
}

// use go-git clone repo
func GitClone(c GitConfig) error {
	// Clone the given repository to the given directory
	Info("git clone %s %s --recursive", c.Url, c.LocalPath)

	_, err := git.PlainClone(c.LocalPath, false, &git.CloneOptions{
		URL:               c.Url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		Error(err)
		return err
	}

	return nil
}
