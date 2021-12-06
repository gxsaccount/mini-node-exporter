package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"mini-node-exporter/services/views"
)

type CustomHandler func(ctx *gin.Context) (httpStatus int, data interface{}, err error)

type Route struct {
	Desc    string
	Method  string
	Path    string
	Handler CustomHandler
}

type MsgView struct {
	Message string `json:"message"`
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)
	// register prometheus metrics
	r.Handle("GET", "/metrics", prometheusHandler())
	// register info apis
	api := r.Group("/info")
	for _, route := range routes {
		api.Handle(route.Method, route.Path, wrappedHandler(route))
	}

	return r
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func wrappedHandler(route *Route) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				msg := MsgView{Message: fmt.Sprintf("%v", p)}
				ctx.JSON(http.StatusInternalServerError, msg)
				panic(p)
			}
		}()

		code, data, err := route.Handler(ctx)
		if err != nil {
			msg := MsgView{Message: err.Error()}
			ctx.JSON(code, msg)
			return
		}
		if data != nil {
			ctx.JSON(code, data)
			return
		}
	}
}

var routes = []*Route{
	{
		Desc:    "get hostname",
		Method:  "GET",
		Path:    "/hostname",
		Handler: views.GetHostname,
	},
	{
		Desc:    "get current node uptime",
		Method:  "GET",
		Path:    "/uptime",
		Handler: views.GetUptime,
	},
	{
		Desc:    "get current load average",
		Method:  "GET",
		Path:    "/load",
		Handler: views.GetLoadAvg,
	},
}
