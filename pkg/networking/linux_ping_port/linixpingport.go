package linux_ping_port

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/ip_helpers"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	"strings"
)

type LinuxPingOptions struct {
	Hosts []string `json:"hosts"`
	Host  string   `json:"host"`
	IsUDP bool     `json:"is_udp"`
	Port  int      `json:"port"`
}

type Response struct {
	Host  string `json:"host"`
	Found bool   `json:"found"`
}

type LinuxPingResponse struct {
	Response []Response
	Error    error
}

func PingPort(network, port string, timeout int, isUDP bool) (message string, err error, ok bool) {
	_timeout := fmt.Sprintf("-w %d", timeout)
	cmd := []string{
		"nc",
		"-zv",
		network,
		port,
		_timeout,
	}
	if isUDP {
		cmd = []string{
			"nc",
			"-vzu",
			network,
			port,
			_timeout,
		}
	}
	out, err := command.Run(cmd...)
	if strings.Contains(out, "succeeded!") {
		ok = true
	}
	return out, err, ok

}

func PingPorts() {
	ipsSequence := []string{"192.168.15.1-224"}
	ips := ip_helpers.GetIpList(ipsSequence)

	for _, ip := range ips {
		fmt.Println(ip.ToString())
		go func() {
			msg, err, ok := PingPort(ip.ToString(), "22", 1, false)
			if err != nil {
				fmt.Println(msg, err)
			}
			if ok {
				fmt.Println(msg, ok)
			}

		}()

	}

	//fmt.Println(ips)
}
