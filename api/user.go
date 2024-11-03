package api

import (
	"context"
	"github.com/ram02z/gh-weekly-log/api/schema"
	"net/http"

	"github.com/cli/go-gh/v2/pkg/api"
)

// UserClient represents api to users REST API endpoints
// See https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28
type UserClient client

func DefaultUserClient() (*UserClient, error) {
	return NewUserClient(api.ClientOptions{})
}

func NewUserClient(opts api.ClientOptions) (*UserClient, error) {
	restClient, err := api.NewRESTClient(opts)
	if err != nil {
		return nil, err
	}

	return &UserClient{
		client: restClient,
		method: "user",
	}, nil
}

// Get retrieves the currently authenticated user
func (c *UserClient) Get(ctx context.Context) (*schema.User, error) {
	user := &schema.User{}
	err := c.client.DoWithContext(ctx, http.MethodGet, c.method, nil, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
