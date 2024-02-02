package postgresql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"gorm.io/gorm"
)

type PriceChangeRuleCollection struct {
	db *gorm.DB
}

func newPriceChangeRuleCollection(db *gorm.DB) *PriceChangeRuleCollection {
	return &PriceChangeRuleCollection{
		db: db,
	}
}

func (ci *PriceChangeRuleCollection) CountAll() (int64, error) {
	var count int64

	err := ci.db.Model(&PriceChangeRuleCollection{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, err
}

func (ci *PriceChangeRuleCollection) CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error) {
	var count int64

	stmt := ci.db.Model(&PriceChangeRuleCollection{}).
		Limit(limit).
		Offset(offset)

	applyFiltersToPriceChangeRuleCollection(stmt, filters)

	err := stmt.Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func (ci *PriceChangeRuleCollection) Find(filters map[string]interface{}, limit, offset int) ([]*domain.PriceChangeRuleCollection, error) {
	var pricechangerules []*domain.PriceChangeRuleCollection

	stmt := ci.db.Model(&PriceChangeRuleCollection{}).
		Limit(limit).
		Offset(offset)

	applyFiltersToPriceChangeRuleCollection(stmt, filters)

	err := stmt.Find(&pricechangerules).Error

	return pricechangerules, err
}

func applyFiltersToPriceChangeRuleCollection(stmt *gorm.DB, filters map[string]interface{}) {
	// TODO
	applyLikeFilterPriceChangeRuleCollection(stmt, filters, "id")
	applyLikeFilterPriceChangeRuleCollection(stmt, filters, "productSerialNumber")
}

func applyLikeFilterPriceChangeRuleCollection(stmt *gorm.DB, filters map[string]interface{}, key string) {
	if v, ok := filters[key]; ok {
		val := "%" + strings.ToLower(v.(string)) + "%"
		stmt = stmt.Where(fmt.Sprintf("LOWER(%s) LIKE ?", toSnakeCase(key)), val)
	}
}

func (ci *PriceChangeRuleCollection) Load(limit, offset int) ([]*domain.PriceChangeRuleCollection, error) {
	var priceChangeRuleCollections []*domain.PriceChangeRuleCollection

	err := ci.db.Limit(limit).Offset(offset).Find(&priceChangeRuleCollections).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return priceChangeRuleCollections, err
}

func (ci *PriceChangeRuleCollection) FindByIDs(ids ...domain.UUID) ([]*domain.PriceChangeRuleCollection, error) {
	var priceChangeRuleCollections []*domain.PriceChangeRuleCollection

	err := ci.db.Table("price_change_rule_collections").Where("id IN ?", ids).Find(&priceChangeRuleCollections).Error

	return priceChangeRuleCollections, err
}

func (ci *PriceChangeRuleCollection) Save(pricechangerulecollection *domain.PriceChangeRuleCollection) error {
	return ci.db.Create(pricechangerulecollection).Error
}

func (ci *PriceChangeRuleCollection) Update(pricechangerulecollection *domain.PriceChangeRuleCollection) error {
	return ci.db.Updates(pricechangerulecollection).Error
}
