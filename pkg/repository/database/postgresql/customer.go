package postgresql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func newCustomer(db *gorm.DB) *Customer {
	return &Customer{
		db: db,
	}
}

func (c *Customer) CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error) {
	var count int64

	stmt := c.db.Model(&Customer{}).
		Limit(limit).
		Offset(offset)

	applyFilters(stmt, filters)

	err := stmt.Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func (c *Customer) Find(filters map[string]interface{}, limit, offset int) ([]*domain.Customer, error) {
	var customers []*domain.Customer

	stmt := c.db.Model(&Customer{}).
		Limit(limit).
		Offset(offset).
		Order("customers.created_at ASC")

	applyFilters(stmt, filters)

	err := stmt.Find(&customers).Error

	return customers, err
}

func applyFilters(stmt *gorm.DB, filters map[string]interface{}) {
	applyLikeFilter(stmt, filters, "first_name")
	applyLikeFilter(stmt, filters, "last_name")
	applyLikeFilter(stmt, filters, "email")
}

func applyLikeFilter(stmt *gorm.DB, filters map[string]interface{}, key string) {
	if v, ok := filters[key]; ok {
		val := "%" + strings.ToLower(v.(string)) + "%"
		stmt = stmt.Where(fmt.Sprintf("LOWER(customers.%s) LIKE ?", key), val)
	}
}

func (c *Customer) Load(limit, offset int) ([]*domain.Customer, error) {
	var customers []*domain.Customer

	err := c.db.Limit(limit).Offset(offset).Find(&customers).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return customers, err
}

func (c *Customer) CountAllCustomers() (int64, error) {
	var count int64

	err := c.db.Model(&Customer{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, err
}

func (c *Customer) Save(customer *domain.Customer) error {
	return c.db.Create(customer).Error
}
