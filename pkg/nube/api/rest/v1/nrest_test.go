package rest

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube_apps"
	"testing"
)

func TestRestL(*testing.T) {

	New(&Service{Url: LocalHost, Port: nube_apps.Services.FlowFramework.Port, Path: "/api/system/ping"}).Request().Log()

}
