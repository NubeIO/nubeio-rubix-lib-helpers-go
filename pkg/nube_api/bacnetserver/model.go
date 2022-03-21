package nube_api_bacnetserver

type ServerPing struct {
	Version        string `json:"version"`
	UpTimeDate     string `json:"up_time_date"`
	UpMin          string `json:"up_min"`
	UpHour         string `json:"up_hour"`
	DeploymentMode string `json:"deployment_mode"`
	Bacnet         struct {
		Enabled                    bool    `json:"enabled"`
		Status                     bool    `json:"status"`
		UsePreSetEthernetInterface bool    `json:"use_pre_set_ethernet_interface"`
		PreSetEthernetInterface    string  `json:"pre_set_ethernet_interface"`
		DefaultPointCov            float32 `json:"default_point_cov"`
	} `json:"bacnet"`
	Mqtt struct {
		Enabled bool `json:"enabled"`
		Status  bool `json:"status"`
	} `json:"mqtt"`
}

type Server struct {
	Ip                string `json:"ip"`
	Port              int    `json:"port"`
	DeviceId          string `json:"device_id"`
	LocalObjName      string `json:"local_obj_name"`
	ModelName         string `json:"model_name"`
	VendorId          string `json:"vendor_id"`
	VendorName        string `json:"vendor_name"`
	EnableIpByNicName bool   `json:"enable_ip_by_nic_name"`
	IpByNicName       string `json:"ip_by_nic_name"` //eth0
}

type BacnetPoint struct {
	AddressUUID          string    `json:"uuid,omitempty"`
	ObjectType           string    `json:"object_type,omitempty"`
	ObjectName           string    `json:"object_name,omitempty"`
	Address              int       `json:"address,omitempty"`
	RelinquishDefault    float64   `json:"relinquish_default"`
	EventState           string    `json:"event_state,omitempty"`
	Units                string    `json:"units,omitempty"`
	Description          string    `json:"description,omitempty"`
	Enable               bool      `json:"enable,omitempty"`
	Fault                bool      `json:"fault,omitempty"`
	DataRound            float64   `json:"data_round,omitempty"`
	DataOffset           float64   `json:"data_offset,omitempty"`
	UseNextAvailableAddr bool      `json:"use_next_available_address,omitempty"`
	COV                  float64   `json:"cov,omitempty"`
	Priority             *Priority `json:"priority_array_write,omitempty"`
}

type Priority struct {
	PointUUID string   `json:"point_uuid,omitempty"`
	P1        *float64 `json:"_1"`
	P2        *float64 `json:"_2"`
	P3        *float64 `json:"_3"`
	P4        *float64 `json:"_4"`
	P5        *float64 `json:"_5"`
	P6        *float64 `json:"_6"`
	P7        *float64 `json:"_7"`
	P8        *float64 `json:"_8"`
	P9        *float64 `json:"_9"`
	P10       *float64 `json:"_10"`
	P11       *float64 `json:"_11"`
	P12       *float64 `json:"_12"`
	P13       *float64 `json:"_13"`
	P14       *float64 `json:"_14"`
	P15       *float64 `json:"_15"`
	P16       *float64 `json:"_16"`
}
