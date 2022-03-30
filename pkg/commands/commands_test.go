package commands

import (
	"fmt"
	"testing"
)

func TestCommands(*testing.T) {

	n := New(&Commands{}).GetHomeDir().AsString()
	fmt.Println(n)

	n = New(&Commands{Command: CommandsList.GetHomeDir}).RunCommand().AsString()

}
