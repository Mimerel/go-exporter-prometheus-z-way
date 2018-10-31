package extraction

import (
	"github.com/apsdehal/go-logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/extractSystemData"
	"go-exporter-prometheus-z-way/extract_data/extractZway"
	"go-exporter-prometheus-z-way/extract_data/models"
	"math"
	"net/http"
	"os"
	"strings"
)

var log, err = logger.New("test", 1, os.Stdout)

func ExtractMetrics(w http.ResponseWriter, r *http.Request, conf *configuration.MainConfig) {
	data := new(Data)
	data.Registry = prometheus.NewRegistry()
	data.Metrics = make(map[string]*prometheus.GaugeVec)
	data.Source = make(map[string]*models.ElementDetails)
	data.Configuration = conf

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
			v.Metric = "zway_" + strings.ToLower(strings.Replace(v.Name, " ", "_", -1)) + "_" + extractZway.Trim(v.Unit)
			data.Source[v.Metric] = &models.ElementDetails{Name: v.Metric, Value: v.Value, Room: v.Room, Type: v.Type, Unit: v.Unit}
		}
	}
	// Creating metrics and populating them
	for index, value := range data.Source {
		if value.Type != "" {
			data.Metrics[index] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: value.Name, Help: value.Name,}, []string{"host", "type", "room", "unit", "name"})
			data.Registry.MustRegister(data.Metrics[index])
			data.Metrics[index].WithLabelValues(data.Configuration.Host, value.Type, value.Room, value.Unit, value.Name).Add(math.Round(value.Value*100) / 100)

		} else {
			data.Metrics[index] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: value.Name, Help: value.Name,}, []string{"host",})
			data.Registry.MustRegister(data.Metrics[index])
			data.Metrics[index].WithLabelValues(data.Configuration.Host).Add(math.Round(value.Value*100) / 100)
		}
	}

	h := promhttp.HandlerFor(data.Registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}
