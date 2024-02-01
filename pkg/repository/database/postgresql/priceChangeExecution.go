package postgresql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"gorm.io/gorm"
)

type PriceChangeExecution struct {
	db *gorm.DB
}

func newPriceChangeExecution(db *gorm.DB) *PriceChangeExecution {
	return &PriceChangeExecution{
		db: db,
	}
}

func (cpe *PriceChangeExecution) CountAll() (int64, error) {
	var count int64

	err := cpe.db.Model(&PriceChangeExecution{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, err
}

func (cpe *PriceChangeExecution) CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error) {
	var count int64

	stmt := cpe.db.Model(&PriceChangeExecution{}).
		Limit(limit).
		Offset(offset)

	applyFiltersToPriceChangeExecution(stmt, filters)

	err := stmt.Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func (cpe *PriceChangeExecution) Find(filters map[string]interface{}, limit, offset int) ([]*domain.PriceChangeExecution, error) {
	var pricechangeExecutions []*domain.PriceChangeExecution

	stmt := cpe.db.Model(&PriceChangeExecution{}).
		Limit(limit).
		Offset(offset)

	applyFiltersToPriceChangeExecution(stmt, filters)

	err := stmt.Find(&pricechangeExecutions).Error

	return pricechangeExecutions, err
}

func applyFiltersToPriceChangeExecution(stmt *gorm.DB, filters map[string]interface{}) {
	// TODO
	applyLikeFilterPriceChangeExecution(stmt, filters, "productSerialNumber")
}

func applyLikeFilterPriceChangeExecution(stmt *gorm.DB, filters map[string]interface{}, key string) {
	if v, ok := filters[key]; ok {
		val := "%" + strings.ToLower(v.(string)) + "%"
		stmt = stmt.Where(fmt.Sprintf("LOWER(%s) LIKE ?", toSnakeCase(key)), val)
	}
}

func (cpe *PriceChangeExecution) Load(limit, offset int) ([]*domain.PriceChangeExecution, error) {
	var priceChangeExecutions []*domain.PriceChangeExecution

	err := cpe.db.Limit(limit).Offset(offset).Find(&priceChangeExecutions).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return priceChangeExecutions, err
}

func (ci *PriceChangeExecution) FindByIDs(ids ...domain.UUID) ([]*domain.PriceChangeExecution, error) {
	var priceChangeExecutions []*domain.PriceChangeExecution

	err := ci.db.Table("price_change_executions").Where("id IN ?", ids).Find(&priceChangeExecutions).Error

	return priceChangeExecutions, err
}

func (cpe *PriceChangeExecution) Save(pricechangeexecution *domain.PriceChangeExecution) error {
	return cpe.db.Create(pricechangeexecution).Error
}
