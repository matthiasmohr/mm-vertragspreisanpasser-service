package priceChangeRule

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

func (c *Creator) Create(ctx context.Context, logEntry logger.Entry, req *dto.CreatePriceChangeRuleRequest) error {
	logEntry.WithField("req", req).Debug("about to create a new price change Rule")

	pcrcUuid, err := domain.ParseUUID(req.PriceChangeRuleCollectionId)
	pcrc, err := c.store.PriceChangeRuleCollection().FindByIDs(pcrcUuid)
	if err != nil || len(pcrc) != 1 {
		logEntry.WithError(err).Warning("unable to retrieve price change rule collection")
		return usecase.ErrDatabaseNotFound
	}

	priceChangeRule, err := domain.NewPriceChangeRule(
		pcrc[0],

		req.RestoreMarginAtSignup,
		req.ChangeBasePriceNetToAmount,
		req.ChangeKwhPriceNetToAmount,
		req.ChangeBasePriceNetByAmount,
		req.ChangeKwhPriceNetByAmount,
		req.ChangeBasePriceNetByFactor,
		req.ChangeKwhPriceNetByFactor,

		req.ValidForProductName,
		req.ValidForCommodity,
		req.ExcludeOrderDateFrom,
		req.ExcludeStartDateFrom,
		req.ExcludeEndDateUntil,
		req.ExcludeLastPriceChangeSince,

		req.LimitToCataloguePriceNet,
		req.LimitToUpperBasePriceNet,
		req.LimitToUpperKwhPriceNet,
		req.LimitToLowerBasePriceNet,
		req.LimitToLowerKwhPriceNet,
		req.LimitToMaxChangeBasePriceNet,
		req.LimitToMaxChangeKwhPriceNet,
		req.LimitToMinChangeBasePriceNet,
		req.LimitToMinChangeKwhPriceNet,
	)
	if err != nil {
		logEntry.WithError(err).Error("unable to create a price change Rule")
		return usecase.ErrDomainInternal
	}

	err = c.store.PriceChangeRule().Save(priceChangeRule)
	if err != nil {
		logEntry.WithContext(ctx).WithError(err).Error("unable to store price change Rule in a db")

		return usecase.ErrDatabaseInternal
	}

	return nil
}
