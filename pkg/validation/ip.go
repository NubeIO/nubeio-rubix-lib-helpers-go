package validation

import (
	"errors"
	netval "github.com/THREATINT/go-net"
)

func IsIPAddr(address string) (bool, error) {
	if netval.IsIPAddr(address) {
		return true, nil
	}
	return false, errors.New("invalid ip address, try 192.168.15.15")
}
