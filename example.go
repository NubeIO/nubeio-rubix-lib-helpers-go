package main

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest"
	"net/http"
	"time"

	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/strings"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/thermistor"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
)

type T struct {
	Health   string `json:"health"`
	Database string `json:"database"`
}

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
	result, err := thermistor.ResistanceToTemperature(1000, thermistor.T210K)
	fmt.Println("1000 Ohm from T2_10K Thermistor = ", result)
	result, err = thermistor.ResistanceToTemperature(1000, thermistor.T310K)
	fmt.Println("1000 Ohm from T3_10K Thermistor = ", result)
	result, err = thermistor.ResistanceToTemperature(87, thermistor.PT100)
	fmt.Println("87 Ohm from PT100 Thermistor = ", result)

	headers := make(http.Header)
	headers.Add("Authorization", "")

	var rb = rest.RequestBuilder{
		Headers:        headers,
		Timeout:        5000 * time.Millisecond,
		BaseURL:        "http://0.0.0.0:1660",
		ContentType:    rest.JSON,
		DisableCache:   false,
		DisableTimeout: false,
	}
	resp := rb.Get("/api/system/ping")
	m := new(T)
	fmt.Println(resp.String())
	json.Unmarshal(resp.Bytes(), &m)
	fmt.Println(m.Health)

}
