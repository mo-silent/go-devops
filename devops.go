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

package devops

import (
	"github.com/mo-silent/go-devops/grafana"
	opsJira "github.com/mo-silent/go-devops/jira"
	"github.com/mo-silent/go-devops/prometheus"
)

// Tools implements all devops tools.
type Tools interface {
	Prometheus() prometheus.MetricsInterface
	AliGrafana() grafana.MetricsInterface
	Grafana() grafana.MetricsInterface
	Jira() opsJira.OpsJira
}

// Devops implements Tools.
type Devops struct{}

// Prometheus implements prometheus.Prometheus.
func (d Devops) Prometheus() prometheus.MetricsInterface {
	return &prometheus.Prometheus{}
}

// AliGrafana implement grafana.AliGrafana.
func (d Devops) AliGrafana() grafana.MetricsInterface {
	return grafana.AliGrafana{}
}

// Grafana implement grafana.Grafana.
func (d Devops) Grafana() grafana.MetricsInterface {
	return grafana.Grafana{}
}

// Jira implement opsJira.Jira.
func (d Devops) Jira() opsJira.OpsJira {
	return &opsJira.Jira{}
}

// NewDevops implements devops method.
func NewDevops() Tools {
	return Devops{}
}
