package common

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type DevopsHttpClient interface {
	Get(ctx context.Context, addr string, headers map[string]string) ([]byte, error)
	Post(ctx context.Context, addr string, payload *strings.Reader, headers map[string]string, params url.Values) ([]byte, error)
}

type newHttp struct {
	Client http.Client
}

// Get is an HTTP GET method that returns a byte slice
// of the body of the GET request.
func (h *newHttp) Get(ctx context.Context, addr string, headers map[string]string) ([]byte, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)

	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyRes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	//defer close(body)
	return bodyRes, nil
}

// Post is an HTTP POST method that returns a byte slice
// of the body of the POST request.
func (h *newHttp) Post(ctx context.Context, addr string, payload *strings.Reader, headers map[string]string, params url.Values) ([]byte, error) {
	paramsUrl, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	paramsUrl.RawQuery = params.Encode()
	urlPathWithParams := paramsUrl.String()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlPathWithParams, payload)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := h.Client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func NewClient(client http.Client) DevopsHttpClient {
	return &newHttp{
		Client: client,
	}
}
