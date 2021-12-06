package proc

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

var (
	LoadTags = []string{"1m", "5m", "15m"}
)

func Load() ([]float64, error) {
	f, err := os.Open("/proc/loadavg")
	if err != nil {
		return nil, err
	}

	var (
		load  float64
		loads []float64
	)
	for i := 0; i < 3; i++ {
		_, err = fmt.Fscan(f, &load)
		if err != nil {
			return nil, err
		}
		loads = append(loads, load)
	}
	if len(loads) < 3 {
		return nil, errors.New("invalid loadavg format")
	}

	return loads, nil
}

func Hostname() (string, error) {
	// Workaround: inject spec.nodeName as env to get physics hostname in k8s environment
	injectNodeName, ok := os.LookupEnv("NODE_NAME")
	if ok {
		return injectNodeName, nil
	}

	f, err := os.Open("/proc/sys/kernel/hostname")
	if err != nil {
		return "", err
	}

	var hostname string
	_, err = fmt.Fscanln(f, &hostname)
	if err != nil {
		return "", err
	}

	return hostname, nil
}

func Uptime() (float64, error) {
	f, err := os.Open("/proc/uptime")
	if err != nil {
		return 0, err
	}

	var up, idle float64
	_, err = fmt.Fscanln(f, &up, &idle)
	if err != nil {
		return 0, err
	}

	return up, nil
}
