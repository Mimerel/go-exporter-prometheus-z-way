package extractZway

import (
	"strings"
)

func trim(value string) (string) {
	return strings.TrimSpace(value)
}

func (data *Data) validTypes(value string) (bool) {
	for _, x := range data.Conf.DeviceTypes {
		if x == trim(value) {
			return true
		}
	}
	return false
}

func BoolToIntensity(value bool) (float64) {
	if value == true {
		return 255.0
	} else {
		return 0.0
	}
}