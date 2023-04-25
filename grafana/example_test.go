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

package grafana_test

import (
	"context"
	"fmt"
	"github.com/mo-silent/go-devops"
	"github.com/mo-silent/go-devops/grafana"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func ExampleGrafana_Query() {
	// An example of an open-source Grafana query

	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// new devops Grafana
	d := devops.NewDevops()
	ag := d.Grafana()
	addr := "https://localhost:3000/api/ds/query"
	token := "Bearer xxxx"
	query := `{"queries":[{"refId":"A","key":"Q-b34b5ac7-c2d7-44fc-bc77-06570400a564-0","instant":true,"range":false,"exemplar":false,"expr":"sum(rate(kube_pod_container_status_restarts_total{environment=~\"pd\"}[$__range])>=0) by (pod, environment)","datasource":{"uid":"yPdqCed7z","type":"prometheus"},"queryType":"timeSeriesQuery","requestId":"Q-b34b5ac7-c2d7-44fc-bc77-06570400a564-0A","utcOffsetSec":28800,"legendFormat":"","interval":"","datasourceId":34,"intervalMs":15000,"maxDataPoints":2512}],"range":{"from":"2023-04-23T02:40:33.146Z","to":"2023-04-23T03:10:33.146Z","raw":{"from":"now-30m","to":"now"}},"from":"1682217633146","to":"1682219433146"}`
	// query metrics
	bytes, err := ag.Query(ctx, addr, token, query, grafana.Options{})
	if err != nil {
		return
	}
	log.Debug(string(bytes))
}

func ExampleAliGrafana_Query() {
	// An example of an Alibaba Cloud Grafana query
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	d := devops.NewDevops()
	ag := d.AliGrafana()
	addr := "https://xxxx.grafana.aliyuncs.com/api/datasources/proxy/:id/api/v1/query"
	token := "Bearer xxxxx"
	query := `sum by (rpc) (sum_over_time(arms_rpc_requests_count{}[%s]))`
	timeframe := "68400s"
	end := time.Now().Local().Unix()

	query = fmt.Sprintf(query, timeframe)

	// query metrics
	res, err := ag.Query(ctx, addr, token, query, grafana.Options{To: end})
	if err != nil {
		log.Errorf("Post metrics data from ali grafana error: %s", err.Error())
	}
	fmt.Println(string(res))
}

func ExampleAliGrafana_QueryRange() {
	// An example of an Alibaba Cloud Grafana range query
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// new devops AliGrafana
	d := devops.NewDevops()
	ag := d.AliGrafana()
	addr := "https://xxxx.grafana.aliyuncs.com/api/datasources/proxy/:id/api/v1/query"
	token := "Bearer xxxxx"
	query := `sum by (rpc) (sum_over_time(arms_rpc_requests_count{}[1m]))`

	m, _ := time.ParseDuration("-15m")
	now := time.Now().Local()
	// options is grafana.Options.
	// The example is in the same directory as the structure,
	// so it does not have a directory beginning
	options := grafana.Options{
		From: now.Add(m).Unix(),
		To:   now.Unix(),
		Step: int64(60),
	}
	// use QueryRange
	res, err := ag.QueryRange(ctx, addr, token, query, options)
	if err != nil {
		log.Errorf("Post metrics data from ali grafana error: %s", err.Error())
	}
	fmt.Println(string(res))
}
