package extractZway

import (
	"encoding/json"
	"fmt"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/models"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (data *Data) GetDataFromZWay() {
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	res, err := client.Get(data.Conf.ZwayServer)
	if err != nil {
		fmt.Println("There was a get site error:", err)
	} else {

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
}

func ExtractZWayMetrics(conf configuration.MainConfig) ([]models.ElementDetails) {
	var data Data
	data.Conf = conf
	data.GetDataFromZWay()
	data.ExtractElements()
	return data.Element
}

func (data *Data) ExtractElements() {
	for _, v := range data.Json.Devices {
		values := strings.Split(v.Data.GivenName.Value, "|")
		if len(values) >= 3 && data.validTypes(values[2]) {
			for _, instanceContent := range v.Instances {
				if instanceContent.CommandClasses.Class50 != (CommandClass50{}) {
					element := new(models.ElementDetails)
					element.Unit = "Watt"
					element.Value = instanceContent.CommandClasses.Class50.Data.Data0.Val.Value
					element.Name = Trim(values[0])
					element.Room = Trim(values[1])
					element.Type = Trim(values[2])
					data.Element = append(data.Element, *element)
				}
				if instanceContent.CommandClasses.Class49 != (CommandClass49{}) {
					if instanceContent.CommandClasses.Class49.Data.Data1 != (CommandClass49DataVal{}) {
						element := new(models.ElementDetails)
						element.Unit = "Degré"
						element.Value = instanceContent.CommandClasses.Class49.Data.Data1.Val.Value
						element.Name = Trim(values[0])
						element.Room = Trim(values[1])
						element.Type = Trim(values[2])
						data.Element = append(data.Element, *element)
					}
					if instanceContent.CommandClasses.Class49.Data.Data5 != (CommandClass49DataVal{}) {
						element := new(models.ElementDetails)
						element.Unit = "Humidité"
						element.Value = instanceContent.CommandClasses.Class49.Data.Data5.Val.Value
						element.Name = Trim(values[0])
						element.Room = Trim(values[1])
						element.Type = Trim(values[2])
						data.Element = append(data.Element, *element)
					}
				}
				if instanceContent.CommandClasses.Class48 != (CommandClass48{}) {
					if instanceContent.CommandClasses.Class48.Data.Data1 != (CommandClass48DataValBool{}) {
						element := new(models.ElementDetails)
						element.Unit = "Alarm"
						element.Value = BoolToIntensity(instanceContent.CommandClasses.Class48.Data.Data1.Level.Value)
						element.Name = Trim(values[0])
						element.Room = Trim(values[1])
						element.Type = Trim(values[2])
						data.Element = append(data.Element, *element)
					}
					if instanceContent.CommandClasses.Class48.Data.Data6 != (CommandClass48DataValBool{}) {
						element := new(models.ElementDetails)
						element.Unit = "Flood"
						element.Value = BoolToIntensity(instanceContent.CommandClasses.Class48.Data.Data6.Level.Value)
						element.Name = Trim(values[0])
						element.Room = Trim(values[1])
						element.Type = Trim(values[2])
						data.Element = append(data.Element, *element)
					}
					if instanceContent.CommandClasses.Class48.Data.Data8 != (CommandClass48DataValBool{}) {
						element := new(models.ElementDetails)
						element.Unit = "Tempered"
						element.Value = BoolToIntensity(instanceContent.CommandClasses.Class48.Data.Data8.Level.Value)
						element.Name = Trim(values[0])
						element.Room = Trim(values[1])
						element.Type = Trim(values[2])
						data.Element = append(data.Element, *element)
					}
				}
				if instanceContent.CommandClasses.Class37 != (CommandClass37{}) {
					element := new(models.ElementDetails)
					element.Unit = "Level"
					element.Value = BoolToIntensity(instanceContent.CommandClasses.Class37.Data.Level.Value)
					element.Name = Trim(values[0])
					element.Room = Trim(values[1])
					element.Type = Trim(values[2])
					data.Element = append(data.Element, *element)
				}
			}
		}
	}
	// fmt.Printf("Devices found : %+v", data.Element)
}
