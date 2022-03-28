package systemctl

import (
	"bytes"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"regexp"
)

var systemctl string

const killed = 130

func init() {
	path, err := exec.LookPath("systemctl")
	if err != nil {
		log.Printf("%v", ErrNotInstalled)
		systemctl = ""
		return
	}
	systemctl = path
}

func execute(ctx context.Context, args []string) (string, string, int, error) {
	var (
		err      error
		stderr   bytes.Buffer
		stdout   bytes.Buffer
		code     int
		output   string
		warnings string
	)

	if systemctl == "" {
		log.Errorln(ErrNotInstalled)
		return "", "", 0, ErrNotInstalled
	}
	cmd := exec.CommandContext(ctx, systemctl, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	output = stdout.String()
	warnings = stderr.String()
	code = cmd.ProcessState.ExitCode()

	customErr := filterErr(warnings)
	if customErr != nil {
		err = customErr
	}
	if code != 0 && err == nil {
		err = fmt.Errorf("received error code %d for stderr `%s`: %w", code, warnings, ErrUnspecified)
	}

	return output, warnings, code, err
}

func filterErr(stderr string) error {
	if matched, _ := regexp.MatchString(`does not exist`, stderr); matched {
		return ErrDoesNotExist
	}
	if matched, _ := regexp.MatchString(`not found.`, stderr); matched {
		return ErrDoesNotExist
	}
	if matched, _ := regexp.MatchString(`not loaded.`, stderr); matched {
		return ErrUnitNotLoaded
	}
	if matched, _ := regexp.MatchString(`No such file or directory`, stderr); matched {
		return ErrDoesNotExist
	}
	if matched, _ := regexp.MatchString(`Interactive authentication required`, stderr); matched {
		return ErrInsufficientPermissions
	}
	if matched, _ := regexp.MatchString(`Access denied`, stderr); matched {
		return ErrInsufficientPermissions
	}
	if matched, _ := regexp.MatchString(`DBUS_SESSION_BUS_ADDRESS`, stderr); matched {
		return ErrBusFailure
	}
	if matched, _ := regexp.MatchString(`is masked`, stderr); matched {
		return ErrMasked
	}
	if matched, _ := regexp.MatchString(`Failed`, stderr); matched {
		return ErrUnspecified
	}
	return nil
}
