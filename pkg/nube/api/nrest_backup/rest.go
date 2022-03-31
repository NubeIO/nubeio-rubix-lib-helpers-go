package nrest_backup

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
	"reflect"
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

var defaultTimeout = 3 * time.Second

var defaultMaxRetries = 100

type ReqOpt struct {
	Timeout time.Duration

	RetryCount       int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration

	Params         map[string]interface{}
	SetQueryString string
	Data           map[string]interface{}
	Headers        map[string]interface{}

	Cookies        map[string]interface{}
	CookiePath     string
	CookieDomain   string
	CookieMaxAge   int
	CookieHttpOnly bool

	Json          interface{}
	FileName      string
	FileParamName string
}

type Reply struct {
	IsError           bool
	Err               error
	Body              []byte
	BodyJson          interface{}
	StatusCode        int
	ApiResponseIsBad  bool //is true response is a 404
	ApiResponseIsJSON bool
	ApiResponseLength int
}

type ApiStdRes struct {
	Code    int
	Message string
	Data    interface{}
}

type Service struct {
	BaseUri         string //0.0.0.0 or nube-io.com
	Proxy           string
	EnableKeepAlive bool
	ReqType         *ReqType
	ReqOpt          *ReqOpt
}

type ReqType struct {
	BaseUri string //0.0.0.0 or nube-io.com
	Port    int    // 80 or 443
	HTTPS   bool
	Path    string //  /api/points
	Method  string
	Debug   bool
	Service string //as in bacnet-server
	LogPath string //in the log message show path or extra message
	LogFunc string
}

func errorMsg(appName string, msg interface{}, e error) (err interface{}) {
	if e != nil && msg != "" {
		err = fmt.Errorf("%s: error:%w  msg:%s", appName, e, msg)
		return
	}
	if e != nil {
		err = fmt.Errorf("%s: error:%w", appName, e)
		return
	}
	if msg != "" {
		err = fmt.Errorf("%s: msg:%s", appName, msg)
		return
	}
	return
}

func isJSON(str string) bool {
	return json.Unmarshal([]byte(str), &json.RawMessage{}) == nil
}

func getJSONLen(str interface{}) (length int) {
	length = reflect.ValueOf(str).Len()
	return
}

// StatusCode2xx method returns true if HTTP status `code >= 200 and <= 299` otherwise false.
func StatusCode2xx(statusCode int) bool {
	return statusCode > 199 && statusCode < 300
}

// StatusCode3xx method returns true if HTTP status `code >= 299 and <= 399` otherwise false.
func StatusCode3xx(statusCode int) bool {
	return statusCode > 299 && statusCode < 399
}

// StatusCode4xx method returns true if HTTP status `code >= 399 and <= 499` otherwise false.
func StatusCode4xx(statusCode int) bool {
	return statusCode > 399 && statusCode < 499
}

// StatusCode5xx method returns true if HTTP status `code >= 499 and <= 599` otherwise false.
func StatusCode5xx(statusCode int) bool {
	return statusCode > 499 && statusCode < 599
}

// StatusCodesAllBad any status for 3xx, 4xx and 5xx
func StatusCodesAllBad(statusCode int) (ok bool) {
	if StatusCode3xx(statusCode) {
		ok = true
	}
	if StatusCode4xx(statusCode) {
		ok = true
	}
	if StatusCode5xx(statusCode) {
		ok = true
	}
	return
}

func DoHTTPReq(r *ReqType, opt *ReqOpt) (response *Reply) {
	host := fmt.Sprintf("http://%s:%d", r.BaseUri, r.Port)
	if r.HTTPS {
		host = fmt.Sprintf("https://%s:%d", r.BaseUri, r.Port)
	}
	s := &Service{
		BaseUri: host,
	}
	if r.Method == "" {
		r.Method = GET
	}
	if r.LogPath == "" {
		r.LogPath = "nube.helpers.nrest"
	}
	response = s.Do(r.Method, r.Path, opt)
	statusCode := response.StatusCode
	logPath := fmt.Sprintf("%s.%s() method: %s host: %s statusCode:%d", r.LogPath, r.LogFunc, strings.ToUpper(r.Method), host+r.Path, statusCode)
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

func (ReqOpt) ParseData(d map[string]interface{}) map[string]string {
	dLen := len(d)
	if dLen == 0 {
		return nil
	}
	data := make(map[string]string, dLen)
	for k, v := range d {
		if val, ok := v.(string); ok {
			data[k] = val
		} else {
			data[k] = fmt.Sprintf("%v", v)
		}
	}
	return data
}

// Do request
// method string  get,post,put,patch,delete,head
// uri    string  BaseUri  /api/whatever
// opt 	  *ReqOpt
func (s *Service) Do(method string, reqUrl string, opt *ReqOpt) *Reply {
	if method == "" {
		return &Reply{
			Err: errors.New("request method is empty"),
		}
	}
	if reqUrl == "" {
		return &Reply{
			Err: errors.New("request url is empty"),
		}
	}
	if opt == nil {
		opt = &ReqOpt{}
	}
	if s.BaseUri != "" {
		reqUrl = strings.TrimRight(s.BaseUri, "/") + reqUrl
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
		if opt.Json != nil {
			req = req.SetBody(opt.Json)
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
				Err: errors.New("read file error: " + err.Error()),
			}
		}
		resp, err := client.R().
			SetFileReader(opt.FileParamName, opt.FileName, bytes.NewReader(b)).
			Post(reqUrl)
		return s.GetResult(resp, err)
	default:
	}

	return &Reply{
		Err: errors.New("request method not support"),
	}
}

//NewRestyClient new resty client
func NewRestyClient() *resty.Client {
	return resty.New()
}

func (s *Service) GetResult(resp *resty.Response, err error) *Reply {
	res := &Reply{}
	if err != nil {
		res.Err = err
		return res
	}
	res.Body = resp.Body()
	if !resp.IsSuccess() {
		if res.AsString() == "" {
			res.ApiResponseIsBad = true
			res.Err = errors.New("request failed -> " + " http StatusCode: " + strconv.Itoa(resp.StatusCode()) + " message: " + resp.Status())
			res.StatusCode = resp.StatusCode()
			return res
		}
	}
	res.StatusCode = resp.StatusCode()
	return res
}

// Status return http status code
func (r *Reply) Status() int {
	return r.StatusCode
}

// AsString return as body as a string
func (r *Reply) AsString() string {
	return string(r.Body)
}

// AsJson return as body as blank interface
func (r *Reply) AsJson() (interface{}, error) {
	var res interface{}
	err := json.Unmarshal(r.Body, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}

// AsJsonNoErr return as body as blank interface and ignore any errors
func (r *Reply) AsJsonNoErr() interface{} {
	var res interface{}
	err := json.Unmarshal(r.Body, &res)
	if err != nil {
		return nil
	}
	return res
}

// ToInterface return as body as a json
func (r *Reply) ToInterface(data interface{}) error {
	if len(r.Body) > 0 {
		err := json.Unmarshal(r.Body, data)
		if err != nil {
			return err
		}
	}
	return nil
}

// ToInterfaceNoErr return as body as a json
func (r *Reply) ToInterfaceNoErr(data interface{}) {
	if len(r.Body) > 0 {
		json.Unmarshal(r.Body, data)
	}

}
