package command_test

import (
	"testing"

	"github.com/arne-zillhardt/todo/pkg/command"
	"github.com/arne-zillhardt/todo/test"
)

func TestGetCommand(t *testing.T) {
	test.AssertEquals("help", command.GetCommand("HeLP").GetActivation(), t)
	test.AssertEquals("add", command.GetCommand("ADd").GetActivation(), t)
	test.AssertEquals("remove", command.GetCommand("RemOVe").GetActivation(), t)
	test.AssertEquals("update", command.GetCommand("UPDaTe").GetActivation(), t)
	test.AssertNull(command.GetCommand(""), t)
}
