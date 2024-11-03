package client

import (
	"context"
	"net/http"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/ram02z/gh-weekly-log/schema"
)

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
