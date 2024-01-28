package test

import (
	"local-build/internal/lblog"
	"local-build/internal/pkg/env"
	"local-build/internal/repo"
	"local-build/internal/service"
	"local-build/internal/store/model"
	"testing"
)

func TestProjectParse(t *testing.T) {
	project := model.Project{
		Path: "/tmp/test/hello",
		RepoConfig: model.RepoConfig{
			AccessType: env.GIT,
		},
	}
	service.ParseProject(project)
}

func TestGitClone(t *testing.T) {
	// https://gitee.com/hotstrip/hello-world-java.git
	// git@gitee.com:hotstrip/hello-world-java.git
	c := repo.GitConfig{
		Url:           "https://gitee.com/hotstrip/hello-world-java.git",
		LocalPath:     "/tmp/test/hello2",
		AccessType:    env.SSH_PRIVATE_KEY,
		SshPrivateKey: "/Users/zhangxin/.ssh/id_rsa",
		KeyPassphrase: "12313",
	}
	err := repo.GitClone(c)
	if err != nil {
		panic(err)
	}
}

func TestGitFetchAll(t *testing.T) {
	c := repo.GitConfig{
		Url:       "https://gitee.com/hotstrip/hello-world-java.git",
		LocalPath: "/tmp/test/hello",
	}
	err := repo.GitFetchAll(c)
	if err != nil {
		panic(err)
	}
}

func TestGitRemoteBranchList(t *testing.T) {
	c := repo.GitConfig{
		Url:       "https://gitee.com/hotstrip/hello-world-java.git",
		LocalPath: "/tmp/test/hello",
	}
	branches, err := repo.GitRemoteBranchList(c)
	if err != nil {
		panic(err)
	}

	for _, b := range branches {
		lblog.Info("branch: %s", b)
	}
}
