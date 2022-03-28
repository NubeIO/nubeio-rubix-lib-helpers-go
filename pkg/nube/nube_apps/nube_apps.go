package nube_apps

type Service struct {
	Name  string //bacnet-server
	Proxy string //rubix-service proxy
	Port  int    //1717
}

var Services = struct {
	Mosquitto     Service
	BacnetServer  Service
	FlowFramework Service
	RubixService  Service
	RubixBios     Service
}{
	Mosquitto:     Service{Name: "mosquitto", Proxy: "", Port: 1883},
	BacnetServer:  Service{Name: "bacnet-server", Proxy: "bacnet", Port: 1717},
	FlowFramework: Service{Name: "flow-framework", Proxy: "ff", Port: 1660},
	RubixService:  Service{Name: "rubix-service", Proxy: "", Port: 1616},
	RubixBios:     Service{Name: "rubix-bios", Proxy: "", Port: 1615},
}
