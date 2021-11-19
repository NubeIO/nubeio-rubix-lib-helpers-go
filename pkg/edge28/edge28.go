package edge28

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/numbers"
	"math"
	"reflect"
	"strconv"
)

//UIGPIOCORRECTIONFACTOR This correction factor is due to the Edge28 Hardware that scales the UI Voltage/Resistance (0-10v) readings to the onboard ADC.
//The voltage on the ADC, when the UI pin voltage is 10vdc, is 1.774v; but the ACD GPIO Output (0-1) is over the range of (0v-1.8v) therefore the ADC (GPIO) value
// has a maximum reading of 0.9544 (on the Edge28 hardware) when there is 10vdc on the UI pin.  This value has been modified slightly based on hardware testing.
var UIGPIOCORRECTIONFACTOR float64 = 0.9544

//CorrectGPIOValueForUIs scales the actual UI GPIO values to account for the correction factor described above in UIGPIOCORRECTIONFACTOR
func CorrectGPIOValueForUIs(value float64) float64 {
	return numbers.Scale(value, 0, UIGPIOCORRECTIONFACTOR, 0, 1)
}

//PercentToGPIOValue scales 0-100% input to BBB GPIO 100-16.666 (16.666-0 is for 10v-12v) .  Note that the input 0-100% input has a 0.9839 scaling factor (this is a software cal for the UOs)
func PercentToGPIOValue(value float64) float64 {
	if value <= 0 {
		return 100
	} else if value >= 100 {
		return 16.666666666666668
	} else {
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
		return numbers.Scale(value, 10, 0, 16.666666666666668, 100)
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
		return 0, nil // 0 is the 12vdc/ON GPIO value
	} else {
		return 100, nil // 100 is the 0vdc/OFF GPIO value
	}
}

//GPIOValueToPercent scales BBB GPIO Value (0-1) to 0-100%
func GPIOValueToPercent(value float64) float64 {
	value = CorrectGPIOValueForUIs(value)
	if value <= 0 {
		return 0
	} else if value >= 1 {
		return 100
	} else {
		result := numbers.Scale(value, 0, 1, 0, 100)
		result = math.Round(result*10) / 10
		return result
	}
}

//GPIOValueToVoltage scales BBB GPIO Value (0-1) to 0-10vdc
func GPIOValueToVoltage(value float64) float64 {
	value = CorrectGPIOValueForUIs(value)
	if value <= 0 {
		return 0
	} else if value >= 1 {
		return 10
	} else {
		result := numbers.Scale(value, 0, 1, 0, 10)
		result = math.Round(result*100) / 100
		return result
	}
}

//ScaleGPIOValueToRange scales a BBB GPIO Value (0-1) input to the specified output range.
func ScaleGPIOValueToRange(value, outputMin, outputMax float64) float64 {
	value = CorrectGPIOValueForUIs(value)
	if value <= 0 {
		return outputMin
	} else if value >= 1 {
		return outputMax
	} else {
		result := numbers.Scale(value, 0, 1, outputMin, outputMax)
		return result
	}
}

//GPIOValueToDigital converts BBB GPIO Value (0-1) to 0 (OFF/Open Circuit) or 1 (ON/Closed Circuit)
func GPIOValueToDigital(value float64) float64 {
	value = CorrectGPIOValueForUIs(value)
	if value < 0.2 {
		return 1 //ON / Closed Circuit
	} else { //previous functions used > 0.6 as an OFF threshold.
		return 0 //OFF / Open Circuit
	}
}

//ScaleGPIOValueTo420ma scales a BBB GPIO Value (0.2-1) input to 4-20mA.
func ScaleGPIOValueTo420ma(value float64) float64 {
	value = CorrectGPIOValueForUIs(value)
	if value <= 0.2 {
		return 4
	} else if value >= 1 {
		return 20
	} else {
		result := numbers.Scale(value, 0.2, 1, 4, 20)
		result = math.Round(result*100) / 100
		return result
	}
}

//ScaleGPIOValueTo420maOrError scales a BBB GPIO Value (0.2-1) input to 4-20mA, or it produces an error if the GPIO value is < 0.2 (less than 4mA).
func ScaleGPIOValueTo420maOrError(value float64) (float64, error) {
	value = CorrectGPIOValueForUIs(value)
	if value <= 0.195 {
		return 0, errors.New("input is below 4mA")
	} else if value >= 1 {
		return 20, nil
	} else {
		result := numbers.Scale(value, 0, 1, 4, 20)
		result = math.Round(result*100) / 100
		return result, nil
	}
}

//ScaleGPIOValueToResistance scales a BBB GPIO Value (0-1) input to Resistance.
func ScaleGPIOValueToResistance(value float64) float64 {
	// CorrectGPIOValueForUIs() is not required here because the equation below takes care of the correction.
	if value <= 0 {
		return 0
	} else if value >= 0.925 { //Upper limit of RAW -> Resistance Equation provided by Craig Burrows (25/10/2021).
		return 1361693
	} else {
		result := (8.65943 * ((value / (0.5555 * 0.9544)) / 0.1774)) / (9.89649 - ((value / (0.5555 * 0.9544)) / 0.1774)) * 1000 //RAW -> Resistance Equation provided by Craig Burrows (25/10/2021).
		result = math.Round(result)
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
		result := numbers.Scale(value, 4, 20, outputMin, outputMax)
		result = math.Round(result*100) / 100
		return result
	}
}
