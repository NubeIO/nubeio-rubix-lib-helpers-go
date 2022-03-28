package pprint

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

func Print(i interface{}) (out string) {
	out = spew.Sdump(i)
	log.Println(spew.Sdump(i))
	return
}

func PrintStrut(i interface{}) {
	fmt.Printf("%+v\n", i)
	return
}
