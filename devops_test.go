package devops

import (
	"github.com/mo-silent/go-devops/prometheus"
	"reflect"
	"testing"
)

func TestDevops_Prometheus(t *testing.T) {
	tests := []struct {
		name string
		want prometheus.MetricsInterface
	}{
		{
			name: "test",
			want: &prometheus.Prometheus{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Devops{}
			if got := d.Prometheus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prometheus() = %v, want %v", got, tt.want)
			}
		})
	}
}
