package ping

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	"strings"
)

//PingPort will ping an address and port, example 192.168.15.10:1515
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
