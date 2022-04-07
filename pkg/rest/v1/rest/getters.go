package rest

/*
PROXY Getters
*/

func (res *ProxyResponse) GetResponse() *response {
	r := res.Response
	return &r
}

func (res *ProxyResponse) GetError() error {
	return res.err
}

func (res *ProxyResponse) GetStatusCode() int {
	return res.Response.StatusCode
}
