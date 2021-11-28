package moniter

import (
	"context"
	"mini-node-exporter/pkg/metrics"
	"mini-node-exporter/pkg/proc"
	"runtime/metrics"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	interval = time.Second * 10
)

func Moniter(ctx context.Context) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			break
		}
		hostname, err := proc.GetHostName()
		if err != nil {
			logrus.WithError(err).Error("fail to get hostname")
			return err
		}
		uptime, err := proc.GetUpTime()
		if err != nil {
			logrus.WithError(err).Error("fail to get uptime")
			return err
		}
		metrics.GaugeNodeUptime.WithLabelValues(hostname).Set(float64(uptime))
		loads, err := proc.GetLoad()
		if err != nil {
			logrus.WithError(err).Error("fail to get load")
			return err
		}
		for interval,load : range loads{
			metrics.GaugeNodeLoadavg.WithLabelValues(hostname, interval).Set(load)
		} 
		
	}
}
