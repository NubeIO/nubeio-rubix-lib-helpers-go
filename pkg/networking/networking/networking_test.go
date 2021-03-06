package networking

import (
	"fmt"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"testing"
)

func TestNetworking(*testing.T) {

	nic := "wlp3s0"

	nets := NewNets()
	all, _ := nets.GetNetworks()
	pprint.PrintStrut(all)
	getIP, err := nets.GetGatewayIP(nic)
	pprint.PrintStrut(err)
	pprint.PrintStrut(getIP)

	names, err := nets.GetNetworksThatHaveGateway()
	if err != nil {
		return
	}

	for _, net := range names {
		fmt.Println(net.Interface)
	}

	pprint.PrintStrut(names)
	//fmt.Println(names)
	//
	//names, err := n.GetValidNetInterfacesForWeb()
	//if err != nil {
	//	return
	//}
	//pprint.PrintStrut(names)

}
