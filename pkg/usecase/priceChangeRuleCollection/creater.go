package priceChangeRuleCollection

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

func (c *Creator) Create(ctx context.Context, logEntry logger.Entry, req *dto.CreatePriceChangeRuleCollectionRequest) error {
	logEntry.WithField("req", req).Debug("about to create a new price change RuleCollection")

	pap := "DUMMY" // TODO: REPLACE BY REAL PRICE ADJUSTMENT PROJECT

	priceChangeRuleCollection, err := domain.NewPriceChangeRuleCollection(
		pap,
	)
	if err != nil {
		logEntry.WithError(err).Error("unable to create a price change RuleCollection")
		return usecase.ErrDomainInternal
	}

	err = c.store.PriceChangeRuleCollection().Save(priceChangeRuleCollection)
	if err != nil {
		logEntry.WithContext(ctx).WithError(err).Error("unable to store price change RuleCollection in a db")

		return usecase.ErrDatabaseInternal
	}

	return nil
}
