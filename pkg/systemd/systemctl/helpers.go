package systemctl

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/systemd/systemctl/properties"
	"strconv"
	"time"
)

const dateFormat = "Mon 2006-01-02 15:04:05 MST"

// GetStartTime Get start time of a service (`systemctl show [unit] --property ExecMainStartTimestamp`) as a `Time` type
func GetStartTime(unit string, opts Options) (time.Time, error) {
	value, err := Show(unit, properties.ExecMainStartTimestamp, opts)

	if err != nil {
		return time.Time{}, err
	}
	// ExecMainStartTimestamp returns an empty string if the unit is not running
	if value == "" {
		return time.Time{}, ErrUnitNotActive
	}
	return time.Parse(dateFormat, value)
}

// GetNumRestarts Get the number of times a process restarted (`systemctl show [unit] --property NRestarts`) as an int
func GetNumRestarts(unit string, opts Options) (int, error) {
	value, err := Show(unit, properties.NRestarts, opts)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(value)
}

// GetMemoryUsage Get current memory in bytes (`systemctl show [unit] --property MemoryCurrent`) an an int
func GetMemoryUsage(unit string, opts Options) (int, error) {
	value, err := Show(unit, properties.MemoryCurrent, opts)
	if err != nil {
		return -1, err
	}
	if value == "[not set]" {
		return -1, ErrValueNotSet
	}
	return strconv.Atoi(value)
}

// GetPID Get the PID of the main process (`systemctl show [unit] --property MainPID`) as an int
func GetPID(unit string, opts Options) (int, error) {
	value, err := Show(unit, properties.MainPID, opts)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(value)
}

//setTimeout limit with the timeout can be
func setTimeout(timeOut int) time.Duration {
	if timeOut <= 0 || timeOut >= 10 {
		timeOut = 10
	}
	return time.Duration(timeOut)
}
