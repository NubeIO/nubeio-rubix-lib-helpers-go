package pprint

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

func Print(i interface{}) (out string) {
	out = spew.Sdump(i)
	log.Println(spew.Sdump(i))
	return
}
