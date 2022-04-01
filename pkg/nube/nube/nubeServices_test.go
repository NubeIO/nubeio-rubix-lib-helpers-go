package nube

import (
	"fmt"
	"testing"
)

func TestNubeServices(*testing.T) {

	fmt.Println(Services.BacnetServer.Name)
	fmt.Println(Services.BacnetServer.Proxy)
	fmt.Println(Services.BacnetServer.Port)

}
