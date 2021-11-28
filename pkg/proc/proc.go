package proc

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

func GetHostName() (string, error) {
	info, error := host.Info()
	return info.Hostname, error
}
func GetUpTime() (uint64, error) {
	info, error := host.Info()
	return info.Uptime, error
}

func GetLoad() (map[string]float64, error) {
	load, error := load.Avg()
	loadMap := make(map[string]float64)
	loadMap["1m"] = load.Load1
	loadMap["5m"] = load.Load5
	loadMap["15m"] = load.Load15
	return loadMap, error
}
