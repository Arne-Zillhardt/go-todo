package mapper_test

import (
	"testing"

	"github.com/arne-zillhardt/todo/pkg/mapper"
	"github.com/arne-zillhardt/todo/test"
)

func TestMapToToDo(t *testing.T) {
	input := []string{"1", "Test", "Just testing", "false"}
	result := mapper.MapToToDo(input)

	test.AssertEquals(1, result.Id, t)
	test.AssertEquals("Test", result.Name, t)
	test.AssertEquals("Just testing", result.Task, t)
	test.AssertFalse(result.Done, t)
}

func TestMapMultipleToToDo(t *testing.T) {
	input := [][]string{{"1", "Test", "Just testing", "false"}, {"2", "Test", "Just another test", "true"}}
	result := mapper.MapMultipleToToDo(input)

	test.AssertEquals(1, result[0].Id, t)
	test.AssertEquals("Test", result[0].Name, t)
	test.AssertEquals("Just testing", result[0].Task, t)
	test.AssertFalse(result[0].Done, t)

	test.AssertEquals(2, result[1].Id, t)
	test.AssertEquals("Test", result[1].Name, t)
	test.AssertEquals("Just another test", result[1].Task, t)
	test.AssertTrue(result[1].Done, t)
}
