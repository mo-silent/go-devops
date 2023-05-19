package devops_jira

import "net/http"

// AuthOptions is a http.RoundTripper that authenticates all requests
// Use HTTP basic authentication and provided username and password
// as well as authentication based on Jira hosting (oauth 2.0 (3lo)).
type AuthOptions struct {
	// Username using HTTP Basic Authentication with the provided username
	// and password.
	Username string
	// Password using HTTP Basic Authentication with the provided  password.
	Password string
	// Token using Jira bearer (oauth 2.0 (3lo)) based authentication.
	Token string
	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}
