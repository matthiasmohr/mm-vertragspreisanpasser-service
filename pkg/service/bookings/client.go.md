package bookings

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/enercity/paymentallocation/pkg/model/domain"
	"github.com/enercity/paymentallocation/pkg/service"
	"github.com/enercity/paymentallocation/pkg/service/bookings/model"
	platformhttp "github.com/enercity/paymentallocation/pkg/service/http"
	httperror "github.com/enercity/paymentallocation/pkg/service/http/errors"
	logger "github.com/enercity/lib-logger/v3"
)

const (
	bookPaymentsEndpoint = "bookings/v2/bookings/allocate-payment"

	// HTTPMaxRetries represents the maximum allowed amount of additional request retries.
	HTTPMaxRetries = 1
	// HTTPRetryWaitTime is the wait time to sleep before retrying request.
	HTTPRetryWaitTime = time.Millisecond * 100
	// HTTPRetryMaxWaitTime is the max wait time to sleep before retrying request.
	HTTPRetryMaxWaitTime = time.Second
	// HTTPRetryTimeout is the max allowed time to wait for the request to finish.
	// This timeout has been increased to handle bigger bookings payload ED4FTR-12428.
	HTTPRetryTimeout = time.Second * 10
)

// Client allows HTTP communication with a bookings adapter server.
type Client struct {
	address     string
	credentials Credentials
	httpClient  service.HTTPClient
	logger      logger.Logger
}

// NewClient returns a new Client.
func NewClient(config Config, client service.HTTPClient, logger logger.Logger) *Client {
	c := &Client{
		address:     config.Host,
		credentials: config.Credentials,
		logger:      logger,
	}

	client.Apply(
		platformhttp.WithBasicAuth(c.credentials.User, c.credentials.Password),
		platformhttp.JSONClient(),
	)

	c.httpClient = client

	return c
}

// BookPayments makes a request to bookings adapter service.
func (c *Client) BookPayments(ctx context.Context, bookings []*domain.Booking) error {
	req := model.BookingRequest{
		Items: make([]model.Booking, len(bookings)),
	}

	for i, b := range bookings {
		req.Items[i] = model.Booking{
			ID:          b.ID.String(),
			Amount:      b.Payment.Amount,
			BookingType: b.PaymentState.PaymentStatus.String(),
			BookingDate: b.Date.Time,
		}
	}

	resp, err := c.httpClient.Post(ctx, c.url(), req)
	if err != nil {
		return httperror.NewHTTPRequestError(err)
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return httperror.NewHTTPError(
		errors.New("the request to book payments was not successful"),
		resp.Body(),
		resp.StatusCode(),
	)
}

func (c *Client) url() string {
	var glue string
	if !strings.HasSuffix(c.address, "/") {
		glue = "/"
	}

	return fmt.Sprintf("%s%s%s", c.address, glue, bookPaymentsEndpoint)
}
