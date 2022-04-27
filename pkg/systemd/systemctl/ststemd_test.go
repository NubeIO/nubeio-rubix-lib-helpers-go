package systemctl

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestJournalCTL(*testing.T) {

	unit := "mosquitto"
	stats, err := IsInstalled(unit, Options{})
	if err != nil {
		log.Fatalf("unable to enable unit %s: %v", "nginx", err)
	}

	fmt.Println(stats)

}
