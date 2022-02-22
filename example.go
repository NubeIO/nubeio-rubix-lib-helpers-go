package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/networking"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/thermistor"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"time"
)

type T struct {
	Health   string `json:"health"`
	Database string `json:"database"`
}

func httpReq(r *nrest.ReqType, opt *nrest.ReqOpt, body interface{}) *nrest.Reply {
	_ip := fmt.Sprintf("http://%s:%s", "0.0.0.0", "1660")
	s := &nrest.Service{
		BaseUri: _ip,
	}
	opt = &nrest.ReqOpt{
		Timeout:          500 * time.Second,
		RetryCount:       0,
		RetryWaitTime:    0 * time.Second,
		RetryMaxWaitTime: 0,
		Json:             body,
	}

	if r.Method == "" {
		r.Method = nrest.GET
	}
	return s.Do(r.Method, r.Path, opt)
}

func main() {

	b, err := bools.Boolean("on on")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	bb, _ := bools.Boolean("0")
	fmt.Println(bb)
	bbb, _ := bools.Boolean("True")
	fmt.Println(bbb)

	uid, _ := uuid.MakeUUID()
	fmt.Println(uid)

	fmt.Println("Testing Temperature Lookup Tables")
	result, err := thermistor.ResistanceToTemperature(1000, thermistor.T210K)
	fmt.Println("1000 Ohm from T2_10K Thermistor = ", result)
	result, err = thermistor.ResistanceToTemperature(1000, thermistor.T310K)
	fmt.Println("1000 Ohm from T3_10K Thermistor = ", result)
	result, err = thermistor.ResistanceToTemperature(87, thermistor.PT100)
	fmt.Println("87 Ohm from PT100 Thermistor = ", result)

	_net, err := networking.GetInterfaceByName("wlp3s0")
	if err != nil {
		//return
	}
	fmt.Println(_net.Interface, _net.Gateway, _net.MacAddress)

	//s := &nrest.Service{
	//	BaseUri: "http://0.0.0.0:1660",
	//}
	//opt := &nrest.ReqOpt{
	//	Timeout:          500 * time.Second,
	//	RetryCount:       0,
	//	RetryWaitTime:    0 * time.Second,
	//	RetryMaxWaitTime: 0,
	//	//Json:             body,
	//}
	//fmt.Println(s.Do("GET", "/api/points", opt).Status())
	//fmt.Println(s.Do("GET", "/api/points", opt).AsString())
	//
	//rt := &nrest.ReqType{
	//	Path: "/api/points",
	//}
	//fmt.Println(httpReq(rt, opt, nil).StatusCode)
}
