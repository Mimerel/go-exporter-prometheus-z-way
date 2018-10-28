package main

import (
	"go-exporter-prometheus-z-way/extract_data"
	"net/http"
)

func main() {
	configuration := extract_data.ReadConfiguration()

	http.HandleFunc("/metrics", func (w http.ResponseWriter, r *http.Request) {
		extract_data.ExtractMetrics(w, r, &configuration)
	})
	http.ListenAndServe(":" + configuration.Port, nil)
}
