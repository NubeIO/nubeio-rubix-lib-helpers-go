package thermistor

import (
	"errors"
	"math"
)

/*
type ThermistorTableType ThermistorTable
const (
	T2_10K = "T2_10K_TempTable"
	T3_10K = "T3_10K_TempTable"
	D_PT100 = "D_PT100_TempTable"
	E_PT100 = "E_PT100_TempTable"
	T1_20K = "T1_20K_TempTable"
)
*/

type resistanceAndTemperaturePair struct {
	resistance  float64
	temperature float64
}

type ThermistorTable []resistanceAndTemperaturePair

func ResistanceToTemperature(resistance float64, tempTable ThermistorTable) (float64, error) {
	if resistance > tempTable[0].resistance || resistance < tempTable[len(tempTable)-1].resistance {
		return -1, errors.New("resistance value is beyond sensor table limits")
	}
	var lowerRes, higherRes, lowerTemp, higherTemp, outValue float64
	for _, pair := range tempTable {
		//fmt.Println("RESISTANCE: ", pair.resistance)
		if resistance > pair.resistance {
			lowerRes = pair.resistance
			lowerTemp = pair.temperature
			break
		}
		higherRes = pair.resistance
		higherTemp = pair.temperature
	}
	//fmt.Println("resistance: ", resistance, "lowerRes: ", lowerRes, "higherRes: ", higherRes, "lowerTemp: ", lowerTemp, "higherTemp: ", higherTemp )
	outValue = Scale(resistance, lowerRes, higherRes, lowerTemp, higherTemp)
	//fmt.Println("RESULT: ", outValue)
	return outValue, nil
}

//TODO: Should use Scale function from Utils once added to flow-framework
//Scale returns the (float64) input value (between inputMin and inputMax) scaled to a value between outputMin and outputMax
func Scale(value float64, inMin float64, inMax float64, outMin float64, outMax float64) float64 {
	scaled := ((value-inMin)/(inMax-inMin))*(outMax-outMin) + outMin
	if scaled > math.Max(outMin, outMax) {
		return math.Max(outMin, outMax)
	} else if scaled < math.Min(outMin, outMax) {
		return math.Min(outMin, outMax)
	} else {
		return scaled
	}
}

//T2_10K_TempTable :Type 2 10K thermistor
var T2_10K_TempTable = ThermistorTable{
	{963849, -55},
	{670166, -50},
	{471985, -45},
	{336479, -40},
	{242681, -35},
	{176974, -30},
	{130421, -25},
	{97081, -20},
	{72957, -15},
	{55329, -10},
	{42327, -5},
	{32650, 0},
	{25392, 5},
	{19901, 10},
	{15712, 15},
	{12493, 20},
	{10000, 25},
	{8057, 30},
	{6531, 35},
	{5326, 40},
	{4368, 45},
	{3602, 50},
	{2986, 55},
	{2488, 60},
	{2083, 65},
	{1752, 70},
	{1480, 75},
	{1255, 80},
	{1070, 85},
	{915.5, 90},
	{786.6, 95},
	{678.6, 100},
	{587.6, 105},
	{510.6, 110},
	{445.3, 115},
	{389.6, 120},
	{341.9, 125},
	{301, 130},
	{265.8, 135},
	{235.3, 140},
	{208.9, 145},
	{186.1, 150},
}

//T3_10K_TempTable :Type 3 10K thermistor
var T3_10K_TempTable = ThermistorTable{
	{607800, -55},
	{441200, -50},
	{323600, -45},
	{239700, -40},
	{179200, -35},
	{135200, -30},
	{102900, -25},
	{78910, -20},
	{61020, -15},
	{47540, -10},
	{37310, -5},
	{29490, 0},
	{23460, 5},
	{18780, 10},
	{15130, 15},
	{12260, 20},
	{10000, 25},
	{8194, 30},
	{6752, 35},
	{5592, 40},
	{4655, 45},
	{3893, 50},
	{3271, 55},
	{2760, 60},
	{2339, 65},
	{1990, 70},
	{1700, 75},
	{1458, 80},
	{1255, 85},
	{1084, 90},
	{939.3, 95},
	{816.8, 100},
	{712.6, 105},
	{623.6, 110},
	{547.3, 115},
	{481.8, 120},
	{425.3, 125},
	{376.4, 130},
	{334, 135},
	{297.2, 140},
	{265.1, 145},
	{237, 150},
}

