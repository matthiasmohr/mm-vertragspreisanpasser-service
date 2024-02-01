package customer

import (
	"context"
	"time"

	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Creator struct {
	store repository.Store
}

func NewCreator(
	store repository.Store,
) *Creator {
	return &Creator{
		store: store,
	}
}

func (c *Creator) Create(ctx context.Context, logEntry logger.Entry, req *dto.CreateCustomerRequest) error {
	logEntry.WithField("req", req).Debug("about to create a new customer")

	customer, err := domain.NewCustomer(time.Now().UTC(), req.FirstName, req.LastName, req.Email)
	if err != nil {
		logEntry.WithError(err).Error("unable to create a customer")

		return usecase.ErrDomainInternal
	}

	err = c.store.Customer().Save(customer)
	if err != nil {
		logEntry.WithContext(ctx).WithError(err).Error("unable to store customer in a db")

		return usecase.ErrDatabaseInternal
	}

	return nil
}
