package domain

import "time"

type ContractInformation struct {
	ID           UUID
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

func NewContractInformation(
	mba string,
	productSerialNumber string,
	productName string,
	inArea bool,
	commodity string,

	orderDate time.Time,
	startDate time.Time,
	endDate time.Time,
	status string,
	priceGuaranteeUntil time.Time,
	priceChangePlanned bool,

	priceValidSince time.Time,
	currentBaseCosts float64,
	currentKwhCosts float64,
	currentBaseMargin float64,
	currentKwhMargin float64,
	currentBasePriceNet float64,
	currentKwhPriceNet float64,
	annualConsumption float64,
) (*ContractInformation, error) {
	id, err := NewUUID()
	now := time.Now()
	if err != nil {
		return nil, err
	}
	if currentBasePriceNet == 0 {
		currentBasePriceNet = currentBaseCosts + currentBaseMargin
	}
	if currentKwhPriceNet == 0 {
		currentKwhPriceNet = currentKwhCosts + currentKwhMargin
	}

	return &ContractInformation{
		ID:           id,
		SnapshotTime: now,

		Mba:                 mba,
		ProductSerialNumber: productSerialNumber,
		ProductName:         productName,
		InArea:              inArea,
		Commodity:           commodity,

		OrderDate:           orderDate,
		StartDate:           startDate,
		EndDate:             endDate,
		Status:              status,
		PriceGuaranteeUntil: priceGuaranteeUntil,
		PriceChangePlanned:  priceChangePlanned,

		PriceValidSince:     priceValidSince,
		CurrentBaseCosts:    currentBaseCosts,
		CurrentKwhCosts:     currentKwhCosts,
		CurrentBaseMargin:   currentBaseMargin,
		CurrentKwhMargin:    currentKwhMargin,
		CurrentBasePriceNet: currentBasePriceNet,
		CurrentKwhPriceNet:  currentKwhPriceNet,
		AnnualConsumption:   annualConsumption,
	}, nil
}
