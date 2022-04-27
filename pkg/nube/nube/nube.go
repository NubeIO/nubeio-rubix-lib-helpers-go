package nube

type service struct {
	Name        string //bacnet-server
	Proxy       string //rubix-service proxy
	Port        int    //1717
	PingPath    string //each service ping path
	ServiceName string //systemctl name
}

var Services = struct {
	Mosquitto     service
	BacnetServer  service
	FlowFramework service
	RubixService  service
	RubixBios     service
	PigPIO        service
}{
	Mosquitto:     service{Name: "mosquitto", Proxy: "", Port: 1883, ServiceName: "mosquitto"},
	BacnetServer:  service{Name: "bacnet-server", Proxy: "bacnet", Port: 1717, PingPath: "/api/system/ping"},
	FlowFramework: service{Name: "flow-framework", Proxy: "ff", Port: 1660, PingPath: "/api/system/ping"},
	RubixService:  service{Name: "rubix-service", Proxy: "", Port: 1616},
	RubixBios:     service{Name: "rubix-bios", Proxy: "", Port: 1615},
	PigPIO:        service{Name: "pigpio", Proxy: "", Port: 8888, ServiceName: "pigpiod"},
}
