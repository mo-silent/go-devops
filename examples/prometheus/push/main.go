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
