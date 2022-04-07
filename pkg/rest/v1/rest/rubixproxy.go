package rest

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

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
ProxyResponse
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
type ProxyResponse struct {
	Response response
	err      error
}

type response struct {
	StatusCode    int         `json:"status_code"`
	GatewayStatus bool        `json:"gateway_status"` //gateway_status: shows the connection between our rubix-service and your app is successful or not (that also means: rubix-service is up or not)? If it’s successful, then the rubix-service is up so gateway_status is true, otherwise it will be false.
	ServiceStatus bool        `json:"service_status"` //will be true if the service is unreachable (as example bacnet-server)
	BadRequest    bool        `json:"bad_request"`    //this is for if the service is online but got a 404
	Message       interface{} `json:"message"`        //"Not Found!",
	BodyType      DataType    `json:"body_type"`      //As an array "rows": [{"name": "point1"}, {"name": "point2"}], { "name": "point1"},
	Body          interface{} `json:"data"`           //As an object
}

//// New returns a new instance of the nube common apis
//func New(nubeRest *NubeRest, rest *Service) *NubeRest {
//	if nubeRest.UseRubixProxy {
//		if nubeRest.RubixPort == 0 {
//			nubeRest.RubixPort = nube.Services.RubixService.Port
//		}
//		if nubeRest.RubixProxyPath == "" {
//			nubeRest.Error = errors.New("proxy path must not be empty")
//			//return nil
//		}
//	}
//	nubeRest.Rest = rest
//	return nubeRest
//}

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

func (res *ProxyResponse) Log() {
	if res.Response.BadRequest {
		log.Errorln(res.Response.Message)
		log.Errorln(res.Response.StatusCode)
	} else {
		pprint.PrintStrut(res.Response.Body)
		log.Println(res.Response.StatusCode)
	}

}

//BuildResponse formats the API resp
func (s *Service) BuildResponse(res *Reply, body interface{}) *ProxyResponse {
	statusCode := res.Status()
	err := res.Error()
	resp := &ProxyResponse{}
	resp.Response.StatusCode = statusCode
	responseIsJSON := res.ApiResponseIsJSON
	resp.Response.ServiceStatus = true
	resp.Response.GatewayStatus = true
	resp.Response.BadRequest = true
	if statusCode == 0 || StatusCodesAllBad(statusCode) { //if status code is 0 it means that either rubix-service is down or a rubix app
		if statusCode == 0 {
			resp.Response.StatusCode = 503
			resp.Response.GatewayStatus = false
			resp.Response.ServiceStatus = false
		}
		if s.NubeProxy.UseRubixProxy {
			if responseIsJSON {
				resp.Response.ServiceStatus = false
				resp.Response.Message = res.AsJsonNoErr()
				return resp
			} else if statusCode == 0 {
				err = fmt.Errorf("rubix-service is offline")
				resp.Response.Message = err.Error()
				return resp
			}

		}
		resp.Response.BodyType = TypeError
		bodyError := ""
		if err != nil {
			bodyError = err.Error()
		} else {
			bodyError = res.AsString()
		}

		if failedBodyMessages(bodyError) {
			if statusCode == 0 { //bacnet-service is offline
				resp.Response.ServiceStatus = false
				err = fmt.Errorf("service: %s is offline", s.AppService)
				resp.Response.Message = err.Error()
			} else { //bacnet-service is online but bad req
				resp.Response.ServiceStatus = true
				err = fmt.Errorf("service: %s bad request", s.AppService)
				resp.Response.Message = err.Error()
			}
			return resp
		} else {
			resp.Response.ServiceStatus = false
			if responseIsJSON {
				resp.Response.Message = res.AsJsonNoErr()
			} else if bodyError != "" {
				resp.Response.Message = bodyError
			} else {
				err = fmt.Errorf("service: %s is offline or bad request", s.AppService)
				resp.Response.Message = err.Error()
			}
			return resp
		}
	}
	getType := types.DetectMapTypes(res.AsJsonNoErr())
	err = res.ToInterface(body)
	noBody := EmptyBody{
		EmptyBody: "no content",
	}
	if getType.IsArray {
		resp.Response.BodyType = TypeArray
		resp.Response.Body = res.AsJsonNoErr()
	} else if getType.IsMap {
		resp.Response.BodyType = TypeObject
		resp.Response.Body = res.AsJsonNoErr()
	} else if getType.IsString {
		resp.Response.BodyType = TypeString
		resp.Response.Message = res.AsJsonNoErr()
	} else if res.AsString() == "" {
		resp.Response.BodyType = TypeString
		resp.Response.Message = noBody
	} else {
		resp.Response.BodyType = TypeString
		resp.Response.Message = noBody
	}
	resp.Response.BadRequest = false
	if statusCode == 204 { //some app's return this when deleting, and it will not return our body so change to 200
		resp.Response.StatusCode = 200
	} else {
		resp.Response.StatusCode = statusCode
	}
	return resp
}

//FixPath will change the nube proxy and the service port ie: from bacnet 1717 to rubix-service port 1616
func (s *Service) FixPath() *Service {
	proxyName := s.NubeProxy.RubixProxyPath
	proxyBacnet := nube.Services.BacnetServer.Proxy
	proxyFF := nube.Services.FlowFramework.Proxy
	if s.NubeProxy.UseRubixProxy {
		if proxyName == proxyFF { //api/bacnet/points
			s.Path = fmt.Sprintf("/%s%s", proxyFF, s.Path)
		} else if proxyName == proxyBacnet {
			s.Path = fmt.Sprintf("/%s%s", proxyBacnet, s.Path)
		}
	}
	return s
}

//LogErr log error messages
func (s *Service) LogErr(errMsg error) {
	if errMsg != nil {
		e := fmt.Sprintf("%s.%s()  error:%v", s.LogPath, s.LogFunc, errMsg)
		log.Errorln(e)
	}
}

// GetToken get rubix-service token
func (s *Service) GetToken() (proxyReturn ProxyReturn) {
	token := s.NubeProxy.RubixToken
	if token == "" || tokenTimeDiffMin(s.NubeProxy.RubixTokenLastUpdate, 15) {
		options := &Options{
			Timeout:          2 * time.Second,
			RetryCount:       2,
			RetryWaitTime:    2 * time.Second,
			RetryMaxWaitTime: 0,
			Body:             map[string]interface{}{"username": s.NubeProxy.RubixUsername, "password": s.NubeProxy.RubixPassword},
		}

		s.Port = s.NubeProxy.RubixPort
		s.Path = "/api/users/login"
		s.Method = POST
		s.Options = options

		resp := s.DoRequest()
		statusCode := resp.Status()
		res := new(TokenResponse)
		err := resp.ToInterface(&res)
		if err != nil || statusCode != 200 || res.AccessToken == "" {
			log.Errorln("failed to get token", resp.AsString(), statusCode)
		}
		s.NubeProxy.RubixToken = res.AccessToken
		proxyReturn.Token = s.NubeProxy.RubixToken
	}
	return

}
