package extract_data

import (
	"fmt"
	"github.com/apsdehal/go-logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strings"
)


var log, err = logger.New("test", 1, os.Stdout)

func ExtractMetrics(w http.ResponseWriter, r *http.Request) {
	data := new(Data)
	data.registry = prometheus.NewRegistry()
	data.host = os.Getenv("EXPORTER_SERVER")
	data.metrics = make(map[string]*prometheus.GaugeVec)
	data.source = make(map[string]*summary)
	if data.host == "" {
		data.host = "Anonymous"
	}

	// Collecting System details
	var systemData []systemDetails
	systemData = GetLocalSystemSituation()
	data.ExtractTotalCpuUsage(systemData)

	// creating metrics and populating them
	for index, value := range data.source {
		data.metrics[index] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: value.metric, Help: value.metric, },[]string{"host",})
		data.registry.MustRegister(data.metrics[index])
		data.metrics[index].WithLabelValues(data.host).Add(math.Round(value.value*100)/100)
	}
	h := promhttp.HandlerFor(data.registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}



func (data *Data) ExtractTotalCpuUsage(systemData []systemDetails) {
	data.source["cpu_total"] = &summary{metric:"Cpu_total", value: 0}
	data.source["cpu_exporter"] = &summary{metric:"Cpu_exporter", value: 0}
	data.source["mem_total"] = &summary{metric:"Mem_total", value: 0}
	data.source["mem_exporter"] = &summary{metric:"Mem_exporter", value: 0}
	for _, value := range systemData {
		data.source["cpu_total"].value = data.source["cpu_total"].value + value.cpu
		data.source["mem_total"].value = data.source["mem_total"].value + value.mem

		if strings.Index(value.command, "go-exporter-prometheus-z-way") != -1 {
			(data.source["cpu_exporter"]).value = value.cpu
			(data.source["mem_exporter"]).value = value.mem
		}
	}
	log.NoticeF("cpu used total : %f percent", (data.source["cpu_total"]).value)
	log.NoticeF("cpu used exporter : %f percent", (data.source["cpu_exporter"]).value)
	log.NoticeF("mem used total : %f percent", (data.source["mem_total"]).value)
	log.NoticeF("mem used exporter : %f percent", (data.source["mem_exporter"]).value)
}

func GetLocalSystemSituation() (data []systemDetails) {
	out, err := exec.Command("ps", "aux").Output()
	if err != nil {
		log.Fatalf("Error unable to execute ps command %s", err)
		return nil
	}
	systemInfo := strings.Split(string(out), "\n")
	systemInfo = systemInfo[1 : len(systemInfo)-1]
	for _, line := range systemInfo {
		var element systemDetails
		_, err = fmt.Sscanf(line,
			"%s %d %f %f %s %s %s %s %s %s %999s",
			&element.name,
			&element.pid,
			&element.cpu,
			&element.mem,
			&element.vsz,
			&element.rss,
			&element.tt,
			&element.stat,
			&element.started,
			&element.time,
			&element.command)
		if err != nil {
			log.Errorf("error %s", err)
		}
		data = append(data, element)
	}
	return data
}
