package priceChangeRuleCollection

import (
	"context"
	"fmt"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
)

type Executer struct {
	store repository.Store
}

func NewExecuter(
	store repository.Store,
) *Executer {
	return &Executer{
		store: store,
	}
}

func (f *Executer) Execute(ctx context.Context, logEntry logger.Entry, req *dto.ExecutePriceChangeRuleCollectionRequest) error {
	logEntry.WithField("req", req).Debug("about to execute price change Rule Collection")

	uuid, err := domain.ParseUUID(req.Id)
	if err != nil {
		logEntry.WithError(err).Warning("unable to convert to UUID")
		return usecase.ErrDomainInternal
	}

	pricechangerulecollection, err := f.store.PriceChangeRuleCollection().FindByIDs(uuid)
	if err != nil || len(pricechangerulecollection) == 0 {
		logEntry.WithError(err).Warning("unable to retrieve database information")
		return usecase.ErrDatabaseNotFound
	}

	for _, pcr := range pricechangerulecollection {
		// TODO: HERE COMES THE CODE TO LOAD AND EXECUTE THE PRICE CHANGE RULE COLLECTION
		fmt.Println("HERE COMES THE CODE.........", pcr)
		if err != nil {
			logEntry.WithError(err).Warning("unable to create new price change execution")
			return usecase.ErrDomainInternal
		}
	}

	return nil
}
