package mock

import (
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/stretchr/testify/mock"
)

type Customer struct {
	mock.Mock
}

func (c *Customer) Find(filters map[string]interface{}, limit, offset int) ([]*domain.Customer, error) {
	args := c.Called(filters, limit, offset)
	if args.Get(0) != nil {
		return args.Get(0).([]*domain.Customer), nil
	}

	return nil, args.Error(1)
}

func (c *Customer) CountAllWithFilters(filters map[string]interface{}, limit, offset int) (int64, error) {
	args := c.Called(filters, limit, offset)

	return args.Get(0).(int64), args.Error(1)
}

func (c *Customer) Load(limit, offset int) ([]*domain.Customer, error) {
	args := c.Called(limit, offset)
	if args.Get(0) != nil {
		return args.Get(0).([]*domain.Customer), nil
	}

	return nil, args.Error(1)
}

func (c *Customer) CountAllCustomers() (int64, error) {
	args := c.Called()

	return args.Get(0).(int64), args.Error(1)
}

func (c *Customer) Save(customer *domain.Customer) error {
	return c.Called(customer).Error(0)
}
