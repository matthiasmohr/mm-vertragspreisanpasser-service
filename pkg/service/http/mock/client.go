package mock

import (
	"context"

	"github.com/enercity/be-service-sample/pkg/service/http"
	"github.com/stretchr/testify/mock"
)

// Client mocks an http.Client object.
type Client struct {
	mock.Mock
}

func (c *Client) Apply(options ...http.ClientOption) {
	c.Called(options)
}

func (c *Client) Get(ctx context.Context, url string, options ...http.RequestOption) (*http.Response, error) {
	args := c.Called(ctx, url, options)
	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	}

	return nil, args.Error(1)
}

func (c *Client) Post(
	ctx context.Context, url string, body interface{}, options ...http.RequestOption,
) (*http.Response, error) {
	args := c.Called(ctx, url, body, options)
	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	}

	return nil, args.Error(1)
}

func (c *Client) Patch(
	ctx context.Context, url string, body interface{}, options ...http.RequestOption,
) (*http.Response, error) {
	args := c.Called(ctx, url, body, options)
	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	}

	return nil, args.Error(1)
}

func (c *Client) Put(
	ctx context.Context, url string, body interface{}, options ...http.RequestOption,
) (*http.Response, error) {
	args := c.Called(ctx, url, body, options)
	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	}

	return nil, args.Error(1)
}

func (c *Client) Delete(
	ctx context.Context, url string, body interface{}, options ...http.RequestOption,
) (*http.Response, error) {
	args := c.Called(ctx, url, body, options)
	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	}

	return nil, args.Error(1)
}
