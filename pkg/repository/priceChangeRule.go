package repository

import "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"

type PriceChangeRule interface {
	Find(filters map[string]interface{}, limit, offset int) ([]*domain.PriceChangeRule, error)
	CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error)
	Load(limit, offset int) ([]*domain.PriceChangeRule, error)
	FindByIDs(ids ...domain.UUID) ([]*domain.PriceChangeRule, error)
	CountAll() (int64, error)
	Save(priceChangeRule *domain.PriceChangeRule) error
	Update(priceChangeRule *domain.PriceChangeRule) error
}
