package validation

import (
	"errors"
	netval "github.com/THREATINT/go-net"
)

func IsDomain(address string) (bool, error) {
	if netval.IsDomain(address) {
		return true, nil
	}
	return false, errors.New("invalid address, nube-io.com")
}
