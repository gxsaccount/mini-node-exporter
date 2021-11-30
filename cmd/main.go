package main

import (
	"context"
	"mini-node-exporter/services/moniter"
	"mini-node-exporter/services/routes"

	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

func main() {
	startMonitor()
	r := routes.NewRouter()
	r.Run(":23333")
}

func startMonitor() {
	go raven.CapturePanic(
		func() {
			err := moniter.Moniter(context.TODO())
			if err != nil {
				logrus.WithError(err).Warn("fail to monitor node")
			}
		}, nil)
}
