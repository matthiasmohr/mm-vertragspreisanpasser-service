package priceChangeExecution

import (
	"context"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Lister struct {
	store repository.Store
}

func NewLister(
	store repository.Store,
) *Lister {
	return &Lister{
		store: store,
	}
}

func (l *Lister) List(ctx context.Context, logEntry logger.Entry, req *dto.ListPriceChangeExecutionRequest) (*dto.ListPriceChangeExecutionResponse, error) {
	logEntry.Debug("about load price change Execution from the db")

	res := &dto.ListPriceChangeExecutionResponse{}

	total, err := l.store.PriceChangeExecution().CountAll()
	if err != nil {
		logEntry.Warning("unable to retrieve total price change Executions")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Limit
	res.Pagination.Offset = req.Offset

	priceChangeExecutions, err := l.store.PriceChangeExecution().Load(req.Limit, req.Offset)
	if err != nil {
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.PriceChangeExecution, 0, len(priceChangeExecutions))
	for _, c := range priceChangeExecutions {
		res.Items = append(res.Items, dto.PriceChangeExecutionFromDomain(c))
	}

	return res, nil
}
