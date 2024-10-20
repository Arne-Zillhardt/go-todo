package file_test

import (
	"testing"

	"github.com/arne-zillhardt/todo/pkg/io/file"
	"github.com/arne-zillhardt/todo/test"
)

func TestReadFile(t *testing.T) {
	result := file.ReadFile()
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
	file.WriteNewTaskToFile(input)

	result := file.ReadFile()
	index := len(result) - 1
	test.AssertEquals("2", result[index][0], t)
	test.AssertEquals("Test", result[index][1], t)
	test.AssertEquals("Test, the writing of documents", result[index][2], t)
	test.AssertEquals("false", result[index][3], t)
}
