package extract_data

import (
	"fmt"
	"os/exec"
	"strings"
)

func (data *Data) ExtractTotalCpuUsage(systemData []SystemDetails) {
	data.source["cpu_total"] = &Summary{metric:"Cpu_total", value: 0}
	data.source["mem_total"] = &Summary{metric: "Mem_total", value: 0}
	for _, services := range data.configuration.FollowedServices {
		data.source["cpu_" + services] = &Summary{metric: "Cpu_" + services, value: 0}
		data.source["mem_" + services] = &Summary{metric: "Mem_" + services, value: 0}
	}
	for _, value := range systemData {
		data.source["cpu_total"].value = data.source["cpu_total"].value + value.cpu
		data.source["mem_total"].value = data.source["mem_total"].value + value.mem

		for _, services := range data.configuration.FollowedServices {

			if strings.Index(strings.ToLower(value.command), strings.ToLower(services)) != -1 {
				(data.source["cpu_" + services]).value = value.cpu
				(data.source["mem_" + services]).value = value.mem
			}
		}

	}
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
