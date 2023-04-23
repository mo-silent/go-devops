package grafana

import "time"

type Dimensions struct {
	DomainName []string `json:"DomainName"`
	NodeID     []string `json:"NodeId"`
}
type Datasource struct {
	Type string `json:"type"`
	UID  string `json:"uid"`
}
type Queries struct {
	IntervalMs       int        `json:"intervalMs"`
	MaxDataPoints    int        `json:"maxDataPoints"`
	QueryMode        string     `json:"queryMode,omitempty"`
	Namespace        string     `json:"namespace,omitempty"`
	MetricName       string     `json:"metricName,omitempty"`
	Expression       string     `json:"expression,omitempty"`
	Dimensions       Dimensions `json:"dimensions,omitempty"`
	Region           string     `json:"region,omitempty"`
	ID               string     `json:"id,omitempty"`
	Alias            string     `json:"alias,omitempty"`
	Statistic        string     `json:"statistic,omitempty"`
	Period           string     `json:"period,omitempty"`
	MetricQueryType  int        `json:"metricQueryType,omitempty"`
	MetricEditorMode int        `json:"metricEditorMode,omitempty"`
	SQLExpression    string     `json:"sqlExpression,omitempty"`
	RefID            string     `json:"refId"`
	Key              string     `json:"key,omitempty"`
	Instant          bool       `json:"instant,omitempty"`
	Range            bool       `json:"range,omitempty"`
	MatchExact       bool       `json:"matchExact,omitempty"`
	Datasource       Datasource `json:"datasource"`
	Type             string     `json:"type,omitempty"`
	Exemplar         bool       `json:"exemplar,omitempty"`
	Expr             string     `json:"expr,omitempty"`
	Format           string     `json:"format,omitempty"`
	Interval         string     `json:"interval,omitempty"`
	IntervalFactor   int        `json:"intervalFactor,omitempty"`
	LegendFormat     string     `json:"legendFormat,omitempty"`
	Metric           string     `json:"metric,omitempty"`
	Step             int        `json:"step,omitempty"`
	QueryType        string     `json:"queryType,omitempty"`
	RequestID        string     `json:"requestId,omitempty"`
	UtcOffsetSec     int        `json:"utcOffsetSec,omitempty"`
	DatasourceID     int        `json:"datasourceId,omitempty"`
	RawSQL           string     `json:"rawSql,omitempty"`
}
type Raw struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type Range struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
	Raw  Raw       `json:"raw"`
}

type Options struct {
	From int64
	To   int64
	Step int64
}
