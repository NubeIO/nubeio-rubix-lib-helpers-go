package str

import "strings"

func RemoveNewLine(in string) string {
	return strings.TrimRight(in, "\n")
}
