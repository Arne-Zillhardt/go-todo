package mapper_test

import (
	"testing"

	"github.com/arne-zillhardt/todo/pkg/command"
	"github.com/arne-zillhardt/todo/pkg/mapper"
	"github.com/arne-zillhardt/todo/test"
)

func TestGetCommandFromString(t *testing.T) {
	input := "hElP"

	command.PopulateCommands()
	output := mapper.GetCommandFromString(input)

	test.AssertEquals("help", output.GetActivation(), t)

	input = ""

	output = mapper.GetCommandFromString(input)

	test.AssertNull(output, t)
}
