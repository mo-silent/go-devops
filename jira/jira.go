package devops_jira

import (
	"github.com/andygrunwald/go-jira"
	"net/http"
	"strings"
)

// OpsJira implements the CRUD operations for Jira issues.
type OpsJira interface {
	Get(issueID string, options *jira.GetQueryOptions) (*jira.Issue, *jira.Response, error)
	Search(jql string, options *jira.SearchOptions) ([]jira.Issue, *jira.Response, error)
	Delete(issueID string) (*jira.Response, error)
}

// Jira use jira.Client.
type Jira struct {
	Client *jira.Client
}

// NewClient creates a new authenticated Jira client.
func (j *Jira) NewClient(addr string, auth AuthOptions) (err error) {
	var c *http.Client
	if auth.Token != "" {
		tp := jira.BearerAuthTransport{
			Token: auth.Token,
		}
		c = tp.Client()
	} else {
		tp := jira.BasicAuthTransport{
			Username: strings.TrimSpace(auth.Username),
			Password: strings.TrimSpace(auth.Password),
		}
		c = tp.Client()
	}

	client, err := jira.NewClient(c, strings.TrimSpace(addr))
	if err != nil {
		return err
	}
	j.Client = client
	return nil
}

// Search is implement github.com/andygrunwald/go-jira Issue.Search
func (j *Jira) Search(jql string, options *jira.SearchOptions) ([]jira.Issue, *jira.Response, error) {
	return j.Client.Issue.Search(jql, options)
}

// Get is implement github.com/andygrunwald/go-jira Issue.Get
func (j *Jira) Get(issueID string, options *jira.GetQueryOptions) (*jira.Issue, *jira.Response, error) {
	return j.Client.Issue.Get(issueID, options)
}

// Delete is implement github.com/andygrunwald/go-jira Issue.Delete
func (j *Jira) Delete(issueID string) (*jira.Response, error) {
	return j.Client.Issue.Delete(issueID)
}
