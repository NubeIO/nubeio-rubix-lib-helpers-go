package nube_api_bacnetserver

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube_api"
)

type RestClient struct {
	NubeRest *nube_api.NubeRest
	Options  *nrest.ReqOpt
}

var err error

// GetPoints get all
func (inst *RestClient) GetPoints() (points []BacnetPoint, response nube_api.RestResponse) {
	inst.NubeRest.Rest.Method = nrest.GET
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.GetPoints)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points", nube_api.DefaultPathBacnet)
	inst.NubeRest = inst.NubeRest.FixPath()
	res, _, err := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.NubeRest.BuildResponse(res, err, &points)
	return
}

// GetPoint get one by its uuid
func (inst *RestClient) GetPoint(uuid string) (points BacnetPoint, response nube_api.RestResponse) {
	inst.NubeRest.Rest.Method = nrest.GET
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.GetPoints)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points/uuid/%s", nube_api.DefaultPathBacnet, uuid)
	inst.NubeRest = inst.NubeRest.FixPath()
	res, _, err := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.NubeRest.BuildResponse(res, err, &points)
	return
}

// AddPoint add one
func (inst *RestClient) AddPoint(body BacnetPoint) (points BacnetPoint, response nube_api.RestResponse) {
	inst.NubeRest.Rest.Method = nrest.POST
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.GetPoints)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points", nube_api.DefaultPathBacnet)
	inst.NubeRest = inst.NubeRest.FixPath()
	inst.Options.Json = body
	res, _, err := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.NubeRest.BuildResponse(res, err, &points)
	return
}
