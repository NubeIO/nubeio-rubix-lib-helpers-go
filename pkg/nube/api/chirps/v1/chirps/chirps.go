package chirps

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest/v1/rest"
)

//admin
//N00BWAN

type ChirpClient struct {
	Rest     *rest.Service
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgID    int    `json:"org_id"`
	Limit    int    `json:"limit"`
}

type Path struct {
	Path string
}

var Paths = struct {
	Login Path
}{
	Login: Path{Path: "/api/internal/login"},
}

var limit = 10
var orgID = 1

// New returns a new instance of the nube common apis
func New(bc *ChirpClient) *ChirpClient {
	return bc
}

func (inst *ChirpClient) builder(method string, logFunc interface{}, path string) *rest.Service {
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

type TokenResponse struct {
	JWT string `json:"jwt"`
}

type Token struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetToken get cs token
func (inst *ChirpClient) GetToken(token *Token) (data *TokenResponse, response *rest.ProxyResponse) {
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(Paths.Login.Path).
		SetBody(token).
		DoRequest()
	response = inst.Rest.BuildResponse(req, data)
	response.GetResponse()
	return
}

// GetDevices get all
func (inst *ChirpClient) GetDevices(token string) (data *Devices, response *rest.ProxyResponse) {
	path := fmt.Sprintf("/api/devices?limit=%d&organizationID=%d", limit, orgID)
	inst.Rest = inst.builder(rest.GET, inst.GetDevices, path)
	inst.Rest.Options.Headers = map[string]interface{}{"Grpc-Metadata-Authorization": token}
	res := inst.Rest.DoRequest()
	response = inst.Rest.BuildResponse(res, &data)
	return
}
