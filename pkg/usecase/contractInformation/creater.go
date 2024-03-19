package contractInformation

import (
	"context"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Creator struct {
	store repository.Store
}

func NewCreator(
	store repository.Store,
) *Creator {
	return &Creator{
		store: store,
	}
}

func (c *Creator) Create(ctx context.Context, logEntry logger.Entry, req *dto.CreateContractInformationRequest) error {
	logEntry.WithField("req", req).Debug("about to create a new contract information")

	contractInformation, err := domain.NewContractInformation(
		req.Mba,
		req.ProductSerialNumber,
		req.ProductName,
		req.InArea,
		req.Commodity,

		req.OrderDate,
		req.StartDate,
		req.EndDate,
		req.Status,
		req.PriceGuaranteeUntil,
		req.PriceChangePlanned,

		req.PriceValidSince,
		req.SignupBaseMargin,
		req.SignupKwhMargin,
		req.CurrentBaseCosts,
		req.CurrentKwhCosts,
		req.CurrentBaseMargin,
		req.CurrentKwhMargin,
		req.CurrentBasePriceNet,
		req.CurrentKwhPriceNet,
		req.AnnualConsumption,
	)
	if err != nil {
		logEntry.WithError(err).Error("unable to create a contract information")
		return usecase.ErrDomainInternal
	}

	err = c.store.ContractInformation().Save(contractInformation)
	if err != nil {
		logEntry.WithContext(ctx).WithError(err).Error("unable to store contract information in a db")

		return usecase.ErrDatabaseInternal
	}

	return nil
}
