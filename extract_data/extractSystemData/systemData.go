package extractSystemData

import (
	"fmt"
	"github.com/prometheus/common/log"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"os/exec"
	"strings"
)

func ExtractTotalCpuUsage(systemData []SystemDetails, conf configuration.MainConfig) (map[string]*Summary) {
	data := make(map[string]*Summary)
	data["cpu_total"] = &Summary{ "Cpu_total",  0}
	data["mem_total"] = &Summary{ "Mem_total", 0}
	for _, services := range conf.FollowedServices {
		data["cpu_" + services] = &Summary{ "Cpu_" + services,  0}
		data["mem_" + services] = &Summary{"Mem_" + services, 0}
	}
	for _, value := range systemData {
		(data["cpu_total"]).Value = data["cpu_total"].Value + value.cpu
		(data["mem_total"]).Value = data["mem_total"].Value + value.mem

		for _, services := range conf.FollowedServices {

			if strings.Index(strings.ToLower(value.command), strings.ToLower(services)) != -1 {
				(data["cpu_" + services]).Value = value.cpu
				(data["mem_" + services]).Value = value.mem
			}
		}

	}
	return data
}

func GetLocalSystemSituation() (data []SystemDetails) {
	out, err := exec.Command("ps", "aux").Output()
	if err != nil {
		log.Fatalf("Error unable to execute ps command %s", err)
		return nil
	}
	systemInfo := strings.Split(string(out), "\n")
	systemInfo = systemInfo[1 : len(systemInfo)-1]
	for _, line := range systemInfo {
		var element SystemDetails
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
