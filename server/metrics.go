package server

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	prometheusMetrics = grpc_prometheus.NewServerMetrics()

	// Create a customized counter metric.
	customizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "demo_server",
		Help: "Total number of RPCs handled on the server.",
	}, []string{"name"})

	httpServer = &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9090)}
)

func init() {
	// Register standard server metrics and customized metrics to registry.
	reg.MustRegister(prometheusMetrics, customizedCounterMetric)
	customizedCounterMetric.WithLabelValues("Server")
}

func startMetrics(){
	if err := httpServer.ListenAndServe(); err != nil {
		logrus.Fatalf("Unable to start a http server.")
	}
}