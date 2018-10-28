package extract_data

type Controller  struct {
	UpdateTime int `json:"updateTime,omitempty"`
	Devices    []Device `json:"devices,omitempty"`
}

type Device struct {
	Instances []DeviceInstances `json:"instances,omitempty"`
	Data DeviceData `json:"data,omitempty"`
}

type DeviceData struct {
	GivenName DeviceValuesString `json:"givenName,omitempty"`
	LastReceived DeviceValuesNumber `json:"lastReceived,omitempty"`
}

type DeviceInstances struct {
	CommandClasses []DeviceInstancesCommandClass `json:"commandClasses,omitempty"`
}

type DeviceInstancesCommandClass struct {
	Name string `json:"name,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type DeviceValuesString struct {
	Value          string `json:"value,omitempty"`
	Type           string `json:"type,omitempty"`
	InvalidateTime int `json:"invalideTime,omitempty"`
	UpdateTime     int `json:"updateTime,omitempty"`
}

type DeviceValuesNumber struct {
	Value          float64 `json:"value,omitempty"`
	Type           string `json:"type,omitempty"`
	InvalidateTime int `json:"invalideTime,omitempty"`
	UpdateTime     int `json:"updateTime,omitempty"`
}

