package repository

// Store holds individual repositories and methods for database transaction management.
type Store interface {
	BeginTransaction() (Store, error)
	Commit() error
	Rollback() error

	Customer() Customer
	ContractInformation() ContractInformation
}
