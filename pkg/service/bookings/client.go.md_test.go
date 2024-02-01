package bookings

import (
	"context"
	stdhttp "net/http"
	"testing"
	"time"

	"github.com/enercity/paymentallocation/pkg/model/domain"
	"github.com/enercity/paymentallocation/pkg/service/bookings/model"
	"github.com/enercity/paymentallocation/pkg/service/http"
	httperror "github.com/enercity/paymentallocation/pkg/service/http/errors"
	httpmock "github.com/enercity/paymentallocation/pkg/service/http/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	user     = "user"
	password = "passpass"
)

func TestClient_BookPayments(t *testing.T) {
	assert := assert.New(t)

	assertRequestOptionsLength := func(args mock.Arguments) {
		assert.Len(args.Get(3).([]http.RequestOption), 0)
	}

	now, err := time.Parse("2006-01-02", time.Now().UTC().Format("2006-01-02"))
	assert.NoError(err)

	bookingDate := &domain.ISODate{Time: now}

	bookings := []*domain.Booking{
		{
			ID: domain.UUID(uuid.New().String()),
			Payment: &domain.Payment{
				Amount: 12.60,
			},
			PaymentState: &domain.PaymentState{
				PaymentStatus: domain.PaymentStatusDuplicate,
			},
			Date: bookingDate,
		},
		{
			ID: domain.UUID(uuid.New().String()),
			Payment: &domain.Payment{
				Amount: 15.65,
			},
			PaymentState: &domain.PaymentState{
				PaymentStatus: domain.PaymentStatusISUAllocated,
			},
			Date: bookingDate,
		},
	}

	t.Run("booking request successfully send", func(t *testing.T) {
		rawResponse := &stdhttp.Response{
			StatusCode: stdhttp.StatusOK,
		}
		response := http.NewResponse(nil, rawResponse)
		requestBody := model.BookingRequest{
			Items: []model.Booking{
				{
					ID:          bookings[0].ID.String(),
					Amount:      bookings[0].Payment.Amount,
					BookingType: "duplicate",
					BookingDate: bookings[0].Date.Time,
				},
				{
					ID:          bookings[1].ID.String(),
					Amount:      bookings[1].Payment.Amount,
					BookingType: "IS-U-allocated",
					BookingDate: bookings[1].Date.Time,
				},
			},
		}
		address := "bookings.test"
		client, httpClient := setupClient(address, assert)
		httpClient.On(
			"Post",
			context.Background(),
			address+"/bookings/v2/bookings/allocate-payment",
			requestBody,
			mock.Anything,
		).Return(response, nil).Run(assertRequestOptionsLength)

		err := client.BookPayments(context.Background(), bookings)
		assert.NoError(err)
		httpClient.AssertExpectations(t)
	})

	t.Run("booking request fails", func(t *testing.T) {
		rawResponse := &stdhttp.Response{
			StatusCode: stdhttp.StatusBadRequest,
		}
		response := http.NewResponse(nil, rawResponse)
		requestBody := model.BookingRequest{
			Items: []model.Booking{
				{
					ID:          bookings[0].ID.String(),
					Amount:      bookings[0].Payment.Amount,
					BookingType: "duplicate",
					BookingDate: bookings[0].Date.Time,
				},
				{
					ID:          bookings[1].ID.String(),
					Amount:      bookings[1].Payment.Amount,
					BookingType: "IS-U-allocated",
					BookingDate: bookings[1].Date.Time,
				},
			},
		}
		address := "bookings.test"
		client, httpClient := setupClient(address, assert)
		httpClient.On(
			"Post",
			context.Background(),
			address+"/bookings/v2/bookings/allocate-payment",
			requestBody,
			mock.Anything,
		).Return(response, nil).Run(assertRequestOptionsLength)

		err := client.BookPayments(context.Background(), bookings)
		assert.Error(err)
		httpError, ok := err.(*httperror.HTTPError)
		assert.True(ok)
		assert.Equal(stdhttp.StatusBadRequest, httpError.StatusCode)
		httpClient.AssertExpectations(t)
	})

	t.Run("validate url consistence ", func(t *testing.T) {
		client, _ := setupClient("bookings.test", assert)
		urlToBeCalled := client.url()
		assert.Equal("bookings.test/bookings/v2/bookings/allocate-payment", urlToBeCalled)

		client, _ = setupClient("bookings.test/", assert)
		urlToBeCalled = client.url()
		assert.Equal("bookings.test/bookings/v2/bookings/allocate-payment", urlToBeCalled)
	})
}

func setupClient(address string, assert *assert.Assertions) (*Client, *httpmock.Client) {
	config := Config{
		Host: address,
		Credentials: Credentials{
			User:     user,
			Password: password,
		},
	}
	httpClient := &httpmock.Client{}
	httpClient.On("Apply", mock.AnythingOfType("[]http.ClientOption")).
		Return().Run(func(args mock.Arguments) {
		assert.Len(args.Get(0).([]http.ClientOption), 2)
	})

	client := NewClient(Config{
		Host:        config.Host,
		Credentials: config.Credentials,
	}, httpClient, nil)

	return client, httpClient
}
