package nube_api

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/api/nrest"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube_apps"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type NubeRest struct {
	Rest                 *nrest.ReqType
	UseRubixProxy        bool   //if true then use rubix-service proxy
	RubixProxyPath       string //the proxy path is what is used in rubix-service to append the url path ps, lora, bacnet
	RubixPort            int
	RubixToken           string
	RubixTokenLastUpdate time.Time
	RubixUsername        string
	RubixPassword        string
}

type DataType string

const (
	TypeObject DataType = "object"
	TypeArray  DataType = "array"
	TypeString DataType = "string"
	TypeError  DataType = "error"
	TypeNull   DataType = "null"
)

var ObjectTypesMap = map[DataType]int8{
	TypeObject: 0, TypeArray: 0, TypeString: 0,
}

func checkType(t string) (DataType, error) {
	if t == "" {
		return TypeObject, nil
	}
	objType := DataType(t)
	if _, ok := ObjectTypesMap[objType]; !ok {
		return "", errors.New("please provide a valid type ie: int, float, array or object")
	}
	return objType, nil
}

/*
Response
status_code: proxy that status code --catch it and send it here
gateway_status: shows the connection between our rubix-service and your app is successful or not (that also means: rubix-service is up or not)? If it’s successful, then the rubix-service is up so gateway_status is true, otherwise it will be false.
status: if it’s on the range 200-299 then status is true.
message:
- If status is on the range 200-299 then it will be null
- Else, we try to parse it on JSON:
-- if data is not parsable into JSON put that string content directly there
-- if data is parsable but doesn’t have the message key put that content directly there
-- if data is parsable and does have the message key, extract message key and put that message key content
type: detect whether that output is JSON object or JSON array, if it’s JSON array, put that content on here --this will be so much easy for user to parse the content accordingly
And just return 200 HTTP status code all the time. Coz, we are using status_code in JSON body.
*/
type Response struct {
	StatusCode    int         `json:"status_code"`
	GatewayStatus bool        `json:"gateway_status"` //gateway_status: shows the connection between our rubix-service and your app is successful or not (that also means: rubix-service is up or not)? If it’s successful, then the rubix-service is up so gateway_status is true, otherwise it will be false.
	ServiceStatus bool        `json:"service_status"` //will be true if the service is unreachable (as example bacnet-server)
	BadRequest    bool        `json:"bad_request"`    //this is for if the service is online but got a 404
	Message       interface{} `json:"message"`        //"Not Found!",
	BodyType      DataType    `json:"body_type"`      //As an array "rows": [{"name": "point1"}, {"name": "point2"}], { "name": "point1"},
	Body          interface{} `json:"data"`           //As an object

}

type RestResponse struct {
	//ApiReply *nrest.Reply
	Response Response
}

func tokenTimeDiffMin(t time.Time, timeDiff float64) (out bool) {
	t1 := time.Now()
	if t1.Sub(t).Minutes() > timeDiff {
		out = true
	}
	return
}

func GetFunctionName(temp interface{}) string {
	s := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
	return s[len(s)-1]
}

// New returns a new instance of the nube common apis
func New(RestClient *NubeRest) *NubeRest {
	return RestClient
}

type TokenBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type TokenResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	Message     *string `json:"message,omitempty"`
}

var err error

type ProxyReturn struct {
	Token string
}

const (
	BaseURL = "0.0.0.0"

	DefaultPathBacnet = "api/bacnet" //main api path's

)

type EmptyBody struct {
	EmptyBody string `json:"empty_body"`
}

func failedBodyMessages(bodyError string) (found bool) {
	re := regexp.MustCompile(`HTML|NOT FOUND|connect: connection refused`)
	found = re.MatchString(bodyError)
	return

}

