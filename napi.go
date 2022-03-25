package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube_api"
	nube_api_bacnetserver "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube_api/bacnetserver"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"time"
)

func main() {

	//inc generic reset client
	rc := &nrest.ReqType{
		BaseUri: "0.0.0.0",
		LogPath: "helpers.nrest",
		Service: "bacnet-server",
	}

	//inc nube rest client
	c := &nube_api.NubeRest{
		Rest:          rc,
		RubixPort:     nube_api.DefaultRubixService,
		RubixUsername: "admin",
		RubixPassword: "N00BWires",
		UseRubixProxy: false,
	}

	//new nube rest client
	nubeRest := nube_api.New(c)
	//nubeRest.GetToken()

	//bacnet client
	options := &nrest.ReqOpt{
		Timeout:          500 * time.Second,
		RetryCount:       0,
		RetryWaitTime:    4 * time.Second,
		RetryMaxWaitTime: 0,
		Headers:          map[string]interface{}{"Authorization": nubeRest.RubixToken},
	}
	rc.LogPath = "helpers.nrest.bacnet.server"
	rc.Port = nube_api.DefaultPortBacnet
	c.RubixProxyPath = nube_api.ProxyBacnet
	bacnetClient := &nube_api_bacnetserver.RestClient{
		NubeRest: nubeRest,
		Options:  options,
	}
	//get points
	pnts, r := bacnetClient.GetPoints()
	pprint.Print(r)
	fmt.Println("pnts", pnts)

}
