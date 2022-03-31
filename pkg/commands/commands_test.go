package commands

import (
	"testing"
)

func TestCommands(*testing.T) {
	New(&Commands{Debug: true}).ChainCommand(CommandsList.PWD).ChainCommand(CommandsList.GetHomeDir).Result().Log()
}
