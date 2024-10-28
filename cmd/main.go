package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/arne-zillhardt/todo/pkg/command"
	"github.com/arne-zillhardt/todo/pkg/io/console"
)

func main() {
	clearConsole()
	for {

		userInput := console.GetUserInput("> ")
		if userInput == "q" {
			fmt.Println("See you later :)")
			os.Exit(0)
		}

		command := command.GetCommand(userInput)
		if command == nil {
			fmt.Println("No command '", userInput, "' found")
			continue
		}
		clearConsole()
		command.Execute()
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
