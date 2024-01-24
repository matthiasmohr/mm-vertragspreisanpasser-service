package contractInformation

import (
	"context"
	"github.com/enercity/be-service-sample/pkg/model/dto"
	"github.com/enercity/be-service-sample/pkg/repository"
	"github.com/enercity/be-service-sample/pkg/usecase"
	logger "github.com/enercity/lib-logger/v3"
)

type Lister struct {
	store repository.Store
}

func NewLister(
	store repository.Store,
) *Lister {
	return &Lister{
		store: store,
	}
}

func (l *Lister) List(ctx context.Context, logEntry logger.Entry, req *dto.ListContractInformationsRequest) (*dto.ListContractInformationsResponse, error) {
	logEntry.Debug("about load contract informations from the db")

	res := &dto.ListContractInformationsResponse{}

	total, err := l.store.ContractInformation().CountAll()
	if err != nil {
		logEntry.Warning("unable to retrieve total contract informations")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Limit
	res.Pagination.Offset = req.Offset

	contractInformations, err := l.store.ContractInformation().Load(req.Limit, req.Offset)
	if err != nil {
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.ContractInformation, 0, len(contractInformations))
	for _, c := range contractInformations {
		res.Items = append(res.Items, dto.ContractInformationFromDomain(c))
	}

	return res, nil
}
