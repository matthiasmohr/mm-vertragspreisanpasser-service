package repository

import "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"

type Customer interface {
	Find(filters map[string]interface{}, limit, offset int) ([]*domain.Customer, error)
	CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error)
	Load(limit, offset int) ([]*domain.Customer, error)
	CountAllCustomers() (int64, error)
	Save(customer *domain.Customer) error
}
