package bsrest

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/api/common/v1/iorest"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"testing"
)

func TestBACnetRest(*testing.T) {
	//commonClient := &nube_api.NubeRest{UseRubixProxy: true}
	commonClient := new(iorest.NubeRest)
	commonClient.UseRubixProxy = true
	commonClient.RubixUsername = "admin"
	commonClient.RubixPassword = "N00BWires"
	commonClient = iorest.New(commonClient)

	newReq := New(&BacnetClient{Url: "0.0.0.0", Port: 1717, IoRest: commonClient})

	ping, res := newReq.Ping()

	fmt.Println(ping.UpHour)
	pprint.PrintStrut(res)
}
