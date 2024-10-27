package main

import (
	"fmt"
	"os"

	"github.com/arne-zillhardt/todo/pkg/command"
	"github.com/arne-zillhardt/todo/pkg/io/console"
)

func main() {
	for {
		userInput := console.GetUserInput("> ")
		if userInput == "q" {
			fmt.Println("See you later :)")
			os.Exit(0)
		}

		command := command.GetCommand(userInput)
		command.Execute()
	}
}
