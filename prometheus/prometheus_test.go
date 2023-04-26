package prometheus

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"time"
)

func TestDataValueVec(t *testing.T) {
	type args struct {
		name  string
		label []string
	}
	tests := []struct {
		name string
		args args
		want *prometheus.GaugeVec
	}{
		{
			name: "t",
			args: args{
				name:  "test",
				label: []string{"tags"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DataValueVec(tt.args.name, tt.args.label); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataValueVec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrometheus_Push(t *testing.T) {
	type args struct {
		ctx  context.Context
		pm   PushMetrics
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "t",
			args: args{
				ctx: context.Background(),
				pm: PushMetrics{
					Name:  "test",
					Label: []string{"test1"},
					Metrics: []PromMetrics{
						{
							Values: []string{"123"},
							Data:   99.9,
						},
						{
							Values: []string{"125"},
							Data:   99.6,
						},
					},
				},
				addr: "localhost:9091",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prometheus{}
			if err := p.Push(tt.args.ctx, tt.args.pm, tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrometheus_Query(t *testing.T) {
	type args struct {
		ctx     context.Context
		client  api.Client
		query   string
		endTime int64
	}
	client, _ := api.NewClient(api.Config{
		Address: "http://10.158.215.90:80",
	})
	tests := []struct {
		name    string
		args    args
		want    ResultValues
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx:     context.Background(),
				client:  client,
				query:   "sum_over_time(api_data_by_channel_response_code_path{channel=\"ca\",job=\"api_data_by_channel_response_code_path_ca\", response_code=\"429\"}[4m])",
				endTime: 1682430985867,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prometheus{}
			got, err := p.Query(tt.args.ctx, tt.args.client, tt.args.query, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			vector, err := got.ConvertByte()
			if err != nil {
				log.Errorf("Convert matrix result to byte error, err: %v", err)
			}
			var result Vector
			if err := jsoniter.Unmarshal(vector, &result); err != nil {
				log.Errorf("error jsoniter unmarshal result to Vector, err: %v", err)
			}
			log.Info(result)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Query() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestPrometheus_QueryRange(t *testing.T) {
	type args struct {
		ctx    context.Context
		client api.Client
		query  string
		r      v1.Range
		opts   []v1.Option
	}
	client, _ := api.NewClient(api.Config{
		Address: "http://10.158.215.90:80",
	})
	tests := []struct {
		name    string
		args    args
		wantRes []MatrixResult
		wantErr bool
	}{
		{
			name: "t",
			args: args{
				ctx:    context.Background(),
				client: client,
				query:  "dynatrace_api_latency",
				r: v1.Range{
					Start: time.Now().Add(-4 * time.Minute),
					End:   time.Now(),
					Step:  time.Minute,
				},
				opts: []v1.Option{
					v1.WithTimeout(5 * time.Second),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prometheus{}
			gotRes, err := p.QueryRange(tt.args.ctx, tt.args.client, tt.args.query, tt.args.r, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			matrix, err := gotRes.ConvertByte()
			if err != nil {
				log.Errorf("Convert matrix result to byte error, err: %v", err)
			}
			var result Matrix
			if err := jsoniter.Unmarshal(matrix, &result); err != nil {
				log.Errorf("error jsoniter unmarshal result to prometheus.Matrix, err: %v", err)
			}
			log.Info(gotRes)
			//if !reflect.DeepEqual(gotRes, tt.wantRes) {
			//	t.Errorf("QueryRange() gotRes = %v, want %v", gotRes, tt.wantRes)
			//}
		})
	}
}
