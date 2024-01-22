package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	logger "github.com/enercity/lib-logger/v3"
	resty "github.com/go-resty/resty/v2"
)

const (
	acceptHeader      = "Accept"
	contentTypeHeader = "Content-Type"
	applicationJSON   = "application/json"
	applicationForm   = "application/x-www-form-urlencoded"
	retryMaxRetries   = 4
	retryWaitTime     = 1000 * time.Millisecond
	retryMaxWaitTime  = 2000 * time.Millisecond
	retryTimeout      = 5000 * time.Millisecond
)

// Client is a configurable HTTP client allowing retries on unsuccessful requests.
type Client struct {
	options               []ClientOption
	retryer               *resty.Client
	beforeRequest         []func(*request) error
	afterRequest          []func(*Response, *request) (*Response, error)
	logger                logger.Logger
	isWithAuthTokenCalled bool
}

// ClientOption represents a functional option used to initialize a Client.
type ClientOption func(*Client)

type tokenHolder interface {
	InvalidateToken()
	RawToken() string
	RefreshToken() error
}

// NewClient initializes a new Client with the provided functional Client options.
// If no options are passed to the constructor, requests will not be retried.
func NewClient(logger logger.Logger, options ...ClientOption) *Client {
	client := &Client{
		retryer: resty.New(),
		options: options,
		logger:  logger,
	}

	for _, option := range options {
		option(client)
	}

	client.withTraceLogging()

	return client
}

// DefaultClient initializes a new Client with the default retryer config values.
func DefaultClient(logger logger.Logger) *Client {
	return NewClient(logger, ConfigureRetryer(DefaultRetryerConfig()))
}

// NewRetryerConfig returns an initialized RetryerConfig.
func NewRetryerConfig(maxRetries, waitTimeInMillis, maxWaitTimeInMillis, timeoutInMillis uint) *RetryerConfig {
	return &RetryerConfig{
		maxRetries:  maxRetries,
		waitTime:    time.Duration(waitTimeInMillis) * time.Millisecond,
		maxWaitTime: time.Duration(maxWaitTimeInMillis) * time.Millisecond,
		timeout:     time.Duration(timeoutInMillis) * time.Millisecond,
	}
}

// DefaultRetryerConfig sets default retryer config options.
func DefaultRetryerConfig() *RetryerConfig {
	return &RetryerConfig{
		maxRetries:  retryMaxRetries,
		waitTime:    retryWaitTime,
		maxWaitTime: retryMaxWaitTime,
		timeout:     retryTimeout,
	}
}

// Apply applies options to the client.
func (c *Client) Apply(options ...ClientOption) {
	options = append(c.options, options...)
	*c = *NewClient(c.logger, options...)
}

// Get performs an HTTP GET request.
func (c *Client) Get(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return c.newRequest(ctx, url, http.MethodGet, nil, options...).execute()
}

// Post performs an HTTP POST request.
func (c *Client) Post(
	ctx context.Context, url string, body interface{}, options ...RequestOption,
) (*Response, error) {
	return c.newRequest(ctx, url, http.MethodPost, body, options...).execute()
}

// Patch performs an HTTP PATCH request.
func (c *Client) Patch(
	ctx context.Context, url string, body interface{}, options ...RequestOption,
) (*Response, error) {
	return c.newRequest(ctx, url, http.MethodPatch, body, options...).execute()
}

// Put performs an HTTP PUT request.
func (c *Client) Put(
	ctx context.Context, url string, body interface{}, options ...RequestOption,
) (*Response, error) {
	return c.newRequest(ctx, url, http.MethodPut, body, options...).execute()
}

// Delete performs an HTTP DELETE request.
func (c *Client) Delete(
	ctx context.Context, url string, body interface{}, options ...RequestOption,
) (*Response, error) {
	return c.newRequest(ctx, url, http.MethodDelete, body, options...).execute()
}

// ConfigureRetryer configures the Client's retryer.
func ConfigureRetryer(config *RetryerConfig) ClientOption {
	return func(c *Client) {
		if config == nil {
			return
		}

		retryStatuses := []int{
			http.StatusRequestTimeout,
			http.StatusInternalServerError,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
		}

		c.retryer.
			EnableTrace().
			SetRetryWaitTime(config.waitTime).
			SetRetryCount(int(config.maxRetries)).
			SetRetryMaxWaitTime(config.maxWaitTime).
			SetTimeout(config.timeout).
			AddRetryCondition(
				func(r *resty.Response, err error) bool {
					for _, status := range retryStatuses {
						if r.StatusCode() == status {
							return true
						}
					}

					return false
				},
			)
	}
}

