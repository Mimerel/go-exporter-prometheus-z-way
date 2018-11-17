package models

type ElementDetails struct {
	Metric string
	Name string
	Room string
	Type string
	Unit string
	Instance string
	Id string
	IdInstance string
	Ignore bool
	Value float64
	Switch string
}

type GlobalDevice struct {
	Name string
	Room string
	Type string
	Switch string
	Instance string
	Id string
	IdInstance string
	Level float64
	Lux float64
	Alarm float64
	Humidity float64
	Temperature float64
	Flood float64
	Amperes float64
	Watts float64
	Volts float64
	Tempered float64
}