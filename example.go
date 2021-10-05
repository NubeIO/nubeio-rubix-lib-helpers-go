package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/strings"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/thermistor"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
)

func main() {

	str := strings.New("what$ up !n the hood ")
	fmt.Println(str.RemoveSpecialCharacter())

	b, err := bools.Boolean("on on")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	bb, _ := bools.Boolean("0")
	fmt.Println(bb)
	bbb, _ := bools.Boolean("True")
	fmt.Println(bbb)

	u, _ := uuid.MakeUUID()
	fmt.Println(u)

	fmt.Println("Testing Temperature Lookup Tables")
	result, err := thermistor.ResistanceToTemperature(1000, thermistor.T2_10K_TempTable)
	fmt.Println("1000 Ohm from T2_10K = ", result)
	result, err = thermistor.ResistanceToTemperature(1000, thermistor.T3_10K_TempTable)
	fmt.Println("1000 Ohm from T3_10K = ", result)
	result, err = thermistor.ResistanceToTemperature(87, thermistor.D_PT100_TempTable)
	fmt.Println("87 Ohm from D_PT100 = ", result)

}
