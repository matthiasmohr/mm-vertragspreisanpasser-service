package dto

import (
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"time"
)

type PriceChangeExecution struct {
	Id                  string
	CreatedAt           time.Time
	ProductSerialNumber string
	PriceChangeOrder    string

	Status               string
	ExecutionTime        time.Time
	PricechangerResponse string

	PriceValidAsOf      time.Time
	CurrentBasePriceNet float64
	FutureBasePriceNet  float64
	CurrentKwhPriceNet  float64
	FutureKwhPriceNet   float64
	AgentHintFlag       bool
	AgentHintText       string
	CommunicationFlag   bool
	CommunicationTime   time.Time
	AnnualConsumption   float64
}

type ListPriceChangeExecutionRequest struct {
	Pagination
}

type ListPriceChangeExecutionResponse struct {
	Pagination Pagination             `json:"pagination"`
	Items      []PriceChangeExecution `json:"items"`
}

type FindPriceChangeExecutionRequest struct {
	Pagination          Pagination `json:"pagination"`
	ProductSerialNumber *string    `json:"productSerialNumber"`
}

type FindPriceChangeExecutionResponse struct {
	Pagination Pagination             `json:"pagination"`
	Items      []PriceChangeExecution `json:"items"`
}

func (cifr *FindPriceChangeExecutionRequest) Map() map[string]interface{} {
	m := map[string]interface{}{}

	if cifr.ProductSerialNumber != nil {
		m["productSerialNumber"] = *cifr.ProductSerialNumber
	}

	return m
}

type ExecutePriceChangeExecutionRequest struct {
	Id string
}

type ExecutePriceChangeExecutionResponse struct {
	Pagination Pagination             `json:"pagination"`
	Items      []PriceChangeExecution `json:"items"`
}

func PriceChangeExecutionFromDomain(c *domain.PriceChangeExecution) PriceChangeExecution {
	return PriceChangeExecution{
		Id:                  c.Id,
		CreatedAt:           c.CreatedAt,
		ProductSerialNumber: c.ProductSerialNumber,
		PriceChangeOrder:    c.PriceChangeOrder,

		Status:               c.Status,
		ExecutionTime:        c.ExecutionTime,
		PricechangerResponse: c.PricechangerResponse,

		PriceValidAsOf:      c.PriceValidAsOf,
		CurrentBasePriceNet: c.CurrentBasePriceNet,
		FutureBasePriceNet:  c.FutureBasePriceNet,
		CurrentKwhPriceNet:  c.CurrentKwhPriceNet,
		FutureKwhPriceNet:   c.FutureKwhPriceNet,
		AgentHintFlag:       c.AgentHintFlag,
		AgentHintText:       c.AgentHintText,
		CommunicationFlag:   c.CommunicationFlag,
		CommunicationTime:   c.CommunicationTime,
		AnnualConsumption:   c.AnnualConsumption,
	}
}
