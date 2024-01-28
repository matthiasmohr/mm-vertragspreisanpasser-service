package handler

import (
	"context"
	"net/http"

	"github.com/enercity/be-service-sample/pkg/model/dto"
	"github.com/enercity/be-service-sample/pkg/server"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/labstack/echo/v4"
)

type (
	priceChangeOrderCreater interface {
		Create(ctx context.Context, logEntry logger.Entry, req *dto.CreatePriceChangeOrderRequest) error
	}

	priceChangeOrderLister interface {
		List(
			ctx context.Context, logEntry logger.Entry, req *dto.ListPriceChangeOrderRequest,
		) (*dto.ListPriceChangeOrderResponse, error)
	}

	priceChangeOrderFinder interface {
		Find(
			ctx context.Context, logEntry logger.Entry, req *dto.FindPriceChangeOrderRequest,
		) (*dto.FindPriceChangeOrderResponse, error)
	}
)

type PriceChangeOrder struct {
	priceChangeOrderCreater priceChangeOrderCreater
	priceChangeOrderLister  priceChangeOrderLister
	priceChangeOrderFinder  priceChangeOrderFinder
	logger                  logger.Logger
}

func NewPriceChangeOrder(
	priceChangeOrderCreatorUsecase priceChangeOrderCreater,
	priceChangeOrderListerUsecase priceChangeOrderLister,
	priceChangeOrderFinderUsecase priceChangeOrderFinder,
	lg logger.Logger,
) *PriceChangeOrder {
	return &PriceChangeOrder{
		priceChangeOrderCreater: priceChangeOrderCreatorUsecase,
		priceChangeOrderLister:  priceChangeOrderListerUsecase,
		priceChangeOrderFinder:  priceChangeOrderFinderUsecase,
		logger:                  lg,
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
func (ci *PriceChangeOrder) Create(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.CreatePriceChangeOrderRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	if err := ci.priceChangeOrderCreater.Create(ctx, logEntry, req); err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.NoContent(http.StatusOK)
}

// Contract Informations returns a list of existing customers.
// @Summary Customers returns a list of existing customers.
// @Description Customers returns a list of existing customers.
// @Tags Customer
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.ListCustomersResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/contractinformation [get].
func (ci *PriceChangeOrder) List(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.ListPriceChangeOrderRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	priceChangeOrders, err := ci.priceChangeOrderLister.List(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, priceChangeOrders)
}

// Customers returns a list of existing customers.
// @Summary Customers returns a list of existing customers.
// @Description Customers returns a list of existing customers.
// @Tags Customer
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Param fisrtName query string false "First Name"
// @Param lastName query string false "Last Name"
// @Param email query string false "Customer Email"
// @Success 200 {object} dto.ListCustomersResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/find [get].
func (ci *PriceChangeOrder) Find(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	productSerialNumber := echoCtx.QueryParam("productSerialNumber")

	req := &dto.FindPriceChangeOrderRequest{ProductSerialNumber: &productSerialNumber}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	priceChangeOrders, err := ci.priceChangeOrderFinder.Find(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, priceChangeOrders)
}