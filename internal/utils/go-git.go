package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
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

// fetch all
func GitFetchAll(c GitConfig) error {
	// 打开本地 Git 仓库
	repo, err := git.PlainOpen(c.LocalPath)
	if err != nil {
		Error(err)
		return err
	}

	// 获取所有远程分支的最新代码
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{
			"+refs/heads/*:refs/heads/*",
		},
	})
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			Info("Repository [%s] is already up-to-date.", c.LocalPath)
			return nil
		}
		Error(err)
	}
	Info("Repository [%s] has been updated.", c.LocalPath)
	return nil
}

// get all remote branches
func GitRemoteBranchList(c GitConfig) ([]string, error) {
	// 打开本地 Git 仓库
	repo, err := git.PlainOpen(c.LocalPath)
	if err != nil {
		Error(err)
		return nil, err
	}

	var branches []string

	refIter, err := repo.References()
	if err != nil {
		Error(err)
		return nil, err
	}

	// 遍历远程分支，输出分支名称
	err = refIter.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}

		// 判断是否为远程分支
		if ref.Name().IsRemote() {
			Info("find a remote branch name: %s", ref.Name().Short())
			branches = append(branches, ref.Name().Short())
		}

		return nil
	})
	if err != nil {
		Error(err)
		return nil, err
	}

	return branches, nil
}
