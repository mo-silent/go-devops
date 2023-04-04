package prometheus

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	log "github.com/sirupsen/logrus"
	"net/http/httptrace"
)

type MetricsInterface interface {
	Push(context.Context, string, string) error
}

type Prometheus struct {
	Pusher  *push.Pusher
	Label   []string
	Metrics []PromMetrics
}

type PromMetrics struct {
	Values []string
	Data   float64
}

var clientTrace = &httptrace.ClientTrace{
	GotConn: func(info httptrace.GotConnInfo) {
		log.Infof("GotConn: %v reused", info)
	},
}

func DataValueVec(name string, label []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf(name),
		Help: fmt.Sprintf("The jobs of %s in dynatrace.", name),
	}, label)
}

func (p *Prometheus) Push(ctx context.Context, name, addr string) error {

	gauge := DataValueVec(name, p.Label)
	for _, k := range p.Metrics {
		gauge.WithLabelValues(k.Values...).Set(k.Data)
	}

	traceCtx := httptrace.WithClientTrace(ctx, clientTrace)
	pn := push.New(addr, name)
	err := pn.AddContext(traceCtx)
	if err != nil {
		log.Errorf("prometheus add context error, err: %v", err)
		return err
	}
	pn.Collector(gauge)
	return pn.Push()
	//return nil
}
