package rest

// go http client support get,post,delete,patch,put,head,file method
// go-resty/resty: https://github.com/go-resty/resty

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	PUT    = "PUT"
	DELETE = "DELETE"
	HEAD   = "HEAD"
	FILE   = "FILE"
)

//New client
func New(s *Service) *Service {
	if s == nil {
		s = &Service{}
	}
	if s.Url == "" {
		s.Url = LocalHost
	}
	url := fmt.Sprintf("http://%s:%d", s.Url, s.Port)
	if s.HTTPS {
		url = fmt.Sprintf("https://%s:%d", s.Url, s.Port)
	}
	if s.Method == "" {
		s.Method = GET
	}
	if s.LogPath == "" {
		s.LogPath = "nube.helpers.nrest"
	}
	s.Url = url
	return s
}

/*Service
Parameters:
	: BaseUri 0.0.0.0 or nube-io.com
	: Proxy
	: Port 1616
	: HTTPS if true will set url to https
	: Path  /api/users
	: Method GET, POST, PATCH, PUT, DELETE, HEAD, FILE
	: Options *Options
Description:

*/
type Service struct {
	Url             string //0.0.0.0 or nube-io.com\
	Path            string //  /api/points
	Port            int    // 80 or 443
	HTTPS           bool
	Method          string
	Debug           bool
	Proxy           string
	AppService      string //as in bacnet-server
	LogPath         string //in the log message show path or extra message
	LogFunc         string
	EnableKeepAlive bool
	Options         *Options
}

type Options struct {
	Timeout          time.Duration
	RetryCount       int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
	Params           map[string]interface{}
	SetQueryString   string
	Data             map[string]interface{}
	Headers          map[string]interface{}
	Cookies          map[string]interface{}
	CookiePath       string
	CookieDomain     string
	CookieMaxAge     int
	CookieHttpOnly   bool
	Body             interface{}
	FileName         string
	FileParamName    string
}

type Reply struct {
	err               error
	body              []byte
	bodyJson          interface{}
	statusCode        int
	IsError           bool
	ApiResponseIsBad  bool //is true response is a 404
	ApiResponseIsJSON bool
	ApiResponseLength int
}

func (s *Service) Request() (response *Reply) {
	response = s.Do()
	statusCode := response.statusCode
	logPath := fmt.Sprintf("%s.%s() method: %s host: %s statusCode:%d", s.LogPath, s.LogFunc, strings.ToUpper(s.Method), s.Url+s.Path, statusCode)
	if response.ApiResponseIsBad {
		log.Errorln(logPath)
	} else {
		log.Println(logPath)
	}
	//check if response is JSON
	isJson := isJSON(response.AsString())
	if isJson {
		response.ApiResponseIsJSON = isJson
		//get response type as in an object or an array
		getType := types.DetectMapTypes(response.AsJsonNoErr())
		if getType.IsArray {
			response.ApiResponseLength = getJSONLen(response.AsJsonNoErr())
		} else {
			response.ApiResponseLength = 1
		}
	}
	return response
}

