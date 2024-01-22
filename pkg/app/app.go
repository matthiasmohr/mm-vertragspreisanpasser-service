package app

import (
	"github.com/enercity/be-service-sample/pkg/repository/database/postgresql"
	"github.com/enercity/be-service-sample/pkg/server"
	"github.com/enercity/be-service-sample/pkg/server/handler"
	"github.com/enercity/be-service-sample/pkg/server/middleware"
	"github.com/enercity/be-service-sample/pkg/service/validation"
	customerUsecases "github.com/enercity/be-service-sample/pkg/usecase/customer"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/pkg/errors"
)

// Run runs the application.
func Run(cfg *Config, lg logger.Logger) error {
	// Init repository.
	store, err := postgresql.New(cfg.Database, lg)
	if err != nil {
		lg.WithError(err).Error("Couldn't connect to the repository")

		return errors.Wrap(err, "could not initialize repository")
	}

	// Init usecases
	customerUsecase := customerUsecases.NewCreator(store)
	customerLoaderUsecase := customerUsecases.NewLoader(store)
	customerFinderUsecase := customerUsecases.NewFinder(store)

	validator, err := validation.NewValidator()
	if err != nil {
		lg.WithError(err).Error("Couldn't create validator")

		return errors.Wrap(err, "error on creating new Validator")
	}

	// Init handler
	statusHandler := handler.NewStatus(
		cfg.Info.Version,
		cfg.Info.BuildDate,
		cfg.Info.Description,
		cfg.Info.CommitHash,
		cfg.Info.CommitDate,
		cfg.Info.BuildBranch,
	)

	customerHandler := handler.NewCustomer(customerUsecase, customerLoaderUsecase, customerFinderUsecase, lg)

	customerServer := server.New(cfg.Server, lg)

	customerServer.SetValidation(validator, &validation.Binder{})
	customerServer.SetErrorHandler(server.ErrorHandler(lg))

	// Set up routes.
	routes := customerServer.SetupRoutes()
	routes.Use(middleware.BuildContext())

	routes.GET("/version", statusHandler.Version)

	v1 := routes.Group("v1")
	customerGroup := v1.Group("/customer")

	customerGroup.POST("", customerHandler.Create)
	customerGroup.GET("", customerHandler.Customers)
	customerGroup.GET("/find", customerHandler.Find)

	return errors.Wrap(customerServer.Run(), "error on customerServer.Run()")
}
