package routes

import (
	"mini-node-exporter/pkg/proc"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//group info
	api := r.Group("info")
	api.Handle("GET", "/hostname", func(c *gin.Context) {
		hostname, _ := proc.GetHostName()
		c.JSON(200, gin.H{
			"message": hostname,
		})
	})
	api.Handle("GET", "/uptime", func(c *gin.Context) {
		uptime, _ := proc.GetUpTime()
		c.JSON(200, gin.H{
			"message": uptime,
		})
	})
	api.Handle("GET", "/load", func(c *gin.Context) {
		load, _ := proc.GetLoad()
		c.JSON(200, gin.H{
			"message": load,
		})
	})

	// promtheus metrics
	r.Handle("GET", "/metrics", prometheusHandler())
	return r
}
func prometheusHandler() gin.HandlerFunc {
	handle := promhttp.Handler()
	return func(c *gin.Context) {
		handle.ServeHTTP(c.Writer, c.Request)
	}
}
