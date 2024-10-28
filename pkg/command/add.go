package command

import (
	"github.com/arne-zillhardt/todo/pkg/io/console"
	"github.com/arne-zillhardt/todo/pkg/io/file"
	"github.com/arne-zillhardt/todo/pkg/mapper"
)

type Add struct{}

func (a Add) GetActivation() string {
	return "add"
}

func (a Add) Execute() {
	todo := mapper.Todo{}
	todo.Id = getNewId()
	todo.Name = console.GetUserInput("Name: ")
	todo.Task = console.GetUserInput("Task: ")
	todo.Done = false

	file.WriteNewTaskToTheTodoFile(mapper.MapToCSV(todo))
}

func getNewId() int {
	todos := file.GetReadFile()
	if len(todos) == 1 {
		return 1
	}

	lastTodo := mapper.MapToToDo(todos[len(todos)-1])
	return lastTodo.Id + 1
}
