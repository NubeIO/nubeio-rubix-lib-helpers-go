package portscanner

import (
	"fmt"
	"testing"
)

func TestPortScanner(t *testing.T) {

	ports := []string{"22", "1414", "1883", "1660", "502", "80", "7"}

	// IP sequence is defined by a '-' between first and last IP address .
	ipsSequence := []string{"192.168.15.1-254"}

	// result returns a map with open ports for each IP address.
	results := IPScanner(ipsSequence, ports, false)

	fmt.Println("-------------HOSTS------------------")
	fmt.Println(results)
	fmt.Println("-------------HOSTS------------------")

}
