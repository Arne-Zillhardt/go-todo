package command

import (
	"fmt"
	"strconv"

	"github.com/arne-zillhardt/todo/pkg/io/console"
	"github.com/arne-zillhardt/todo/pkg/io/file"
	"github.com/arne-zillhardt/todo/pkg/mapper"
)

type Show struct{}

func (s Show) GetActivation() string {
	return "show"
}

func (s Show) Execute() {
	id := console.GetUserInput("Id of task (leave blank for all): ")
	if id == "" {
		showAll()
		return
	}

	idAsInt, _ := strconv.Atoi(id)
	showId(idAsInt)
}

func showAll() {
	readFile := file.GetReadFile()
	var toParse [][]string

	for i := 0; i < len(readFile); i++ {
		if i == 0 {
			continue
		}

		toParse = append(toParse, readFile[i])
	}

	parsedTodos := mapper.MapMultipleToToDo(toParse)
	for _, todo := range parsedTodos {
		printTodo(todo)
	}
}

func showId(id int) {
	rowId := findRowWithId(id)
	if rowId == -1 {
		fmt.Println("!No todo with this id found!")
		return
	}

	todo := mapper.MapToToDo(file.GetReadFile()[rowId])
	printTodo(todo)
}

func printTodo(todo mapper.Todo) {
	fmt.Printf("ID: %d\n", todo.Id)
	fmt.Printf("Name: %s\n", todo.Name)
	fmt.Printf("Task: %s\n", todo.Task)
	fmt.Printf("Done: %t\n", todo.Done)
	fmt.Println()
}
