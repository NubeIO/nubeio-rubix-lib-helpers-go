package ip_helpers

import (
	"fmt"
	"testing"
)

func TestNetIP(*testing.T) {

	fmt.Println(GetIPSubnet("192.168.15.1", "255.255.0.0"))
	fmt.Println(SubnetString(6))
}
