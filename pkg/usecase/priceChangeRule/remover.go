package priceChangeRule

import (
	"context"
	"fmt"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Remover struct {
	store repository.Store
}

func NewRemover(
	store repository.Store,
) *Remover {
	return &Remover{
		store: store,
	}
}

func (f *Remover) Remove(ctx context.Context, logEntry logger.Entry, req *dto.RemovePriceChangeRuleRequest) error {
	logEntry.WithField("req", req).Debug("about to execute price change Rule")

	uuid, err := domain.ParseUUID(req.Id)
	if err != nil {
		logEntry.WithError(err).Warning("unable to convert to UUID")
		return usecase.ErrDomainInternal
	}

	pricechangerules, err := f.store.PriceChangeRule().FindByIDs(uuid)
	if err != nil || len(pricechangerules) == 0 {
		logEntry.WithError(err).Warning("unable to retrieve database information")
		return usecase.ErrDatabaseNotFound
	}

	for _, pcr := range pricechangerules {
		// TODO: HERE COMES THE CODE TO REMOVE A PRICE CHANGE RULE
		fmt.Println(pcr)
		if err != nil {
			logEntry.WithError(err).Warning("unable to create new price change execution")
			return usecase.ErrDomainInternal
		}
	}

	return nil
}
