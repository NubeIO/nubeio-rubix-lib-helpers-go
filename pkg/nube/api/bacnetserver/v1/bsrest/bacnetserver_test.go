package bsrest

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest/v1/rest"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
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

	ping, resp := bacnetClient.Ping()
	fmt.Println(ping)
	if resp.GetError() != nil || ping == nil {
		fmt.Println(resp.GetError())
		fmt.Println(resp.GetStatusCode())
		return
	}

	fmt.Println(ping.UpHour)
	pprint.PrintStrut(resp)

	point := &model.Point{}

	//var addr *int
	value := 0
	point.AddressID = &value

	fmt.Println(nils.IntIsNil(point.AddressID))

	bacnetPoint := &BacnetPoint{}
	bacnetPoint.Description = nils.RandomString()
	bacnetPoint.ObjectName = nils.RandomString()
	bacnetPoint.Enable = true
	bacnetPoint.UseNextAvailableAddr = false
	bacnetPoint.Address = nils.IntIsNil(point.AddressID)
	bacnetPoint.ObjectType = "analogOutput"
	bacnetPoint.COV = 0
	bacnetPoint.EventState = "normal"
	bacnetPoint.Units = "noUnits"

	addpoint, r := bacnetClient.AddPoint(bacnetPoint)
	fmt.Println(addpoint)
	pprint.PrintStrut(r.GetResponse())
	fmt.Println(r.Response.Message)

}
