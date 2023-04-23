package grafana

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestAliGrafana(t *testing.T) {
	type args struct {
		ctx   context.Context
		addr  string
		token string
		query string
		ops   Options
	}
	m, _ := time.ParseDuration("-15m")
	now := time.Now().Local()
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx:   context.Background(),
				addr:  "localhost:3000",
				token: "BBearer xxxx",
				query: "sum by (callType) (sum_over_time[1m])",
				ops: Options{
					From: now.Add(m).Unix(),
					To:   now.Unix(),
					Step: int64(60),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ag := AliGrafana{}
			got, err := ag.Query(tt.args.ctx, tt.args.addr, tt.args.token, tt.args.query, tt.args.ops)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() got = %v, want %v", got, tt.want)
			}

			got2, err := ag.QueryRange(tt.args.ctx, tt.args.addr, tt.args.token, tt.args.query, tt.args.ops)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("QueryRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrafana_Query(t *testing.T) {
	type args struct {
		ctx   context.Context
		addr  string
		token string
		query string
		ops   Options
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "t",
			args: args{
				ctx:   context.Background(),
				addr:  "localhost:3000",
				token: "BBearer xxxx",
				query: `{"queries":[{"refId":"A","key":"Q-b34b5ac7-c2d7-44fc-bc77-06570400a564-0","instant":true,"range":false,"exemplar":false,"expr":"sum(rate(kube_pod_container_status_restarts_total{environment=~\"pd\"}[$__range])>=0) by (pod, environment)","datasource":{"uid":"yPdqCed7z","type":"prometheus"},"queryType":"timeSeriesQuery","requestId":"Q-b34b5ac7-c2d7-44fc-bc77-06570400a564-0A","utcOffsetSec":28800,"legendFormat":"","interval":"","datasourceId":34,"intervalMs":15000,"maxDataPoints":2512}],"range":{"from":"2023-04-23T02:40:33.146Z","to":"2023-04-23T03:10:33.146Z","raw":{"from":"now-30m","to":"now"}},"from":"1682217633146","to":"1682219433146"}`,
				ops:   Options{To: time.Now().Unix()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ag := Grafana{}
			got, err := ag.Query(tt.args.ctx, tt.args.addr, tt.args.token, tt.args.query, tt.args.ops)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() got = %v, want %v", got, tt.want)
			}
		})
	}
}
