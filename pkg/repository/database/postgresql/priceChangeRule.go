package postgresql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"gorm.io/gorm"
)

type PriceChangeRule struct {
	db *gorm.DB
}

func newPriceChangeRule(db *gorm.DB) *PriceChangeRule {
	return &PriceChangeRule{
		db: db,
	}
}

func (ci *PriceChangeRule) CountAll() (int64, error) {
	var count int64

	err := ci.db.Model(&PriceChangeRule{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, err
}

func (ci *PriceChangeRule) CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error) {
	var count int64

	stmt := ci.db.Model(&PriceChangeRule{}).
		Limit(limit).
		Offset(offset)

	applyFiltersToPriceChangeRule(stmt, filters)

	err := stmt.Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func (ci *PriceChangeRule) Find(filters map[string]interface{}, limit, offset int) ([]*domain.PriceChangeRule, error) {
	var pricechangerules []*domain.PriceChangeRule

	stmt := ci.db.Model(&PriceChangeRule{}).
		Limit(limit).
		Offset(offset)

	applyFiltersToPriceChangeRule(stmt, filters)

	err := stmt.Find(&pricechangerules).Error

	return pricechangerules, err
}

func applyFiltersToPriceChangeRule(stmt *gorm.DB, filters map[string]interface{}) {
	// TODO
	applyLikeFilterPriceChangeRule(stmt, filters, "id")
	applyLikeFilterPriceChangeRule(stmt, filters, "productSerialNumber")
}

func applyLikeFilterPriceChangeRule(stmt *gorm.DB, filters map[string]interface{}, key string) {
	if v, ok := filters[key]; ok {
		val := "%" + strings.ToLower(v.(string)) + "%"
		stmt = stmt.Where(fmt.Sprintf("LOWER(%s) LIKE ?", toSnakeCase(key)), val)
	}
}

func (ci *PriceChangeRule) Load(limit, offset int) ([]*domain.PriceChangeRule, error) {
	var priceChangeRules []*domain.PriceChangeRule

	err := ci.db.Limit(limit).Offset(offset).Find(&priceChangeRules).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return priceChangeRules, err
}

func (ci *PriceChangeRule) FindByIDs(ids ...domain.UUID) ([]*domain.PriceChangeRule, error) {
	var priceChangeRules []*domain.PriceChangeRule

	err := ci.db.Table("price_change_rules").Where("id IN ?", ids).Find(&priceChangeRules).Error

	return priceChangeRules, err
}

func (ci *PriceChangeRule) FindByFindByPriceChangeRuleCollectionId(id domain.UUID) ([]*domain.PriceChangeRule, error) {
	var priceChangeRules []*domain.PriceChangeRule

	//err := ci.db.Find(&priceChangeRules).Error
	err := ci.db.Where("price_change_rule_collection_id = ?", id).Find(&priceChangeRules).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return priceChangeRules, err
}

func (ci *PriceChangeRule) Save(pricechangerule *domain.PriceChangeRule) error {
	return ci.db.Create(pricechangerule).Error
}

func (ci *PriceChangeRule) Update(pricechangerule *domain.PriceChangeRule) error {
	return ci.db.Updates(pricechangerule).Error
}
