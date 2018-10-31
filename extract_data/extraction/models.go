package extraction

import (
	"github.com/prometheus/client_golang/prometheus"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/models"
)

type Data struct {
	Registry *prometheus.Registry
	Host string
	ZWayPath string
	Source map[string]*models.ElementDetails
	Metrics map[string]*prometheus.GaugeVec
	Configuration *configuration.MainConfig
}
