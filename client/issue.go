package client

import (
	"context"
	"net/http"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/ram02z/gh-weekly-log/schema"
)

// IssueClient represents client to issues REST API endpoints
// See https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28
type IssueClient client

func DefaultIssueClient() (*IssueClient, error) {
	return NewIssueClient(api.ClientOptions{})
}

func NewIssueClient(opts api.ClientOptions) (*IssueClient, error) {
	client, err := api.NewRESTClient(opts)
	if err != nil {
		return nil, err
	}
	return &IssueClient{client: client, method: "issues"}, nil
}

// IssueListOptions specifies the optional parameters to the
// IssueClient.List method
type IssueListOptions struct {
	// Filter specifies which issues to list. Possible values are: assigned,
	// created, mentioned, subscribed, all. Default is "assigned".
	Filter string `url:"filter,omitempty"`

	// State filters issues based on their state. Possible values are: open,
	// closed, all. Default is "open".
	State string `url:"state,omitempty"`

	// Labels filters issues based on their label.
	Labels []string `url:"labels,comma,omitempty"`

	// Sort specifies how to sort issues. Possible values are: created, updated,
	// and comments. Default value is "created".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort issues. Possible values are: asc, desc.
	// Default is "desc".
	Direction string `url:"direction,omitempty"`

	// Since filters issues by time.
	Since time.Time `url:"since,omitempty"`

	ListOptions
}

// List gets all issues for authenticated user
func (c *IssueClient) List(ctx context.Context, opts *IssueListOptions) ([]schema.Issue, error) {
	m, err := addOptions(c.method, opts)
	println(m)
	if err != nil {
		return nil, err
	}
	var issues []schema.Issue
	err = c.client.DoWithContext(ctx, http.MethodGet, m, nil, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
