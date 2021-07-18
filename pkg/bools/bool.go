package bools

import (
	"errors"
	"strings"
)


const (
	InvalidLogicalString = "invalid string value to test boolean value"
)

var _false = []string{"off", "no", "0", "false", "False"}
var _true = []string{"on", "yes", "1", "true", "True"}


// Boolean func returns boolean value of string value like on, off, 0, 1, yes, no
// returns boolean value of string input. You can chain this function on other function
func Boolean(input string) bool {
	inputLower := strings.ToLower(input)
	off := contains(_false, inputLower)
	if off {
		return false
	}
	on := contains(_true, inputLower)
	if on {
		return true
	}
	panic(errors.New(InvalidLogicalString))
}



func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}


