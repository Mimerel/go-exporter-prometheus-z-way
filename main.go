package main

import (
	"go-exporter-prometheus-z-way/extract_data"
	"net/http"
)

func main() {
	http.HandleFunc("/metrics", func (w http.ResponseWriter, r *http.Request) {
		extract_data.ExtractMetrics(w, r)
	})
	http.ListenAndServe(":2112", nil)
}
