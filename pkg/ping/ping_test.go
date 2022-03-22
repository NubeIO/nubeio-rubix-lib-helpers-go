package ping

import (
	"fmt"
	"testing"
)

func TestLinuxPing(*testing.T) {
	ping, err := LinuxPing("192.168.15.1", 1, 1, 1)
	if err != nil {
		return
	}
	fmt.Println(ping, err)
}
