package extraction

import (
	"github.com/prometheus/client_golang/prometheus"
	"go-exporter-prometheus-z-way/extract_data/configuration"
)

type Data struct {
	Registry *prometheus.Registry
	Host string
	ZWayPath string
	Source map[string]*Summary
	Metrics map[string]*prometheus.GaugeVec
	Configuration *configuration.MainConfig
}

type Summary struct {
	Metric string
	Value  float64
}
