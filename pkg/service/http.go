package service

import (
	"context"

	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/service/http"
)

// HTTPClient is a custom HTTP client used for communication with external services.
type HTTPClient interface {
	Apply(options ...http.ClientOption)
	Get(ctx context.Context, url string, options ...http.RequestOption) (*http.Response, error)
	Put(ctx context.Context, url string, body interface{}, options ...http.RequestOption) (*http.Response, error)
	Post(ctx context.Context, url string, body interface{}, options ...http.RequestOption) (*http.Response, error)
	Patch(ctx context.Context, url string, body interface{}, options ...http.RequestOption) (*http.Response, error)
	Delete(ctx context.Context, url string, body interface{}, options ...http.RequestOption) (*http.Response, error)
}
