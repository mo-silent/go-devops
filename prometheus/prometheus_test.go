package prometheus

import (
	"context"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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

func TestPrometheus_QueryRange(t *testing.T) {
	type args struct {
		ctx    context.Context
		client api.Client
		query  string
		r      v1.Range
		opts   []Option
	}
	client, _ := api.NewClient(api.Config{
		Address: "http://localhost:80",
	})
	tests := []struct {
		name    string
		args    args
		wantRes []MatrixResult
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "t",
			args: args{
				ctx:    context.Background(),
				client: client,
				query:  "test",
				r: v1.Range{
					Start: time.Now().Add(-4 * time.Minute),
					End:   time.Now(),
					Step:  time.Minute,
				},
				opts: []Option{
					WithTimeout(5 * time.Second),
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
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("QueryRange() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
