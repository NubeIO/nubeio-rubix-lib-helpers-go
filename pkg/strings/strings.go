package strings

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)


const (
	First                = "first"
	Last                 = "last"
	OddError             = "odd number rule provided please provide in even count"
	SelectCapital        = "([a-z])([A-Z])"
	ReplaceCapital       = "$1 $2"
	LengthError          = "passed length cannot be greater than input length"
	InvalidLogicalString = "invalid string value to test boolean value"
)


// input is struct that holds input form user and result
type input struct {
	Input  string
	Result string
}


// StringManipulation is an interface that holds all abstract methods to manipulate strings
type StringManipulation interface {
	Get() string
	Between(start, end string) StringManipulation
	CamelCase(rule ...string) string
	RemoveSpecialCharacter() string
	Reverse() string
	ToLower() string
	ToUpper() string
	ReplaceFirst(search, replace string) string
	ReplaceLast(search, replace string) string
}

// New func returns pointer to input struct
func New(val string) StringManipulation {
	return &input{Input: val}
}

func (i *input) Get() string {
	return getInput(*i)
}

// ReplaceFirst takes two param search and replace. It returns string by searching search
// sub string and replacing it with replace substring on first occurrence it can be chained
// on function which return StringManipulation interface.
func (i *input) ReplaceFirst(search, replace string) string {
	input := getInput(*i)
	return replaceStr(input, search, replace, First)
}

// ReplaceLast takes two param search and replace
// it return string by searching search sub string and replacing it
// with replace substring on last occurrence
// it can be chained on function which return StringManipulation interface
func (i *input) ReplaceLast(search, replace string) string {
	input := getInput(*i)
	return replaceStr(input, search, replace, Last)
}

// CamelCase is variadic function which takes one Param rule i.e slice of strings and it returns
// input type string in camel case form and rule helps to omit character you want to omit from string.
// By default special characters like "_", "-","."," " are l\treated like word separator and treated
// accordingly by default and you dont have to worry about it
// Example input: hello user
// Result : HelloUser
func (i *input) CamelCase(rule ...string) string {
	input := getInput(*i)
	// removing excess space
	wordArray := caseHelper(input, true, rule...)
	for i, word := range wordArray {
		wordArray[i] = ucfirst(word)
	}
	return strings.Join(wordArray, "")
}


// Between takes two string params start and end which and returns
// value which is in middle of start and end part of input. You can
// chain to upper which with make result all upercase or ToLower which
// will make result all lower case or Get which will return result as it is
func (i *input) Between(start, end string) StringManipulation {
	if (start == "" && end == "") || i.Input == "" {
		return i
	}
	input := strings.ToLower(i.Input)
	lcStart := strings.ToLower(start)
	lcEnd := strings.ToLower(end)
	var startIndex, endIndex int

	if len(start) > 0 && strings.Contains(input, lcStart) {
		startIndex = len(start)
	}
	if len(end) > 0 && strings.Contains(input, lcEnd) {
		endIndex = strings.Index(input, lcEnd)
	} else if len(input) > 0 {
		endIndex = len(input)
	}
	i.Result = strings.TrimSpace(i.Input[startIndex:endIndex])
	return i
}

func getInput(i input) (input string) {
	if i.Result != "" {
		input = i.Result
	} else {
		input = i.Input
	}
	return
}


// RemoveSpecialCharacter removes all special characters and returns the string
// it can be chained on function which return StringManipulation interface
func (i *input) RemoveSpecialCharacter() string {
	input := getInput(*i)
	var result strings.Builder
	for i := 0; i < len(input); i++ {
		b := input[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func (i *input) Reverse() string {
	input := getInput(*i)
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}


// ToLower makes all string of user input to lowercase
// it can be chained on function which return StringManipulation interface
func (i *input) ToLower() (result string) {
	input := getInput(*i)
	return strings.ToLower(input)
}

// ToUpper makes all string of user input to uppercase
// it can be chained on function which return StringManipulation interface
func (i *input) ToUpper() string {
	input := getInput(*i)
	return strings.ToUpper(input)
}


// UcFirst makes first word of user input to uppercase
// it can be chained on function which return StringManipulation interface
func (i *input) UcFirst() string {
	input := getInput(*i)
	return ucfirst(input)
}





func caseHelper(input string, isCamel bool, rule ...string) []string {
	if !isCamel {
		re := regexp.MustCompile(SelectCapital)
		input = re.ReplaceAllString(input, ReplaceCapital)
	}
	input = strings.Join(strings.Fields(strings.TrimSpace(input)), " ")
	if len(rule) > 0 && len(rule)%2 != 0 {
		panic(errors.New(OddError))
	}
	rule = append(rule, ".", " ", "_", " ", "-", " ")

	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	return words
}


func replaceStr(input, search, replace, types string) string {
	lcInput := strings.ToLower(input)
	lcSearch := strings.ToLower(search)
	if input == "" || !strings.Contains(lcInput, lcSearch) {
		return input
	}
	var start int
	if types == "last" {
		start = strings.LastIndex(lcInput, lcSearch)
	} else {
		start = strings.Index(lcInput, lcSearch)
	}
	end := start + len(search)
	return input[:start] + replace + input[end:]
}

func ucfirst(val string) string {
	for i, v := range val {
		return string(unicode.ToUpper(v)) + val[i+1:]
	}
	return ""
}


