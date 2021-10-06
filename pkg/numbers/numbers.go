package numbers

import (
	"math"
	"math/rand"
	"time"
)

//THE FOLLOWING NIL CHECK FUNCTIONS COULD PROBABLY BE REDUCED TO 2 TEMPLATE FUNCTIONS

//GetFloat64ValueOrZero returns 0 if float64 input pointer is nil, otherwise it returns the value at the pointer.
func GetFloat64ValueOrZero(b *float64) float64 {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

//Float64PointerIsNil returns true if the float64 pointer is nil or false if the pointer is valid.
func Float64PointerIsNil(b *float64) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

//GetIntValueOrZero returns 0 if int input pointer is nil, otherwise it returns the value at the pointer.
func GetIntValueOrZero(b *int) int {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

//IntPointerIsNil returns true if the int pointer is nil or false if the pointer is valid.
func IntPointerIsNil(b *int) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

//GetFloat32ValueOrZero returns 0 if float32 input pointer is nil, otherwise it returns the value at the pointer.
func GetFloat32ValueOrZero(b *float32) float32 {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

//Float32PointerIsNil returns true if the float32 pointer is nil or false if the pointer is valid.
func Float32PointerIsNil(b *float32) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

//GetUint16ValueOrZero returns 0 if uint16 input pointer is nil, otherwise it returns the value at the pointer.
func GetUint16ValueOrZero(b *uint16) uint16 {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

//Uint16PointerIsNil returns true if the uint16 pointer is nil or false if the pointer is valid.
func Uint16PointerIsNil(b *uint16) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

//GetUint32ValueOrZero returns 0 if uint32 input pointer is nil, otherwise it returns the value at the pointer.
func GetUint32ValueOrZero(b *uint32) uint32 {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

//Uint32PointerIsNil returns true if the uint32 pointer is nil or false if the pointer is valid.
func Uint32PointerIsNil(b *uint32) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

//LimitToRange returns the input value clamped within the specified range
func LimitToRange(value float64, range1 float64, range2 float64) float64 {
	if range1 == range2 {
		return range1
	}
	var min, max float64
	if range1 > range2 {
		max = range1
		min = range2
	} else {
		max = range2
		min = range1
	}
	return math.Min(math.Max(value, min), max)
}

//RoundTo returns the input value rounded to the specified number of decimal places.
func RoundTo(value float64, decimals uint32) float64 {
	if decimals < 0 {
		return value
	}
	return math.Round(value*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

//RandInt returns a random int within the specified range.
func RandInt(range1, range2 int) int {
	if range1 == range2 {
		return range1
	}
	var min, max int
	if range1 > range2 {
		max = range1
		min = range2
	} else {
		max = range2
		min = range1
	}
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

//RandFloat returns a random float64 within the specified range.
func RandFloat(range1, range2 float64) float64 {
	if range1 == range2 {
		return range1
	}
	var min, max float64
	if range1 > range2 {
		max = range1
		min = range2
	} else {
		max = range2
		min = range1
	}
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

//Scale returns the (float64) input value (between inputMin and inputMax) scaled to a value between outputMin and outputMax
func Scale(value, inMin, inMax, outMin, outMax float64) float64 {
	scaled := ((value-inMin)/(inMax-inMin))*(outMax-outMin) + outMin
	if scaled > math.Max(outMin, outMax) {
		return math.Max(outMin, outMax)
	} else if scaled < math.Min(outMin, outMax) {
		return math.Min(outMin, outMax)
	} else {
		return scaled
	}
}
