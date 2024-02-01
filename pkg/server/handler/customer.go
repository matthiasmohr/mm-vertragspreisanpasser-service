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
	customerCreater interface {
		Create(ctx context.Context, logEntry logger.Entry, req *dto.CreateCustomerRequest) error
	}

	customerLoader interface {
		Load(
			ctx context.Context, logEntry logger.Entry, req *dto.ListCustomersRequest,
		) (*dto.ListCustomersResponse, error)
	}

	customerFinder interface {
		Find(
			ctx context.Context, logEntry logger.Entry, req *dto.CustomerFindRequest,
		) (*dto.ListCustomersResponse, error)
	}
)

type Customer struct {
	customerCreater customerCreater
	customerLoader  customerLoader
	customerFinder  customerFinder
	logger          logger.Logger
}

func NewCustomer(
	customerCreatorUsecase customerCreater,
	customerLoaderUsecase customerLoader,
	customerFinderUsecase customerFinder,
	lg logger.Logger,
) *Customer {
	return &Customer{
		customerCreater: customerCreatorUsecase,
		customerLoader:  customerLoaderUsecase,
		customerFinder:  customerFinderUsecase,
		logger:          lg,
	}
}

// Create customer creates a new customer.
// @Summary Creates a new customer
// @Description Creates a new customer.
// @Tags Customer
// @Accept json
// @Produce json
// @Param firstName body string true "Customer Firstname"
// @Param lastName body string true "Customer Lastname"
// @Param email body string true "Customer Email"
// @Success 200 {string} string "Validation passed and blocking order was sent"
// @Failure 400 {string} string "Bad request payload"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/customer [post].
func (c *Customer) Create(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := c.logger.WithContext(ctx)

	req := &dto.CreateCustomerRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	if err := c.customerCreater.Create(ctx, logEntry, req); err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.NoContent(http.StatusOK)
}

// Customers returns a list of existing customers.
// @Summary Customers returns a list of existing customers.
// @Description Customers returns a list of existing customers.
// @Tags Customer
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.ListCustomersResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/customers [get].
func (c *Customer) Customers(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := c.logger.WithContext(ctx)

	req := &dto.ListCustomersRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	customers, err := c.customerLoader.Load(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, customers)
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
func (c *Customer) Find(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := c.logger.WithContext(ctx)

	req := &dto.CustomerFindRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	customers, err := c.customerFinder.Find(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, customers)
}
