package devops

import (
	"github.com/mo-silent/go-devops/prometheus"
	"reflect"
	"testing"
)

func TestNewPrometheus(t *testing.T) {
	var tests []struct {
		name string
		want prometheus.MetricsInterface
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrometheus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrometheus() = %v, want %v", got, tt.want)
			}
		})
	}
}
