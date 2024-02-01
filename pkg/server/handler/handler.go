package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/server"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/server/middleware"
)

func bindAndValidate(request interface{}, ctx echo.Context) error {
	if err := ctx.Bind(request); err != nil {
		return fmt.Errorf("error while binding request: %w", err)
	}

	if err := ctx.Validate(request); err != nil {
		return fmt.Errorf("error while validating request: %w", err)
	}

	return nil
}

func loadContext(echoCtx echo.Context) (context.Context, error) {
	ctx, ok := echoCtx.Get(middleware.ContextKeyNativeContext).(context.Context)

	if !ok {
		return nil, server.NewHTTPError(errors.New("unable to load context.Context"))
	}

	return ctx, nil
}