//D_PT100_TempTable :Type D-PT100 thermistor
var D_PT100_TempTable = ThermistorTable{
	{157.33, 150},
	{155.46, 145},
	{153.58, 140},
	{151.71, 135},
	{149.83, 130},
	{147.95, 125},
	{146.07, 120},
	{144.18, 115},
	{142.29, 110},
	{140.4, 105},
	{138.51, 100},
	{136.61, 95},
	{134.71, 90},
	{132.8, 85},
	{130.9, 80},
	{128.99, 75},
	{127.08, 70},
	{125.16, 65},
	{123.24, 60},
	{121.32, 55},
	{119.4, 50},
	{117.47, 45},
	{117.47, 45},
	{115.54, 40},
	{113.61, 35},
	{111.67, 30},
	{109.74, 25},
	{107.79, 20},
	{105.85, 15},
	{103.9, 10},
	{101.95, 5},
	{100, 0},
	{98.04, -5},
	{96.09, -10},
	{94.12, -15},
	{92.16, -20},
	{90.19, -25},
	{88.22, -30},
	{86.25, -35},
	{84.27, -40},
	{82.29, -45},
	{80.31, -50},
	{78.32, -55},
}

//E_PT100_TempTable :Type E-PT100 thermistor
var E_PT100_TempTable = ThermistorTable{
	{1573.3, 150},
	{1554.6, 145},
	{1535.8, 140},
	{1517.1, 135},
	{1498.3, 130},
	{1479.5, 125},
	{1460.7, 120},
	{1441.8, 115},
	{1422.9, 110},
	{1404, 105},
	{1385.1, 100},
	{1366.1, 95},
	{1347.1, 90},
	{1328, 85},
	{1309, 80},
	{1289.9, 75},
	{1270.8, 70},
	{1251.6, 65},
	{1232.4, 60},
	{1213.2, 55},
	{1194, 50},
	{1174.7, 45},
	{1155.4, 40},
	{1136.1, 35},
	{1116.7, 30},
	{1097.4, 25},
	{1077.9, 20},
	{1058.5, 15},
	{1039, 10},
	{1019.5, 5},
	{1000, 0},
	{980.4, -5},
	{960.9, -10},
	{941.2, -15},
	{921.6, -20},
	{901.9, -25},
	{882.2, -30},
	{862.5, -35},
	{842.7, -40},
	{822.9, -45},
	{803.1, -50},
	{783.2, -55},
}

//T1_20K_TempTable :Type 1 20K thermistor
var T1_20K_TempTable = ThermistorTable{
	{2394000, -55},
	{1646200, -50},
	{1145800, -45},
	{806800, -40},
	{574400, -35},
	{413400, -30},
	{300400, -25},
	{220600, -20},
	{163500, -15},
	{122280, -10},
	{92240, -5},
	{70160, 0},
	{53780, 5},
	{41560, 10},
	{32340, 15},
	{25360, 20},
	{20000, 25},
	{15892, 30},
	{12704, 35},
	{10216, 40},
	{8264, 45},
	{6722, 50},
	{5498, 55},
	{4520, 60},
	{3734, 65},
	{3100, 70},
	{2586, 75},
	{2166, 80},
	{1822.6, 85},
	{1540, 90},
	{1306.4, 95},
	{1112.6, 100},
	{951, 105},
	{815.8, 110},
	{702.2, 115},
	{606.4, 120},
	{525.6, 125},
	{0, 130},
}
