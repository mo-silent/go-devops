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
