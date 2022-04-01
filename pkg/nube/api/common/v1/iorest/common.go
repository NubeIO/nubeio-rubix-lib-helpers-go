package iorest

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest/v1/rest"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

type NubeRest struct {
	RubixApp             nube.Service
	UseRubixProxy        bool   //if true then use rubix-service proxy
	RubixProxyPath       string //the proxy path is what is used in rubix-service to append the url path ps, lora, bacnet
	RubixPort            int
	RubixToken           string
	RubixTokenLastUpdate time.Time
	RubixUsername        string
	RubixPassword        string
	Error                error
	Rest                 *rest.Service
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
RestResponse
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
type RestResponse struct {
	StatusCode    int         `json:"status_code"`
	GatewayStatus bool        `json:"gateway_status"` //gateway_status: shows the connection between our rubix-service and your app is successful or not (that also means: rubix-service is up or not)? If it’s successful, then the rubix-service is up so gateway_status is true, otherwise it will be false.
	ServiceStatus bool        `json:"service_status"` //will be true if the service is unreachable (as example bacnet-server)
	BadRequest    bool        `json:"bad_request"`    //this is for if the service is online but got a 404
	Message       interface{} `json:"message"`        //"Not Found!",
	BodyType      DataType    `json:"body_type"`      //As an array "rows": [{"name": "point1"}, {"name": "point2"}], { "name": "point1"},
	Body          interface{} `json:"data"`           //As an object
	err           error
}

// New returns a new instance of the nube common apis
func New(nubeRest *NubeRest) *NubeRest {
	if nubeRest.UseRubixProxy {
		if nubeRest.RubixPort == 0 {
			nubeRest.RubixPort = nube.Services.RubixService.Port
		}
		if nubeRest.RubixProxyPath == "" {
			nubeRest.Error = errors.New("proxy path must not be empty")
			//return nil
		}
	}
	return nubeRest
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

type ProxyReturn struct {
	Token string
}

type EmptyBody struct {
	EmptyBody string `json:"empty_body"`
}

func failedBodyMessages(bodyError string) (found bool) {
	re := regexp.MustCompile(`HTML|NOT FOUND|connect: connection refused`)
	found = re.MatchString(bodyError)
	return

}

func (res *RestResponse) Log() {

	if res.BadRequest {
		log.Errorln(res.Message)
		log.Errorln(res.StatusCode)
	} else {
		pprint.PrintStrut(res.Body)
		log.Println(res.StatusCode)
	}

}

func (res *RestResponse) GetResponse() *RestResponse {
	return res
}

func (res *RestResponse) GetError() error {
	return res.err
}

func (res *RestResponse) GetStatusCode() int {
	return res.StatusCode
}

//BuildResponse formats the API response
func (inst *NubeRest) BuildResponse(res *rest.Reply, body interface{}) *RestResponse {
	statusCode := res.Status()
	err := res.Error()
	response := &RestResponse{}
	response.StatusCode = statusCode
	responseIsJSON := res.ApiResponseIsJSON
	response.ServiceStatus = true
	response.GatewayStatus = true
	response.BadRequest = true
	if statusCode == 0 || rest.StatusCodesAllBad(statusCode) { //if status code is 0 it means that either rubix-service is down or a rubix app
		if statusCode == 0 {
			response.StatusCode = 503
			response.GatewayStatus = false
			response.ServiceStatus = false
		}
		if inst.UseRubixProxy {
			if responseIsJSON {
				response.ServiceStatus = false
				response.Message = res.AsJsonNoErr()
				return response
			} else if statusCode == 0 {
				err = fmt.Errorf("rubix-service is offline")
				response.Message = err.Error()
				return response
			}

		}
		response.BodyType = TypeError
		bodyError := ""
		if err != nil {
			bodyError = err.Error()
		} else {
			bodyError = res.AsString()
		}

		if failedBodyMessages(bodyError) {
			if statusCode == 0 { //bacnet-service is offline
				response.ServiceStatus = false
				err = fmt.Errorf("service: %s is offline", inst.Rest.AppService)
				response.Message = err.Error()
			} else { //bacnet-service is online but bad req
				response.ServiceStatus = true
				err = fmt.Errorf("service: %s bad request", inst.Rest.AppService)
				response.Message = err.Error()
			}
			return response
		} else {
			response.ServiceStatus = false
			if responseIsJSON {
				response.Message = res.AsJsonNoErr()
			} else if bodyError != "" {
				response.Message = bodyError
			} else {
				err = fmt.Errorf("service: %s is offline or bad request", inst.Rest.AppService)
				response.Message = err.Error()
			}
			return response
		}
	}
	getType := types.DetectMapTypes(res.AsJsonNoErr())
	err = res.ToInterface(&body)
	noBody := EmptyBody{
		EmptyBody: "no content",
	}
	if getType.IsArray {
		response.BodyType = TypeArray
		response.Body = res.AsJsonNoErr()
	} else if getType.IsMap {
		response.BodyType = TypeObject
		response.Body = res.AsJsonNoErr()
	} else if getType.IsString {
		response.BodyType = TypeString
		response.Message = res.AsJsonNoErr()
	} else if res.AsString() == "" {
		response.BodyType = TypeString
		response.Message = noBody
	} else {
		response.BodyType = TypeString
		response.Message = noBody
	}
	response.BadRequest = false
	if statusCode == 204 { //some app's return this when deleting, and it will not return our body so change to 200
		response.StatusCode = 200
	} else {
		response.StatusCode = statusCode
	}
	return response
}

//FixPath will change the nube proxy and the service port ie: from bacnet 1717 to rubix-service port 1616
func (inst *NubeRest) FixPath() *NubeRest {
	proxyName := inst.RubixProxyPath
	proxyBacnet := nube.Services.BacnetServer.Proxy
	proxyFF := nube.Services.FlowFramework.Proxy
	if inst.UseRubixProxy {
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
		options := &rest.Options{
			Timeout:          2 * time.Second,
			RetryCount:       2,
			RetryWaitTime:    2 * time.Second,
			RetryMaxWaitTime: 0,
			Body:             map[string]interface{}{"username": inst.RubixUsername, "password": inst.RubixPassword},
		}

		inst.Rest.Port = inst.RubixPort
		inst.Rest.Path = "/api/users/login"
		inst.Rest.Method = rest.POST
		inst.Rest.Options = options

		response := inst.Rest.Request()
		statusCode := response.Status()
		res := new(TokenResponse)
		err := response.ToInterface(&res)
		if err != nil || statusCode != 200 || res.AccessToken == "" {
			log.Errorln("failed to get token", response.AsString(), statusCode)
		}
		inst.RubixToken = res.AccessToken
		proxyReturn.Token = inst.RubixToken
	}
	return

}
