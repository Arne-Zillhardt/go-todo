package command

import (
	"fmt"
	"strconv"

	"github.com/arne-zillhardt/todo/pkg/io/console"
	"github.com/arne-zillhardt/todo/pkg/io/file"
	"github.com/arne-zillhardt/todo/pkg/mapper"
)

type Update struct{}

func (u Update) GetActivation() string {
	return "update"
}

func (u Update) Execute() {
	idToUpdate, _ := strconv.Atoi(console.GetUserInput("Id to update: "))
	id := findRowWithId(idToUpdate)
	if id == -1 {
		fmt.Println("!No todo with this id found!")
		return
	}

	todo := mapper.MapToToDo(file.GetReadFile()[id])

	name := console.GetUserInput("New Name (leave blank to keep original): ")
	task := console.GetUserInput("New Task (leave blank to keep original): ")
	done := console.GetUserInput("Done (leave blank to keep original): ")

	updateName(name, &todo)
	updateTask(task, &todo)
	err := updateDone(done, &todo)
	if err != nil {
		fmt.Println("Wrong input. Boolean expected")
		return
	}

	file.UpdateTaskInTheTodoFile(mapper.MapToCSV(todo))
}

func updateName(name string, todo *mapper.Todo) {
	if keepOriginal(name) {
		return
	}

	todo.Name = name
}

func updateTask(task string, todo *mapper.Todo) {
	if keepOriginal(task) {
		return
	}

	todo.Task = task
}

func updateDone(done string, todo *mapper.Todo) error {
	if keepOriginal(done) {
		return nil
	}

	newDone, err := strconv.ParseBool(done)
	if err != nil {
		return err
	}

	todo.Done = newDone
	return nil
}

func keepOriginal(input string) bool {
	return input == ""
}
