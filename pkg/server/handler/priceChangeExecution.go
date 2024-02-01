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
	priceChangeExecutionLister interface {
		List(
			ctx context.Context, logEntry logger.Entry, req *dto.ListPriceChangeExecutionRequest,
		) (*dto.ListPriceChangeExecutionResponse, error)
	}

	priceChangeExecutionFinder interface {
		Find(
			ctx context.Context, logEntry logger.Entry, req *dto.FindPriceChangeExecutionRequest,
		) (*dto.FindPriceChangeExecutionResponse, error)
	}
	priceChangeExecutionExecuter interface {
		Execute(
			ctx context.Context, logEntry logger.Entry, req *dto.ExecutePriceChangeExecutionRequest) error
	}
)

type PriceChangeExecution struct {
	priceChangeExecutionLister   priceChangeExecutionLister
	priceChangeExecutionFinder   priceChangeExecutionFinder
	priceChangeExecutionExecuter priceChangeExecutionExecuter
	logger                       logger.Logger
}

func NewPriceChangeExecution(
	priceChangeExecutionListerUsecase priceChangeExecutionLister,
	priceChangeExecutionFinderUsecase priceChangeExecutionFinder,
	priceChangeExecutionExecuterUsecase priceChangeExecutionExecuter,
	lg logger.Logger,
) *PriceChangeExecution {
	return &PriceChangeExecution{
		priceChangeExecutionLister:   priceChangeExecutionListerUsecase,
		priceChangeExecutionFinder:   priceChangeExecutionFinderUsecase,
		priceChangeExecutionExecuter: priceChangeExecutionExecuterUsecase,
		logger:                       lg,
	}
}

// Contract Informations returns a list of existing Executions.
// @Summary Executions returns a list of existing Executions.
// @Description Executions returns a list of existing Executions.
// @Tags Execution
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.ListExecutionsResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/contractinformation [get].
func (ci *PriceChangeExecution) List(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.ListPriceChangeExecutionRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	priceChangeExecutions, err := ci.priceChangeExecutionLister.List(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, priceChangeExecutions)
}

// Executions returns a list of existing Executions.
// @Summary Executions returns a list of existing Executions.
// @Description Executions returns a list of existing Executions.
// @Tags Execution
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Param fisrtName query string false "First Name"
// @Param lastName query string false "Last Name"
// @Param email query string false "Execution Email"
// @Success 200 {object} dto.ListExecutionsResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/find [get].
func (ci *PriceChangeExecution) Find(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	productSerialNumber := echoCtx.QueryParam("productSerialNumber")

	req := &dto.FindPriceChangeExecutionRequest{ProductSerialNumber: &productSerialNumber}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	priceChangeExecutions, err := ci.priceChangeExecutionFinder.Find(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, priceChangeExecutions)
}

func (ci *PriceChangeExecution) Execute(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	id := echoCtx.QueryParam("id")

	req := &dto.ExecutePriceChangeExecutionRequest{id}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	err = ci.priceChangeExecutionExecuter.Execute(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.NoContent(http.StatusOK)
}
