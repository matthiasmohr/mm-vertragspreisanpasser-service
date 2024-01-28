package app

import (
	"github.com/enercity/be-service-sample/pkg/repository/database/postgresql"
	"github.com/enercity/be-service-sample/pkg/server"
	"github.com/enercity/be-service-sample/pkg/server/handler"
	"github.com/enercity/be-service-sample/pkg/server/middleware"
	"github.com/enercity/be-service-sample/pkg/service/validation"
	contractInformationUsecases "github.com/enercity/be-service-sample/pkg/usecase/contractInformation"
	customerUsecases "github.com/enercity/be-service-sample/pkg/usecase/customer"
	priceChangeOrderUsecases "github.com/enercity/be-service-sample/pkg/usecase/priceChangeOrder"
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
	customerCreateUsecase := customerUsecases.NewCreator(store)
	customerLoaderUsecase := customerUsecases.NewLoader(store)
	customerFinderUsecase := customerUsecases.NewFinder(store)

	contractInformationCreateUseCase := contractInformationUsecases.NewCreator(store)
	contractInformationListerUseCase := contractInformationUsecases.NewLister(store)
	contractInformationFinderUseCase := contractInformationUsecases.NewFinder(store)

	priceChangeOrderCreateUseCase := priceChangeOrderUsecases.NewCreator(store)
	priceChangeOrderListerUseCase := priceChangeOrderUsecases.NewLister(store)
	priceChangeOrderFinderUseCase := priceChangeOrderUsecases.NewFinder(store)

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

	customerHandler := handler.NewCustomer(customerCreateUsecase, customerLoaderUsecase, customerFinderUsecase, lg)
	contractInformationHandler := handler.NewContractInformation(contractInformationCreateUseCase, contractInformationListerUseCase, contractInformationFinderUseCase, lg)
	priceChangeOrderHandler := handler.NewPriceChangeOrder(priceChangeOrderCreateUseCase, priceChangeOrderListerUseCase, priceChangeOrderFinderUseCase, lg)

	mmServer := server.New(cfg.Server, lg)

	mmServer.SetValidation(validator, &validation.Binder{})
	mmServer.SetErrorHandler(server.ErrorHandler(lg))

	// Set up routes.
	routes := mmServer.SetupRoutes()
	routes.Use(middleware.BuildContext())

	routes.GET("/version", statusHandler.Version)

	v1 := routes.Group("v1")

	customerGroup := v1.Group("/customer")
	customerGroup.GET("", customerHandler.Customers)
	customerGroup.GET("/find", customerHandler.Find)
	customerGroup.POST("", customerHandler.Create)

	contractInformationGroup := v1.Group("/contractinformation")
	contractInformationGroup.GET("", contractInformationHandler.List)
	contractInformationGroup.GET("/find", contractInformationHandler.Find)
	contractInformationGroup.POST("", contractInformationHandler.Create)

	priceChangeOrderGroup := v1.Group("/pricechangeorder")
	priceChangeOrderGroup.GET("", priceChangeOrderHandler.List)
	priceChangeOrderGroup.GET("/find", priceChangeOrderHandler.Find)
	priceChangeOrderGroup.POST("", priceChangeOrderHandler.Create)

	return errors.Wrap(mmServer.Run(), "error on customerServer.Run()")
}
