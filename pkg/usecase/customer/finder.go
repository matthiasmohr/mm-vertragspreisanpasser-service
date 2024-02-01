package customer

import (
	"context"

	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Finder struct {
	store repository.Store
}

func NewFinder(
	store repository.Store,
) *Finder {
	return &Finder{
		store: store,
	}
}

func (f *Finder) Find(
	ctx context.Context, logEntry logger.Entry, req *dto.CustomerFindRequest,
) (*dto.ListCustomersResponse, error) {
	logEntry.WithField("req", req).Debug("about search for customers from the db by given request data")

	res := &dto.ListCustomersResponse{}

	total, err := f.store.Customer().CountAllWithFilters(req.Map(), req.Pagination.Limit, req.Pagination.Offset)
	if err != nil {
		logEntry.WithError(err).Warning("unable to retrieve total customers")

		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Pagination.Limit
	res.Pagination.Offset = req.Pagination.Offset

	customers, err := f.store.Customer().Find(req.Map(), req.Pagination.Limit, req.Pagination.Offset)
	if err != nil {
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.Customer, 0, len(customers))
	for _, c := range customers {
		res.Items = append(res.Items, dto.CustomerFromDomain(c))
	}

	return res, nil
}
