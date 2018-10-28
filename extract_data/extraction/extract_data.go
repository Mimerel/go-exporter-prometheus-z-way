package extraction

import (
	"github.com/apsdehal/go-logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/extractSystemData"
	"go-exporter-prometheus-z-way/extract_data/extractZway"
	"math"
	"net/http"
	"os"
)


var log, err = logger.New("test", 1, os.Stdout)

func ExtractMetrics(w http.ResponseWriter, r *http.Request, conf *configuration.MainConfig) {
	data := new(Data)
	data.Registry = prometheus.NewRegistry()
	data.Metrics = make(map[string]*prometheus.GaugeVec)
	data.Source = make(map[string]*Summary)
	data.Configuration = conf

	if data.Configuration.Host == "" {
		data.Configuration.Host = "Anonymous"
	}

	// Collecting System details
	for k, v := range extractSystemData.ExtractTotalCpuUsage(*conf) {
			data.Source[k] = &Summary{v.Metric, v.Value}
	}

	// Collecting Z-way metrics
	for k,v := range extractZway.ExtractZWayMetrics(*conf) {
		data.Source[k] = &Summary{v.Metric, v.Value}
	}

	// Creating metrics and populating them
	for index, value := range data.Source {
		data.Metrics[index] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: value.Metric, Help: value.Metric, },[]string{"host",})
		data.Registry.MustRegister(data.Metrics[index])
		data.Metrics[index].WithLabelValues(data.Configuration.Host).Add(math.Round(value.Value*100)/100)
	}
	h := promhttp.HandlerFor(data.Registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

