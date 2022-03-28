package journalctl

import (
	"fmt"
	"testing"
)

func TestJournalCTL(*testing.T) {

	getLogs, err := NewJournalCTL().EntriesAfter("mosquitto.service", "", "3")

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, entry := range getLogs {
		fmt.Println(entry.Message)
		fmt.Println(i)

	}

}
