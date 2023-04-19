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
	Metric string `json:"metric""`
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
