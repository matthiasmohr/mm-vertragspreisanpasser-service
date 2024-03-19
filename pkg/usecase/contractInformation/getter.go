package contractInformation

import (
	"context"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Getter struct {
	store repository.Store
}

func NewGetter(
	store repository.Store,
) *Getter {
	return &Getter{
		store: store,
	}
}

func (g *Getter) Get(ctx context.Context, logEntry logger.Entry, id string) (*dto.GetContractInformationResponse, error) {
	logEntry.WithField("id: ", id).Debug("about to fetch a contractInformation")

	res := &dto.GetContractInformationResponse{}

	// Get the respective priceChangeRuleCollection
	uuid, err := domain.ParseUUID(id)
	if err != nil {
		logEntry.WithError(err).Warning("unable to convert to UUID")
		return nil, usecase.ErrDomainInternal
	}

	contractInformation, err := g.store.ContractInformation().FindByID(uuid)
	if err != nil {
		logEntry.WithContext(ctx).WithError(err).Error("unable to fetch contractInformation in a db")

		return nil, usecase.ErrDatabaseInternal
	}

	res.Item = dto.ContractInformationFromDomain(contractInformation)

	return res, nil
}
