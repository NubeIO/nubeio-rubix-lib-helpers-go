package validation

import (
	"errors"
	"net"
)

func IsIPSubnet(address string) (bool, error) {
	mask := net.IPMask(net.ParseIP(address).To4()) // If you have the mask as a string
	prefixSize, _ := mask.Size()
	if prefixSize == 0 {
		return false, errors.New("invalid subnet address, 255.255.255.0")
	}
	return true, nil
}
