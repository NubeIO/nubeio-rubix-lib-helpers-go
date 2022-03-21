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

type Response struct {
	StatusCode    int
	ServerOffline bool
	ErrorBody     interface{}
	Body          interface{}
}

type RestResponse struct {
	ApiReply *nrest.Reply
	Response Response
}

var err error

func (inst *RestClient) BuildResponse(res *nrest.Reply, err error, body interface{}) (response RestResponse) {
	response.ApiReply = res
	response.Response.StatusCode = res.StatusCode
	response.Response.ServerOffline = res.RemoteServerOffline
	inst.NubeRest.LogErr(err)
	if res.ApiResponseIsBad {
		response.Response.ErrorBody = err.Error()
	}
	if !res.ApiResponseIsBad {
		err = res.ToInterface(&body)
		inst.NubeRest.LogErr(err)
	}
	return
}

// GetPoints get all
func (inst *RestClient) GetPoints() (points []BacnetPoint, response RestResponse) {
	inst.NubeRest.Rest.Method = nrest.GET
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.GetPoints)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/api/bacnet/points")
	inst.NubeRest = inst.NubeRest.FixPath()
	res, _, err := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.BuildResponse(res, err, &points)
	return
}

// GetPoint get one by its uuid
func (inst *RestClient) GetPoint(uuid string) (points []BacnetPoint, response RestResponse) {
	inst.NubeRest.Rest.Method = nrest.GET
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.GetPoints)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/api/bacnet/points")
	inst.NubeRest = inst.NubeRest.FixPath()
	res, _, err := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.BuildResponse(res, err, &points)
	return
}
