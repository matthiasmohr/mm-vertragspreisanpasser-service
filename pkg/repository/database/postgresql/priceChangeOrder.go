package postgresql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/enercity/be-service-sample/pkg/model/domain"
	"gorm.io/gorm"
)

type PriceChangeOrder struct {
	db *gorm.DB
}

func newPriceChangeOrder(db *gorm.DB) *PriceChangeOrder {
	return &PriceChangeOrder{
		db: db,
	}
}

func (c *PriceChangeOrder) CountAll() (int64, error) {
	var count int64

	err := c.db.Model(&PriceChangeOrder{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, err
}

func (ci *PriceChangeOrder) CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error) {
	var count int64

	stmt := ci.db.Model(&PriceChangeOrder{}).
		Limit(limit).
		Offset(offset)

	applyFilters(stmt, filters)

	err := stmt.Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func (ci *PriceChangeOrder) Find(filters map[string]interface{}, limit, offset int) ([]*domain.PriceChangeOrder, error) {
	var pricechangeorders []*domain.PriceChangeOrder

	stmt := ci.db.Model(&PriceChangeOrder{}).
		Limit(limit).
		Offset(offset)

	applyFiltersToPriceChangeOrder(stmt, filters)

	err := stmt.Find(&pricechangeorders).Error

	return pricechangeorders, err
}

func applyFiltersToPriceChangeOrder(stmt *gorm.DB, filters map[string]interface{}) {
	// TODO
	applyLikeFilterPriceChangeOrder(stmt, filters, "productSerialNumber")
}

func applyLikeFilterPriceChangeOrder(stmt *gorm.DB, filters map[string]interface{}, key string) {
	if v, ok := filters[key]; ok {
		val := "%" + strings.ToLower(v.(string)) + "%"
		stmt = stmt.Where(fmt.Sprintf("LOWER(%s) LIKE ?", toSnakeCase(key)), val)
	}
}

func (ci *PriceChangeOrder) Load(limit, offset int) ([]*domain.PriceChangeOrder, error) {
	var priceChangeOrders []*domain.PriceChangeOrder

	err := ci.db.Limit(limit).Offset(offset).Find(&priceChangeOrders).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return priceChangeOrders, err
}

func (ci *PriceChangeOrder) Save(pricechangeorder *domain.PriceChangeOrder) error {
	return ci.db.Create(pricechangeorder).Error
}
