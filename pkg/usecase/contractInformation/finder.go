package contractInformation

import (
	"context"
	"github.com/enercity/be-service-sample/pkg/model/dto"
	"github.com/enercity/be-service-sample/pkg/repository"
	"github.com/enercity/be-service-sample/pkg/usecase"
	logger "github.com/enercity/lib-logger/v3"
)

type Finder struct {
	store repository.Store
}

func NewFinder(
	store repository.Store,
) *Finder {
	return &Finder{
		store: store,
	}
}

func (f *Finder) Find(
	ctx context.Context, logEntry logger.Entry, req *dto.FindContractInformationRequest,
) (*dto.FindContractInformationsResponse, error) {
	logEntry.WithField("req", req).Debug("about search for contract informations from the db by given request data")

	res := &dto.FindContractInformationsResponse{}

	total, err := f.store.ContractInformation().CountAllWithFilters(req.Map(), req.Pagination.Limit, req.Pagination.Offset)

	if err != nil {
		logEntry.WithError(err).Warning("unable to retrieve total contract informations")

		return nil, usecase.ErrDatabaseInternal
	}

	res.Pagination.TotalRows = total
	res.Pagination.Limit = req.Pagination.Limit
	res.Pagination.Offset = req.Pagination.Offset

	contractInformations, err := f.store.ContractInformation().Find(req.Map(), req.Pagination.Limit, req.Pagination.Offset)
	if err != nil {
		logEntry.WithError(err).Warning("unable to retrieve database information")
		return nil, usecase.ErrDatabaseInternal
	}

	res.Items = make([]dto.ContractInformation, 0, len(contractInformations))
	for _, c := range contractInformations {
		res.Items = append(res.Items, dto.ContractInformationFromDomain(c))
	}

	return res, nil
}
