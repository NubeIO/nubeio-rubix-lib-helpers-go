package ping

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/validation"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func numBetween(num int) int {
	if num <= 0 {
		return 1
	}
	if num >= 10 {
		return 10
	}
	return 1
}

type LinuxPingOptions struct {
	Hosts       []string `json:"hosts"`
	Host        string   `json:"host"`
	PingCount   int      `json:"ping_count"`
	IntervalSec int      `json:"interval_sec"`
	TimeoutSec  int      `json:"timeout_sec"`
	IsUDP       bool     `json:"is_udp"`
	Port        int      `json:"port"`
}

type Response struct {
	Host  string `json:"host"`
	Found bool   `json:"found"`
}

type LinuxPingResponse struct {
	Response []Response
	Error    error
}

func LinuxPingHost(opts LinuxPingOptions) (res Response, err error) {
	host := opts.Host
	_, err = validation.IsIPAddr(host)
	if err != nil {
		errMsg := fmt.Sprintf("nubeio.helpers.ping.LinuxPingHosts() failed on validation of ip on host:%s", host)
		log.Errorln(errMsg)
		return res, errors.New(errMsg)
	}
	/*
		-c = ping count
		-i = interval between ping's
		-w = timeout of wait for ping result
	*/
	count := fmt.Sprintf("%d", numBetween(opts.PingCount))
	interval := fmt.Sprintf("%d", numBetween(opts.IntervalSec))
	timeout := fmt.Sprintf("%d", numBetween(opts.TimeoutSec))
	cmd := fmt.Sprintf("ping -c %s -i %s -w %s  %s > /dev/null && echo true || echo false", count, interval, timeout, host)
	out := exec.Command("/bin/sh", "-c", cmd)
	output, err := out.Output()
	res.Host = host
	if strings.Contains(string(output), "true") {
		res.Found = true
		return res, err
	} else {
		return res, err
	}

}

func LinuxPingHosts(opts LinuxPingOptions) (res LinuxPingResponse) {
	/*
		-c = ping count, as in ping 3 times
		-i = interval between ping's
		-w = timeout of wait for ping result
	*/
	//count := fmt.Sprintf("-c %d", numBetween(opts.PingCount))
	//interval := fmt.Sprintf("-i %d", numBetween(opts.IntervalSec))
	//timeout := fmt.Sprintf("-w %d", numBetween(opts.TimeoutSec))
	//res.Response
	for _, host := range opts.Hosts {
		opts.Host = host
		h, err := LinuxPingHost(opts)
		if err != nil {
			res.Response = append(res.Response, Response{Host: h.Host, Found: h.Found})
		} else {
			res.Response = append(res.Response, Response{Host: h.Host, Found: h.Found})
		}
	}

	return res

}
