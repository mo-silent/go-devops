package grafana

import (
	"context"
	"reflect"
	"testing"
)

func TestAliGrafana_Query(t *testing.T) {
	type args struct {
		ctx   context.Context
		addr  string
		token string
		query string
		end   int64
	}
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
				addr:  "localhost:80",
				token: "BBearer xxxx",
				query: "sum by (callType) (sum_over_time[1m])",
				end:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ag := AliGrafana{}
			got, err := ag.Query(tt.args.ctx, tt.args.addr, tt.args.token, tt.args.query, tt.args.end)
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
