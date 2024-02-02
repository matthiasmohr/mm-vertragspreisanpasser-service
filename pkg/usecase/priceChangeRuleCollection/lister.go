package priceChangeRuleCollection

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

func (l *Lister) List(ctx context.Context, logEntry logger.Entry, req *dto.ListPriceChangeRuleCollectionRequest) (*dto.ListPriceChangeRuleCollectionResponse, error) {
	logEntry.Debug("about load price change Rule Collection from the db")

	res := &dto.ListPriceChangeRuleCollectionResponse{}

	total, err := l.store.PriceChangeRuleCollection().CountAll()
	if err != nil {
		logEntry.Warning("unable to retrieve total price change Rule Collections")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Limit
	res.Pagination.Offset = req.Offset

	priceChangeRuleCollections, err := l.store.PriceChangeRuleCollection().Load(req.Limit, req.Offset)
	if err != nil {
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.PriceChangeRuleCollection, 0, len(priceChangeRuleCollections))
	for _, c := range priceChangeRuleCollections {
		res.Items = append(res.Items, dto.PriceChangeRuleCollectionFromDomain(c))
	}

	return res, nil
}