// JSONClient ensures that requests will be expected to send and receive JSON content, based on their headers.
// JSONClient sets the accept and contentType headers.
func JSONClient() ClientOption {
	return func(c *Client) {
		c.retryer.
			SetHeader(acceptHeader, applicationJSON).
			SetHeader(contentTypeHeader, applicationJSON)
	}
}

// FormClient sets the accept and contentType headers.
func FormClient() ClientOption {
	return func(c *Client) {
		c.retryer.
			SetHeader(acceptHeader, applicationJSON).
			SetHeader(contentTypeHeader, applicationForm)
	}
}

// WithAuthToken configures the client in a way that allows authorization tokens to be set
// before every request. The provided tokenHolder implements the necessary token handling operations.
// If a request results in an "uanuthorized" response, the client refreshes the token and attempts
// to perform the request one more time.
// Requests with the `skipAuthToken` flag set to true will skip the token validation & fetching process,
// and won't include the token in the request.
func WithAuthToken(holder tokenHolder) ClientOption {
	return func(c *Client) {
		if !c.isWithAuthTokenCalled {
			c.withAuthToken(holder)
		}
	}
}

// WithTraceLogging traces outgoing requests and responses.
func WithTraceLogging() ClientOption {
	return func(c *Client) {
		c.withTraceLogging()
	}
}

func (c *Client) withTraceLogging() {
	after := func(response *Response, request *request) (*Response, error) {
		traceInfo := request.baseRequest.TraceInfo()

		c.logger.WithFields(map[string]interface{}{
			"response_time": traceInfo.ResponseTime,
			"total_time":    traceInfo.TotalTime,
			"url":           request.url,
			"method":        request.method,
			"status_code":   response.StatusCode(),
		}).Debug("response")

		return response, nil
	}
	before := func(request *request) error {
		ctx := request.baseRequest.Context()
		c.logger.WithContext(ctx).WithFields(map[string]interface{}{
			"url":    request.url,
			"method": request.method,
		}).Debug("outgoing request")

		return nil
	}

	c.afterRequest = append(c.afterRequest, after)
	c.beforeRequest = append(c.beforeRequest, before)
}

func (c *Client) withAuthToken(holder tokenHolder) {
	before := func(request *request) error {
		if request.skipAuthToken {
			return nil
		}

		if err := holder.RefreshToken(); err != nil {
			return fmt.Errorf("refresh token: %w", err)
		}

		request.setAuthToken(holder.RawToken())

		return nil
	}
	c.beforeRequest = append(c.beforeRequest, before)

	after := func(response *Response, request *request) (*Response, error) {
		if request.skipAuthToken {
			return response, nil
		}

		if response.StatusCode() != http.StatusUnauthorized {
			return response, nil
		}

		holder.InvalidateToken()

		if !request.shouldRepeat {
			return response, nil
		}

		request.shouldRepeat = false

		return request.execute()
	}
	c.afterRequest = append(c.afterRequest, after)
	c.isWithAuthTokenCalled = true
}

func WithBasicAuth(username string, password string) ClientOption {
	return func(c *Client) {
		c.withBasicAuth(username, password)
	}
}

func (c *Client) withBasicAuth(username string, password string) {
	before := func(request *request) error {
		if username == "" || password == "" {
			return errors.New("username and/or password are/is empty")
		}

		request.baseRequest.SetBasicAuth(username, password)

		return nil
	}
	c.beforeRequest = append(c.beforeRequest, before)
}

// WithBasicAuthToken configures the client in a way that allows authorization tokens to be set
// before every request, by using basic auth to obtain them.
func WithBasicAuthToken(token string) ClientOption {
	return func(c *Client) {
		if !c.isWithAuthTokenCalled {
			c.withBasicAuthToken(token)
		}
	}
}

func (c *Client) withBasicAuthToken(token string) {
	before := func(request *request) error {
		if request.skipAuthToken {
			return nil
		}

		request.baseRequest.SetHeader("Authorization", "Basic "+token)

		return nil
	}
	c.beforeRequest = append(c.beforeRequest, before)

	after := func(response *Response, request *request) (*Response, error) {
		if request.skipAuthToken {
			return response, nil
		}

		if response.StatusCode() != http.StatusUnauthorized {
			return response, nil
		}

		if !request.shouldRepeat {
			return response, nil
		}

		request.shouldRepeat = false

		return request.execute()
	}
	c.afterRequest = append(c.afterRequest, after)
	c.isWithAuthTokenCalled = true
}
