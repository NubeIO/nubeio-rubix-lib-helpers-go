package ping

import (
	"fmt"
	"testing"
)

func TestLinuxPingHost(*testing.T) {
	opts := LinuxPingOptions{
		Host: "192.168.15.1",
	}
	ping, err := LinuxPingHost(opts)
	if err != nil {
		fmt.Println("LinuxPingHost", err)
		return
	}
	fmt.Println("-----------HOST------------")
	fmt.Println(ping.Host, ping.Found)
}

func TestLinuxPingHosts(*testing.T) {
	opts := LinuxPingOptions{
		Hosts: []string{"192.168.15.1", "192.168.15.0"},
	}
	ping := LinuxPingHosts(opts)
	if ping.Error != nil {
		fmt.Println("LinuxPingHosts", ping.Error)
		return
	}
	fmt.Println("-----------HOSTS------------")
	fmt.Println(ping.Response)
}

//TestLinuxPingPort can ping a udp
func TestLinuxPingPort(*testing.T) {
	opts := LinuxPingOptions{
		Host:       "192.168.15.1",
		Port:       99,
		TimeoutSec: 2,
	}
	found, err := PingPort(opts)
	fmt.Println("-----------PING PORT------------")
	fmt.Println(found, err)
}
