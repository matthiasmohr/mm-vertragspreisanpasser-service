package middleware

import (
	"github.com/enercity/be-service-sample/pkg/contextbuilder"
	"github.com/labstack/echo/v4"
)

const (
	ContextKeyNativeContext = "native-context"
)

func BuildContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			ctx := echoCtx.Request().Context()

			reqID := echoCtx.Response().Header().Get(echo.HeaderXRequestID)
			ctx = contextbuilder.NewContextWithRequestID(ctx, reqID)

			echoCtx.Set(ContextKeyNativeContext, ctx)

			return next(echoCtx)
		}
	}
}
