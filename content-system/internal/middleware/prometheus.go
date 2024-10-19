package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

func Prometheus() gin.HandlerFunc {

	rspCodeTotal := prometheus.NewCounterVec(prometheus.CounterOpts{

		Name: "http_requests_total",
		Help: "Total number of http request code",
	}, []string{"method", "path", "code"})

	reqDuration := prometheus.NewSummaryVec(prometheus.SummaryOpts{

		Name: "response_time_total",
		Help: "Response time of http request",
		Objectives: map[float64]float64{
			0.5:  0.05, // 50%以上的请求响应时间
			0.90: 0.01, // 90%以上的请求响应时间
			0.99: 0.01, // 99%以上的请求响应时间
		},
	}, []string{
		"method",
		"path",
	})

	prometheus.MustRegister(rspCodeTotal)
	prometheus.MustRegister(reqDuration)

	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.FullPath()

		c.Next()
		elapsed := time.Since(start)
		reqDuration.WithLabelValues(method, path).Observe(elapsed.Seconds())

		statusCode := c.Writer.Status()
		rspCodeTotal.WithLabelValues(method, path, strconv.Itoa(statusCode)).Inc()
	}
}
