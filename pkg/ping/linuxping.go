package ping

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/validation"
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

func LinuxPing(ip string, pingCount, intervalSec, timeoutSec int) (found bool, err error) {
	_, err = validation.IsIPAddr(ip)
	if err != nil {
		return false, err
	}
	count := fmt.Sprintf("-c %d", numBetween(pingCount))
	interval := fmt.Sprintf("-i %d", numBetween(intervalSec))
	timeout := fmt.Sprintf("-w %d", numBetween(timeoutSec))
	/*
		-c = ping count
		-i = interval between ping's
		-w = timeout of wait for ping result
	*/
	out := exec.Command("ping", ip, count, interval, timeout)
	//fmt.Println(out.String())
	output, err := out.Output()
	if strings.Contains(string(output), "Destination Host Unreachable") {
		return false, err
	} else {
		return true, err
	}

}
