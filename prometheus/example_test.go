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

package prometheus_test

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mo-silent/go-devops"
	"github.com/mo-silent/go-devops/prometheus"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// ExamplePrometheus_Push demonstrates how to use Prometheus.Push.
// Open the link to see an example: https://go.dev/play/p/ViGSBJbtGKz
func ExamplePrometheus_Push() {
	// prometheus push metrics example
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	p := devops.NewDevops().Prometheus()
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

// ExamplePrometheus_Query_vector demonstrates how to
// query vector data using Prometheus.Query.
// Open the link to see an example: https://go.dev/play/p/cbqzfN1JDOE
func ExamplePrometheus_Query_vector() {
	// An example of how to query Prometheus metrics by range.
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// new prometheus client
	client, err := api.NewClient(api.Config{
		Address: "http://10.158.215.90:80",
	})
	if err != nil {
		log.Errorf("Error creating client: %v\n", err)
		return
	}

	end := time.Now().UTC().Unix()

	p := devops.NewDevops().Prometheus()
	res, err := p.Query(ctx, client, "dynatrace_api_latency", end)
	if err != nil {
		log.Errorf("Error querying Prometheus: %v\n", err)
	}
	vector, err := res.ConvertByte()
	if err != nil {
		log.Errorf("Convert matrix result to byte error, err: %v", err)
	}
	var result prometheus.Vector
	if err := jsoniter.Unmarshal(vector, &result); err != nil {
		log.Errorf("error jsoniter unmarshal result to prometheus.Matrix, err: %v", err)
	}
	fmt.Println(result)
}

// ExamplePrometheus_Query_vector demonstrates how to
// query scalar data using Prometheus.Query.
// Open the link to see an example: https://go.dev/play/p/cbqzfN1JDOE
func ExamplePrometheus_Query_scalar() {
	// An example of how to query Prometheus metrics by range.
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// new prometheus client
	client, err := api.NewClient(api.Config{
		Address: "http://10.158.215.90:80",
	})
	if err != nil {
		log.Errorf("Error creating client: %v\n", err)
		return
	}

	end := time.Now().UTC().Unix()

	p := devops.NewDevops().Prometheus()
	res, err := p.Query(ctx, client, "dynatrace_api_latency", end)
	if err != nil {
		log.Errorf("Error querying Prometheus: %v\n", err)
	}
	vector, err := res.ConvertByte()
	if err != nil {
		log.Errorf("Convert matrix result to byte error, err: %v", err)
	}
	var result prometheus.Vector
	if err := jsoniter.Unmarshal(vector, &result); err != nil {
		log.Errorf("error jsoniter unmarshal result to prometheus.Matrix, err: %v", err)
	}
	fmt.Println(result)
}

// ExamplePrometheus_QueryRange demonstrates how to use Prometheus.QueryRange.
// Open the link to see an example: https://go.dev/play/p/32LVSjFs2hU
func ExamplePrometheus_QueryRange() {
	// An example of how to query Prometheus metrics by range.
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// new prometheus client
	client, err := api.NewClient(api.Config{
		Address: "http://10.158.215.90:80",
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
	log.Debugf("start: %v, end time: %v", end.Add(-4*time.Minute), end)

	p := devops.NewDevops().Prometheus()
	res, err := p.QueryRange(ctx, client, "dynatrace_api_latency", r, v1.WithTimeout(5*time.Second))
	if err != nil {
		log.Errorf("Error querying Prometheus: %v\n", err)
	}
	matrix, err := res.ConvertByte()
	if err != nil {
		log.Errorf("Convert matrix result to byte error, err: %v", err)
	}
	var result prometheus.Matrix
	if err := jsoniter.Unmarshal(matrix, &result); err != nil {
		log.Errorf("error jsoniter unmarshal result to prometheus.Matrix, err: %v", err)
	}
	fmt.Println(result)

}
