package dto

import (
	"github.com/enercity/be-service-sample/pkg/model/domain"
	"time"
)

type ContractInformation struct {
	ID           string
	SnapshotTime time.Time

	Mba                 string
	ProductSerialNumber string
	ProductName         string
	InArea              bool
	Commodity           string

	OrderDate           time.Time
	StartDate           time.Time
	EndDate             time.Time
	Status              string
	PriceGuaranteeUntil time.Time
	PriceChangePlanned  bool

	PriceValidSince     time.Time
	CurrentBaseCosts    float64
	CurrentKwhCosts     float64
	CurrentBaseMargin   float64
	CurrentKwhMargin    float64
	CurrentBasePriceNet float64
	CurrentKwhPriceNet  float64
	AnnualConsumption   float64
}

type CreateContractInformationRequest struct {
	Mba                 string `json:"mba" validate:"required"`
	ProductSerialNumber string `json:"productSerialNumber" validate:"required"`

	// TODO: Define which ones are required
	ProductName string `json:"productName"`
	InArea      bool   `json:"inArea"`
	Commodity   string `json:"commodity"`

	OrderDate           time.Time `json:"orderDate"`
	StartDate           time.Time `json:"startDate"`
	EndDate             time.Time `json:"endDate"`
	Status              string    `json:"status"`
	PriceGuaranteeUntil time.Time `json:"priceGuaranteeUntil"`
	PriceChangePlanned  bool      `json:"priceChangePlanned"`

	PriceValidSince     time.Time `json:"priceValidSince"`
	CurrentBaseCosts    float64   `json:"currentBaseCosts"`
	CurrentKwhCosts     float64   `json:"currentKwhCosts"`
	CurrentBaseMargin   float64   `json:"currentBaseMargin"`
	CurrentKwhMargin    float64   `json:"currentKwhMargin"`
	CurrentBasePriceNet float64   `json:"currentBasePriceNet"`
	CurrentKwhPriceNet  float64   `json:"currentKwhPriceNet"`
	AnnualConsumption   float64   `json:"annualConsumption"`
}

type ListContractInformationsRequest struct {
	Pagination
}

type ListContractInformationsResponse struct {
	Pagination Pagination            `json:"pagination"`
	Items      []ContractInformation `json:"items"`
}

type FindContractInformationRequest struct {
	Pagination          Pagination `json:"pagination"`
	Mba                 *string    `json:"mba" validate:"required"`
	ProductSerialNumber *string    `json:"productSerialNumber"`
}

type FindContractInformationsResponse struct {
	Pagination Pagination            `json:"pagination"`
	Items      []ContractInformation `json:"items"`
}

func (cifr *FindContractInformationRequest) Map() map[string]interface{} {
	m := map[string]interface{}{}

	if cifr.Mba != nil {
		m["mba"] = *cifr.Mba
	}

	if cifr.ProductSerialNumber != nil {
		m["productSerialNumber"] = *cifr.ProductSerialNumber
	}

	return m
}

func ContractInformationFromDomain(c *domain.ContractInformation) ContractInformation {
	return ContractInformation{
		ID:                  c.ID.String(),
		Mba:                 c.Mba,
		ProductSerialNumber: c.ProductSerialNumber,
	}
}
