package freeport

import (
	"fmt"
	"testing"
)

func TestLinuxPingPort(*testing.T) {

	port, err := FindFreePort(1883)

	fmt.Println(port)
	fmt.Println(err)

}
