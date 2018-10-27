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

type summaryFinal struct {
	cpu_total    summary
	mem_total    summary
	cpu_exporter summary
	mem_exporter summary
}

var log, err = logger.New("test", 1, os.Stdout)

func ExtractMetrics(w http.ResponseWriter, r *http.Request) {
	registry := prometheus.NewRegistry()
	var concernedServer = os.Getenv("EXPORTER_SERVER")
	if concernedServer == "" {
		concernedServer = "Anonymous"
	}
	var final summaryFinal
	var data []systemDetails
	data = GetLocalSystemSituation()
	final.ExtractTotalCpuUsage(data)
	cpuTotal := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "CPU_Total",
		Help: "The total amount of CPU Used",
	},[]string{"host",})
	cpuExporter := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "CPU_Exporter",
		Help: "The Exporter amount of CPU Used",
	},[]string{"host",})
	memTotal := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "MEM_Total",
		Help: "The total amount of MEMORY Used",
	},[]string{"host",})
	memExporter := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "MEM_exporter",
		Help: "The Exporter amount of MEMORY Used",
	},[]string{"host",})
	registry.MustRegister(cpuTotal)
	registry.MustRegister(cpuExporter)
	registry.MustRegister(memTotal)
	registry.MustRegister(memExporter)
	cpuTotal.WithLabelValues(concernedServer).Add(math.Round(final.cpu_total.value*100)/100)
	memTotal.WithLabelValues(concernedServer).Add(math.Round(final.mem_total.value*100)/100)
	cpuExporter.WithLabelValues(concernedServer).Add(math.Round(final.cpu_exporter.value*100)/100)
	memExporter.WithLabelValues(concernedServer).Add(math.Round(final.mem_exporter.value*100)/100)
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}



func (final *summaryFinal) ExtractTotalCpuUsage(data []systemDetails) {
	for _, value := range data {
		final.cpu_total.value = final.cpu_total.value + value.cpu
		final.mem_total.value = final.mem_total.value + value.mem

		if strings.Index(value.command, "go-exporter-prometheus-z-way") != -1 {
			final.cpu_exporter.value = value.cpu
			final.mem_exporter.value = value.mem
		}
	}
	log.NoticeF("cpu used total : %f percent", final.cpu_total.value)
	log.NoticeF("cpu used exporter : %f percent", final.cpu_exporter.value)
	log.NoticeF("mem used total : %f percent", final.mem_total.value)
	log.NoticeF("mem used exporter : %f percent", final.mem_exporter.value)
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
