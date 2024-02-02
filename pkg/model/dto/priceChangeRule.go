package dto

import (
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"time"
)

type PriceChangeRule struct {
	Id                          string
	PriceChangeRuleCollectionId string
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

type CreatePriceChangeRuleRequest struct {
	PriceChangeRuleCollectionId string `json:"PriceChangeRuleCollectionId" validate:"required"`

	RestoreMarginAtSignup      bool    `json:"restoreMarginAtSignup"`
	ChangeBasePriceNetToAmount float64 `json:"changeBasePriceNetToAmount"`
	ChangeKwhPriceNetToAmount  float64 `json:"changeKwhPriceNetToAmount"`
	ChangeBasePriceNetByAmount float64 `json:"changeBasePriceNetByAmount"`
	ChangeKwhPriceNetByAmount  float64 `json:"changeKwhPriceNetByAmount"`
	ChangeBasePriceNetByFactor float64 `json:"changeBasePriceNetByFactor"`
	ChangeKwhPriceNetByFactor  float64 `json:"changeKwhPriceNetByFactor"`

	ValidForProductNames        string    `json:"validForProductNames"`
	ValidForCommodity           string    `json:"validForCommodity" validate:"required"`
	ExcludeOrderDateFrom        time.Time `json:"excludeOrderDateFrom"`
	ExcludeStartDateFrom        time.Time `json:"excludeStartDateFrom"`
	ExcludeEndDateUntil         time.Time `json:"excludeEndDateUntil"`
	ExcludeLastPriceChangeSince time.Time `json:"excludeLastPriceChangeSince"`

	LimitToCataloguePriceNet     bool    `json:"limitToCataloguePriceNet"`
	LimitToUpperBasePriceNet     float64 `json:"limitToUpperBasePriceNet"`
	LimitToUpperKwhPriceNet      float64 `json:"limitToUpperKwhPriceNet"`
	LimitToLowerBasePriceNet     float64 `json:"limitToLowerBasePriceNet"`
	LimitToLowerKwhPriceNet      float64 `json:"limitToLowerKwhPriceNet"`
	LimitToMaxChangeBasePriceNet float64 `json:"limitToMaxChangeBasePriceNet"`
	LimitToMaxChangeKwhPriceNet  float64 `json:"limitToMaxChangeKwhPriceNet"`
	LimitToMinChangeBasePriceNet float64 `json:"limitToMinChangeBasePriceNet"`
	LimitToMinChangeKwhPriceNet  float64 `json:"limitToMinChangeKwhPriceNet"`
}

type ListPriceChangeRuleRequest struct {
	Pagination
}

type ListPriceChangeRuleResponse struct {
	Pagination Pagination        `json:"pagination"`
	Items      []PriceChangeRule `json:"items"`
}

type RemovePriceChangeRuleRequest struct {
	Id string
}

func PriceChangeRuleFromDomain(c *domain.PriceChangeRule) PriceChangeRule {
	// TODO
	return PriceChangeRule{
		Id:                          c.Id.String(),
		PriceChangeRuleCollectionId: c.PriceChangeRuleCollectionId.String(),
		//*PriceChangeRuleCollection:
		RestoreMarginAtSignup:            c.RestoreMarginAtSignup,
		ChangeBasePriceNetToAmount:       c.ChangeBasePriceNetToAmount,
		ChangeKwhPriceNetToAmount:        c.ChangeKwhPriceNetByAmount,
		ChangeBasePriceNetByAmount:       c.ChangeBasePriceNetByAmount,
		ChangeKwhPriceNetByAmount:        c.ChangeKwhPriceNetToAmount,
		ChangeBasePriceNetByFactor:       c.ChangeBasePriceNetByFactor,
		ChangeKwhPriceNetByFactor:        c.ChangeKwhPriceNetByFactor,
		ValidForProductNames:             c.ValidForProductNames,
		ValidForCommodity:                c.ValidForCommodity,
		ExcludeOrderDateFrom:             c.ExcludeOrderDateFrom,
		ExcludeStartDateFrom:             c.ExcludeStartDateFrom,
		ExcludeEndDateUntil:              c.ExcludeEndDateUntil,
		ExcludeLastPriceChangeSince:      c.ExcludeLastPriceChangeSince,
		LimitToCataloguePriceNet:         c.LimitToCataloguePriceNet,
		LimitToUpperBasePriceNet:         c.LimitToUpperBasePriceNet,
		LimitToUpperKwhPriceNet:          c.LimitToUpperKwhPriceNet,
		LimitToLowerBasePriceNet:         c.LimitToLowerBasePriceNet,
		LimitToLowerKwhPriceNet:          c.LimitToLowerKwhPriceNet,
		LimitToMaxChangeBasePriceNet:     c.LimitToMaxChangeBasePriceNet,
		LimitToMaxChangeKwhPriceNet:      c.LimitToMaxChangeKwhPriceNet,
		LimitToMinChangeBasePriceNet:     c.LimitToMinChangeBasePriceNet,
		LimitToMinChangeKwhPriceNet:      c.LimitToMinChangeKwhPriceNet,
		OrderInPriceChangeRuleCollection: c.OrderInPriceChangeRuleCollection,
		CreatedAt:                        c.CreatedAt,
	}
}
