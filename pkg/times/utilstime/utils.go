package utilstime

import (
	"strings"
	"time"
)

func clean(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		i := strings.Trim(s[idx:], ":")
		i = strings.Join(strings.Fields(strings.TrimSpace(i)), " ")
		return i
	}
	return s
}

func TimeStamp() (hostTime string) {
	hostTime = time.Now().Format(time.RFC850)
	return
}
