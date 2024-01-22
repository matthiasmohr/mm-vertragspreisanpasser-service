package server

import (
	"errors"
	"net/http"

	"github.com/enercity/be-service-sample/pkg/service/validation"
	"github.com/enercity/be-service-sample/pkg/usecase"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/labstack/echo/v4"
)

// HTTPError represents a handler error.
type HTTPError struct {
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Code      int         `json:"code,omitempty"` // Unique application error code.
	HTTPCode  int         `json:"-"`              // HTTP status code.
	Err       error       `json:"-"`
	RequestID string      `json:"request_id,omitempty"`
}

// Error returns the error message.
func (e *HTTPError) Error() string {
	return e.Message
}

// GenericInternalServerError represents an internal server error with a generic error message.
func GenericInternalServerError() *HTTPError {
	return &HTTPError{
		HTTPCode: http.StatusInternalServerError,
		Message:  "internal server error",
	}
}

// NewHTTPError initializes an internal server error with the provided error.
func NewHTTPError(err error) error {
	httpErr := GenericInternalServerError()
	httpErr.Err = err

	var uErr *usecase.Error
	if errors.As(err, &uErr) {
		switch {
		case errors.Is(uErr, usecase.ErrDatabaseCustomerNotFound):
			httpErr.Message = uErr.Message
			httpErr.HTTPCode = http.StatusNotFound
		default:
			httpErr.Message = uErr.Message
		}
	}

	return httpErr
}

// ErrorHandler is a custom error handler able to differentiate between error types.
func ErrorHandler(lg logger.Logger) echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		// Prevent double execution of the error handler.
		if ctx.Response().Committed {
			return
		}

		id := ctx.Request().Header.Get(echo.HeaderXRequestID)
		if id == "" {
			id = ctx.Response().Header().Get(echo.HeaderXRequestID)
		}

		logEntry := lg.WithField("request_id", id)

		httpErr := GenericInternalServerError()

		var (
			echoErr *echo.HTTPError
			appErr  *HTTPError
		)

		if errors.As(err, &echoErr) {
			// Echo middleware errors (e.g. basic auth middleware).
			logEntry.WithField("origin", "echo").WithError(echoErr).Info("handler received error")
			httpErr.HTTPCode = echoErr.Code

			httpErr.Message = echoErr.Error()
			if msg, ok := echoErr.Message.(string); ok {
				httpErr.Message = msg
			}
		}

		if errors.As(err, &appErr) {
			// Application specific errors.
			logEntry.WithField("origin", "app").WithError(appErr.Err).Info("handler received error")

			httpErr.HTTPCode = appErr.HTTPCode
			httpErr.Message = appErr.Message

			var vErr *validation.Error
			if errors.As(appErr.Err, &vErr) {
				httpErr.Message = "validation error"
				httpErr.Details = vErr
				httpErr.HTTPCode = http.StatusBadRequest
			}
		}

		httpErr.RequestID = ctx.Response().Header().Get(echo.HeaderXRequestID)

		var responseErr error
		if ctx.Request().Method == http.MethodHead {
			responseErr = ctx.NoContent(httpErr.HTTPCode)
		} else {
			responseErr = ctx.JSON(httpErr.HTTPCode, httpErr)
		}

		if responseErr != nil {
			logEntry.WithError(err).Warning("unable to send response")
		}
	}
}
