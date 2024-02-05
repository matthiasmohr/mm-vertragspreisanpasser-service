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

func (g *Getter) Get(ctx context.Context, logEntry logger.Entry, req *dto.GetPriceChangeRuleCollectionRequest) (*dto.GetPriceChangeRuleCollectionResponse, error) {
	logEntry.WithField("req", req).Debug("about to fetch a price change RuleCollection")

	res := &dto.GetPriceChangeRuleCollectionResponse{}

	// Get the respective priceChangeRuleCollection
	uuid, err := domain.ParseUUID(req.Id)
	if err != nil {
		logEntry.WithError(err).Warning("unable to convert to UUID")
		return nil, usecase.ErrDomainInternal
	}

	priceChangeRuleCollections, err := g.store.PriceChangeRuleCollection().FindByIDs(uuid)
	if err != nil || len(priceChangeRuleCollections) != 1 {
		logEntry.WithContext(ctx).WithError(err).Error("unable to fetch price change RuleCollection in a db")

		return nil, usecase.ErrDatabaseInternal
	}

	res.RuleCollection = dto.PriceChangeRuleCollectionFromDomain(priceChangeRuleCollections[0])

	// Get the nested price Change Rules
	priceChangeRules, err := g.store.PriceChangeRule().FindByFindByPriceChangeRuleCollectionId(uuid)
	if err != nil {
		logEntry.WithContext(ctx).WithError(err).Error("unable to fetch price change Rules nested in the price change rule collection")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Rules = make([]dto.PriceChangeRule, 0, len(priceChangeRules))
	for _, c := range priceChangeRules {
		res.Rules = append(res.Rules, dto.PriceChangeRuleFromDomain(c))
	}

	return res, nil
}
