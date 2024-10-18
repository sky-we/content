package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

func PrometheusMiddleware() gin.HandlerFunc {
	/*
		Prometheus 指标名称规则
		只能包含 ASCII 字符：指标名称只能包含 ASCII 字符。
		不能以数字开头：指标名称不能以数字开头。
		只能包含字母、数字和下划线：指标名称只能包含字母（a-z, A-Z）、数字（0-9）和下划线（_）。
		不能包含空格或特殊字符：指标名称不能包含空格或特殊字符（如 -, :, . 等）。
		长度限制：指标名称的最大长度是 256 个字符。
	*/
	reqTotals := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "ContentSystem",
		Subsystem: "http",
		Name:      "requests_total",
		Help:      "Total number of http request",
	}, []string{"method", "path"})

	rspCodeTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "ContentSystem",
		Subsystem: "http",
		Name:      "request_code_total",
		Help:      "Total number of http request code",
	}, []string{"method", "path", "code"})

	reqDuration := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "ContentSystem",
		Subsystem: "http",
		Name:      "response_time_total",
		Help:      "Response time of http request",
		Objectives: map[float64]float64{
			0.5:  0.05, // 50%以上的请求响应时间
			0.90: 0.01, // 90%以上的请求响应时间
			0.99: 0.01, // 99%以上的请求响应时间
		},
	}, []string{
		"method",
		"path",
	})

	prometheus.MustRegister(reqTotals)
	prometheus.MustRegister(rspCodeTotal)
	prometheus.MustRegister(reqDuration)

	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.FullPath()
		reqTotals.WithLabelValues(method, path).Inc()

		c.Next()
		elapsed := time.Since(start)
		reqDuration.WithLabelValues(method, path).Observe(elapsed.Seconds())

		statusCode := c.Writer.Status()
		rspCodeTotal.WithLabelValues(method, path, strconv.Itoa(statusCode)).Inc()
	}
}
