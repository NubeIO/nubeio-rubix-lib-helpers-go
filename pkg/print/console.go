package pprint

import (
	"github.com/davecgh/go-spew/spew"
)

func Print(i interface{}) string {
	return spew.Sdump(i)
}
