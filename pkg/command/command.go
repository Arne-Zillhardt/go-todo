package command

type Command interface {
	Execute()
	GetActivation() string
}

var commands []Command

func PopulateCommands() {
	commands = []Command{
		Help{},
	}
}

func GetCommands() []Command {
	return commands
}
