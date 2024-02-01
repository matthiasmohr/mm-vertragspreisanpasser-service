package customer

import (
	"context"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Loader struct {
	store repository.Store
}

func NewLoader(
	store repository.Store,
) *Loader {
	return &Loader{
		store: store,
	}
}

func (l *Loader) Load(
	ctx context.Context, logEntry logger.Entry, req *dto.ListCustomersRequest,
) (*dto.ListCustomersResponse, error) {
	logEntry.Debug("about load customers from the db")

	res := &dto.ListCustomersResponse{}

	total, err := l.store.Customer().CountAllCustomers()
	if err != nil {
		logEntry.Warning("unable to retrieve total customers")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Limit
	res.Pagination.Offset = req.Offset

	customers, err := l.store.Customer().Load(req.Limit, req.Offset)
	if err != nil {
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.Customer, 0, len(customers))
	for _, c := range customers {
		res.Items = append(res.Items, dto.CustomerFromDomain(c))
	}

	return res, nil
}
