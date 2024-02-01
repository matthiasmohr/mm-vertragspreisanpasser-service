package dto

import (
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"time"
)

type PriceChangeOrder struct {
	Id              string
	Created_at      time.Time
	PriceChangeRule string

	ProductSerialNumber string
	Status              string

	PriceValidSince     time.Time
	CurrentBaseCosts    float64
	CurrentKwhCosts     float64
	CurrentBaseMargin   float64
	CurrentKwhMargin    float64
	CurrentBasePriceNet float64
	CurrentKwhPriceNet  float64
	AnnualConsumption   float64

	PriceValidAsOf     time.Time
	FutureBaseCosts    float64
	FutureKwhCosts     float64
	FutureKwhMargin    float64
	FutureBaseMargin   float64
	FutureBasePriceNet float64
	FutureKwhPriceNet  float64
	AgentHintFlag      bool
	AgentHintText      string
	CommunicationFlag  bool
	CommunicationTime  time.Time
}

type CreatePriceChangeOrderRequest struct {
	//PriceChangeRule string `json:"priceChangeRule"`
	ProductSerialNumber string `json:"productSerialNumber" validate:"required"`

	PriceValidSince     time.Time `json:"priceValidSince" validate:"required"`
	CurrentBaseCosts    float64   `json:"currentBaseCosts"`
	CurrentKwhCosts     float64   `json:"currentKwhCosts"`
	CurrentBaseMargin   float64   `json:"currentBaseMargin"`
	CurrentKwhMargin    float64   `json:"currentKwhMargin"`
	CurrentBasePriceNet float64   `json:"currentBasePriceNet" validate:"required"`
	CurrentKwhPriceNet  float64   `json:"currentKwhPriceNet" validate:"required"`
	AnnualConsumption   float64   `json:"annualConsumption" validate:"required"`

	PriceValidAsOf     time.Time `json:"priceValidAsOf" validate:"required"`
	FutureBaseCosts    float64   `json:"futureBaseCosts"`
	FutureKwhCosts     float64   `json:"futureKwhCosts"`
	FutureKwhMargin    float64   `json:"futureKwhMargin"`
	FutureBaseMargin   float64   `json:"futureBaseMargin"`
	FutureBasePriceNet float64   `json:"futureBasePriceNet" validate:"required"`
	FutureKwhPriceNet  float64   `json:"futureKwhPriceNet" validate:"required"`
	AgentHintFlag      bool      `json:"agentHintFlag"`
	AgentHintText      string    `json:"agentHintText"`
	CommunicationFlag  bool      `json:"communicationFlag"`
	CommunicationTime  time.Time `json:"communicationTime"`
}

type ListPriceChangeOrderRequest struct {
	Pagination
}

type ListPriceChangeOrderResponse struct {
	Pagination Pagination         `json:"pagination"`
	Items      []PriceChangeOrder `json:"items"`
}

type FindPriceChangeOrderRequest struct {
	Pagination          Pagination `json:"pagination"`
	ProductSerialNumber *string    `json:"productSerialNumber"`
}

type FindPriceChangeOrderResponse struct {
	Pagination Pagination         `json:"pagination"`
	Items      []PriceChangeOrder `json:"items"`
}

func (cifr *FindPriceChangeOrderRequest) Map() map[string]interface{} {
	m := map[string]interface{}{}

	if cifr.ProductSerialNumber != nil {
		m["productSerialNumber"] = *cifr.ProductSerialNumber
	}

	return m
}

type ExecutePriceChangeOrderRequest struct {
	Id string
}

func PriceChangeOrderFromDomain(c *domain.PriceChangeOrder) PriceChangeOrder {
	return PriceChangeOrder{
		Id:                  c.Id.String(),
		Created_at:          c.Created_at,
		PriceChangeRule:     c.PriceChangeRule,
		Status:              c.Status,
		ProductSerialNumber: c.ProductSerialNumber,
		PriceValidSince:     c.PriceValidSince,
		CurrentBaseCosts:    c.CurrentBaseCosts,
		CurrentKwhCosts:     c.CurrentKwhCosts,
		CurrentBaseMargin:   c.CurrentBaseMargin,
		CurrentKwhMargin:    c.CurrentKwhMargin,
		CurrentBasePriceNet: c.CurrentBasePriceNet,
		CurrentKwhPriceNet:  c.CurrentKwhPriceNet,
		AnnualConsumption:   c.AnnualConsumption,
		PriceValidAsOf:      c.PriceValidAsOf,
		FutureBaseCosts:     c.FutureBaseCosts,
		FutureKwhCosts:      c.FutureKwhCosts,
		FutureKwhMargin:     c.FutureKwhMargin,
		FutureBaseMargin:    c.FutureBaseMargin,
		FutureBasePriceNet:  c.FutureBasePriceNet,
		FutureKwhPriceNet:   c.FutureKwhPriceNet,
		AgentHintFlag:       c.AgentHintFlag,
		AgentHintText:       c.AgentHintText,
		CommunicationFlag:   c.CommunicationFlag,
		CommunicationTime:   c.CommunicationTime,
	}
}
