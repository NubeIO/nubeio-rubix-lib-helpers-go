package freeport

import (
	"fmt"
	linixpingport "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/linuxpingport"
	log "github.com/sirupsen/logrus"
)

func FindFreePort(port int) (int, error) {
	tryUsePort := port
	count := 1
	outPort := port
	for {
		p := fmt.Sprintf("%d", tryUsePort)
		_, err, foundPort := linixpingport.PingPort("0.0.0.0", p, 1, false)
		if err != nil {
			return port, err
		}
		if !foundPort {
			log.Infoln("nubeio.helpers-FindFreePort() PORT TO USE", p)
			outPort = tryUsePort
			break
		}
		count += 1
		tryUsePort = tryUsePort + count
		log.Infoln("nubeio.helpers-FindFreePort() PORT IN USE SO TRY", tryUsePort)
	}

	return outPort, nil

}
