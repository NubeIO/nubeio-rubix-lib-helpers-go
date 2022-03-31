package rest_bacnet

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/api/nube_api"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/api/rest/v1"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube_apps"
)

type BacnetClient struct {
	Url     string
	Port    int
	nubeApi *nube_api.NubeRest
}

type Path struct {
	Name string
}

var Paths = struct {
	Points Path
}{
	Points: Path{Name: "/api/bacnet/points"},
}

//type NewClient struct {
//	Url      string
//	Port     int
//	NubeRest *nube_api.NubeRest
//}

//var client *RestClient

// New returns a new instance of the nube common apis
func New(inst *BacnetClient) *BacnetClient {

	if inst.nubeApi.UseRubixProxy {
		inst.Port = nube_apps.Services.RubixService.Port
	}
	if inst.Port == 0 {
		inst.Port = nube_apps.Services.BacnetServer.Port
	}
	client := rest.New(&rest.Service{Url: inst.Url, Port: inst.Port})
	inst.nubeApi.RubixProxyPath = nube_apps.Services.BacnetServer.Proxy
	inst.nubeApi.Rest = client
	nr := &BacnetClient{
		nubeApi: inst.nubeApi,
	}
	return nr
}

func (inst *BacnetClient) builder(method string, logFunc interface{}, path string) *rest.Service {
	inst.nubeApi.Rest.Method = method
	inst.nubeApi.Rest.Path = path
	inst.nubeApi.Rest.LogFunc = nube_api.GetFunctionName(logFunc)
	inst.nubeApi = inst.nubeApi.FixPath()
	return inst.nubeApi.Rest

}

// Ping get all
func (inst *BacnetClient) Ping() (response *nube_api.RestResponse) {
	path := nube_apps.Services.BacnetServer.PingPath
	inst.nubeApi.Rest = inst.builder(rest.GET, inst.Ping, path)
	res := inst.nubeApi.Rest.Request()
	response = inst.nubeApi.BuildResponse(res, nil)
	return
}

//// GetPoints get all
//func (inst *RestClient) GetPoints() (points []BacnetPoint, response nube_api.RestResponse) {
//	path := fmt.Sprintf("/%s/points", nube_api.DefaultPathBacnet)
//	inst.NubeRest.Rest = inst.builder(rest.GET, inst.GetPoints, path)
//	res := inst.NubeRest.Rest.Request()
//	response = inst.NubeRest.BuildResponse(res, res.Err, &points)
//	return
//}

//
//// GetPoint get one by its uuid
//func (inst *RestClient) GetPoint(uuid string) (points BacnetPoint, response nube_api.RestResponse) {
//	inst.NubeRest.Rest.Method = rest.GET
//	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.GetPoint)
//	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points/uuid/%s", nube_api.DefaultPathBacnet, uuid)
//	inst.NubeRest = inst.NubeRest.FixPath()
//	res := rest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
//	response = inst.NubeRest.BuildResponse(res, res.Err, &points)
//	return
//}

// AddPoint add one object
//func (inst *RestClient) AddPoint(body BacnetPoint) (points BacnetPoint, response *nube_api.RestResponse) {
//	path := nube_apps.Services.BacnetServer.PingPath
//	inst.nubeRest.Rest = inst.builder(rest.GET, inst.Ping, path)
//	inst.nubeRest.Rest.Options.Json = body
//	res := inst.nubeRest.Rest.Request()
//	response = inst.nubeRest.BuildResponse(res, &points)
//	return
//}

//
//// UpdatePoint update one object
//func (inst *RestClient) UpdatePoint(uuid string, body BacnetPoint) (points BacnetPoint, response nube_api.RestResponse) {
//	inst.NubeRest.Rest.Method = rest.PATCH
//	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.UpdatePoint)
//	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points/uuid/%s", nube_api.DefaultPathBacnet, uuid)
//	inst.NubeRest = inst.NubeRest.FixPath()
//	inst.Options.Json = body
//	res := rest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
//	response = inst.NubeRest.BuildResponse(res, res.Err, &points)
//	return
//}
//
//// DeletePoint delete one by its uuid
//func (inst *RestClient) DeletePoint(uuid string) (response nube_api.RestResponse, notFound bool, deletedOk bool) {
//	inst.NubeRest.Rest.Method = rest.DELETE
//	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.DeletePoint)
//	inst.NubeRest.Rest.Path = fmt.Sprintf("/%s/points/uuid/%s", nube_api.DefaultPathBacnet, uuid)
//	inst.NubeRest = inst.NubeRest.FixPath()
//	res := rest.DoHTTPReq(inst.NubeRest.Rest, inst.Options)
//	if strings.Contains(res.AsString(), "uuid") {
//		notFound = true
//	} else if res.StatusCode == 204 {
//		deletedOk = true
//	}
//	response = inst.NubeRest.BuildResponse(res, res.Err, nil)
//	return
//}
//
//// DropPoints delete all objects
//func (inst *RestClient) DropPoints() (response nube_api.RestResponse) {
//	inst.NubeRest.Rest.LogFunc = nube_api.GetFunctionName(inst.DropPoints)
//	points, response := inst.GetPoints()
//	statusCode := response.Response.StatusCode
//	if rest.StatusCode2xx(statusCode) {
//		count := 0
//		for _, pnt := range points {
//			count++
//			inst.DeletePoint(pnt.UUID)
//		}
//		response.Response.BodyType = nube_api.TypeString
//		response.Response.Body = ""
//		response.Response.Message = fmt.Sprintf("points delete %d", count)
//	}
//
//	return
//}
