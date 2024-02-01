package priceChangeOrder

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
	ctx context.Context, logEntry logger.Entry, req *dto.FindPriceChangeOrderRequest,
) (*dto.FindPriceChangeOrderResponse, error) {
	logEntry.WithField("req", req).Debug("about search for price change order from the db by given request data")

	res := &dto.FindPriceChangeOrderResponse{}

	total, err := f.store.PriceChangeOrder().CountAllWithFilters(req.Map(), req.Pagination.Limit, req.Pagination.Offset)

	if err != nil {
		logEntry.WithError(err).Warning("unable to retrieve total price Change Orders")

		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Pagination.Limit
	res.Pagination.Offset = req.Pagination.Offset

	priceChangeOrders, err := f.store.PriceChangeOrder().Find(req.Map(), req.Pagination.Limit, req.Pagination.Offset)
	if err != nil {
		logEntry.WithError(err).Warning("unable to retrieve database information")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.PriceChangeOrder, 0, len(priceChangeOrders))
	for _, c := range priceChangeOrders {
		res.Items = append(res.Items, dto.PriceChangeOrderFromDomain(c))
	}

	return res, nil
}
