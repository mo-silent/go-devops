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

package prometheus

// A MetricLabel is a prometheus metric label structure.
type MetricLabel struct {
	Label string
	Value string
}

// MetricValues gets the metric value from the prometheus query.
type MetricValues struct {
	Timestamp int64
	Value     float64
}

// MatrixResult obtains the matrix result y from the prometheus query.
type MatrixResult struct {
	Metric string `json:"metric"`
	//Labels []MetricLabel  `json:"labels"`
	Values []MetricValues `json:"values"`
}

type VectorResult struct {
	Metric string       `json:"metric"`
	Values MetricValues `json:"values"`
}

type ResultT interface {
	MetricValues | []MatrixResult | []VectorResult
}
