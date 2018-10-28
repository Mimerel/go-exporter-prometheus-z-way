package extract_data

import (
	"github.com/apsdehal/go-logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
	"net/http"
	"os"
)


var log, err = logger.New("test", 1, os.Stdout)

func ExtractMetrics(w http.ResponseWriter, r *http.Request, configuration *MainConfig) {
	data := new(Data)
	data.registry = prometheus.NewRegistry()
	data.metrics = make(map[string]*prometheus.GaugeVec)
	data.source = make(map[string]*Summary)
	data.configuration = configuration

	if data.configuration.Host == "" {
		data.configuration.Host = "Anonymous"
	}

	// Collecting System details
	var systemData []SystemDetails
	systemData = GetLocalSystemSituation()
	data.ExtractTotalCpuUsage(systemData)

	// Creating metrics and populating them
	for index, value := range data.source {
		data.metrics[index] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: value.metric, Help: value.metric, },[]string{"host",})
		data.registry.MustRegister(data.metrics[index])
		data.metrics[index].WithLabelValues(data.configuration.Host).Add(math.Round(value.value*100)/100)
	}
	h := promhttp.HandlerFor(data.registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

