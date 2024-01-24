package postgresql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/enercity/be-service-sample/pkg/model/domain"
	"gorm.io/gorm"
)

type ContractInformation struct {
	db *gorm.DB
}

func newContractInformation(db *gorm.DB) *ContractInformation {
	return &ContractInformation{
		db: db,
	}
}

func (c *ContractInformation) CountAll() (int64, error) {
	var count int64

	err := c.db.Model(&ContractInformation{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, err
}

func (ci *ContractInformation) CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error) {
	var count int64

	stmt := ci.db.Model(&ContractInformation{}).
		Limit(limit).
		Offset(offset)

	applyFilters(stmt, filters)

	err := stmt.Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func (ci *ContractInformation) Find(filters map[string]interface{}, limit, offset int) ([]*domain.ContractInformation, error) {
	var contractinformations []*domain.ContractInformation

	stmt := ci.db.Model(&ContractInformation{}).
		Limit(limit).
		Offset(offset).
		Order("contractinformation.created_at ASC")

	applyFilters(stmt, filters)

	err := stmt.Find(&contractinformations).Error

	return contractinformations, err
}

func applyFiltersToContractInformation(stmt *gorm.DB, filters map[string]interface{}) {
	// TODO
	applyLikeFilterContractInformation(stmt, filters, "first_name")
	applyLikeFilterContractInformation(stmt, filters, "last_name")
	applyLikeFilterContractInformation(stmt, filters, "email")
}

func applyLikeFilterContractInformation(stmt *gorm.DB, filters map[string]interface{}, key string) {
	if v, ok := filters[key]; ok {
		val := "%" + strings.ToLower(v.(string)) + "%"
		stmt = stmt.Where(fmt.Sprintf("LOWER(contractinformation.%s) LIKE ?", key), val)
	}
}

func (ci *ContractInformation) Load(limit, offset int) ([]*domain.ContractInformation, error) {
	var contractInformations []*domain.ContractInformation

	err := ci.db.Limit(limit).Offset(offset).Find(&contractInformations).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return contractInformations, err
}

func (ci *ContractInformation) Save(contractinformation *domain.ContractInformation) error {
	return ci.db.Create(contractinformation).Error
}
