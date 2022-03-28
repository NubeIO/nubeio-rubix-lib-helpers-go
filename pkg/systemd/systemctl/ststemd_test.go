package systemctl

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestJournalCTL(*testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Equivalent to `systemctl enable nginx` with a 10 second timeout
	opts := Options{UserMode: false}
	unit := "mosquitto"
	stats, err := Stats(ctx, unit, opts)
	if err != nil {
		log.Fatalf("unable to enable unit %s: %v", "nginx", err)
	}

	fmt.Println(stats.State)
	fmt.Println(stats.SubState)
	fmt.Println(stats.ActiveState)

}
