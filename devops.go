package devops

import (
	"github.com/mo-silent/go-devops/prometheus"
	log "github.com/sirupsen/logrus"
	"os"
)

func NewPrometheus() prometheus.MetricsInterface {
	return &prometheus.Prometheus{}
}

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}
