package rest_bacnet

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/api/nube_api"
	"testing"
)

func TestBACnetRest(*testing.T) {

	commonClient := &nube_api.NubeRest{UseRubixProxy: false}
	newReq := New(&BacnetClient{"", 0, commonClient})
	newReq.Ping()

}
