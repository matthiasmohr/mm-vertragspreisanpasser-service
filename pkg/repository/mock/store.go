package mock

import (
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/stretchr/testify/mock"
)

type Store struct {
	mock.Mock

	CustomerMock *Customer
}

// New returns an initialized mock for use in tests.
func New() *Store {
	return &Store{
		CustomerMock: &Customer{},
	}
}

func (s *Store) BeginTransaction() (repository.Store, error) {
	args := s.Called()

	if args.Get(0) != nil {
		return args.Get(0).(repository.Store), args.Error(1)
	}

	return nil, args.Error(1)
}

// Commit mocks implementation of the real method.
func (s *Store) Commit() error {
	return s.Called().Error(0)
}

// Rollback mocks implementation of the real method.
func (s *Store) Rollback() error {
	return s.Called().Error(0)
}

// Case mocks implementation of the real method.
func (s *Store) Customer() repository.Customer {
	return s.CustomerMock
}
