package cli

import (
	"bufio"
	"os"
	"strings"

	"github.com/jahidxuddin/git-fast-clone/internal/utils"
	"github.com/rivo/tview"
)

func PromptGitHubUsername() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	print("Enter your GitHub username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(username), nil
}

func PromptRepository(repositories []utils.Repository) (string, error) {
	app := tview.NewApplication()

	list := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Select a repository", "", 0, nil)

	for _, repository := range repositories {
		list.AddItem(repository.Name, "", 0, nil)
	}

	var selectedRepository string
	list.SetSelectedFunc(func(i int, _ string, _ string, _ rune) {
		if i == 0 {
			selectedRepository = repositories[0].CloneURL
		} else {
			selectedRepository = repositories[i-1].CloneURL
		}

		app.Stop()
	})

	if err := app.SetRoot(list, true).Run(); err != nil {
		return "", err
	}

	return selectedRepository, nil
}
