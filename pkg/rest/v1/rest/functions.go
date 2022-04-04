package rest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var defaultTimeout = 5 * time.Second

var defaultMaxRetries = 100

const (
	LocalHost = "0.0.0.0"
)

func isJSON(str string) bool {
	return json.Unmarshal([]byte(str), &json.RawMessage{}) == nil
}

func getJSONLen(str interface{}) (length int) {
	length = reflect.ValueOf(str).Len()
	return
}

// StatusCode2xx method returns true if HTTP status `code >= 200 and <= 299` otherwise false.
func StatusCode2xx(statusCode int) bool {
	return statusCode > 199 && statusCode < 300
}

// StatusCode3xx method returns true if HTTP status `code >= 299 and <= 399` otherwise false.
func StatusCode3xx(statusCode int) bool {
	return statusCode > 299 && statusCode < 399
}

// StatusCode4xx method returns true if HTTP status `code >= 399 and <= 499` otherwise false.
func StatusCode4xx(statusCode int) bool {
	return statusCode > 399 && statusCode < 499
}

// StatusCode5xx method returns true if HTTP status `code >= 499 and <= 599` otherwise false.
func StatusCode5xx(statusCode int) bool {
	return statusCode > 499 && statusCode < 599
}

// StatusCodesAllBad any status for 3xx, 4xx and 5xx
func StatusCodesAllBad(statusCode int) (ok bool) {
	if StatusCode3xx(statusCode) {
		ok = true
	}
	if StatusCode4xx(statusCode) {
		ok = true
	}
	if StatusCode5xx(statusCode) {
		ok = true
	}
	return
}

func errorMsg(appName string, msg interface{}, e error) (err interface{}) {
	if e != nil && msg != "" {
		err = fmt.Errorf("%s: error:%w  msg:%s", appName, e, msg)
		return
	}
	if e != nil {
		err = fmt.Errorf("%s: error:%w", appName, e)
		return
	}
	if msg != "" {
		err = fmt.Errorf("%s: msg:%s", appName, msg)
		return
	}
	return
}

func (Options) ParseData(d map[string]interface{}) map[string]string {
	dLen := len(d)
	if dLen == 0 {
		return nil
	}
	data := make(map[string]string, dLen)
	for k, v := range d {
		if val, ok := v.(string); ok {
			data[k] = val
		} else {
			data[k] = fmt.Sprintf("%v", v)
		}
	}
	return data
}

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
