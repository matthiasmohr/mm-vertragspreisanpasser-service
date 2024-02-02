package priceChangeOrder

import (
	"context"
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

func (c *Creator) Create(ctx context.Context, logEntry logger.Entry, req *dto.CreatePriceChangeOrderRequest) error {
	logEntry.WithField("req", req).Debug("about to create a new price change order")

	// TODO: CHECK IF THERE IS ALREADY AN CREATED PRICE CHANGE ORDER ON THE CONTRACT INFORMATION ID NOT EXECUTED

	priceChangeOrder, err := domain.NewPriceChangeOrder(
		"00000000-0000-0000-0000-000000000000",
		req.ContractInformationId,
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
		req.CommunicationTime,
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
