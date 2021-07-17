package bool


import "strconv"

// S2b converts a string true/T/"1" to a true/false
func S2b(value string)  bool {
	r, _ := strconv.ParseBool(value)
	return r
}

func I2b(b int) bool {
	if b == 1 {
		return true
	}
	return false
}

func B2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}



