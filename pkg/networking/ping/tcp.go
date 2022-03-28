package ping

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/command"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/validation"
	log "github.com/sirupsen/logrus"
	"strings"
)

//PingPort will ping an address and port, example 192.168.15.10:1515
func PingPort(opts LinuxPingOptions) (found bool, err error) {
	host := opts.Host
	_, err = validation.IsIPAddr(host)
	if err != nil {
		errMsg := fmt.Sprintf("nubeio.helpers.ping.PingPort() failed on validation of ip on host:%s", host)
		err = errors.New(errMsg)
		log.Errorln(errMsg)
		return
	}
	port := fmt.Sprintf("%d", opts.Port)
	isUDP := opts.IsUDP
	timeout := fmt.Sprintf("-w %d", opts.TimeoutSec)

	cmd := []string{
		"nc",
		"-zv",
		host,
		port,
		timeout,
	}
	if isUDP {
		cmd = []string{
			"nc",
			"-vzu",
			host,
			port,
			timeout,
		}
	}
	out, err := command.Run(cmd...)
	if err != nil {
		errMsg := fmt.Sprintf("nubeio.helpers.ping.PingPort() failed to run ping on host::%s", host)
		err = errors.New(errMsg)
		log.Errorln(errMsg)
	} else {
		if strings.Contains(out, "succeeded!") {
			found = true
		}
	}
	return
}
