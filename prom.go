package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var RequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "demo_request_total",
		Help: "Number of requests",
	},
	[]string{"method", "handler", "status"},
)

func main() {
	registry := prometheus.NewRegistry()
	registry.MustRegister(RequestTotal)

	g := gin.Default()
	g.Use(func(c *gin.Context) {
		c.Next()
		RequestTotal.WithLabelValues(c.Request.Method, c.Request.URL.Path, strconv.Itoa(c.Writer.Status())).Inc()
	})
	g.GET("/metrics", gin.WrapH(promhttp.HandlerFor(registry, promhttp.HandlerOpts{})))
	g.GET("/echo/:code", func(c *gin.Context) {
		code, err := strconv.Atoi(c.Param("code"))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "bad request",
			})
			return
		}
		c.JSON(code, gin.H{
			"message": "hello worldd",
		})
		return
	})
	fmt.Println(g.Run(":8080"))
}
