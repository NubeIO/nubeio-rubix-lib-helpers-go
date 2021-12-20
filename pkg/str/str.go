package str

import "strings"

//RemoveNewLine will remove /n from the end of a string
func RemoveNewLine(in string) string {
	return strings.TrimRight(in, "\n")
}
