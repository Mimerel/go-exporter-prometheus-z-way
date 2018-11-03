package main

import (
	"fmt"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/extraction"
	"go-exporter-prometheus-z-way/extract_data/logs"
	"net/http"
)

func main() {
	configuration := configuration.ReadConfiguration()
	logs.Info(configuration.Logger, configuration.Host, fmt.Sprint("Application started"))

	http.HandleFunc("/metrics", func (w http.ResponseWriter, r *http.Request) {
		extraction.ExtractMetrics(w, r, &configuration)
	})
	http.ListenAndServe(":" + configuration.Port, nil)
}
