package logs

import (
	"io/ioutil"
	"log"
)

func DisableLogging(enable bool) {
	if enable {
		log.Print("INIT APP: LOGGING IS DISABLED")
		log.SetOutput(ioutil.Discard)
	}

}


