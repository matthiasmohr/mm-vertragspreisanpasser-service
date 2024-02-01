package domain

import (
	"time"
)

type PriceChangeOrder struct {
	Id              UUID
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

func NewPriceChangeOrder(
	priceChangeRule string,
	productSerialNumber string,

	priceValidSince time.Time,
	currentBaseCosts float64,
	currentKwhCosts float64,
	currentBaseMargin float64,
	currentKwhMargin float64,
	currentBasePriceNet float64,
	currentKwhPriceNet float64,
	annualConsumption float64,

	priceValidAsOf time.Time,
	futureBaseCosts float64,
	futureKwhCosts float64,
	futureKwhMargin float64,
	futureBaseMargin float64,
	futureBasePriceNet float64,
	futureKwhPriceNet float64,
	agentHintFlag bool,
	agentHintText string,
	communicationFlag bool,
	communicationTime time.Time,
) (*PriceChangeOrder, error) {
	id, err := NewUUID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	status := "new"

	return &PriceChangeOrder{
		Id:         id,
		Created_at: now,

		PriceChangeRule: priceChangeRule,

		ProductSerialNumber: productSerialNumber,
		Status:              status,

		PriceValidSince:     priceValidSince,
		CurrentBaseCosts:    currentBaseCosts,
		CurrentKwhCosts:     currentKwhCosts,
		CurrentBaseMargin:   currentBaseMargin,
		CurrentKwhMargin:    currentKwhMargin,
		CurrentBasePriceNet: currentBasePriceNet,
		CurrentKwhPriceNet:  currentKwhPriceNet,
		AnnualConsumption:   annualConsumption,

		PriceValidAsOf:     priceValidAsOf,
		FutureBaseCosts:    futureBaseCosts,
		FutureKwhCosts:     futureKwhCosts,
		FutureKwhMargin:    futureKwhMargin,
		FutureBaseMargin:   futureBaseMargin,
		FutureBasePriceNet: futureBasePriceNet,
		FutureKwhPriceNet:  futureKwhPriceNet,
		AgentHintFlag:      agentHintFlag,
		AgentHintText:      agentHintText,
		CommunicationFlag:  communicationFlag,
		CommunicationTime:  communicationTime,
	}, nil
}
