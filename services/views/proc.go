package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"mini-node-exporter/pkg/proc"
)

// GetHostname returns local hostname
func GetHostname(ctx *gin.Context) (int, interface{}, error) {
	name, err := proc.Hostname()
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "fail to get hostname")
	}

	return http.StatusOK, name, nil
}

// GetUptime returns current host uptime
func GetUptime(ctx *gin.Context) (int, interface{}, error) {
	uptime, err := proc.Uptime()
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "fail to get uptime")
	}

	return http.StatusOK, uptime, nil
}

// GetLoadAvg returns current host uptime
func GetLoadAvg(ctx *gin.Context) (int, interface{}, error) {
	avg, err := proc.Load()
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "fail to get uptime")
	}

	load := make(map[string]float64)
	for i, tag := range proc.LoadTags {
		load[tag] = avg[i]
	}

	return http.StatusOK, load, nil
}
