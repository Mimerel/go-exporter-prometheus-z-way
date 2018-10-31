package main

import (
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/extraction"
	"net/http"
)

func main() {
	configuration := configuration.ReadConfiguration()

	http.HandleFunc("/metrics", func (w http.ResponseWriter, r *http.Request) {
		extraction.ExtractMetrics(w, r, &configuration)
	})
	http.ListenAndServe(":" + configuration.Port, nil)
}
