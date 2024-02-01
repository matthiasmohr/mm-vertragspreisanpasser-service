package priceChangeExecution

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
	ctx context.Context, logEntry logger.Entry, req *dto.FindPriceChangeExecutionRequest,
) (*dto.FindPriceChangeExecutionResponse, error) {
	logEntry.WithField("req", req).Debug("about search for price change Execution from the db by given request data")

	res := &dto.FindPriceChangeExecutionResponse{}

	total, err := f.store.PriceChangeExecution().CountAllWithFilters(req.Map(), req.Pagination.Limit, req.Pagination.Offset)

	if err != nil {
		logEntry.WithError(err).Warning("unable to retrieve total price Change Executions")

		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Pagination.Limit
	res.Pagination.Offset = req.Pagination.Offset

	priceChangeExecutions, err := f.store.PriceChangeExecution().Find(req.Map(), req.Pagination.Limit, req.Pagination.Offset)
	if err != nil {
		logEntry.WithError(err).Warning("unable to retrieve database information")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.PriceChangeExecution, 0, len(priceChangeExecutions))
	for _, c := range priceChangeExecutions {
		res.Items = append(res.Items, dto.PriceChangeExecutionFromDomain(c))
	}

	return res, nil
}
