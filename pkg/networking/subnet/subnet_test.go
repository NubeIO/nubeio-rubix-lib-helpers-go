package subnet

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	ipv4, err := SubnetIPV4("192.168.15.1", 24)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ipv4)
	}

}
