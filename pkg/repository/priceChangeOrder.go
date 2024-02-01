package repository

import "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"

type PriceChangeOrder interface {
	Find(filters map[string]interface{}, limit, offset int) ([]*domain.PriceChangeOrder, error)
	CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error)
	Load(limit, offset int) ([]*domain.PriceChangeOrder, error)
	FindByIDs(ids ...domain.UUID) ([]*domain.PriceChangeOrder, error)
	CountAll() (int64, error)
	Save(priceChangeOrder *domain.PriceChangeOrder) error
	Update(priceChangeOrder *domain.PriceChangeOrder) error
}
