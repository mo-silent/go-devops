package grafana

import (
	"context"
	"crypto/tls"
	"github.com/mo-silent/go-devops/common"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// MetricsInterface is the interface that
// implements grafana query metrics.
type MetricsInterface interface {
	Query(ctx context.Context, addr, token, query string, end int64) ([]byte, error)
	QueryRange(ctx context.Context, addr, token, query string, end, start, step int64) ([]byte, error)
}

type Grafana struct {
	//AliGrafana
}

type AliGrafana struct {
}

func (ag AliGrafana) Query(ctx context.Context, addr, token, query string, end int64) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS13,
				InsecureSkipVerify: true,
			}},
	}
	newHttp := common.NewClient(client)
	payload := strings.NewReader("")

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = token
	params := url.Values{}
	params.Set("time", strconv.FormatInt(end, 10))
	//log.Debugf("query: %s", query)
	params.Set("query", query)

	res, err := newHttp.Post(ctx, addr, payload, header, params)
	if err != nil {
		//log.Errorf("Post metrics data from grafana error: %s", err.Error())
		return nil, err
	}
	return res, nil

}

func (ag AliGrafana) QueryRange(ctx context.Context, addr, token, query string, end, start, step int64) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS13,
				InsecureSkipVerify: true,
			}},
	}
	newHttp := common.NewClient(client)
	payload := strings.NewReader("")

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = token
	params := url.Values{}
	params.Set("end", strconv.FormatInt(end, 10))
	params.Set("start", strconv.FormatInt(start, 10))
	params.Set("step", strconv.FormatInt(step, 10))
	//log.Debugf("query: %s", query)
	params.Set("query", query)

	res, err := newHttp.Post(ctx, addr, payload, header, params)
	if err != nil {
		//log.Errorf("Post metrics data from grafana error: %s", err.Error())
		return nil, err
	}
	return res, nil

}
