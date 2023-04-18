package prometheus

// A MetricLabel is a collection of string and string pairs.  The LabelSet
// may be fully-qualified down to the point where it may resolve to a single
// Metric in the data store or not.  All operations that occur within the realm
// of a MetricLabel can emit a vector of Metric entities to which the MetricLabel may
// match.
type MetricLabel map[string]string

type MatrixResult struct {
	Metric MetricLabel    `json:"metric"`
	Values []MatrixValues `json:"values"`
}

type MatrixValues []interface{}
