package main

import (
	"context"
	"fmt"
	"github.com/mo-silent/go-devops"
	"github.com/mo-silent/go-devops/prometheus"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// An example of how to query Prometheus metrics by range.
func main() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	client, err := api.NewClient(api.Config{
		Address: "http://localhost:80",
	})
	if err != nil {
		log.Errorf("Error creating client: %v\n", err)
		return
	}

	end := time.Now().UTC()

	r := v1.Range{
		Start: end.Add(-4 * time.Minute),
		End:   end,
		Step:  time.Minute,
	}
	log.Debugf("start: %v, end time: %v", end.Add(-6*time.Minute), end)

	p := devops.NewPrometheus()
	res, err := p.QueryRange(ctx, client, "test", r, prometheus.WithTimeout(5*time.Second))
	if err != nil {
		log.Errorf("Error querying Prometheus: %v\n", err)
	}
	fmt.Println(res)
}
