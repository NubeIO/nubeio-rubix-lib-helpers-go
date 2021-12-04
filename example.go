package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/ssh"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/admin"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/systemd"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/ufw"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/thermistor"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
)

type T struct {
	Health   string `json:"health"`
	Database string `json:"database"`
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

	h := ssh.Host{
		IP:          "123.209.74.192",
		Port:        2022,
		Username:    "pi",
		Password:    "N00BRCRC",
		IsLocalhost: false,
		CommandOpts: ssh.CommandOpts{
			CMD: "pwd",
		},
	}
	command, _, err := h.RunCommand()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(command)

	a := admin.Admin{
		Host: h,
	}

	arch, _, err := a.DetectArch()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", arch)

	node := admin.Admin{
		Host: h,
	}

	nodeVersion, _, err := node.NodeGetVersion()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", nodeVersion)

	u := ufw.UFW{
		Host: h,
	}

	status, isInstalled, err := u.UWFStatus()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("UFW STATUS:", status, isInstalled)

	reset, err := u.UWFReset()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("UFW reset:", reset)

	profile, err := u.UFWLoadProfile(true)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(profile.PortsCurrentState)

	s := systemd.DefaultService{
		Name: "mosquitto",
		Host: h,
	}
	isInstalled, err = s.IsInstalled()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("isInstalled", isInstalled)
	start, err := s.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(start)

}
