package chirps

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/rest/v1/rest"
	"testing"
)

func TestChirpsPing(*testing.T) {

	restService := &rest.Service{}
	restService.Url = "192.168.15.13"
	restService.Port = 8080
	restOptions := &rest.Options{}
	restService.Options = restOptions
	restService = rest.New(restService)

	nubeProxy := &rest.NubeProxy{}
	nubeProxy.UseRubixProxy = false
	nubeProxy.RubixUsername = "admin"
	nubeProxy.RubixPassword = "N00BWires"
	restService.NubeProxy = nubeProxy

	client := New(&ChirpClient{Rest: restService})

	token, res := client.GetToken()
	fmt.Println("token", res.StatusCode)
	fmt.Println("token", token)

	//devices, res := client.GetDevices(token.JWT)
	//fmt.Println("devices", res.StatusCode)
	//fmt.Println("devices", *devices)

}
