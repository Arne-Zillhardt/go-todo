package file_test

import (
	"os"
	"testing"

	"github.com/arne-zillhardt/todo/pkg/io/file"
	"github.com/arne-zillhardt/todo/test"
)

var filePath string = "../../../../assets/todo.csv"

func TestMain(m *testing.M) {
	setUp()
	ex := m.Run()
	teardown()
	os.Exit(ex)
}

func setUp() {
	test.SaveFile(filePath)
}

func teardown() {
	test.WriteSavedFile(filePath)
}

func TestReadFile(t *testing.T) {
	result := file.GetReadFile()
	test.AssertEquals("id", result[0][0], t)
	test.AssertEquals("name", result[0][1], t)
	test.AssertEquals("task", result[0][2], t)
	test.AssertEquals("done", result[0][3], t)

	test.AssertEquals("1", result[1][0], t)
	test.AssertEquals("Test", result[1][1], t)
	test.AssertEquals("Test, the reading of documents", result[1][2], t)
	test.AssertEquals("false", result[1][3], t)
}

func TestWriteNewTaskToFile(t *testing.T) {
	input := []string{"2", "Test", "Test, the writing of documents", "false"}
	file.WriteNewTaskToTheTodoFile(input)

	result := file.GetReadFile()
	index := len(result) - 1
	test.AssertEquals("2", result[index][0], t)
	test.AssertEquals("Test", result[index][1], t)
	test.AssertEquals("Test, the writing of documents", result[index][2], t)
	test.AssertEquals("false", result[index][3], t)
}

func TestUpdateTaskInTheTodoFile(t *testing.T) {
	input := []string{"3", "Test", "Test, the writing of documents", "false"}
	file.WriteNewTaskToTheTodoFile(input)

	input[2] = "Test the updating of documents"
	file.UpdateTaskInTheTodoFile(input)
	result := file.GetReadFile()
	index := len(result) - 1
	test.AssertEquals("3", result[index][0], t)
	test.AssertEquals("Test", result[index][1], t)
	test.AssertEquals("Test the updating of documents", result[index][2], t)
	test.AssertEquals("false", result[index][3], t)
}

func TestRemoveTaskInTheTodoFile(t *testing.T) {
	input := []string{"4", "Test", "Test, the writing of documents", "false"}
	file.WriteNewTaskToTheTodoFile(input)

	file.RemoveTaskInTheTodoFile(4)
	result := file.GetReadFile()
	index := len(result) - 1
	test.AssertEquals("3", result[index][0], t)
	test.AssertEquals("Test", result[index][1], t)
	test.AssertEquals("Test the updating of documents", result[index][2], t)
	test.AssertEquals("false", result[index][3], t)
}
