package linixpingport

import (
	"fmt"
	"testing"
)

func TestLinuxPingPort(*testing.T) {

	//port, err, foundPort := PingPort("0.0.0.0", "1885", 1, false)
	//if err != nil {
	//	log.Errorln("foundPort", foundPort)
	//	return
	//}
	//fmt.Println(port, err, "foundPort", foundPort)
	//
	//found := false
	//for ok := true; ok; ok = !found {
	//
	//}
	tryUsePort := 1883
	count := 1
	for {
		port := fmt.Sprintf("%d", tryUsePort)
		_, _, foundPort := PingPort("0.0.0.0", port, 1, false)
		fmt.Println("PORT IN USE", foundPort)
		if !foundPort {
			fmt.Println("YAY USE PORT:", port)
			break
		}
		count += 1
		tryUsePort = tryUsePort + count
		fmt.Println("new port", tryUsePort)
	}

}
