package domain

import (
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

func NewPriceChangeExecution(
	productSerialNumber string,
	priceChangeOrder string,

	priceValidAsOf time.Time,
	currentBasePriceNet float64,
	futureBasePriceNet float64,
	currentKwhPriceNet float64,
	futureKwhPriceNet float64,
	agentHintFlag bool,
	agentHintText string,
	communicationFlag bool,
	communicationTime time.Time,
	annualConsumption float64,
) (*PriceChangeExecution, error) {
	id, err := NewUUID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	status := "new"

	return &PriceChangeExecution{
		Id:                  id.String(),
		CreatedAt:           now,
		ProductSerialNumber: productSerialNumber,
		PriceChangeOrder:    priceChangeOrder,

		Status: status,

		PriceValidAsOf:      priceValidAsOf,
		CurrentBasePriceNet: currentBasePriceNet,
		FutureBasePriceNet:  futureBasePriceNet,
		CurrentKwhPriceNet:  currentKwhPriceNet,
		FutureKwhPriceNet:   futureKwhPriceNet,
		AgentHintFlag:       agentHintFlag,
		AgentHintText:       agentHintText,
		CommunicationFlag:   communicationFlag,
		CommunicationTime:   communicationTime,
		AnnualConsumption:   annualConsumption,
	}, nil
}
