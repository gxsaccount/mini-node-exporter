package views

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"mini-node-exporter/pkg/metrics"
	"mini-node-exporter/pkg/proc"
)

var interval = time.Second * 10

func MonitorNode(ctx context.Context) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	hostname, err := proc.Hostname()
	if err != nil {
		return err
	}

	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			break
		}

		uptime, err := proc.Uptime()
		if err != nil {
			logrus.WithError(err).Error("fail to get uptime")
			return err
		}
		metrics.GaugeNodeUptime.WithLabelValues(hostname).Set(uptime)

		load, err := proc.Load()
		if err != nil {
			logrus.WithError(err).Error("fail to get load average")
			return err
		}
		for i, tag := range proc.LoadTags {
			metrics.GaugeNodeLoadavg.WithLabelValues(hostname, tag).Set(load[i])
		}
	}
}
