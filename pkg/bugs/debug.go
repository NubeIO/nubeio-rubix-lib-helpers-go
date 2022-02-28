package bugs

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func DebugPrint(use, fun interface{}, err error, args ...interface{}) (out string) {
	funcName, _ := GetFuncName(fun)
	if err == nil {
		out = fmt.Sprintf("%s: funcName:%s  msg:%s", use, funcName, strings.Trim(fmt.Sprint(args), "[]"))
	} else {
		out = fmt.Sprintf("%s: funcName:%s  msg:%s  error:%s", use, funcName, err.Error())
	}
	return
}

func GetFuncName(temp interface{}) (funcName, fileName string) {
	s := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
	if len(s) >= 1 {
		return s[len(s)-1], s[0]
	} else {
		return "", ""
	}
}
