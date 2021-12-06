package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	GaugeNodeLoadavg = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "node_load",
			Help: "node avg load in 1min, 5min and 15min",
		},
		[]string{"hostname", "duration"},
	)

	GaugeNodeUptime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "node_uptime",
			Help: "node uptime",
		},
		[]string{"hostname"},
	)
)
