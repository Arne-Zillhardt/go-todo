package command

import (
	"fmt"
	"strconv"

	"github.com/arne-zillhardt/todo/pkg/io/console"
	"github.com/arne-zillhardt/todo/pkg/io/file"
)

type Remove struct{}

func (r Remove) GetActivation() string {
	return "remove"
}

func (r Remove) Execute() {
	idToRemove, err := strconv.Atoi(console.GetUserInput("Id to remove: "))
	if err != nil {
		fmt.Println("Wrong input. Number expected")
		return
	}

	id := findRowWithId(idToRemove)
	if id == -1 {
		fmt.Println("!No todo with this id found!")
		return
	}

	file.RemoveTaskInTheTodoFile(id)
}

func findRowWithId(id int) int {
	rows := file.GetReadFile()

	for i, line := range rows {
		if line[0] == strconv.Itoa(id) {
			return i
		}
	}

	return -1
}
