package extract_data

import "github.com/prometheus/client_golang/prometheus"

type SystemDetails struct {
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

type Summary struct {
	metric string
	value  float64
}


type Data struct {
	registry *prometheus.Registry
	host string
	zWayPath string
	source map[string]*Summary
	metrics map[string]*prometheus.GaugeVec
	configuration *MainConfig
}

type MainConfig struct{
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	FollowedServices map[string]string `yaml:"followedServices,omitempty"`
	ActivatedModules []string `yaml:"activatedModules,omitempty"`
}

