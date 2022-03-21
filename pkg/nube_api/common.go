package nube_api

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	log "github.com/sirupsen/logrus"
	"reflect"
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
)

var ObjectTypesMap = map[DataType]int8{
	//bacnet
	TypeObject: 0, TypeArray: 0,
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
	Status        bool        `json:"status"`         //status: if it’s on the range 200-299 then status is true.
	BodyCount     int         `json:"count"`
	Message       interface{} `json:"message"` //"Not Found!",
	Type          DataType    `json:"type"`    //As an array "rows": [{"name": "point1"}, {"name": "point2"}], { "name": "point1"},
	Body          interface{} `json:"data"`    //As an object
}

type RestResponse struct {
	ApiReply *nrest.Reply
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
	DefaultPathBacnet = "api/bacnet" //main api path's

	ProxyFF           = "ff" // rubix proxy path's
	ProxyBacnet       = "bacnet"
	ProxyBacnetMaster = "rbm"
	ProxyLora         = "lora"
	ProxyPointServer  = "ps"

	DefaultPortFlow         = 1660 // rubix default ports
	DefaultPortBacnet       = 1717
	DefaultPortBacnetMaster = 1718
	DefaultPortLoRa         = 1919
	DefaultPointServer      = 1515
	DefaultModbus           = 1516
	DefaultRubixService     = 1616
	DefaultRubixBios        = 1615
)

//BuildResponse formats the API response
func (inst *NubeRest) BuildResponse(res *nrest.Reply, err error, body interface{}) (response RestResponse) {
	response.ApiReply = res
	response.Response.StatusCode = res.StatusCode
	response.Response.GatewayStatus = res.RemoteServerOffline
	if err != nil {
		response.Response.Status = false
		inst.LogErr(err)
	}

	getType := types.DetectMapTypes(res.AsJsonNoErr())
	fmt.Println("ApiResponseIsBad", res.ApiResponseIsBad)
	fmt.Println("er", err)
	fmt.Println("res.StatusCode", res.StatusCode)

	if res.ApiResponseIsBad {
		response.Response.Status = true
		response.Response.Message = err.Error()
		if res.ApiResponseIsJSON { //if response is json then pass on the body response
			response.Response.Message = response.ApiReply.AsJsonNoErr()
		}
	}

	if !res.ApiResponseIsBad {
		err = res.ToInterface(&body)
		if err != nil {
			response.Response.Status = false
			inst.LogErr(err)
		} else {
			if getType.IsArray {
				response.Response.Type = TypeArray
				response.Response.Body = response.ApiReply.AsJsonNoErr()
			} else if getType.IsMap {
				response.Response.Type = TypeObject
				response.Response.BodyCount = response.ApiReply.ApiResponseLength
				response.Response.Body = response.ApiReply.AsJsonNoErr()
			}
		}
	}
	return
}

//FixPath will change the nube proxy and the service port ie: from bacnet 1717 to rubix-service port 1616
func (inst *NubeRest) FixPath() *NubeRest {
	proxyName := inst.RubixProxyPath
	if inst.UseRubixProxy {
		inst.Rest.Port = inst.RubixPort
		if proxyName == ProxyFF { //api/bacnet/points
			inst.Rest.Path = fmt.Sprintf("/%s%s", ProxyFF, inst.Rest.Path)
		} else if proxyName == ProxyBacnet {
			inst.Rest.Path = fmt.Sprintf("/%s%s", ProxyBacnet, inst.Rest.Path)
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
		response, statusCode, err := nrest.DoHTTPReq(inst.Rest, options)
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
