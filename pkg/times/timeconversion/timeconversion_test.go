package timeconversion

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestAdjustTime(*testing.T) {
	now := time.Now()
	timeFormat := "3:04 PM on January 2, 2006"

	fmt.Printf("Now is %s\n", now.Format(timeFormat))

	adjustedTime, err := AdjustTime(now, "2s")
	if err != nil {
		fmt.Printf("error adjusting time: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Adjusted is %s\n", adjustedTime.Format(timeFormat))
}
