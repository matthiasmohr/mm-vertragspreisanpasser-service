package priceChangeOrder

import (
	"context"
	"github.com/enercity/be-service-sample/pkg/model/domain"
	"github.com/enercity/be-service-sample/pkg/model/dto"
	"github.com/enercity/be-service-sample/pkg/repository"
	"github.com/enercity/be-service-sample/pkg/usecase"
	logger "github.com/enercity/lib-logger/v3"
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

func (c *Creator) Create(ctx context.Context, logEntry logger.Entry, req *dto.CreatePriceChangeOrderRequest) error {
	logEntry.WithField("req", req).Debug("about to create a new price change order")

	priceChangeOrder, err := domain.NewPriceChangeOrder(
		"manual",
		req.ProductSerialNumber,
		req.PriceValidSince,
		req.CurrentBaseCosts,
		req.CurrentKwhCosts,
		req.CurrentBaseMargin,
		req.CurrentKwhMargin,
		req.CurrentBasePriceNet,
		req.CurrentKwhPriceNet,
		req.AnnualConsumption,
		req.PriceValidAsOf,
		req.FutureBaseCosts,
		req.FutureKwhCosts,
		req.FutureKwhMargin,
		req.FutureBaseMargin,
		req.FutureBasePriceNet,
		req.FutureKwhPriceNet,
		req.AgentHintFlag,
		req.AgentHintText,
		req.CommunicationFlag,
		req.CommunictionTime,
	)
	if err != nil {
		logEntry.WithError(err).Error("unable to create a price change order")
		return usecase.ErrDomainInternal
	}

	err = c.store.PriceChangeOrder().Save(priceChangeOrder)
	if err != nil {
		logEntry.WithContext(ctx).WithError(err).Error("unable to store price change order in a db")

		return usecase.ErrDatabaseInternal
	}

	return nil
}
