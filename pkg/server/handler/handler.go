package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/enercity/be-service-sample/pkg/server"
	"github.com/enercity/be-service-sample/pkg/server/middleware"
	"github.com/labstack/echo/v4"
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
