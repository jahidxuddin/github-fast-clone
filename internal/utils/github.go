package utils

import (
	"encoding/json"
	"net/http"
)

type Repository struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
}

func FetchRepositories(token string) ([]Repository, error) {
	var request *http.Request

	url := "https://api.github.com/user/repos?type=all"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(request)
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
