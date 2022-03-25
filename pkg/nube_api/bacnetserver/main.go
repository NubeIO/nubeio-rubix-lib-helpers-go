package nube_api_bacnetserver

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube_api"
	"strings"
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
	res := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.NubeRest.BuildResponse(res, res.Err, &points)
	return
}

// GetPoint get one by its uuid
func (inst *RestClient) GetPoint(uuid string) (points BacnetPoint, response nube_api.RestResponse) {
	inst.NubeRest.Rest.Method = nrest.GET
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.GetPoint)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points/uuid/%s", nube_api.DefaultPathBacnet, uuid)
	inst.NubeRest = inst.NubeRest.FixPath()
	res := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.NubeRest.BuildResponse(res, res.Err, &points)
	return
}

// AddPoint add one object
func (inst *RestClient) AddPoint(body BacnetPoint) (points BacnetPoint, response nube_api.RestResponse) {
	inst.NubeRest.Rest.Method = nrest.POST
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.AddPoint)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points", nube_api.DefaultPathBacnet)
	inst.NubeRest = inst.NubeRest.FixPath()
	inst.Options.Json = body
	res := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.NubeRest.BuildResponse(res, res.Err, &points)
	return
}

// UpdatePoint update one object
func (inst *RestClient) UpdatePoint(uuid string, body BacnetPoint) (points BacnetPoint, response nube_api.RestResponse) {
	inst.NubeRest.Rest.Method = nrest.PATCH
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.UpdatePoint)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points/uuid/%s", nube_api.DefaultPathBacnet, uuid)
	inst.NubeRest = inst.NubeRest.FixPath()
	inst.Options.Json = body
	res := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	response = inst.NubeRest.BuildResponse(res, res.Err, &points)
	return
}

// DeletePoint delete one by its uuid
func (inst *RestClient) DeletePoint(uuid string) (response nube_api.RestResponse, notFound bool, deletedOk bool) {
	inst.NubeRest.Rest.Method = nrest.DELETE
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.DeletePoint)
	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points/uuid/%s", nube_api.DefaultPathBacnet, uuid)
	inst.NubeRest = inst.NubeRest.FixPath()
	res := nrest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
	if strings.Contains(res.AsString(), "uuid") {
		notFound = true
	} else if res.StatusCode == 204 {
		deletedOk = true
	}
	response = inst.NubeRest.BuildResponse(res, res.Err, nil)
	return
}

// DropPoints delete all objects
func (inst *RestClient) DropPoints() (response nube_api.RestResponse) {
	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.DropPoints)
	points, response := inst.GetPoints()
	statusCode := response.Response.StatusCode
	if nrest.StatusCode2xx(statusCode) {
		count := 0
		for _, pnt := range points {
			count++
			inst.DeletePoint(pnt.UUID)
		}
		response.Response.BodyType = nube_api.TypeString
		response.Response.Body = ""
		response.Response.Message = fmt.Sprintf("points delete %d", count)
	}

	return
}
