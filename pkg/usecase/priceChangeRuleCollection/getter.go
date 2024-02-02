package priceChangeRuleCollection

import (
	"context"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Getter struct {
	store repository.Store
}

func NewGetter(
	store repository.Store,
) *Getter {
	return &Getter{
		store: store,
	}
}

func (c *Creator) Get(ctx context.Context, logEntry logger.Entry, req *dto.GetPriceChangeRuleCollectionRequest) (*dto.GetPriceChangeRuleCollectionResponse, error) {
	logEntry.WithField("req", req).Debug("about to fetch a price change RuleCollection")

	res := &dto.GetPriceChangeRuleCollectionResponse{}

	// Get the respective priceChangeRuleCollection
	uuid, err := domain.ParseUUID(req.Id)
	if err != nil {
		logEntry.WithError(err).Warning("unable to convert to UUID")
		return nil, usecase.ErrDomainInternal
	}

	priceChangeRuleCollections, err := c.store.PriceChangeRuleCollection().FindByIDs(uuid)
	if err != nil || len(priceChangeRuleCollections) != 1 {
		logEntry.WithContext(ctx).WithError(err).Error("unable to fetch price change RuleCollection in a db")

		return nil, usecase.ErrDatabaseInternal
	}

	res.RuleCollection = dto.PriceChangeRuleCollectionFromDomain(priceChangeRuleCollections[0])

	// Get the nested priceChangeRules for the collection

	return res, nil
}
