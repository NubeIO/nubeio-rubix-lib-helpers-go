package bsrest

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest/v1/rest"
	"strings"
)

type BacnetClient struct {
	Rest *rest.Service
}

type Path struct {
	Path string
}

var Paths = struct {
	Points Path
}{
	Points: Path{Path: "/api/bacnet/points"},
}

// New returns a new instance of the nube common apis
func New(bc *BacnetClient) *BacnetClient {
	return bc
}

func (inst *BacnetClient) builder(method string, logFunc interface{}, path string) *rest.Service {
	//get token if using proxy
	if inst.Rest.NubeProxy.UseRubixProxy {
		r := inst.Rest.GetToken()
		inst.Rest.Options.Headers = map[string]interface{}{"Authorization": r.Token}
	}
	inst.Rest.Method = method
	inst.Rest.Path = path
	inst.Rest.LogFunc = rest.GetFunctionName(logFunc)
	inst.Rest = inst.Rest.FixPath()
	return inst.Rest

}

// Ping Ping server
func (inst *BacnetClient) Ping() (ping ServerPing, response *rest.RestResponse) {
	path := nube.Services.BacnetServer.PingPath
	inst.Rest = inst.builder(rest.GET, inst.Ping, path)
	res := inst.Rest.Request()
	response = inst.Rest.BuildResponse(res, &ping)
	return
}

// GetPoints  get all
func (inst *BacnetClient) GetPoints() (points []BacnetPoint, response *rest.RestResponse) {
	path := fmt.Sprintf("%s/points", Paths.Points.Path)
	inst.Rest = inst.builder(rest.GET, inst.GetPoints, path)
	res := inst.Rest.Request()
	response = inst.Rest.BuildResponse(res, &points)
	return
}

//
// GetPoint get one by its uuid
func (inst *BacnetClient) GetPoint(uuid string) (points BacnetPoint, response *rest.RestResponse) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	inst.Rest = inst.builder(rest.GET, inst.GetPoint, path)
	res := inst.Rest.Request()
	response = inst.Rest.BuildResponse(res, &points)
	return
}

// AddPoint add one object
func (inst *BacnetClient) AddPoint(body *BacnetPoint) (points BacnetPoint, response *rest.RestResponse) {
	path := Paths.Points.Path
	inst.Rest = inst.builder(rest.POST, inst.AddPoint, path)
	inst.Rest.Options.Body = body
	res := inst.Rest.Request()
	response = inst.Rest.BuildResponse(res, &points)
	return
}

// UpdatePoint update one object
func (inst *BacnetClient) UpdatePoint(uuid string, body BacnetPoint) (points BacnetPoint, response *rest.RestResponse) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	inst.Rest = inst.builder(rest.PATCH, inst.UpdatePoint, path)
	inst.Rest.Options.Body = body
	res := inst.Rest.Request()
	response = inst.Rest.BuildResponse(res, &points)
	return
}

// UpdatePointValue do a point write
func (inst *BacnetClient) UpdatePointValue(uuid string, body BacnetPoint) (points BacnetPoint, response *rest.RestResponse) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	inst.Rest = inst.builder(rest.PATCH, inst.UpdatePointValue, path)
	inst.Rest.Options.Body = body
	res := inst.Rest.Request()
	response = inst.Rest.BuildResponse(res, &points)
	return
}

// DeletePoint delete one by its uuid
func (inst *BacnetClient) DeletePoint(uuid string) (response *rest.RestResponse, notFound bool, deletedOk bool) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	inst.Rest = inst.builder(rest.DELETE, inst.DeletePoint, path)
	res := inst.Rest.Request()
	response = inst.Rest.BuildResponse(res, nil)
	if strings.Contains(res.AsString(), "uuid") {
		notFound = true
	} else if res.Status() == 204 {
		deletedOk = true
	} else if res.Status() == 404 {
		notFound = true
	}
	return
}

//
// DropPoints delete all objects
func (inst *BacnetClient) DropPoints() (response *rest.RestResponse) {
	inst.Rest.LogFunc = rest.GetFunctionName(inst.DropPoints)
	points, response := inst.GetPoints()
	statusCode := response.GetStatusCode()
	if rest.StatusCode2xx(statusCode) {
		count := 0
		for _, pnt := range points {
			count++
			inst.DeletePoint(pnt.UUID)
		}
		response.BodyType = rest.TypeString
		response.Body = ""
		response.Message = fmt.Sprintf("points delete %d", count)
	}

	return
}
