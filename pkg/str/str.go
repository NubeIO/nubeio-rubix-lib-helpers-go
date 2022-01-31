package str

import (
	"strings"
	"unicode"
)

//RemoveNewLine will remove /n from the end of a string
func RemoveNewLine(in string) string {
	return strings.TrimRight(in, "\n")
}

func UcFirst(val string) string {
	for i, v := range val {
		return string(unicode.ToUpper(v)) + val[i+1:]
	}
	return ""
}

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

//StandardizeSpaces will remove all extra white spaces in text but will leave one white space between a word or letter
func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
