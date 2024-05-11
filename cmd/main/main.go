package main

import (
	"os"

	"github.com/jahidxuddin/git-fast-clone/internal/cli"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		println("Too many arguments.")
		return
	}

	if err := godotenv.Load(); err != nil {
		println("Error loading .env file: " + err.Error())
	}

	cli.HandleCommands(args)
}
