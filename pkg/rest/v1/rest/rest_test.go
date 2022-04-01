package rest

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"testing"
)

func TestRestL(*testing.T) {

	New(&Service{Url: LocalHost, Port: nube.Services.BacnetServer.Port, Path: nube.Services.BacnetServer.PingPath}).Request().Log()

}
