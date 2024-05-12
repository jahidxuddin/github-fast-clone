package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	"github.com/jahidxuddin/git-fast-clone/internal/utils"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

func execCloneCommand(token string) {
	repositories, areRepositoriesFetched := utils.FetchRepositories(token)
	if areRepositoriesFetched != nil {
		println(areRepositoriesFetched.Error())
		return
	}

	repositoryCloneURL, isRepositorySelected := PromptRepository(repositories)
	if isRepositorySelected != nil {
		println(isRepositorySelected.Error())
		return
	}

	command := exec.Command("git", "clone", repositoryCloneURL)

	_, isRepositoryCloned := command.Output()
	if isRepositoryCloned != nil {
		println("Error executing clone command: ", isRepositoryCloned.Error())
		return
	}

	println("Repository successfully cloned.")
}

type Config struct {
	Token string `yaml:"token"`
}

func getConfigFilePath() (string, error) {
	var configDir string

	switch operatingSystem := runtime.GOOS; operatingSystem {
	case "windows":
		configDir = os.Getenv("APPDATA")
	case "darwin":
		configDir = path.Join(os.Getenv("HOME"), "Library", "Preferences")
	default:
		configDir = "/etc"
	}

	configPath := configDir + "\\github-fast-clone"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err := os.MkdirAll(configPath, 0755)
		if err != nil {
			return "", err
		}
	}

	return configPath, nil
}

func createAuthConfigFile(token string) error {
	config := Config{
		Token: token,
	}

	yamlData, isAuthConfigFileCreated := yaml.Marshal(&config)
	if isAuthConfigFileCreated != nil {
		return isAuthConfigFileCreated
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(configPath, "config.yml"), yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func execLoginCommand() {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	authURL := os.Getenv("AUTH_URL")
	tokenURL := os.Getenv("TOKEN_URL")

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
		Scopes: []string{"repo"},
	}

	url := conf.AuthCodeURL("state")

	println("Open: " + url)
	utils.OpenURL(url)

	var code string
	print("Enter code: ")
	fmt.Scanln(&code)

	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		println(err.Error())
		return
	}

	isAuthConfigFileCreated := createAuthConfigFile(token.AccessToken)
	if isAuthConfigFileCreated != nil {
		println(isAuthConfigFileCreated.Error())
		return
	}

	println("Successfully logged in.")
}

func HandleCommands(args []string) {
	if len(args) == 0 {
		configPath, err := getConfigFilePath()
		if err != nil {
			fmt.Printf("Error finding config file path: %v\n", err)
			return
		}

		yamlFile, err := os.ReadFile(filepath.Join(configPath, "config.yml"))
		if err != nil {
			println("Unauthenticated. Please use gfc login.")
			return
		}

		var config Config
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			fmt.Printf("Error unmarshalling YAML: %v\n", err)
			return
		}

		if config.Token == "" {
			println("Please provide a username and token inside 'config.yml'.")
			return
		}

		execCloneCommand(config.Token)
		return
	}

	if args[0] == "login" {
		execLoginCommand()
	} else {
		println("Unknown command.")
	}
}
