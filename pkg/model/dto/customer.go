package dto

import "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"

type CreateCustomerRequest struct {
	FirstName string `json:"firstName" validate:"required,gt=0"`
	LastName  string `json:"lastName" validate:"required,gt=0"`
	Email     string `json:"email" validate:"required,email"`
}

type LoadCustomerRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func CustomerFromDomain(c *domain.Customer) Customer {
	return Customer{
		ID:        c.ID.String(),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
	}
}
