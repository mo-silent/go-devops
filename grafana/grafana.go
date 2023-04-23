package grafana

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/mo-silent/go-devops/common"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	client = http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS13,
				InsecureSkipVerify: true,
			}},
	}
	payload = strings.NewReader("")
)

// MetricsInterface is the interface that
// implements grafana metrics query and
// range query methods.
type MetricsInterface interface {
	Query(ctx context.Context, addr, token, query string, options Options) ([]byte, error)
	QueryRange(ctx context.Context, addr, token, query string, options Options) ([]byte, error)
}

// Grafana implements open-source grafana
// query and range query methods.
type Grafana struct {
}

// The Query method is used to query open-source grafana
// indicator data and return byte slices.
func (ag Grafana) Query(ctx context.Context, addr, token, query string, _ Options) ([]byte, error) {
	newHttp := common.NewClient(client)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = token

	params := url.Values{}

	data := []byte(query)
	body := bytes.NewBuffer(data)

	res, err := newHttp.Post(ctx, addr, body, header, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QueryRange implements Grafana.Query.
func (ag Grafana) QueryRange(ctx context.Context, addr, token, query string, options Options) ([]byte, error) {
	//res, err := ag.Query(ctx, addr, token, query, 0)
	return ag.Query(ctx, addr, token, query, options)
}

// AliGrafana implements Alibaba Cloud grafana
// query and range query methods.
type AliGrafana struct {
}

// The Query method is used to query Alibaba Cloud grafana
// indicator data and return byte slices.
func (ag AliGrafana) Query(ctx context.Context, addr, token, query string, options Options) ([]byte, error) {
	newHttp := common.NewClient(client)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = token

	params := url.Values{}
	params.Set("time", strconv.FormatInt(options.To, 10))
	params.Set("query", query)

	res, err := newHttp.Post(ctx, addr, payload, header, params)
	if err != nil {
		return nil, err
	}
	return res, nil

}

// The QueryRange method is used to query Alibaba Cloud grafana
// indicator data range and return byte slices.
func (ag AliGrafana) QueryRange(ctx context.Context, addr, token, query string, options Options) ([]byte, error) {
	newHttp := common.NewClient(client)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = token
	params := url.Values{}
	params.Set("end", strconv.FormatInt(options.To, 10))
	params.Set("start", strconv.FormatInt(options.From, 10))
	params.Set("step", strconv.FormatInt(options.Step, 10))
	params.Set("query", query)

	res, err := newHttp.Post(ctx, addr, payload, header, params)
	if err != nil {
		return nil, err
	}
	return res, nil

}
