package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Warnings is an array of non-critical errors
type Warnings []string

const (
	apiPrefix    = "/api/v1"
	epQueryRange = apiPrefix + "/query_range"
)

// apiClient wraps a regular client and processes successful API responses.
// Successful also includes responses that errored at the API level.
type apiClient interface {
	URL(ep string, args map[string]string) *url.URL
	Do(context.Context, *http.Request) (*http.Response, []byte, Warnings, error)
	DoGetFallback(ctx context.Context, u *url.URL, args url.Values) (*http.Response, []byte, Warnings, error)
}

type apiClientImpl struct {
	client api.Client
}

type apiResponse struct {
	Status    string          `json:"status"`
	Data      json.RawMessage `json:"data"`
	ErrorType v1.ErrorType    `json:"errorType"`
	Error     string          `json:"error"`
	Warnings  []string        `json:"warnings,omitempty"`
}

func (h *apiClientImpl) URL(ep string, args map[string]string) *url.URL {
	return h.client.URL(ep, args)
}

func (h *apiClientImpl) Do(ctx context.Context, req *http.Request) (*http.Response, []byte, Warnings, error) {
	resp, body, err := h.client.Do(ctx, req)
	if err != nil {
		return resp, body, nil, err
	}

	code := resp.StatusCode
	if code/100 != 2 && !apiError(code) {
		errorType, errorMsg := errorTypeAndMsgFor(resp)
		return resp, body, nil, &v1.Error{
			Type:   errorType,
			Msg:    errorMsg,
			Detail: string(body),
		}
	}

	var result apiResponse

	if http.StatusNoContent != code {
		if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
			return resp, body, nil, &v1.Error{
				Type: v1.ErrBadResponse,
				Msg:  jsonErr.Error(),
			}
		}
	}
	if apiError(code) && result.Status == "success" {
		err = &v1.Error{
			Type: v1.ErrBadResponse,
			Msg:  "inconsistent body for response code",
		}
	}

	if result.Status == "error" {
		err = &v1.Error{
			Type: result.ErrorType,
			Msg:  result.Error,
		}
	}

	return resp, result.Data, result.Warnings, err
}

// DoGetFallback will attempt to do the request as-is, and on a 405 or 501 it
// will fall back to a GET request.
func (h *apiClientImpl) DoGetFallback(ctx context.Context, u *url.URL, args url.Values) (*http.Response, []byte, Warnings, error) {
	encodedArgs := args.Encode()
	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(encodedArgs))
	if err != nil {
		return nil, nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Following comment originates from https://pkg.go.dev/net/http#Transport
	// Transport only retries a request upon encountering a network error if the request is
	// idempotent and either has not body or has its Request.GetBody defined. HTTP requests
	// are considered idempotent if they have HTTP methods GET, HEAD, OPTIONS, or TRACE; or
	// if their Header map contains an "Idempotency-Key" or "X-Idempotency-Key" entry. If the
	// idempotency key value is a zero-length slice, the request is treated as idempotent but
	// the header is not sent on the wire.
	req.Header["Idempotency-Key"] = nil

	resp, body, warnings, err := h.Do(ctx, req)
	if resp != nil && (resp.StatusCode == http.StatusMethodNotAllowed || resp.StatusCode == http.StatusNotImplemented) {
		u.RawQuery = encodedArgs
		req, err = http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, nil, warnings, err
		}
		return h.Do(ctx, req)
	}
	return resp, body, warnings, err
}

func apiError(code int) bool {
	// These are the codes that Prometheus sends when it returns an error.
	return code == http.StatusUnprocessableEntity || code == http.StatusBadRequest
}

func errorTypeAndMsgFor(resp *http.Response) (v1.ErrorType, string) {
	switch resp.StatusCode / 100 {
	case 4:
		return v1.ErrClient, fmt.Sprintf("client error: %d", resp.StatusCode)
	case 5:
		return v1.ErrServer, fmt.Sprintf("server error: %d", resp.StatusCode)
	}
	return v1.ErrBadResponse, fmt.Sprintf("bad response code %d", resp.StatusCode)
}

// RewriteAPI provides bindings for refactoring the v1 API for Prometheus.
type RewriteAPI interface {
	// QueryRange performs a query for the given range.
	QueryRange(ctx context.Context, query string, r v1.Range, opts ...Option) (any, Warnings, error)
}

type httpAPI struct {
	client apiClient
}

type apiOptions struct {
	timeout time.Duration
}

type Option func(c *apiOptions)

// WithTimeout can be used to provide an optional query evaluation timeout for Query and QueryRange.
// https://prometheus.io/docs/prometheus/latest/querying/api/#instant-queries
func WithTimeout(timeout time.Duration) Option {
	return func(o *apiOptions) {
		o.timeout = timeout
	}
}

// queryResult contains result data for a query.
type queryResult struct {
	Type   model.ValueType `json:"resultType"`
	Result interface{}     `json:"result"`

	// The decoded value.
	v any
}

func (qr *queryResult) UnmarshalJSON(b []byte) error {
	v := struct {
		Type   model.ValueType `json:"resultType"`
		Result json.RawMessage `json:"result"`
	}{}

	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case model.ValScalar:
		var sv model.Scalar
		err = json.Unmarshal(v.Result, &sv)
		qr.v = &sv

	case model.ValVector:
		var vv model.Vector
		err = json.Unmarshal(v.Result, &vv)
		fmt.Println(vv)
		qr.v = vv

	case model.ValMatrix:
		var mv []MatrixResult
		err = json.Unmarshal(v.Result, &mv)
		fmt.Println(mv[0])
		qr.v = mv

	default:
		err = fmt.Errorf("unexpected value type %q", v.Type)
	}
	return err
}

// NewAPI returns a new API for the client.
//
// It is safe to use the returned API from multiple goroutines.
func NewAPI(c api.Client) RewriteAPI {
	return &httpAPI{
		client: &apiClientImpl{
			client: c,
		},
	}
}

func (h *httpAPI) QueryRange(ctx context.Context, query string, r v1.Range, opts ...Option) (any, Warnings, error) {
	u := h.client.URL(epQueryRange, nil)
	q := u.Query()

	q.Set("query", query)
	q.Set("start", formatTime(r.Start))
	q.Set("end", formatTime(r.End))
	q.Set("step", strconv.FormatFloat(r.Step.Seconds(), 'f', -1, 64))

	opt := &apiOptions{}
	for _, o := range opts {
		o(opt)
	}

	d := opt.timeout
	if d > 0 {
		q.Set("timeout", d.String())
	}
	_, body, warnings, err := h.client.DoGetFallback(ctx, u, q)
	if err != nil {
		return nil, warnings, err
	}

	var queryRes queryResult
	//fmt.Println(string(body))
	return queryRes.v, warnings, json.Unmarshal(body, &queryRes)
}

func formatTime(t time.Time) string {
	return strconv.FormatFloat(float64(t.Unix())+float64(t.Nanosecond())/1e9, 'f', -1, 64)
}
