package command

import "strings"

type Command interface {
	Execute()
	GetActivation() string
}

var commands []Command

func populateCommands() {
	if len(commands) != 0 {
		return
	}

	commands = []Command{
		Help{},
		Add{},
		Remove{},
		Update{},
		Show{},
	}
}

func GetCommand(input string) Command {
	populateCommands()
	userInput := strings.ToLower(input)
	for _, command := range commands {
		activation := strings.ToLower(command.GetActivation())

		if activation == userInput {
			return command
		}
	}
	return nil
}

func GetCommands() []Command {
	return commands
}