// Do request
// method string  get,post,put,patch,delete,head
// uri    string  BaseUri  /api/whatever
// opt 	  *ReqOpt
func (s *Service) Do() *Reply {
	reqUrl := s.Url
	method := s.Method
	path := s.Path
	if method == "" {
		return &Reply{
			err: errors.New("request method is empty"),
		}
	}
	if reqUrl == "" {
		return &Reply{
			err: errors.New("request url is empty"),
		}
	}
	opt := s.Options
	if opt == nil {
		opt = &Options{}
	}
	if path != "" {
		reqUrl = strings.TrimRight(reqUrl, "/") + path
	}
	if opt.Timeout == 0 {
		opt.Timeout = defaultTimeout
	}
	client := resty.New()
	client = client.SetTimeout(opt.Timeout) //timeout
	if !s.EnableKeepAlive {
		client = client.SetHeader("Connection", "close")
	}

	if s.Proxy != "" {
		client = client.SetProxy(s.Proxy)
	}
	if opt.RetryCount > 0 {
		if opt.RetryCount > defaultMaxRetries {
			opt.RetryCount = defaultMaxRetries
		}

		client = client.SetRetryCount(opt.RetryCount)

		if opt.RetryWaitTime != 0 {
			client = client.SetRetryWaitTime(opt.RetryWaitTime)
		}

		if opt.RetryMaxWaitTime != 0 {
			client = client.SetRetryMaxWaitTime(opt.RetryMaxWaitTime)
		}
	}
	if cLen := len(opt.Cookies); cLen > 0 {
		cookies := make([]*http.Cookie, cLen)
		for k, _ := range opt.Cookies {
			cookies = append(cookies, &http.Cookie{
				Name:     k,
				Value:    fmt.Sprintf("%v", opt.Cookies[k]),
				Path:     opt.CookiePath,
				Domain:   opt.CookieDomain,
				MaxAge:   opt.CookieMaxAge,
				HttpOnly: opt.CookieHttpOnly,
			})
		}

		client = client.SetCookies(cookies)
	}

	if len(opt.Headers) > 0 {
		client = client.SetHeaders(opt.ParseData(opt.Headers))
	}

	var resp *resty.Response
	var err error

	method = strings.ToLower(method)
	switch method {
	case "get", "delete", "head":
		client = client.SetQueryParams(opt.ParseData(opt.Params))
		if method == "get" {
			resp, err = client.R().SetQueryString(opt.SetQueryString).Get(reqUrl)
			return s.GetResult(resp, err)
		}

		if method == "delete" {
			resp, err = client.R().Delete(reqUrl)
			return s.GetResult(resp, err)
		}

		if method == "head" {
			resp, err = client.R().Head(reqUrl)
			return s.GetResult(resp, err)
		}

	case "post", "put", "patch":
		req := client.R().SetQueryString(opt.SetQueryString)
		if len(opt.Data) > 0 {
			// SetFormData method sets Form parameters and their values in the current request.
			// It's applicable only HTTP method `POST` and `PUT` and requests content type would be
			// set as `application/x-www-form-urlencoded`.

			req = req.SetFormData(opt.ParseData(opt.Data))
		}

		//setBody: for struct and map data type defaults to 'application/json'
		// SetBody method sets the request body for the request. It supports various realtime needs as easy.
		// We can say its quite handy or powerful. Supported request body data types is `string`,
		// `[]byte`, `struct`, `map`, `slice` and `io.Reader`. Body value can be pointer or non-pointer.
		// Automatic marshalling for JSON and XML content type, if it is `struct`, `map`, or `slice`.
		if opt.Body != nil {
			req = req.SetBody(opt.Body)
		}

		if method == "post" {
			resp, err = req.Post(reqUrl)
			return s.GetResult(resp, err)
		}

		if method == "put" {
			resp, err = req.Put(reqUrl)
			return s.GetResult(resp, err)
		}

		if method == "patch" {
			resp, err = req.Patch(reqUrl)
			return s.GetResult(resp, err)
		}
	case "file":
		b, err := ioutil.ReadFile(opt.FileName)
		if err != nil {
			return &Reply{
				err: errors.New("read file error: " + err.Error()),
			}
		}
		resp, err := client.R().
			SetFileReader(opt.FileParamName, opt.FileName, bytes.NewReader(b)).
			Post(reqUrl)
		return s.GetResult(resp, err)
	default:
	}

	return &Reply{
		err: errors.New("request method not support"),
	}
}

//NewRestyClient new resty client
func NewRestyClient() *resty.Client {
	return resty.New()
}

var reply = &Reply{}

func (s *Service) GetResult(resp *resty.Response, err error) *Reply {
	if err != nil {
		reply.err = err
		return reply
	}
	reply.body = resp.Body()
	if !resp.IsSuccess() {
		if reply.AsString() == "" {
			reply.ApiResponseIsBad = true
			reply.err = errors.New("request failed -> " + " http StatusCode: " + strconv.Itoa(resp.StatusCode()) + " message: " + resp.Status())
			reply.statusCode = resp.StatusCode()
			return reply
		}
	}
	reply.statusCode = resp.StatusCode()
	return reply
}

// Log log for debugging
func (r *Reply) Log() {
	log.Println(reply.Status())
	log.Println(reply.err)

}

// Status return http status code
func (r *Reply) Status() int {
	return r.statusCode
}

// Status return http status code
func (r *Reply) Error() error {
	return r.err
}

// AsString return as body as a string
func (r *Reply) AsString() string {
	return string(r.body)
}

// AsJson return as body as blank interface
func (r *Reply) AsJson() (interface{}, error) {
	var res interface{}
	err := json.Unmarshal(reply.body, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}

// AsJsonNoErr return as body as blank interface and ignore any errors
func (r *Reply) AsJsonNoErr() interface{} {
	var res interface{}
	err := json.Unmarshal(r.body, &res)
	if err != nil {
		return nil
	}
	return res
}

// ToInterface return as body as a json
func (r *Reply) ToInterface(data interface{}) error {
	if len(r.body) > 0 {
		err := json.Unmarshal(r.body, data)
		if err != nil {
			return err
		}
	}
	return nil
}

// ToInterfaceNoErr return as body as a json
func (r *Reply) ToInterfaceNoErr(data interface{}) {
	if len(r.body) > 0 {
		json.Unmarshal(r.body, data)
	}

}
