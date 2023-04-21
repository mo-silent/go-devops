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
	"github.com/mo-silent/go-devops/common"
	"github.com/mo-silent/go-devops/grafana"
	"github.com/mo-silent/go-devops/prometheus"
	"net/http"
)

// Tools implements all devops tools.
type Tools interface {
	Prometheus() prometheus.MetricsInterface
	Http(client http.Client) common.DevopsHttpClient
	AliGrafana() grafana.MetricsInterface
}

// Devops implements Tools.
type Devops struct{}

// Prometheus implements prometheus method.
func (d Devops) Prometheus() prometheus.MetricsInterface {
	return &prometheus.Prometheus{}
}

func (d Devops) Http(client http.Client) common.DevopsHttpClient {
	return common.NewClient(client)
}

func (d Devops) AliGrafana() grafana.MetricsInterface {
	return grafana.AliGrafana{}
}

// NewDevops implements devops method.
func NewDevops() Tools {
	return Devops{}
}
