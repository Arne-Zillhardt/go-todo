package command

import "fmt"

type Help struct {
}

func (h Help) Execute() {
	for _, command := range GetCommands() {
		fmt.Println(command.GetActivation())
	}
}

func (h Help) GetActivation() string {
	return "help"
}
