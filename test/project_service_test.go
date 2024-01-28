package test

import (
	"local-build/internal/lblog"
	"local-build/internal/pkg/env"
	"local-build/internal/service"
	"local-build/internal/store/model"
	"local-build/internal/utils"
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
	c := utils.GitConfig{
		Url:           "https://gitee.com/hotstrip/hello-world-java.git",
		LocalPath:     "/tmp/test/hello2",
		AccessType:    env.SSH_PRIVATE_KEY,
		SshPrivateKey: "/Users/zhangxin/.ssh/id_rsa",
		KeyPassphrase: "12313",
	}
	err := utils.GitClone(c)
	if err != nil {
		panic(err)
	}
}

func TestGitFetchAll(t *testing.T) {
	c := utils.GitConfig{
		Url:       "https://gitee.com/hotstrip/hello-world-java.git",
		LocalPath: "/tmp/test/hello",
	}
	err := utils.GitFetchAll(c)
	if err != nil {
		panic(err)
	}
}

func TestGitRemoteBranchList(t *testing.T) {
	c := utils.GitConfig{
		Url:       "https://gitee.com/hotstrip/hello-world-java.git",
		LocalPath: "/tmp/test/hello",
	}
	branches, err := utils.GitRemoteBranchList(c)
	if err != nil {
		panic(err)
	}

	for _, b := range branches {
		lblog.Info("branch: %s", b)
	}
}
