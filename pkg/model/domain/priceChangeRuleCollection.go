package domain

import (
	"time"
)

type PriceChangeRuleCollection struct {
	Id                       UUID
	CreatedAt                time.Time
	PriceAdjustmentProjectId string
}

func NewPriceChangeRuleCollection(priceAdjustmentProject string) (*PriceChangeRuleCollection, error) {
	id, err := NewUUID()
	if err != nil {
		return nil, err
	}
	now := time.Now()

	return &PriceChangeRuleCollection{
		Id:                       id,
		CreatedAt:                now,
		PriceAdjustmentProjectId: priceAdjustmentProject,
	}, nil
}
