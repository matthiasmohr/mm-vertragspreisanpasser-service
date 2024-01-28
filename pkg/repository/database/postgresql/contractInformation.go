package postgresql

import (
	"errors"
	"fmt"
	"regexp"
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
		Offset(offset)

	applyFiltersToContractInformation(stmt, filters)

	err := stmt.Find(&contractinformations).Error

	return contractinformations, err
}

func applyFiltersToContractInformation(stmt *gorm.DB, filters map[string]interface{}) {
	// TODO
	applyLikeFilterContractInformation(stmt, filters, "mba")
	applyLikeFilterContractInformation(stmt, filters, "productSerialNumber")
}

func applyLikeFilterContractInformation(stmt *gorm.DB, filters map[string]interface{}, key string) {
	if v, ok := filters[key]; ok {
		val := "%" + strings.ToLower(v.(string)) + "%"
		stmt = stmt.Where(fmt.Sprintf("LOWER(%s) LIKE ?", toSnakeCase(key)), val)
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
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
