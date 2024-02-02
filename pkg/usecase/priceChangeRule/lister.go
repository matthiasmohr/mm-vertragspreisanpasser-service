package priceChangeRule

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

func (l *Lister) List(ctx context.Context, logEntry logger.Entry, req *dto.ListPriceChangeRuleRequest) (*dto.ListPriceChangeRuleResponse, error) {
	logEntry.Debug("about load price change Rule from the db")

	res := &dto.ListPriceChangeRuleResponse{}

	total, err := l.store.PriceChangeRule().CountAll()
	if err != nil {
		logEntry.Warning("unable to retrieve total price change Rules")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Limit
	res.Pagination.Offset = req.Offset

	priceChangeRules, err := l.store.PriceChangeRule().Load(req.Limit, req.Offset)
	if err != nil {
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.PriceChangeRule, 0, len(priceChangeRules))
	for _, c := range priceChangeRules {
		res.Items = append(res.Items, dto.PriceChangeRuleFromDomain(c))
	}

	return res, nil
}
