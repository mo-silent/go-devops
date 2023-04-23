package main

import (
	"context"
	"github.com/mo-silent/go-devops"
	"github.com/mo-silent/go-devops/grafana"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	d := devops.NewDevops()
	ag := d.Grafana()
	addr := "https://localhost:3000/api/ds/query"
	token := "Bearer xxxx"
	query := `{"queries":[{"refId":"A","key":"Q-b34b5ac7-c2d7-44fc-bc77-06570400a564-0","instant":true,"range":false,"exemplar":false,"expr":"sum(rate(kube_pod_container_status_restarts_total{environment=~\"pd\"}[$__range])>=0) by (pod, environment)","datasource":{"uid":"yPdqCed7z","type":"prometheus"},"queryType":"timeSeriesQuery","requestId":"Q-b34b5ac7-c2d7-44fc-bc77-06570400a564-0A","utcOffsetSec":28800,"legendFormat":"","interval":"","datasourceId":34,"intervalMs":15000,"maxDataPoints":2512}],"range":{"from":"2023-04-23T02:40:33.146Z","to":"2023-04-23T03:10:33.146Z","raw":{"from":"now-30m","to":"now"}},"from":"1682217633146","to":"1682219433146"}`
	bytes, err := ag.Query(ctx, addr, token, query, grafana.Options{})
	if err != nil {
		return
	}
	log.Debug(string(bytes))

}
