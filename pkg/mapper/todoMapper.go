package mapper

import (
	"log"
	"strconv"
)

type Todo struct {
	Id   int
	Name string
	Task string
	Done bool
}

func MapToToDo(csvTodo []string) (todo Todo) {
	var err error
	todo.Id, err = strconv.Atoi(csvTodo[0])
	if err != nil {
		log.Fatal(err)
	}

	todo.Name = csvTodo[1]
	todo.Task = csvTodo[2]

	todo.Done, err = strconv.ParseBool(csvTodo[3])
	if err != nil {
		log.Fatal(err)

	}

	return
}

func MapMultipleToToDo(csvTodos [][]string) (todos []Todo) {
	for _, line := range csvTodos {
		todos = append(todos, MapToToDo(line))
	}

	return
}

func MapToCSV(todo Todo) (csvTodo []string) {
	csvTodo = append(csvTodo, strconv.Itoa(todo.Id))
	csvTodo = append(csvTodo, todo.Name)
	csvTodo = append(csvTodo, todo.Task)
	csvTodo = append(csvTodo, strconv.FormatBool(todo.Done))

	return
}
