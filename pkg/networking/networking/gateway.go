package networking

import (
	"net"
	"os/exec"
	"strings"
)

//GetGatewayIP Get gateway IP address
func (nets *nets) GetGatewayIP(iFaceName string) (ip string, err error) {
	cmd := exec.Command("ip", "route", "show", "dev", iFaceName)
	d, err := cmd.Output()
	if err != nil || cmd.ProcessState.ExitCode() != 0 {
		return
	}
	fields := strings.Fields(string(d))
	if len(fields) < 3 || fields[0] != "default" {
		return
	}
	getIP := net.ParseIP(fields[2])
	if getIP == nil {
		return
	}

	return getIP.String(), nil
}
