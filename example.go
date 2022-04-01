package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bugs"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/subnet"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/thermistor"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"net"
	"time"
)

type T struct {
	Health   string `json:"health"`
	Database string `json:"database"`
}

func printTime(t time.Time) {
	zone, offset := t.Zone()
	fmt.Println(t.Format(time.Kitchen), "Zone:", zone, "Offset UTC:", offset)
}

func main() {

	fmt.Println(bugs.GetFuncName(printTime))

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

	printTime(time.Now())
	printTime(time.Now().UTC())

	loc, _ := time.LoadLocation("America/New_York")
	printTime(time.Now().In(loc))

	//Or try
	//https://github.com/brotherpowers/ipsubnet

	var ip = &subnet.Subnet{}
	err = ip.Calculate("10.0.1.8/16")

	fmt.Println("SUBNET----ERROR", err)
	fmt.Printf("CIDR: %d\n", ip.CIDR)
	fmt.Printf("IP Address uint32: 0x%08X\n", ip.IPUINT32)
	fmt.Printf("IP Address: %v\n", ip.IP)
	fmt.Printf("Broadcast Address uint32: 0x%08X\n", ip.BroadcastAddressUINT32)
	fmt.Printf("Broadcast Address: %v\n", ip.BroadcastAddress)
	fmt.Printf("Network Address uint32: 0x%08X\n", ip.NetworkAddressUINT32)
	fmt.Printf("Network Address: %v\n", ip.NetworkAddress)
	fmt.Printf("Subnet Mask uint32: 0x%08X\n", ip.SubnetMaskUINT32)
	fmt.Printf("Subnet Mask: %v\n", ip.SubnetMask)
	fmt.Printf("Wildcard uint32: 0x%08X\n", ip.WildcardUINT32)
	fmt.Printf("Wildcard: %v\n", ip.Wildcard)
	fmt.Printf("Subnet Bitmap: %q\n", ip.SubnetBitmap)
	fmt.Printf("Number of Hosts: %d\n", ip.HostsMAX)

	//convert 255.255.255.0 to 24
	mask := net.IPMask(net.ParseIP("255.255.0.0").To4()) // If you have the mask as a string
	prefixSize, _ := mask.Size()
	fmt.Println(prefixSize)

}
