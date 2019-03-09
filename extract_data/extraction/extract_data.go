package extraction

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/extractSystemData"
	"go-exporter-prometheus-z-way/extract_data/extractZway"
	"go-exporter-prometheus-z-way/extract_data/logs"
	"go-exporter-prometheus-z-way/extract_data/models"
	"math"
	"net/http"
	"strings"
)

func ExtractMetrics(w http.ResponseWriter, r *http.Request, conf *configuration.MainConfig) {



	data := new(Data)
	data.Registry = prometheus.NewRegistry()
	data.Metrics = make(map[string]*prometheus.GaugeVec)
	data.Source = make(map[string]*models.ElementDetails)
	data.Configuration = conf

	keys := r.URL.Query()["json"]
	if len(keys) != 0 {
		data.Json = true
	}


	if data.Configuration.Host == "" {
		data.Configuration.Host = "Anonymous"
	}

	runZway := false
	runCpu := false

	for _, v := range data.Configuration.ActivatedModules {
		if v == "zway" {
			runZway = true
		}
		if v == "systemData" {
			runCpu = true
		}
	}

	// Collecting System details
	if runCpu {
		for _, v := range extractSystemData.ExtractTotalCpuUsage(*conf) {
			v.Metric = strings.ToLower(strings.Replace(v.Name, " ", "_", -1))
			data.Source[v.Metric] = &models.ElementDetails{Name: v.Metric, Value: v.Value}
		}
	}
	// Collecting Z-way metrics
	if runZway {
		for _, v := range extractZway.ExtractZWayMetrics(*conf) {
			overrideValues(*conf, &v)
			v.Metric = "zway_" + strings.ToLower(strings.Replace(v.Name, " ", "_", -1)) + "_" + extractZway.Trim(v.Unit) + "_" + v.Instance
			data.Source[v.Metric] = &models.ElementDetails{Metric: v.Metric, Name: strings.Replace(v.Name, " ", "_", -1), Value: v.Value, Room: v.Room, Type: v.Type, Unit: v.Unit, Instance: v.Instance, Switch: v.Switch,
			IdInstance: v.Id + "_" + v.Instance, Id: v.Id, Ignore: v.Ignore}
		}
	}

	if !data.Json {
		// Creating metrics and populating them
		for index, value := range data.Source {
			if value.Type != "" && value.Ignore == false {
				data.Metrics[index] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
					Name: index, Help: index,}, []string{"host", "type", "room", "unit", "name", "instance", "id", "switch"})
				data.Registry.MustRegister(data.Metrics[index])
				data.Metrics[index].WithLabelValues(data.Configuration.Host, value.Type, value.Room, value.Unit, value.Name, value.Instance, value.Id, value.Switch).
					Set(math.Round(value.Value*100) / 100)
			} else if value.Type == "" {
				data.Metrics[index] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
					Name: index, Help: index,}, []string{"host","name"})
				data.Registry.MustRegister(data.Metrics[index])
				data.Metrics[index].WithLabelValues(data.Configuration.Host, value.Name).Set(math.Round(value.Value*100) / 100)
			}
		}
		h := promhttp.HandlerFor(data.Registry, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	} else {
		js, err := json.Marshal(organizedValues(conf, data.Source))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	logs.Info(conf.Logger, conf.Host, fmt.Sprint("New Request for metrics Successful"))
}


func organizedValues(conf *configuration.MainConfig, source map[string]*models.ElementDetails) map[string]*models.GlobalDevice {
	final := make(map[string]*models.GlobalDevice)
	for _, value := range source {
		if value.Ignore == false {
			if final[value.Id] == nil {
				final[value.Id] = new(models.GlobalDevice)
				final[value.Id].Alarm = -9999.0
				final[value.Id].Level = -9999.0
				final[value.Id].Lux = -9999.0
				final[value.Id].Tempered = -9999.0
				final[value.Id].Flood = -9999.0
				final[value.Id].Amperes = -9999.0
				final[value.Id].Volts = -9999.0
				final[value.Id].Watts = -9999.0
				final[value.Id].Humidity = -9999.0
				final[value.Id].Temperature = -9999.0
			}
			final[value.Id].Name = value.Name
			final[value.Id].IdInstance = value.IdInstance
			final[value.Id].Instance = value.Instance
			final[value.Id].Id = value.Id
			final[value.Id].HostIp = conf.ZwayServer
			final[value.Id].Host = conf.Host
			final[value.Id].Room = value.Room
			final[value.Id].Type = value.Type
			switch value.Unit {
			case "degré":
				final[value.Id].Temperature = value.Value
			case "Humidity":
				final[value.Id].Humidity = value.Value
			case "Watt":
				final[value.Id].Watts = value.Value
			case "Volts":
				final[value.Id].Volts = value.Value
			case "Ampères":
				final[value.Id].Amperes = value.Value
			case "Flood":
				final[value.Id].Flood = value.Value
			case "Tempered":
				final[value.Id].Tempered = value.Value
			case "Lux":
				final[value.Id].Lux = value.Value
			case "Level":
				final[value.Id].Level = value.Value
				final[value.Id].Switch = value.Switch
			case "Alarm":
				final[value.Id].Alarm = value.Value
			}
		}
	}
	return final
}

func overrideValues(conf configuration.MainConfig, value *models.ElementDetails) {
	if conf.DeviceConfiguration[value.IdInstance]!= (configuration.DeviceConf{}) {
		overrideElements := conf.DeviceConfiguration[value.IdInstance]

		if overrideElements.Type != "" {
			value.Type = extractZway.Trim(overrideElements.Type)
		}
		if overrideElements.Unit != "" {
			value.Unit = extractZway.Trim(overrideElements.Unit)
		}
		if overrideElements.Name != "" {
			value.Name = extractZway.Trim(overrideElements.Name)
		}
		if overrideElements.Room != "" {
			value.Room = extractZway.Trim(overrideElements.Room)
		}
		if overrideElements.Ignore == true {
			value.Ignore = true
		}
	}
}