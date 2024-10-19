package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	reporter "github.com/openzipkin/zipkin-go/reporter/http"
	"net/http"
)

func OpenTracing() gin.HandlerFunc {
	// 创建上报节点
	report := reporter.NewReporter("http://localhost:9411/api/v2/spans")
	// 创建本地节点
	endpoint, err := zipkin.NewEndpoint("content-system", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// zipkin tracer
	tracer, err := zipkin.NewTracer(report,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithTraceID128Bit(true),
	)
	if err != nil {
		panic(err)
	}

	zipTracer := zipkinot.Wrap(tracer)
	opentracing.SetGlobalTracer(zipTracer)
	return ginhttp.Middleware(zipTracer, ginhttp.OperationNameFunc(func(r *http.Request) string {
		return r.URL.Path
	}))
}
