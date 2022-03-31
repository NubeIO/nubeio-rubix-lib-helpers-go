package rest_bacnet

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/api/nube_api"
	"testing"
)

func TestBACnetRest(*testing.T) {
	commonClient := &nube_api.NubeRest{UseRubixProxy: true}
	//commonClient := new(nube_api.NubeRest)
	//commonClient.UseRubixProxy = true
	commonClient = nube_api.New(commonClient)
	//
	newReq := New(&BacnetClient{"", 0, commonClient})

	newReq.Ping()

}
