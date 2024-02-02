package dto

import (
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"time"
)

type PriceChangeRuleCollection struct {
	Id                       string
	PriceAdjustmentProjectId string
	CreatedAt                time.Time
}

type CreatePriceChangeRuleCollectionRequest struct {
	PriceAdjustmentProjectId string `json:"priceAdjustmentProjectId" validate:"required"`
}
type ListPriceChangeRuleCollectionRequest struct {
	Pagination
}

type ListPriceChangeRuleCollectionResponse struct {
	Pagination Pagination                  `json:"pagination"`
	Items      []PriceChangeRuleCollection `json:"items"`
}

type GetPriceChangeRuleCollectionRequest struct {
	Id string
}

type GetPriceChangeRuleCollectionResponse struct {
	RuleCollection PriceChangeRuleCollection
	Rules          []PriceChangeRule
}

type ExecutePriceChangeRuleCollectionRequest struct {
	Id string
}

func PriceChangeRuleCollectionFromDomain(c *domain.PriceChangeRuleCollection) PriceChangeRuleCollection {
	return PriceChangeRuleCollection{
		Id:                       c.Id.String(),
		PriceAdjustmentProjectId: c.PriceAdjustmentProjectId,
		CreatedAt:                c.CreatedAt,
	}
}
