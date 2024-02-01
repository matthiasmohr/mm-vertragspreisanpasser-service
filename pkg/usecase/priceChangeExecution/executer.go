package priceChangeExecution

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

func (f *Executer) Execute(ctx context.Context, logEntry logger.Entry, req *dto.ExecutePriceChangeExecutionRequest) error {
	logEntry.WithField("req", req).Debug("about to execute price change execution")

	uuid, err := domain.ParseUUID(req.Id)
	if err != nil {
		logEntry.WithError(err).Warning("unable to convert to UUID")
		return usecase.ErrDomainInternal
	}

	// Execute against pricechanger service
	pricechangeexecutions, err := f.store.PriceChangeExecution().FindByIDs(uuid)
	if err != nil || len(pricechangeexecutions) == 0 {
		logEntry.WithError(err).Warning("unable to retrieve database information")
		return usecase.ErrDatabaseNotFound
	}

	for _, pce := range pricechangeexecutions {
		if pce.Status == "Execution Created" {
			return usecase.ErrNotAllowed
		}

		fmt.Println("MMMM: ", "CODE COMES HERE", pce)
	}
	// Update Price Change Execution with result

	return nil
}
