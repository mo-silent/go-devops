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
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/common/model"
	"net/http/httptrace"
	"time"
)

// MetricsInterface is the interface that
// implements prometheus push metrics
// and range queries.
type MetricsInterface interface {
	Push(context.Context, PushMetrics, string) error
	Query(ctx context.Context, client api.Client, query string, endTime int64) (ResultValues, error)
	QueryRange(ctx context.Context, client api.Client, query string, r v1.Range, opts ...v1.Option) (ResultValues, error)
}

// Prometheus implements MetricsInterface.
type Prometheus struct{}

// PushMetrics implements a new Prometheus metrics.
type PushMetrics struct {
	Name    string        // description of metrics name
	Label   []string      // description of metrics label
	Metrics []PromMetrics // values of metrics label and metrics value
}

// PromMetrics is a prometheus metrics values.
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

// Push implements a new prometheus metrics pushed to PushGateway.
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

func (p *Prometheus) Query(ctx context.Context, client api.Client, query string, endTime int64) (ResultValues, error) {
	end := time.Unix(0, endTime*int64(time.Millisecond)).UTC()
	v1api := v1.NewAPI(client)
	res, warnings, err := v1api.Query(ctx, query, end, v1.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	return convertValue(res), nil
}

// QueryRange implements prometheus range query.
func (p *Prometheus) QueryRange(ctx context.Context, client api.Client, query string, r v1.Range, opts ...v1.Option) (ResultValues, error) {
	v1api := v1.NewAPI(client)
	v1Res, warnings, err := v1api.QueryRange(ctx, query, r, opts...)
	if err != nil {
		return nil, err
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	return convertValue(v1Res), nil
}

func convertValue(value model.Value) (res ResultValues) {
	switch value.Type() {
	case model.ValScalar:
		v := value.(*model.Scalar)
		res = decodeScalar(v)
	case model.ValVector:
		v, _ := value.(model.Vector)
		res = decodeVector(v)
	case model.ValMatrix:
		mv := value.(model.Matrix)
		res = decodeMatrix(mv)
	default:
		fmt.Printf("Unknow Type")
	}
	return
}

func decodeMatrix(matrix model.Matrix) (res Matrix) {

	for _, ss := range matrix {
		mr := MatrixResult{
			Metric: fmt.Sprint(ss.Metric),
			Values: nil,
		}
		for _, sp := range ss.Values {
			v := MetricValues{
				Timestamp: int64(sp.Timestamp),
				Value:     float64(sp.Value),
			}
			mr.Values = append(mr.Values, v)
		}
		res = append(res, mr)
	}
	return
}

func decodeVector(vector model.Vector) (res Vector) {
	for _, mv := range vector {
		v := VectorResult{
			Metric: fmt.Sprint(mv.Metric),
			Values: MetricValues{
				Timestamp: int64(mv.Timestamp),
				Value:     float64(mv.Value),
			},
		}
		res = append(res, v)
	}
	return
}

func decodeScalar(scalar *model.Scalar) Scalar {
	return Scalar{
		Value: MetricValues{
			Value:     float64(scalar.Value),
			Timestamp: int64(scalar.Timestamp),
		},
	}
}
