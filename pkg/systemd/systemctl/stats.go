package systemctl

import (
	"context"
	"strings"
	"time"
)

// UnitFileState is state type for Systemctl
type UnitFileState string

// ActiveState is state type for Systemctl
type ActiveState string

// SubState is state type for Systemctl
type SubState string

const (
	// Enabled is a state reported by systemctl
	Enabled = UnitFileState("enabled")

	// Disabled is a state reported by systemctl
	Disabled = UnitFileState("disabled")

	// Active is a state reported by systemctl
	Active = ActiveState("active")

	// Inactive is a state reported by systemctl
	Inactive = ActiveState("inactive")

	// Running is a substate reported by systemctl
	Running = SubState("running")

	// Dead is a substate reported by systemctl
	Dead = SubState("dead")
)

type SystemStats struct {
	State                  UnitFileState //enabled, disabled
	ActiveState            ActiveState   // active, inactive
	SubState               SubState      //running, //dead
	ActiveEnterTimestamp   string
	InactiveEnterTimestamp string
	Restarts               string //NRestarts number of restart
}

// Stats get status
func Stats(unit string, opts Options) (SystemStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(opts.Timeout)*time.Second)
	defer cancel()
	var args = []string{"show", unit, "--no-page"}
	if opts.UserMode {
		args[1] = "--user"
	}
	stdout, _, _, err := execute(ctx, args)
	stats := SystemStats{}

	unitState := UnitFileState("")
	activeState := ActiveState("")
	subState := SubState("")

	for _, line := range strings.Split(stdout, "\n") {
		fields := strings.SplitN(line, "=", 2)
		if len(fields) != 2 {
			continue
		}
		switch fields[0] {
		case "UnitFileState":
			unitState = UnitFileState(fields[1])
		case "ActiveState":
			activeState = ActiveState(fields[1])
		case "SubState":
			subState = SubState(fields[1])
		case "ActiveEnterTimestamp":
			stats.ActiveEnterTimestamp = fields[1]
		case "InactiveEnterTimestamp":
			stats.InactiveEnterTimestamp = fields[1]
		case "NRestarts":
			stats.Restarts = fields[1]
		default:
			// ignore
		}
	}

	stats.State = unitState
	stats.ActiveState = activeState
	stats.SubState = subState
	return stats, err
}
