package main

import (
	"context"
	"fmt"
	"github.com/mo-silent/go-devops"
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
	ag := d.AliGrafana()
	addr := "https://xxxx.grafana.aliyuncs.com/api/datasources/proxy/:id/api/v1/query"
	token := "Bearer xxxxx"
	query := `sum by (rpc) (sum_over_time(arms_rpc_requests_count{}[%s]))`
	timeframe := "68400s"
	end := time.Now().Local().Unix()

	query = fmt.Sprintf(query, timeframe)
	res, err := ag.Query(ctx, addr, token, query, end)
	if err != nil {
		log.Errorf("Post metrics data from ali grafana error: %s", err.Error())
	}
	fmt.Println(string(res))
}
