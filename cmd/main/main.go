package main

import (
	"os"
	"os/exec"

	"github.com/jahidxuddin/git-fast-clone/internal/cli"
	"github.com/jahidxuddin/git-fast-clone/internal/utils"
)

func main() {
	args := os.Args[1:]

	gitHubUsername, isGitHubUsernameParsed := cli.ParseGitHubUsernameFromArgs(args)
	if isGitHubUsernameParsed != nil {
		println(isGitHubUsernameParsed.Error())
		return
	}

	repositories, areRepositoriesFetched := utils.FetchRepositories(gitHubUsername)
	if areRepositoriesFetched != nil {
		println(areRepositoriesFetched.Error())
		return
	}

	repositoryCloneURL, isRepositorySelected := cli.PromptRepository(repositories)
	if isRepositorySelected != nil {
		println(isRepositorySelected.Error())
		return
	}

	command := exec.Command("git", "clone", repositoryCloneURL)

	_, isRepositoryCloned := command.Output()
	if isRepositoryCloned != nil {
		println("Error executing command: ", isRepositoryCloned.Error())
		return
	}

	println("Repository successfully cloned.")
}
