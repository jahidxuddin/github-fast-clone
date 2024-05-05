package cli

import "github.com/jahidxuddin/git-fast-clone/internal/utils"

func ParseGitHubUsernameFromArgs(args []string) (string, error) {
	if len(args) > 1 {
		return "", utils.NewError("Too many arguments. Expected one.")
	}

	var gitHubUsername string
	if len(args) == 0 {
		name, err := PromptGitHubUsername()
		if err != nil {
			return "", err
		}
		gitHubUsername = name
	} else {
		gitHubUsername = args[0]
	}

	if err := utils.DoesGitHubUserExist(gitHubUsername); err != nil {
		return "", err
	}

	return gitHubUsername, nil
}
