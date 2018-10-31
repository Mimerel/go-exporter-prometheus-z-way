package extractZway

import (
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/models"
)

type Controller  struct {
	UpdateTime int `json:"updateTime,omitempty"`
	Devices    map[string]Device `json:"devices"`
}

type Device struct {
	Instances map[string]DeviceInstances `json:"instances,omitempty"`
	Data DeviceData `json:"data,omitempty"`
}

type DeviceData struct {
	GivenName DeviceValuesString `json:"givenName,omitempty"`
	LastReceived DeviceValuesNumber `json:"lastReceived,omitempty"`
}

type DeviceInstances struct {
	CommandClasses DeviceInstancesCommandClass `json:"commandClasses,omitempty"`
}

type DeviceInstancesCommandClass struct {
	Class37 CommandClass37 `json:"37,omitempty"`
	Class48 CommandClass48 `json:"48,omitempty"`
	Class49 CommandClass49 `json:"49,omitempty"`
	Class50 CommandClass50 `json:"50,omitempty"`
}

type CommandClass48 struct {
	Name string `json:"name,omitempty"`
	Data CommandClass48Data `json:"data,omitempty"`
}

type CommandClass48Data struct {
	Data1 CommandClass48DataValBool `json:"1,omitempty"`
	Data6 CommandClass48DataValBool `json:"6,omitempty"`
	Data8 CommandClass48DataValBool `json:"8,omitempty"`
}

type CommandClass48DataValBool struct {
	Type string `json:"type,omitempty"`
	Level CommandClassBoolValues `json:"level,omitempty"`
}

type CommandClass49 struct {
	Name string `json:"name,omitempty"`
	Data CommandClass49Data `json:"data,omitempty"`
}

type CommandClass49Data struct {
	Data1 CommandClass49DataVal `json:"1,omitempty"`
	Data5 CommandClass49DataVal `json:"5,omitempty"`
}

type CommandClass49DataVal struct {
	Type string `json:"type,omitempty"`
	Val CommandClassFloatValues `json:"val,omitempty"`
}

type CommandClass50 struct {
	Name string `json:"name,omitempty"`
	Data CommandClass50Data `json:"data,omitempty"`
}

type CommandClass50Data struct {
	Data2 CommandClass50DataVal `json:"2,omitempty"`
}

type CommandClass50DataVal struct {
	Val CommandClassFloatValues `json:"val,omitempty"`
}

type CommandClass37 struct {
	Name string `json:"name,omitempty"`
	Data CommandClass37Data `json:"data,omitempty"`
}

type CommandClass37Data struct {
	Level CommandClassBoolValues `json:"level,omitempty"`
}

type CommandClassBoolValues struct {
	Value bool `json:"value,omitempty"`
}

type CommandClassFloatValues struct {
	Value float64 `json:"value,omitempty"`
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

type Data struct {
	Json Controller
	Conf configuration.MainConfig
	Element []models.ElementDetails
}
