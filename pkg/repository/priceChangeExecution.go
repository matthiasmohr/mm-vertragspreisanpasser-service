package repository

import "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"

type PriceChangeExecution interface {
	Find(filters map[string]interface{}, limit, offset int) ([]*domain.PriceChangeExecution, error)
	CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error)
	Load(limit, offset int) ([]*domain.PriceChangeExecution, error)
	FindByIDs(ids ...domain.UUID) ([]*domain.PriceChangeExecution, error)
	CountAll() (int64, error)
	Save(priceChangeExecution *domain.PriceChangeExecution) error
}
