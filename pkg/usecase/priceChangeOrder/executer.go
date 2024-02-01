package priceChangeOrder

import (
	"context"
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

func (f *Executer) Execute(ctx context.Context, logEntry logger.Entry, req *dto.ExecutePriceChangeOrderRequest) error {
	logEntry.WithField("req", req).Debug("about to execute price change order")

	uuid, err := domain.ParseUUID(req.Id)
	if err != nil {
		logEntry.WithError(err).Warning("unable to convert to UUID")
		return usecase.ErrDomainInternal
	}

	pricechangeorders, err := f.store.PriceChangeOrder().FindByIDs(uuid)
	if err != nil || len(pricechangeorders) == 0 {
		logEntry.WithError(err).Warning("unable to retrieve database information")
		return usecase.ErrDatabaseNotFound
	}

	for _, pco := range pricechangeorders {
		// Check if there was an execution already created
		if pco.Status == "Execution Created" {
			return usecase.ErrNotAllowed
		}
		// Create Price Change Execution Item
		priceChangeExecution, err := domain.NewPriceChangeExecution(
			pco.ProductSerialNumber,
			pco.Id.String(),
			pco.PriceValidAsOf,
			pco.CurrentBasePriceNet,
			pco.FutureBasePriceNet,
			pco.CurrentKwhPriceNet,
			pco.FutureKwhPriceNet,
			pco.AgentHintFlag,
			pco.AgentHintText,
			pco.CommunicationFlag,
			pco.CommunicationTime,
			pco.AnnualConsumption,
		)
		if err != nil {
			logEntry.WithError(err).Warning("unable to create new price change execution")
			return usecase.ErrDomainInternal
		}

		// Store Price Change Execution
		err = f.store.PriceChangeExecution().Save(priceChangeExecution)
		if err != nil {
			logEntry.WithError(err).Warning("unable to store new price change execution to database")
			return usecase.ErrDatabaseInternal
		}

		// Update Price Change Order as execution created
		pco.Status = "Execution Created"
		err = f.store.PriceChangeOrder().Update(pco)
		if err != nil {
			logEntry.WithError(err).Warning("unable to update Status of price change order as execution created")
			return usecase.ErrDatabaseInternal
		}
	}

	return nil
}