//BuildResponse formats the API response
func (inst *NubeRest) BuildResponse(res *nrest.Reply, err error, body interface{}) (response RestResponse) {
	statusCode := res.StatusCode
	response.Response.StatusCode = statusCode
	responseIsJSON := res.ApiResponseIsJSON
	response.Response.ServiceStatus = true
	response.Response.GatewayStatus = true
	response.Response.BadRequest = true
	if statusCode == 0 || nrest.StatusCodesAllBad(statusCode) { //if status code is 0 it means that either rubix-service is down or a rubix app
		if statusCode == 0 {
			response.Response.StatusCode = 503
			response.Response.GatewayStatus = false
			response.Response.ServiceStatus = false
		}
		if inst.UseRubixProxy {
			if responseIsJSON {
				response.Response.ServiceStatus = false
				response.Response.Message = res.AsJsonNoErr()
				return
			} else if statusCode == 0 {
				err = fmt.Errorf("rubix-service is offline")
				response.Response.Message = err.Error()
				return
			}

		}
		response.Response.BodyType = TypeError
		bodyError := ""
		if err != nil {
			bodyError = err.Error()
		} else {
			bodyError = res.AsString()
		}

		if failedBodyMessages(bodyError) {
			if statusCode == 0 { //bacnet-service is offline
				response.Response.ServiceStatus = false
				err = fmt.Errorf("service: %s is offline", inst.Rest.Service)
				response.Response.Message = err.Error()
			} else { //bacnet-service is online but bad req
				response.Response.ServiceStatus = true
				err = fmt.Errorf("service: %s bad request", inst.Rest.Service)
				response.Response.Message = err.Error()
			}
			return
		} else {
			response.Response.ServiceStatus = false
			if responseIsJSON {
				response.Response.Message = res.AsJsonNoErr()
			} else if bodyError != "" {
				response.Response.Message = bodyError
			} else {
				err = fmt.Errorf("service: %s is offline or bad request", inst.Rest.Service)
				response.Response.Message = err.Error()
			}
			return
		}
	}
	getType := types.DetectMapTypes(res.AsJsonNoErr())
	err = res.ToInterface(&body)
	noBody := EmptyBody{
		EmptyBody: "no content",
	}
	if getType.IsArray {
		response.Response.BodyType = TypeArray
		response.Response.Body = res.AsJsonNoErr()
	} else if getType.IsMap {
		response.Response.BodyType = TypeObject
		response.Response.Body = res.AsJsonNoErr()
	} else if getType.IsString {
		response.Response.BodyType = TypeString
		response.Response.Message = res.AsJsonNoErr()
	} else if res.AsString() == "" {
		response.Response.BodyType = TypeString
		response.Response.Message = noBody
	} else {
		response.Response.BodyType = TypeString
		response.Response.Message = noBody
	}
	response.Response.BadRequest = false
	if statusCode == 204 { //some app's return this when deleting, and it will not return our body so change to 200
		response.Response.StatusCode = 200
	} else {
		response.Response.StatusCode = statusCode
	}
	return
}

//FixPath will change the nube proxy and the service port ie: from bacnet 1717 to rubix-service port 1616
func (inst *NubeRest) FixPath() *NubeRest {
	proxyName := inst.RubixProxyPath
	proxyBacnet := nube_apps.Services.BacnetServer.Proxy
	proxyFF := nube_apps.Services.FlowFramework.Proxy
	if inst.UseRubixProxy {
		inst.Rest.Port = inst.RubixPort
		if proxyName == proxyFF { //api/bacnet/points
			inst.Rest.Path = fmt.Sprintf("/%s%s", proxyFF, inst.Rest.Path)
		} else if proxyName == proxyBacnet {
			inst.Rest.Path = fmt.Sprintf("/%s%s", proxyBacnet, inst.Rest.Path)
		}
	}
	return inst
}

//LogErr log error messages
func (inst *NubeRest) LogErr(errMsg error) {
	if errMsg != nil {
		e := fmt.Sprintf("%s.%s()  error:%v", inst.Rest.LogPath, inst.Rest.LogFunc, errMsg)
		log.Errorln(e)
	}
}

// GetToken get rubix-service token
func (inst *NubeRest) GetToken() (proxyReturn ProxyReturn) {
	token := inst.RubixToken
	if token == "" || tokenTimeDiffMin(inst.RubixTokenLastUpdate, 15) {
		options := &nrest.ReqOpt{
			Timeout:          2 * time.Second,
			RetryCount:       2,
			RetryWaitTime:    2 * time.Second,
			RetryMaxWaitTime: 0,
			Json:             map[string]interface{}{"username": inst.RubixUsername, "password": inst.RubixPassword},
		}

		inst.Rest.Port = inst.RubixPort
		inst.Rest.Path = "/api/users/login"
		inst.Rest.Method = nrest.POST
		response := nrest.DoHTTPReq(inst.Rest, options)
		statusCode := response.StatusCode
		res := new(TokenResponse)
		err = response.ToInterface(&res)
		if err != nil || statusCode != 200 || res.AccessToken == "" {
			log.Errorln("failed to get token", response.AsString(), statusCode)
		}
		inst.RubixToken = res.AccessToken
		proxyReturn.Token = inst.RubixToken
	}
	return

}
