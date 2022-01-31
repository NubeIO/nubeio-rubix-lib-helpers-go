package validation

import (
	"errors"
	netval "github.com/THREATINT/go-net"
)

func IsURL(address string) (bool, error) {
	if netval.IsURL(address) {
		return true, nil
	}
	return false, errors.New("invalid address, www.nube-io.com")
}
