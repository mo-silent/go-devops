package devops

import (
	"github.com/mo-silent/go-devops/prometheus"
)

func NewPrometheus() prometheus.MetricsInterface {
	return &prometheus.Prometheus{}
}
