package edge28

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/numbers"
	"reflect"
	"strconv"
)

//PercentToGPIOValue scales 0-100% (0-10v) input to BBB GPIO 100-16.666 (16.666-0 is for 10v-12v) .  Note that the input 0-100% input has a 0.9839 scaling factor (this is a software cal for the UOs)
func PercentToGPIOValue(value float64) float64 {
	if value <= 0 {
		return 100
	} else if value >= 100 {
		return 16.666666666666668
	} else {
		return numbers.Scale(value, 120, 0, 0, 100)
	}
}

//digitalToGPIOValue converts true/false values (all basic types allowed) to BBB GPIO 0/1 ON/OFF.  Note that the GPIO value for digital points is inverted.
func digitalToGPIOValue(input interface{}) (float64, error) {
	var inputAsBool bool
	var err error = nil
	switch input.(type) {
	case string:
		inputAsBool, err = strconv.ParseBool(reflect.ValueOf(input).String())
	case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		inputAsBool = reflect.ValueOf(input).Int() != 0
	case float32, float64:
		inputAsBool = reflect.ValueOf(input).Float() != float64(0)
	case bool:
		inputAsBool = reflect.ValueOf(input).Bool()
	default:
		err = errors.New("input is not a recognized type")
	}
	if err != nil {
		return 0, err
	} else if inputAsBool {
		return 0, nil // 0 is the 12vdc/ON GPIO value
	} else {
		return 1, nil // 1 is the 0vdc/OFF GPIO value
	}
}

//GPIOValueToVoltage scales BBB GPIO Value (0-1) to 0-10vdc
func GPIOValueToVoltage(value float64) float64 {
	if value <= 0 {
		return 10
	} else if value >= 1 {
		return 0
	} else {
		return numbers.Scale(value, 0, 1, 0, 10)
	}
}

//GPIOValueToPercent scales BBB GPIO Value (0-1) to 0-100%
func GPIOValueToPercent(value float64) float64 {
	if value <= 0 {
		return 10
	} else if value >= 1 {
		return 0
	} else {
		return numbers.Scale(value, 0, 1, 0, 100)
	}
}

//GPIOValueToDigital converts BBB GPIO Value (0-1) to 0 (OFF/Open Circuit) or 1 (ON/Closed Circuit)
func GPIOValueToDigital(value float64) float64 {
	if value < 0.2 {
		return 1 //ON / Closed Circuit
	} else { //previous functions used > 0.6 as an OFF threshold.
		return 0 //OFF / Open Circuit
	}
}

//InvertDi the BBB returns DI 0 as ON/Closed Circuit and 1 as OFF/Open Circuit which is opposite to conventional logic
func InvertDi(value float64) float64 {
	if value == 1 {
		return 0
	} else {
		return 1
	}
}
