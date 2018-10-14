package system_info

import (
	"fmt"
	"github.com/apsdehal/go-logger"
	"os"
	"os/exec"
	"strings"
)


type systemDetails struct {
	name string
	pid  int
	cpu float64
	mem float64
	vsz string
	rss string
	tt string
	stat string
	started string
	time string
	command string
}

var log, err = logger.New("test", 1, os.Stdout)


func GetLocalSystemSituation() (data []systemDetails) {
	out, err := exec.Command("ps", "aux").Output()
	if err != nil {
		log.Fatalf("Error unable to execute ps command %s", err)
		return nil
	}
	systemInfo := strings.Split(string(out), "\n")
	systemInfo = systemInfo[1:len(systemInfo)-1]
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