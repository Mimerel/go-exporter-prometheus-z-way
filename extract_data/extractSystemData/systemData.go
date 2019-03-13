package extractSystemData

import (
	"fmt"
	"go-exporter-prometheus-z-way/extract_data/configuration"
	"go-exporter-prometheus-z-way/extract_data/models"
	"os/exec"
	"strings"
)

func updateValue(data *[]models.ElementDetails, name string, value float64) {
	for key, v := range *data {
		if v.Name == name {
			(*data)[key].Value = v.Value + value
		}
	}
}

func ExtractTotalCpuUsage(conf configuration.MainConfig) ([]models.ElementDetails) {
	systemData := GetLocalSystemSituation(conf)
	var data []models.ElementDetails
	data = append(data, models.ElementDetails{ Name: "Cpu_total",  Value:0})
	data = append(data, models.ElementDetails{ Name: "Mem_total", Value:0})
	for _, services := range conf.FollowedServices {
		data = append(data, models.ElementDetails{ Name:"Cpu_" + services,  Value: 0})
		data = append(data, models.ElementDetails{Name:"Mem_" + services, Value: 0})
	}
	for _, value := range systemData {
		data[0].Value = data[0].Value + value.cpu
		data[1].Value = data[1].Value + value.mem
		for _, services := range conf.FollowedServices {
			//conf.Logger.Info("value : %s ||| ", value.command, services )
			if strings.Index(strings.ToLower(value.command), strings.ToLower(services)) != -1 {
				conf.Logger.Info("value found %s, %s, %s", value, value.cpu, value.mem)
				updateValue(&data ,"Cpu_" + services, value.cpu )
				updateValue(&data ,"Mem_" + services, value.mem )
			}
		}

	}
	return data
}

func GetLocalSystemSituation(conf configuration.MainConfig) (data []SystemDetails) {
	out, err := exec.Command("ps", "aux").Output()
	if err != nil {
		conf.Logger.Error("Error unable to execute ps command %s", err)
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
			conf.Logger.Error("error decryting ps - aux elements : %s", err)
		}
		data = append(data, element)
	}
	return data
}
