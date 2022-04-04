package bsrest

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest/v1/rest"
	"testing"
)

func TestBACnetRest(*testing.T) {
	//commonClient := &nube_api.NubeRest{UseRubixProxy: true}

	restService := &rest.Service{}
	restService.Port = 1717
	restOptions := &rest.Options{}
	restService.Options = restOptions
	restService = rest.New(restService)

	nubeProxy := &rest.NubeProxy{}
	nubeProxy.UseRubixProxy = false
	nubeProxy.RubixUsername = "admin"
	nubeProxy.RubixPassword = "N00BWires"
	restService.NubeProxy = nubeProxy

	bacnetClient := New(&BacnetClient{Rest: restService})

	ping, res := bacnetClient.Ping()

	fmt.Println(ping.UpHour)
	pprint.PrintStrut(res)

	bacnetPoint := &BacnetPoint{}
	bacnetPoint.Description = nils.RandomString()
	bacnetPoint.ObjectName = nils.RandomString()
	bacnetPoint.Enable = true
	bacnetPoint.UseNextAvailableAddr = true
	//bacnetPoint.Address = nils.RandomInt(0, 20000)
	bacnetPoint.ObjectType = "analogOutput"
	bacnetPoint.COV = 0
	bacnetPoint.EventState = "normal"
	bacnetPoint.Units = "noUnits"

	addpoint, r := bacnetClient.AddPoint(bacnetPoint)
	fmt.Println(addpoint)
	fmt.Println(r.StatusCode)

}
