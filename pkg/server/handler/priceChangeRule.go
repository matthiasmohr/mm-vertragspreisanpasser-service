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
	priceChangeRuleLister interface {
		List(
			ctx context.Context, logEntry logger.Entry, req *dto.ListPriceChangeRuleRequest,
		) (*dto.ListPriceChangeRuleResponse, error)
	}

	priceChangeRuleCreater interface {
		Create(
			ctx context.Context, logEntry logger.Entry, req *dto.CreatePriceChangeRuleRequest,
		) error
	}
	priceChangeRuleRemover interface {
		Remove(
			ctx context.Context, logEntry logger.Entry, req *dto.RemovePriceChangeRuleRequest) error
	}
)

type PriceChangeRule struct {
	priceChangeRuleLister  priceChangeRuleLister
	priceChangeRuleCreater priceChangeRuleCreater
	priceChangeRuleRemover priceChangeRuleRemover
	logger                 logger.Logger
}

func NewPriceChangeRule(
	priceChangeRuleListerUsecase priceChangeRuleLister,
	priceChangeRuleCreaterUsecase priceChangeRuleCreater,
	priceChangeRuleRemoverUsecase priceChangeRuleRemover,
	lg logger.Logger,
) *PriceChangeRule {
	return &PriceChangeRule{
		priceChangeRuleLister:  priceChangeRuleListerUsecase,
		priceChangeRuleCreater: priceChangeRuleCreaterUsecase,
		priceChangeRuleRemover: priceChangeRuleRemoverUsecase,
		logger:                 lg,
	}
}

// Contract Informations returns a list of existing Rules.
// @Summary Rules returns a list of existing Rules.
// @Description Rules returns a list of existing Rules.
// @Tags Rule
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.ListRulesResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/contractinformation [get].
func (ci *PriceChangeRule) List(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.ListPriceChangeRuleRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	priceChangeRules, err := ci.priceChangeRuleLister.List(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.JSON(http.StatusOK, priceChangeRules)
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
func (ci *PriceChangeRule) Create(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	req := &dto.CreatePriceChangeRuleRequest{}
	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	if err := ci.priceChangeRuleCreater.Create(ctx, logEntry, req); err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.NoContent(http.StatusOK)
}

func (ci *PriceChangeRule) Delete(echoCtx echo.Context) error {
	ctx, err := loadContext(echoCtx)
	if err != nil {
		return server.NewHTTPError(err)
	}

	logEntry := ci.logger.WithContext(ctx)

	id := echoCtx.QueryParam("id")
	req := &dto.RemovePriceChangeRuleRequest{id}

	if err := bindAndValidate(req, echoCtx); err != nil {
		return server.NewHTTPError(err)
	}

	err = ci.priceChangeRuleRemover.Remove(ctx, logEntry, req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return echoCtx.NoContent(http.StatusOK)
}
