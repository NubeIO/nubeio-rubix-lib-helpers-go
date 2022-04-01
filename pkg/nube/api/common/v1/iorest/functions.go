package iorest

import (
	"reflect"
	"runtime"
	"strings"
	"time"
)

func tokenTimeDiffMin(t time.Time, timeDiff float64) (out bool) {
	t1 := time.Now()
	if t1.Sub(t).Minutes() > timeDiff {
		out = true
	}
	return
}

func GetFunctionName(temp interface{}) string {
	s := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
	return s[len(s)-1]
}
