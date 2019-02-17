package client

import (
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()
	// Create some standard client metrics.
	prometheusMetrics = grpc_prometheus.NewClientMetrics()
)

func init() {
	// Register client metrics to registry.
	reg.MustRegister(prometheusMetrics)
}
