package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
}

func DoesGitHubUserExist(username string) error {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	
	if resp.StatusCode == http.StatusOK {
		return nil
	} else if resp.StatusCode == http.StatusNotFound {
		return NewError("GitHub user does not exist.")
	} else {
		return NewError(resp.Status)
	}
}

func FetchRepositories(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var repositories []Repository
	err = json.NewDecoder(resp.Body).Decode(&repositories)
	if err != nil {
		return nil, err
	}

	return repositories, nil
}
