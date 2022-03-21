package nube_api

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nrest"
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

type RestResponse struct {
	ApiReply *nrest.Reply
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
	ProxyFF           = "ff"
	ProxyBacnet       = "bacnet"
	ProxyBacnetMaster = "rbm"
	ProxyLora         = "lora"
	ProxyPointServer  = "ps"

	DefaultPortFlow         = 1660
	DefaultPortBacnet       = 1717
	DefaultPortBacnetMaster = 1718
	DefaultPortLoRa         = 1919
	DefaultPointServer      = 1515
	DefaultModbus           = 1516
	DefaultRubixService     = 1616
	DefaultRubixBios        = 1615
)

func (inst *NubeRest) FixPath() *NubeRest {
	proxyName := inst.RubixProxyPath
	if inst.UseRubixProxy {
		inst.Rest.Port = inst.RubixPort
		if proxyName == ProxyFF {
			inst.Rest.Path = fmt.Sprintf("/%s%s", ProxyFF, inst.Rest.Path)
		} else if proxyName == ProxyBacnet {
			inst.Rest.Path = fmt.Sprintf("/%s%s", ProxyBacnet, inst.Rest.Path)
		}
	}
	return inst
}

func (inst *NubeRest) LogErr(errMsg error) {
	if errMsg != nil {
		e := fmt.Sprintf("%s.%s()  error:%v", inst.Rest.LogPath, inst.Rest.LogFunc, errMsg)
		log.Errorln(e)
	}
}

// GetToken all points
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
