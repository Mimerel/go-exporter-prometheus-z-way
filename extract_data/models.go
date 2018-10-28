package extract_data

import "github.com/prometheus/client_golang/prometheus"

type systemDetails struct {
	name    string
	pid     int
	cpu     float64
	mem     float64
	vsz     string
	rss     string
	tt      string
	stat    string
	started string
	time    string
	command string
}

type summary struct {
	metric string
	value  float64
}


type Data struct {
	registry *prometheus.Registry
	host string
	zWayPath string
	source map[string]*summary
	metrics map[string]*prometheus.GaugeVec
}