package main

import (
	"context"
	"github.com/mo-silent/go-devops"
	"github.com/mo-silent/go-devops/prometheus"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// An example of how to push Prometheus metrics with labels.
func main() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	p := devops.NewPrometheus()
	pm := prometheus.PushMetrics{
		Name:  "test",
		Label: []string{"sample1", "sample2"},
		Metrics: []prometheus.PromMetrics{
			{
				Values: []string{"s1", "s2"},
				Data:   99.99,
			},
		},
	}
	if err := p.Push(ctx, pm, "localhost:9091"); err != nil {
		log.Errorf("push metrics error, err:  %v", err)
	}
}
