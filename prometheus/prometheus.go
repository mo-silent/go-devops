// Copyright © 2023  silent mo <1916393131@qq.com>.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"net/http/httptrace"
)

// MetricsInterface is the interface that
// implements prometheus push metrics
// and range queries
type MetricsInterface interface {
	Push(context.Context, PushMetrics, string) error
	QueryRange(context.Context, api.Client, string, v1.Range, ...Option) ([]MatrixResult, error)
}

// Prometheus implemented MetricsInterface
type Prometheus struct{}

// PushMetrics implemented a new Prometheus metrics
type PushMetrics struct {
	Name    string        // description of metrics name
	Label   []string      // description of metrics label
	Metrics []PromMetrics // values of metrics label and metrics value
}

type PromMetrics struct {
	Values []string // values of metrics label
	Data   float64  // metrics value
}

var clientTrace = &httptrace.ClientTrace{
	GotConn: func(info httptrace.GotConnInfo) {
		fmt.Printf("GotConn: %v reused\n", info)
	},
}

// DataValueVec implements a new prometheus metrics
// that includes label.
func DataValueVec(name string, label []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf(name),
		Help: fmt.Sprintf("The jobs of %s in dynatrace.", name),
	}, label)
}

// Push implements a new prometheus metrics pushed to PushGateway
func (p *Prometheus) Push(ctx context.Context, pm PushMetrics, addr string) error {

	gauge := DataValueVec(pm.Name, pm.Label)
	for _, k := range pm.Metrics {
		gauge.WithLabelValues(k.Values...).Set(k.Data)
	}

	traceCtx := httptrace.WithClientTrace(ctx, clientTrace)
	pn := push.New(addr, pm.Name)
	err := pn.AddContext(traceCtx)
	if err != nil {
		fmt.Printf("prometheus add context error, err: %v", err)
		return err
	}
	pn.Collector(gauge)
	return pn.Push()
	//return nil
}

// QueryRange implements prometheus range query.
func (p *Prometheus) QueryRange(ctx context.Context, client api.Client, query string, r v1.Range, opts ...Option) (res []MatrixResult, err error) {
	v1api := NewAPI(client)
	v1Res, warnings, err := v1api.QueryRange(ctx, query, r, opts...)
	if err != nil {
		return nil, err
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	if err = json.Unmarshal(v1Res, &res); err != nil {
		fmt.Printf("un marshal query range value error, err: %v, value: %s\n", err, string(v1Res))
	}

	return
}
