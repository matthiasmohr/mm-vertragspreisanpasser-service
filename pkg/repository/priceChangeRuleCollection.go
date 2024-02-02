package repository

import "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"

type PriceChangeRuleCollection interface {
	Find(filters map[string]interface{}, limit, offset int) ([]*domain.PriceChangeRuleCollection, error)
	CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error)
	Load(limit, offset int) ([]*domain.PriceChangeRuleCollection, error)
	FindByIDs(ids ...domain.UUID) ([]*domain.PriceChangeRuleCollection, error)
	CountAll() (int64, error)
	Save(priceChangeRuleCollection *domain.PriceChangeRuleCollection) error
	Update(priceChangeRuleCollection *domain.PriceChangeRuleCollection) error
}
