package repository

import "github.com/enercity/be-service-sample/pkg/model/domain"

type ContractInformation interface {
	Find(filters map[string]interface{}, limit, offset int) ([]*domain.ContractInformation, error)
	CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error)
	Load(limit, offset int) ([]*domain.ContractInformation, error)
	CountAll() (int64, error)
	Save(customer *domain.ContractInformation) error
}
