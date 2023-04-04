package prometheus

import (
	"context"
	"encoding/json"
	"github.com/prometheus/client_golang/api"
	"net/http"
	"net/url"
	"strings"
)

// Warnings is an array of non-critical errors
type Warnings []string

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
	ErrorType string          `json:"errorType"`
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
		return resp, body, nil, nil
	}

	var result apiResponse

	if http.StatusNoContent != code {
		if jsonErr := json.Unmarshal(body, &result); jsonErr != nil {
			return resp, body, nil, nil
		}
	}
	return resp, body, nil, nil
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
	// idempotent and either has no body or has its Request.GetBody defined. HTTP requests
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
