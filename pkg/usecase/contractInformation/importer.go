package contractInformation

import (
	"context"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/gocarina/gocsv"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/domain"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase"
	"os"
)

type Importer struct {
	store repository.Store
}

func NewImporter(
	store repository.Store,
) *Importer {
	return &Importer{
		store: store,
	}
}

func (imp *Importer) Import(ctx context.Context, logEntry logger.Entry, req *dto.ImportContractInformationRequest) (*dto.ImportContractInformationResponse, error) {
	logEntry.WithField("req", req).Debug("about to import new contract information")

	res := &dto.ImportContractInformationResponse{}

	// Assign Standard to filename Input
	var filename string
	if len(req.File) == 0 {
		filename = "dataInput/Stage_Data_10.csv"
	} else {
		filename = "dataInput/" + req.File
	}
	res.File = req.File

	// Open File in OS
	f, err := os.Open(filename)
	if err != nil {
		logEntry.WithError(err).Error("unable to open file")
		return nil, usecase.ErrDatabaseInternal
	}
	defer f.Close()

	// Parse file into Contract Information Slices
	var contractFile []domain.ContractInformation
	err = gocsv.UnmarshalFile(f, &contractFile)
	if err != nil {
		logEntry.WithError(err).Error("unable to unmarshal import file into struct")
		return nil, usecase.ErrDatabaseInternal
	}

	// iterate through Contract Slice and store to database
	for i, _ := range contractFile {
		ci, err := domain.NewContractInformation(
			contractFile[i].Mba,
			contractFile[i].ProductSerialNumber,
			contractFile[i].ProductName,
			contractFile[i].InArea,
			contractFile[i].Commodity,

			contractFile[i].OrderDate,
			contractFile[i].StartDate,
			contractFile[i].EndDate,
			contractFile[i].Status,
			contractFile[i].PriceGuaranteeUntil,
			contractFile[i].PriceChangePlanned,

			contractFile[i].PriceValidSince,
			contractFile[i].CurrentBaseCosts,
			contractFile[i].CurrentKwhCosts,
			contractFile[i].CurrentBaseMargin,
			contractFile[i].CurrentKwhMargin,
			contractFile[i].CurrentBasePriceNet,
			contractFile[i].CurrentKwhPriceNet,
			contractFile[i].AnnualConsumption,
		)
		if err != nil {
			logEntry.WithError(err).Error("unable to create a contract information")
			return nil, usecase.ErrDomainInternal
		}
		err = imp.store.ContractInformation().Save(ci)
		if err != nil {
			logEntry.WithContext(ctx).WithError(err).Error("unable to store contract information in a db")
			return nil, usecase.ErrDatabaseInternal
		}
	}

	res.Imported = len(contractFile)

	return res, nil
}
