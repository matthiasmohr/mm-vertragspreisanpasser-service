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
	SignupKwhMargin     float64
	SignupBaseMargin    float64
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
	signupKwhMargin float64,
	signupBaseMargin float64,
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

		SignupKwhMargin:     signupKwhMargin,
		SignupBaseMargin:    signupBaseMargin,
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

func (ci *ContractInformation) ProposedBasePriceNet() float64 {
	if ci.SignupBaseMargin != 0 && ci.CurrentBaseCosts != 0 {
		return ci.SignupBaseMargin + ci.CurrentBaseCosts
	} else {
		return 0
	}
}

func (ci *ContractInformation) ProposedKwhPriceNet() float64 {
	if ci.SignupKwhMargin != 0 && ci.CurrentKwhCosts != 0 {
		return ci.SignupKwhMargin + ci.CurrentKwhCosts
	} else {
		return 0
	}
}
