package tracing

import (
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

const (
	ZIPKIN_HTTP_ENDPOINT      = "http://127.0.0.1:9411/api/v1/spans"
	ZIPKIN_RECORDER_HOST_PORT = "127.0.0.1:9000"
)

func NewZipkinTracer(serviceName string) (opentracing.Tracer, error) {
	collector, err := zipkin.NewHTTPCollector(ZIPKIN_HTTP_ENDPOINT)
	if err != nil {
		return nil, err
	}

	recorder := zipkin.NewRecorder(collector, true, ZIPKIN_RECORDER_HOST_PORT, serviceName)

	return zipkin.NewTracer(recorder, zipkin.ClientServerSameSpan(false))
}
