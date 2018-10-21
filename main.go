package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"go-exporter-prometheus-z-way/system-info"
)

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})


func init() {
	prometheus.MustRegister(cpuTemp)
}



func main() {
	system_info.AnalyseSystemInfo()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
