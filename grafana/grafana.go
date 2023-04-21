package grafana

import (
	"context"
	"crypto/tls"
	"github.com/mo-silent/go-devops"
	log "github.com/sirupsen/logrus"
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
	newHttp := devops.NewDevops().Http(client)
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
		log.Errorf("Post metrics data from grafana error: %s", err.Error())
		return nil, err
	}
	return res, nil

}
