package main

import (
	"context"

	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
	"mini-node-exporter/services/views"

	"mini-node-exporter/services/routes"
)

func main() {
	startMonitor()
	route := routes.NewRouter()
	logrus.Fatal(route.Run(":23333"))

}

func startMonitor() {
	go raven.CapturePanic(
		func() {
			err := views.MonitorNode(context.TODO())
			if err != nil {
				logrus.WithError(err).Warn("fail to monitor node")
			}
		}, nil)
}
