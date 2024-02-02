package domain

import (
	"time"
)

type PriceChangeRule struct {
	Id                          UUID
	PriceChangeRuleCollectionId UUID
	*PriceChangeRuleCollection

	RestoreMarginAtSignup      bool
	ChangeBasePriceNetToAmount float64
	ChangeKwhPriceNetToAmount  float64
	ChangeBasePriceNetByAmount float64
	ChangeKwhPriceNetByAmount  float64
	ChangeBasePriceNetByFactor float64
	ChangeKwhPriceNetByFactor  float64

	ValidForProductNames        string
	ValidForCommodity           string
	ExcludeOrderDateFrom        time.Time
	ExcludeStartDateFrom        time.Time
	ExcludeEndDateUntil         time.Time
	ExcludeLastPriceChangeSince time.Time

	LimitToCataloguePriceNet         bool
	LimitToUpperBasePriceNet         float64
	LimitToUpperKwhPriceNet          float64
	LimitToLowerBasePriceNet         float64
	LimitToLowerKwhPriceNet          float64
	LimitToMaxChangeBasePriceNet     float64
	LimitToMaxChangeKwhPriceNet      float64
	LimitToMinChangeBasePriceNet     float64
	LimitToMinChangeKwhPriceNet      float64
	OrderInPriceChangeRuleCollection int
	CreatedAt                        time.Time
}

func NewPriceChangeRule(
	priceChangeRuleCollection *PriceChangeRuleCollection,

	restoreMarginAtSignup bool,
	changeBasePriceNetToAmount float64,
	changeKwhPriceNetToAmount float64,
	changeBasePriceNetByAmount float64,
	changeKwhPriceNetByAmount float64,
	changeBasePriceNetByFactor float64,
	changeKwhPriceNetByFactor float64,

	validForProductNames string,
	validForCommodity string,
	excludeOrderDateFrom time.Time,
	excludeStartDateFrom time.Time,
	excludeEndDateUntil time.Time,
	excludeLastPriceChangeSince time.Time,

	limitToCataloguePriceNet bool,
	limitToUpperBasePriceNet float64,
	limitToUpperKwhPriceNet float64,
	limitToLowerBasePriceNet float64,
	limitToLowerKwhPriceNet float64,
	limitToMaxChangeBasePriceNet float64,
	limitToMaxChangeKwhPriceNet float64,
	limitToMinChangeBasePriceNet float64,
	limitToMinChangeKwhPriceNet float64,
) (*PriceChangeRule, error) {
	id, err := NewUUID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	orderInPriceChangeRuleCollection := 1 // TODO: MAKE DYNAMIC

	return &PriceChangeRule{
		Id:                          id,
		PriceChangeRuleCollectionId: priceChangeRuleCollection.Id,

		RestoreMarginAtSignup:      restoreMarginAtSignup,
		ChangeBasePriceNetToAmount: changeBasePriceNetToAmount,
		ChangeKwhPriceNetToAmount:  changeKwhPriceNetToAmount,
		ChangeBasePriceNetByAmount: changeBasePriceNetByAmount,
		ChangeKwhPriceNetByAmount:  changeKwhPriceNetByAmount,
		ChangeBasePriceNetByFactor: changeBasePriceNetByFactor,
		ChangeKwhPriceNetByFactor:  changeKwhPriceNetByFactor,

		ValidForProductNames:        validForProductNames,
		ValidForCommodity:           validForCommodity,
		ExcludeOrderDateFrom:        excludeOrderDateFrom,
		ExcludeStartDateFrom:        excludeStartDateFrom,
		ExcludeEndDateUntil:         excludeEndDateUntil,
		ExcludeLastPriceChangeSince: excludeLastPriceChangeSince,

		LimitToCataloguePriceNet:         limitToCataloguePriceNet,
		LimitToUpperBasePriceNet:         limitToUpperBasePriceNet,
		LimitToUpperKwhPriceNet:          limitToUpperKwhPriceNet,
		LimitToLowerBasePriceNet:         limitToLowerBasePriceNet,
		LimitToLowerKwhPriceNet:          limitToLowerKwhPriceNet,
		LimitToMaxChangeBasePriceNet:     limitToMaxChangeBasePriceNet,
		LimitToMaxChangeKwhPriceNet:      limitToMaxChangeKwhPriceNet,
		LimitToMinChangeBasePriceNet:     limitToMinChangeBasePriceNet,
		LimitToMinChangeKwhPriceNet:      limitToMinChangeKwhPriceNet,
		OrderInPriceChangeRuleCollection: orderInPriceChangeRuleCollection,
		CreatedAt:                        now,
	}, nil
}
