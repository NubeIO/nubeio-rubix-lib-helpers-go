package nsystemd

import (
	"context"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/systemd/systemctl"

	"time"
)

type SystemD struct {
	ServiceName string
	Timeout     int
}

type Result struct {
	active bool
	err    error
}

func New() *SystemD {
	return &SystemD{}
}

func (service *SystemD) IsActive() (result *Result) {
	result = &Result{}
	ctx, cancel := context.WithTimeout(context.Background(), setTimeout(service.Timeout)*time.Second)
	defer cancel()
	opts := systemctl.Options{UserMode: false}
	unit := service.ServiceName
	result.active, result.err = systemctl.IsActive(ctx, unit, opts)
	return result

}

// Active returns
func (e *Result) Active() bool {
	return e.active
}

// Error returns
func (e *Result) Error() error {
	return e.err
}

//setTimeout limit with the timeout can be
func setTimeout(timeOut int) time.Duration {
	if timeOut <= 0 || timeOut >= 10 {
		timeOut = 10
	}
	return time.Duration(timeOut)
}
