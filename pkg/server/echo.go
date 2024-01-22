package server

import (
	"strconv"

	logger "github.com/enercity/lib-logger/v3"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

const baseInt = 10

// REST represents a REST server.
type REST struct {
	address string
	engine  *echo.Echo
	logger  logger.Logger
}

// New creates a new REST web server.
// @title Blocking service
// @description Service that deals with customers meter blocking.
// @license.name Closed
// @schemes https
// @host localhost:8080
func New(cfg Config, lg logger.Logger) *REST {
	docs.SwaggerInfo.Version = cfg.Version

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(loggerMiddleware(lg))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: cfg.CORS.AllowCredentials,
		AllowHeaders:     cfg.CORS.Headers,
		AllowMethods:     cfg.CORS.Methods,
		AllowOrigins:     cfg.CORS.Origins,
	}))

	// Set up swagger documentation.
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Server.ReadTimeout = cfg.ReadTimeout
	e.Server.WriteTimeout = cfg.WriteTimeout

	e.Debug = cfg.Debug
	e.HideBanner = true

	e.Server.Addr = cfg.Address

	server := &REST{
		address: cfg.Address,
		engine:  e,
		logger:  lg,
	}

	return server
}

// Since logrus & echo loggers are incompatible (interface-wise).
func loggerMiddleware(lg logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			bytesIn := "0"

			contentLength := req.Header.Get(echo.HeaderContentLength)
			if contentLength != "" {
				bytesIn = contentLength
			}

			// Proceed handling the request, but return afterwards to log what happened.
			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}

			logFields := map[string]interface{}{
				"request_id": id,
				"host":       req.Host,
				"remote_ip":  c.RealIP(),
				"protocol":   req.Proto,
				"user_agent": req.UserAgent(),
				"status":     res.Status,
				"uri":        req.RequestURI,
				"method":     req.Method,
				"bytes_in":   bytesIn,
				"bytes_out":  strconv.FormatInt(res.Size, baseInt),
			}

			lg.WithFields(logFields).Info()

			return err
		}
	}
}

// SetErrorHandler sets the error handler.
func (r *REST) SetErrorHandler(errorHandler echo.HTTPErrorHandler) {
	r.engine.HTTPErrorHandler = errorHandler
}

// SetValidation sets the validator and binder that validate incoming payload.
func (r *REST) SetValidation(validator echo.Validator, binder echo.Binder) {
	r.engine.Validator = validator
	r.engine.Binder = binder
}

// SetBasicAuth sets basic authentication.
func (r *REST) SetBasicAuth(validator middleware.BasicAuthValidator, skipper middleware.Skipper) {
	r.engine.Use(
		middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
			Validator: validator,
			Skipper:   skipper,
		}),
	)
}

// SetupRoutes setups the server routes.
func (r *REST) SetupRoutes() *echo.Group {
	return r.engine.Group("")
}

// Run runs the REST server.
func (r *REST) Run() error {
	r.logger.WithField("address", r.address).Info("run server")

	return errors.Wrap(gracehttp.Serve(r.engine.Server), "error on gracehttp.Serve")
}
