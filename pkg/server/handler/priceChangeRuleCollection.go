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
	priceChangeRuleCollectionLister interface {
		List(
			ctx context.Context, logEntry logger.Entry, req *dto.ListPriceChangeRuleCollectionRequest,
		) (*dto.ListPriceChangeRuleCollectionResponse, error)
	}

	priceChangeRuleCollectionCreater interface {
		Create(
			ctx context.Context, logEntry logger.Entry, req *dto.CreatePriceChangeRuleCollectionRequest,
		) error
	}
	priceChangeRuleCollectionExecuter interface {
		Execute(
			ctx context.Context, logEntry logger.Entry, req *dto.ExecutePriceChangeRuleCollectionRequest) error
	}
	priceChangeRuleCollectionGetter interface {
		Get(
			ctx context.Context, logEntry logger.Entry, req *dto.GetPriceChangeRuleCollectionRequest) (*dto.GetPriceChangeRuleCollectionResponse, error)
	}
)

type PriceChangeRuleCollection struct {
	priceChangeRuleCollectionLister   priceChangeRuleCollectionLister
	priceChangeRuleCollectionCreater  priceChangeRuleCollectionCreater
	priceChangeRuleCollectionExecuter priceChangeRuleCollectionExecuter
	priceChangeRuleCollectionGetter   priceChangeRuleCollectionGetter
	logger                            logger.Logger
}

func NewPriceChangeRuleCollection(
	priceChangeRuleCollectionListerUsecase priceChangeRuleCollectionLister,
	priceChangeRuleCollectionCreaterUsecase priceChangeRuleCollectionCreater,
	priceChangeRuleCollectionExecuterUsecase priceChangeRuleCollectionExecuter,
	priceChangeRuleCollectionGetterUsecase priceChangeRuleCollectionGetter,
	lg logger.Logger,
) *PriceChangeRuleCollection {
	return &PriceChangeRuleCollection{
		priceChangeRuleCollectionLister:   priceChangeRuleCollectionListerUsecase,
		priceChangeRuleCollectionCreater:  priceChangeRuleCollectionCreaterUsecase,
		priceChangeRuleCollectionExecuter: priceChangeRuleCollectionExecuterUsecase,
		priceChangeRuleCollectionGetter:   priceChangeRuleCollectionGetterUsecase,
		logger:                            lg,
	}
}

// Contract Informations returns a list of existing RuleCollections.
// @Summary RuleCollections returns a list of existing RuleCollections.
// @Description RuleCollections returns a list of existing RuleCollections.
// @Tags RuleCollection
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.ListRuleCollectionsResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/contractinformation [get].
func (ci *PriceChangeRuleCollection) List(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.ListPriceChangeRuleCollectionRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	priceChangeRuleCollections, err := ci.priceChangeRuleCollectionLister.List(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, priceChangeRuleCollections)
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
func (ci *PriceChangeRuleCollection) Create(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.CreatePriceChangeRuleCollectionRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	if err := ci.priceChangeRuleCollectionCreater.Create(ctx, logEntry, req); err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.NoContent(http.StatusOK)
}

func (ci *PriceChangeRuleCollection) Execute(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	id := echoCtx.QueryParam("id")
	req := &dto.ExecutePriceChangeRuleCollectionRequest{id}

	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	err = ci.priceChangeRuleCollectionExecuter.Execute(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.NoContent(http.StatusOK)
}

func (ci *PriceChangeRuleCollection) Get(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	id := echoCtx.QueryParam("id")
	req := &dto.GetPriceChangeRuleCollectionRequest{id}

	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	priceChangeRuleCollection, err := ci.priceChangeRuleCollectionGetter.Get(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, priceChangeRuleCollection)
}
