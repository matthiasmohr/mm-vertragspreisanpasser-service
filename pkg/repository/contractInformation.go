package repository

import "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"

type ContractInformation interface {
	Find(filters map[string]interface{}, limit, offset int) ([]*domain.ContractInformation, error)
	CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error)
	Load(limit, offset int) ([]*domain.ContractInformation, error)
	CountAll() (int64, error)
	FindByID(id domain.UUID) (*domain.ContractInformation, error)

	Save(contractInformation *domain.ContractInformation) error
}
