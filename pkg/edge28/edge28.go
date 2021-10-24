package edge28

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/numbers"
	"reflect"
	"strconv"
)

// THESE FUNCTIONS ALL NEED TO BE VERIFIED AGAINST THE EDGE28 API RESPONSE PAYLOAD

//PercentToGPIOValue scales 0-100% input to BBB GPIO 100-16.666 (16.666-0 is for 10v-12v) .  Note that the input 0-100% input has a 0.9839 scaling factor (this is a software cal for the UOs)
func PercentToGPIOValue(value float64) float64 {
	if value <= 0 {
		return 100
	} else if value >= 100 {
		return 16.666666666666668
	} else {
		//value = value * 0.9839 //TODO: IS THIS REQUIRED/CORRECT??
		return numbers.Scale(value, 100, 0, 16.666666666666668, 100)
	}
}

//VoltageToGPIOValue scales 0-10vdc input to BBB GPIO 100-16.666 (16.666-0 is for 10v-12v) .  Note that the input 0-100% input has a 0.9839 scaling factor (this is a software cal for the UOs)
func VoltageToGPIOValue(value float64) float64 {
	if value <= 0 {
		return 100
	} else if value >= 10 {
		return 16.666666666666668
	} else {
		//value = value * 0.9839 //TODO: IS THIS REQUIRED/CORRECT??
		return numbers.Scale(value, 10, 0, 16.666666666666668, 10)
	}
}

//DigitalToGPIOValue converts true/false values (all basic types allowed) to BBB GPIO 0/1 ON/OFF.  Note that the GPIO value for digital points is inverted.
func DigitalToGPIOValue(input interface{}) (float64, error) {
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
		return 1, nil // 1 is the 12vdc/ON GPIO value
	} else {
		return 0, nil // 0 is the 0vdc/OFF GPIO value
	}
}

//GPIOValueToPercent scales BBB GPIO Value (0-1) to 0-100%
func GPIOValueToPercent(value float64) float64 {
	if value <= 0 {
		return 0
	} else if value >= 1 {
		return 100
	} else {
		return numbers.Scale(value, 0, 1, 0, 100)
	}
}

//GPIOValueToVoltage scales BBB GPIO Value (0-1) to 0-10vdc
func GPIOValueToVoltage(value float64) float64 {
	if value <= 0 {
		return 0
	} else if value >= 1 {
		return 10
	} else {
		return numbers.Scale(value, 0, 1, 0, 10)
	}
}

//ScaleGPIOValueToRange scales a BBB GPIO Value (0-1) input to the specified output range.
func ScaleGPIOValueToRange(value, outputMin, outputMax float64) float64 {
	if value <= 0 {
		return outputMin
	} else if value >= 1 {
		return outputMax
	} else {
		return numbers.Scale(value, 0, 1, outputMin, outputMax)
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

//ScaleGPIOValueTo420ma scales a BBB GPIO Value (0-1) input to 4-20mA.
func ScaleGPIOValueTo420ma(value float64) float64 {
	if value <= 0 { //TODO: is this correct? should there be another value for 4mA? and 0 would be 0mA?
		return 4
	} else if value >= 1 {
		return 20
	} else {
		return numbers.Scale(value, 0, 1, 4, 20)
	}
}

//ScaleGPIOValueToResistance scales a BBB GPIO Value (0-1) input to Resistance.
func ScaleGPIOValueToResistance(value float64) float64 {
	if value <= 0 { //TODO: is this correct? should there be another value for 4mA? and 0 would be 0mA?
		return 0
	} else if value >= 0.96 { //Upper limit of RAW -> Resistance Equation provided by Craig Burrows
		return 544884.73
	} else {
		result := (8.65943 * ((value / 0.5555) / 0.1774)) / (9.89649 - (value / 0.5555 / 0.1774)) * 1000 //RAW -> Resistance Equation provided by Craig Burrows
		return result
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

//Scale420maToRange scales a 4-20mA input to the specified output range.
func Scale420maToRange(value, outputMin, outputMax float64) float64 {
	if value <= 4 {
		return outputMin
	} else if value >= 20 {
		return outputMax
	} else {
		return numbers.Scale(value, 4, 20, outputMin, outputMax)
	}
}
