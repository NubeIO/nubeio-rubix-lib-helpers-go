package validation

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	ipv4, err := IsIPSubnet("255.255.255.0")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ipv4)
	}

}
