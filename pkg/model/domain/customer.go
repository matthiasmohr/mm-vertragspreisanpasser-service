package domain

import "time"

type Customer struct {
	ID        UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(
	now time.Time,
	firstName string,
	lastName string,
	email string,
) (*Customer, error) {
	id, err := NewUUID()
	if err != nil {
		return nil, err
	}

	return &Customer{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
