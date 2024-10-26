package mapper

import (
	"strings"

	c "github.com/arne-zillhardt/todo/pkg/command"
)

func GetCommandFromString(input string) c.Command {
	userInput := strings.ToLower(input)
	for _, command := range c.GetCommands() {
		activation := strings.ToLower(command.GetActivation())

		if activation == userInput {
			return command
		}
	}
	return nil
}
