package dto

import (
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
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
	Mba                 string `json:"mba" validate:"required" validate:"required"`
	ProductSerialNumber string `json:"productSerialNumber" validate:"required"`

	// TODO: Define which ones are required
	ProductName string `json:"productName" validate:"required"`
	InArea      bool   `json:"inArea"`
	Commodity   string `json:"commodity" validate:"required"`

	OrderDate           time.Time `json:"orderDate" validate:"required"`
	StartDate           time.Time `json:"startDate" validate:"required"`
	EndDate             time.Time `json:"endDate"`
	Status              string    `json:"status" validate:"required"`
	PriceGuaranteeUntil time.Time `json:"priceGuaranteeUntil"`
	PriceChangePlanned  bool      `json:"priceChangePlanned"`

	PriceValidSince     time.Time `json:"priceValidSince" validate:"required"`
	CurrentBaseCosts    float64   `json:"currentBaseCosts" validate:"required"`
	CurrentKwhCosts     float64   `json:"currentKwhCosts" validate:"required"`
	CurrentBaseMargin   float64   `json:"currentBaseMargin" validate:"required"`
	CurrentKwhMargin    float64   `json:"currentKwhMargin" validate:"required"`
	CurrentBasePriceNet float64   `json:"currentBasePriceNet" validate:"required"`
	CurrentKwhPriceNet  float64   `json:"currentKwhPriceNet" validate:"required"`
	AnnualConsumption   float64   `json:"annualConsumption" validate:"required"`
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

	// TODO: Extent if required

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
		ProductName:         c.ProductName,
		InArea:              c.InArea,
		Commodity:           c.Commodity,
		OrderDate:           c.OrderDate,
		StartDate:           c.StartDate,
		EndDate:             c.EndDate,
		Status:              c.Status,
		PriceGuaranteeUntil: c.PriceGuaranteeUntil,
		PriceChangePlanned:  c.PriceChangePlanned,
		PriceValidSince:     c.PriceValidSince,
		CurrentBaseCosts:    c.CurrentBaseCosts,
		CurrentKwhCosts:     c.CurrentKwhCosts,
		CurrentBaseMargin:   c.CurrentBaseMargin,
		CurrentKwhMargin:    c.CurrentKwhMargin,
		CurrentBasePriceNet: c.CurrentBasePriceNet,
		CurrentKwhPriceNet:  c.CurrentKwhPriceNet,
		AnnualConsumption:   c.AnnualConsumption,
	}
}
