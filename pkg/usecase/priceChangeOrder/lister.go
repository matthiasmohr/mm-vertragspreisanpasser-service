package priceChangeOrder

import (
	"context"
	"github.com/enercity/be-service-sample/pkg/model/dto"
	"github.com/enercity/be-service-sample/pkg/repository"
	"github.com/enercity/be-service-sample/pkg/usecase"
	logger "github.com/enercity/lib-logger/v3"
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

func (l *Lister) List(ctx context.Context, logEntry logger.Entry, req *dto.ListPriceChangeOrderRequest) (*dto.ListPriceChangeOrderResponse, error) {
	logEntry.Debug("about load price change order from the db")

	res := &dto.ListPriceChangeOrderResponse{}

	total, err := l.store.PriceChangeOrder().CountAll()
	if err != nil {
		logEntry.Warning("unable to retrieve total price change orders")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Limit
	res.Pagination.Offset = req.Offset

	priceChangeOrders, err := l.store.PriceChangeOrder().Load(req.Limit, req.Offset)
	if err != nil {
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.PriceChangeOrder, 0, len(priceChangeOrders))
	for _, c := range priceChangeOrders {
		res.Items = append(res.Items, dto.PriceChangeOrderFromDomain(c))
	}

	return res, nil
}
