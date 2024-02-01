package handler

import (
	"context"
	"net/http"

	logger "github.com/enercity/lib-logger/v3"
	"github.com/labstack/echo/v4"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/model/dto"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/server"
)

type (
	contractInformationCreater interface {
		Create(ctx context.Context, logEntry logger.Entry, req *dto.CreateContractInformationRequest) error
	}

	contractInformationLister interface {
		List(
			ctx context.Context, logEntry logger.Entry, req *dto.ListContractInformationsRequest,
		) (*dto.ListContractInformationsResponse, error)
	}

	contractInformationFinder interface {
		Find(
			ctx context.Context, logEntry logger.Entry, req *dto.FindContractInformationRequest,
		) (*dto.FindContractInformationsResponse, error)
	}
)

type ContractInformation struct {
	contractInformationCreater contractInformationCreater
	contractInformationLister  contractInformationLister
	contractInformationFinder  contractInformationFinder
	logger                     logger.Logger
}

func NewContractInformation(
	contractInformationCreatorUsecase contractInformationCreater,
	contractInformationListerUsecase contractInformationLister,
	contractInformationFinderUsecase contractInformationFinder,
	lg logger.Logger,
) *ContractInformation {
	return &ContractInformation{
		contractInformationCreater: contractInformationCreatorUsecase,
		contractInformationLister:  contractInformationListerUsecase,
		contractInformationFinder:  contractInformationFinderUsecase,
		logger:                     lg,
	}
}

// Create contractInformation creates a new contractInformation.
// @Summary creates a new contractInformation
// @Description creates a new contractInformation
// @Tags ContractInformation
// @Accept json
// @Produce json
// @Param xx body string true "xxx"
// @Param yy body string true "yyy"
// @Param zz body string true "zzz"
// @Success 200 {string} string "Validation passed and blocking order was sent"
// @Failure 400 {string} string "Bad request payload"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/contractinformation [post].
func (ci *ContractInformation) Create(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.CreateContractInformationRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	if err := ci.contractInformationCreater.Create(ctx, logEntry, req); err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.NoContent(http.StatusOK)
}

// Contract Informations returns a list of existing ContractInformations.
// @Summary ContractInformations returns a list of existing ContractInformations.
// @Description ContractInformations returns a list of existing ContractInformations.
// @Tags ContractInformation
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.ListContractInformationsResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/contractinformation [get].
func (ci *ContractInformation) List(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.ListContractInformationsRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	contractinformations, err := ci.contractInformationLister.List(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, contractinformations)
}

// ContractInformations returns a list of existing ContractInformations.
// @Summary ContractInformations returns a list of existing ContractInformations.
// @Description ContractInformations returns a list of existing ContractInformations.
// @Tags ContractInformation
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Param fisrtName query string false "First Name"
// @Param lastName query string false "Last Name"
// @Param email query string false "ContractInformation Email"
// @Success 200 {object} dto.ListContractInformationsResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/find [get].
func (ci *ContractInformation) Find(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	mba := echoCtx.QueryParam("mba")
	productSerialNumber := echoCtx.QueryParam("productSerialNumber")

	req := &dto.FindContractInformationRequest{Mba: &mba, ProductSerialNumber: &productSerialNumber}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	contractinformations, err := ci.contractInformationFinder.Find(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, contractinformations)
}
