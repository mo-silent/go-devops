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

import "encoding/json"

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

// MatrixResult obtains the matrix result from the prometheus query.
type MatrixResult struct {
	Metric string `json:"metric"`
	//Labels []MetricLabel  `json:"labels"`
	Values []MetricValues `json:"values"`
}

// VectorResult obtains the vector result from the prometheus query.
type VectorResult struct {
	Metric string       `json:"metric"`
	Values MetricValues `json:"values"`
}

// ScalarResult obtains the scalar result from the prometheus query.
type ScalarResult struct {
	Value MetricValues `json:"value"`
}

// Matrix is a list of time series.
type Matrix []MatrixResult

// Vector is a list of vector result from the prometheus query
type Vector []VectorResult

// Scalar is scalar result
type Scalar ScalarResult

type ResultValues interface {
	ConvertByte() ([]byte, error)
}

func (m Matrix) ConvertByte() ([]byte, error) {
	return json.Marshal(m)
}

func (v Vector) ConvertByte() ([]byte, error) {
	return json.Marshal(v)
}

func (s Scalar) ConvertByte() ([]byte, error) {
	return json.Marshal(s)
}
