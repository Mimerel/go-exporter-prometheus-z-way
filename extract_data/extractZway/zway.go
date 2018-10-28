package extractZway

import (
	"encoding/json"
	"fmt"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"io/ioutil"
	"net/http"
	"strings"
)

func (data *Data) GetDataFromZWay() {
	res, err := http.Get(data.Conf.ZwayServer)
	if err != nil {
		fmt.Println("There was a get site error:", err)
	}

	temp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("There was a read error:", err)
	}

	res.Body.Close()

	err = json.Unmarshal(temp, &data.Json)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
}

func ExtractZWayMetrics(conf configuration.MainConfig) (map[string]*Summary) {
	var data Data
	data.Conf = conf
	data.GetDataFromZWay()
	data.ExtractElements()
	return nil
}

func (data *Data) ExtractElements() {
	for _, v := range data.Json.Devices {
		values := strings.Split(v.Data.GivenName.Value, "|")
		if len(values) >= 3 && data.validTypes(values[2]) {
			for _, instanceContent := range v.Instances {
				if instanceContent.CommandClasses.Class50 != (CommandClass50{}) {
					element := new(ElementDetails)
					element.Unit = "Watt"
					element.value = instanceContent.CommandClasses.Class50.Data.Data0.Val.Value
					element.Name = trim(values[0])
					element.Room = trim(values[1])
					element.Type = trim(values[2])
					data.Element = append(data.Element, *element)
				}
				if instanceContent.CommandClasses.Class49 != (CommandClass49{}) {
					if instanceContent.CommandClasses.Class49.Data.Data1 != (CommandClass49DataVal{}) {
						element := new(ElementDetails)
						element.Unit = "Degré"
						element.value = instanceContent.CommandClasses.Class49.Data.Data1.Val.Value
						element.Name = trim(values[0])
						element.Room = trim(values[1])
						element.Type = trim(values[2])
						data.Element = append(data.Element, *element)
					}
					if instanceContent.CommandClasses.Class49.Data.Data5 != (CommandClass49DataVal{}) {
						element := new(ElementDetails)
						element.Unit = "Humidité"
						element.value = instanceContent.CommandClasses.Class49.Data.Data5.Val.Value
						element.Name = trim(values[0])
						element.Room = trim(values[1])
						element.Type = trim(values[2])
						data.Element = append(data.Element, *element)
					}
				}
				if instanceContent.CommandClasses.Class48 != (CommandClass48{}) {
					if instanceContent.CommandClasses.Class48.Data.Data1 != (CommandClass48DataValBool{}) {
						element := new(ElementDetails)
						element.Unit = "Alarm"
						fmt.Println("Alarm : %+v", instanceContent.CommandClasses.Class48.Data.Data1.Level.Value)
						element.value = BoolToIntensity(instanceContent.CommandClasses.Class48.Data.Data1.Level.Value)
						element.Name = trim(values[0])
						element.Room = trim(values[1])
						element.Type = trim(values[2])
						data.Element = append(data.Element, *element)
					}
					if instanceContent.CommandClasses.Class48.Data.Data6 != (CommandClass48DataValBool{}) {
						element := new(ElementDetails)
						element.Unit = "Flood"
						fmt.Println("Flood : %+v", instanceContent.CommandClasses.Class48.Data.Data6.Level.Value)
						element.value = BoolToIntensity(instanceContent.CommandClasses.Class48.Data.Data6.Level.Value)
						element.Name = trim(values[0])
						element.Room = trim(values[1])
						element.Type = trim(values[2])
						data.Element = append(data.Element, *element)
					}
					if instanceContent.CommandClasses.Class48.Data.Data8 != (CommandClass48DataValBool{}) {
						element := new(ElementDetails)
						element.Unit = "Tempered"
						fmt.Println("Tempered : %+v", instanceContent.CommandClasses.Class48.Data.Data8.Level.Value)
						element.value = BoolToIntensity(instanceContent.CommandClasses.Class48.Data.Data8.Level.Value)
						element.Name = trim(values[0])
						element.Room = trim(values[1])
						element.Type = trim(values[2])
						data.Element = append(data.Element, *element)
					}
			}
				if instanceContent.CommandClasses.Class37 != (CommandClass37{}) {
					element := new(ElementDetails)
					element.Unit = "Level"
					element.value = BoolToIntensity(instanceContent.CommandClasses.Class37.Data.Level.Value)
					element.Name = trim(values[0])
					element.Room = trim(values[1])
					element.Type = trim(values[2])
					data.Element = append(data.Element, *element)
				}
			}
		}
	}
	fmt.Printf("Devices found : %+v", data.Element)
}
