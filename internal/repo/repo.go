package repo

import (
	"local-build/internal/lblog"
	"local-build/internal/pkg/env"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type GitConfig struct {
	Url           string
	LocalPath     string
	AccessType    string
	UserName      string
	Password      string
	SshPrivateKey string
	KeyPassphrase string
	AccessToken   string
}

// update git repo
func UpdateGitRepo(c GitConfig) error {
	// check if ${path}/.git exist
	dir := filepath.Join(c.LocalPath, ".git")
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			lblog.Warning("local git repo [%s] does not exist on [%s].", c.Url, c.LocalPath)
			err := GitClone(c)
			if err != nil {
				lblog.Error("git clone [%s] [%s] error: %s", c.Url, c.LocalPath, err)
				return err
			}
		} else {
			lblog.Error("the repo local path [%s] error: %s ", c.LocalPath, err)
			return err
		}
	}

	// git fetch all
	err = GitFetchAll(c)
	if err != nil {
		lblog.Error("git fetch all repo [%s] [%s] error: %s ", c.Url, c.LocalPath, err)
		return err
	}
	return nil
}

// init auth
func getAuth(c GitConfig) (transport.AuthMethod, error) {
	lblog.Info("use [%s] to access repo [%s]", c.AccessType, c.Url)
	switch c.AccessType {
	case env.CREDENTIALS:
		return &http.BasicAuth{
			Username: c.UserName,
			Password: c.Password,
		}, nil
	case env.SSH_PRIVATE_KEY:
		if c.UserName == "" {
			c.UserName = "git"
			lblog.Info("use default user: [git] to SSH private key access repo")
		}
		auth, err := ssh.NewPublicKeysFromFile(c.UserName, c.SshPrivateKey, c.KeyPassphrase)
		if err != nil {
			lblog.Error("ssh private key auth error: %s", err)
			return nil, err
		}
		return auth, nil
	case env.ACCESS_TOKEN:
		return &http.TokenAuth{
			Token: c.AccessToken,
		}, nil
	}
	return nil, nil
}

// use go-git clone repo
func GitClone(c GitConfig) error {
	// Clone the given repository to the given directory
	lblog.Info("git clone [%s] [%s] --recursive", c.Url, c.LocalPath)

	auth, err := getAuth(c)
	if err != nil {
		lblog.Error("get git auth error: %s", err)
		return err
	}

	_, err = git.PlainClone(c.LocalPath, false, &git.CloneOptions{
		URL:               c.Url,
		Auth:              auth,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		lblog.Error("git plain clone error: %s", err)
		return err
	}

	lblog.Info("Repository [%s] has been cloned into [%s].", c.Url, c.LocalPath)
	return nil
}

// fetch all
func GitFetchAll(c GitConfig) error {
	// 打开本地 Git 仓库
	repo, err := git.PlainOpen(c.LocalPath)
	if err != nil {
		lblog.Error("git plan open error: %s", err)
		return err
	}

	auth, err := getAuth(c)
	if err != nil {
		lblog.Error("get git auth error: %s", err)
		return err
	}

	// 获取所有远程分支的最新代码
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{
			"+refs/heads/*:refs/heads/*",
		},
		Auth: auth,
	})
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			lblog.Info("Repository [%s] is already up-to-date.", c.LocalPath)
			return nil
		}
		lblog.Error("git fetch error: %s", err)
		return err
	}
	lblog.Info("Repository [%s] has been updated.", c.LocalPath)
	return nil
}

// get all remote branches
func GitRemoteBranchList(c GitConfig) ([]string, error) {
	// 打开本地 Git 仓库
	repo, err := git.PlainOpen(c.LocalPath)
	if err != nil {
		lblog.Error("git plan open error: %s", err)
		return nil, err
	}

	var branches []string

	refIter, err := repo.References()
	if err != nil {
		lblog.Error("get repo references error: %s", err)
		return nil, err
	}

	// 遍历远程分支，输出分支名称
	err = refIter.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}

		// 判断是否为远程分支
		if ref.Name().IsRemote() {
			lblog.Info("find a remote branch name: %s", ref.Name().Short())
			branches = append(branches, ref.Name().Short())
		}

		return nil
	})
	if err != nil {
		lblog.Error("foreach references error: %s", err)
		return nil, err
	}

	return branches, nil
}
