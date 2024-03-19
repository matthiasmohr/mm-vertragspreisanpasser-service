package app

import (
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/repository/database/postgresql"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/server"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/server/handler"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/server/middleware"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/service/validation"
	contractInformationUsecases "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase/contractInformation"
	customerUsecases "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase/customer"
	priceChangeExecutionUsecases "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase/priceChangeExecution"
	priceChangeOrderUsecases "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase/priceChangeOrder"
	priceChangeRuleUsecases "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase/priceChangeRule"
	priceChangeRuleCollectionUsecases "github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/usecase/priceChangeRuleCollection"
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
	contractInformationImporterUseCase := contractInformationUsecases.NewImporter(store)
	contractInformationGetterUseCase := contractInformationUsecases.NewGetter(store)

	priceChangeRuleCollectionListerUseCase := priceChangeRuleCollectionUsecases.NewLister(store)
	priceChangeRuleCollectionCreaterUseCase := priceChangeRuleCollectionUsecases.NewCreator(store)
	priceChangeRuleCollectionExecuterUseCase := priceChangeRuleCollectionUsecases.NewExecuter(store)
	priceChangeRuleCollectionGetterUseCase := priceChangeRuleCollectionUsecases.NewGetter(store)

	priceChangeRuleListerUseCase := priceChangeRuleUsecases.NewLister(store)
	priceChangeRuleCreaterUseCase := priceChangeRuleUsecases.NewCreator(store)
	priceChangeRuleRemoverUseCase := priceChangeRuleUsecases.NewRemover(store)

	priceChangeOrderCreateUseCase := priceChangeOrderUsecases.NewCreator(store)
	priceChangeOrderListerUseCase := priceChangeOrderUsecases.NewLister(store)
	priceChangeOrderFinderUseCase := priceChangeOrderUsecases.NewFinder(store)
	priceChangeOrderExecuteUseCase := priceChangeOrderUsecases.NewExecuter(store)

	priceChangeExecutionListerUseCase := priceChangeExecutionUsecases.NewLister(store)
	priceChangeExecutionFinderUseCase := priceChangeExecutionUsecases.NewFinder(store)
	priceChangeExecutionExecuteUseCase := priceChangeExecutionUsecases.NewExecuter(store)

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
	contractInformationHandler := handler.NewContractInformation(contractInformationCreateUseCase, contractInformationListerUseCase, contractInformationFinderUseCase, contractInformationImporterUseCase, contractInformationGetterUseCase, lg)
	priceChangeRuleCollectionHandler := handler.NewPriceChangeRuleCollection(priceChangeRuleCollectionListerUseCase, priceChangeRuleCollectionCreaterUseCase, priceChangeRuleCollectionExecuterUseCase, priceChangeRuleCollectionGetterUseCase, lg)
	priceChangeRuleHandler := handler.NewPriceChangeRule(priceChangeRuleListerUseCase, priceChangeRuleCreaterUseCase, priceChangeRuleRemoverUseCase, lg)
	priceChangeOrderHandler := handler.NewPriceChangeOrder(priceChangeOrderCreateUseCase, priceChangeOrderListerUseCase, priceChangeOrderFinderUseCase, priceChangeOrderExecuteUseCase, lg)
	priceChangeExecutionHandler := handler.NewPriceChangeExecution(priceChangeExecutionListerUseCase, priceChangeExecutionFinderUseCase, priceChangeExecutionExecuteUseCase, lg)

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
	contractInformationGroup.GET("/:id", contractInformationHandler.Get)
	contractInformationGroup.GET("/find", contractInformationHandler.Find)
	contractInformationGroup.POST("", contractInformationHandler.Create)
	contractInformationGroup.POST("/import", contractInformationHandler.Import)

	priceChangeRuleCollectionGroup := v1.Group("/pricechangerulecollection")
	priceChangeRuleCollectionGroup.GET("", priceChangeRuleCollectionHandler.List)
	priceChangeRuleCollectionGroup.GET("/get", priceChangeRuleCollectionHandler.Get)
	priceChangeRuleCollectionGroup.POST("", priceChangeRuleCollectionHandler.Create)
	priceChangeRuleCollectionGroup.POST("/execute", priceChangeRuleCollectionHandler.Execute)

	priceChangeRuleGroup := v1.Group("/pricechangerule")
	priceChangeRuleGroup.GET("", priceChangeRuleHandler.List)
	priceChangeRuleGroup.POST("", priceChangeRuleHandler.Create)
	priceChangeRuleGroup.DELETE("", priceChangeRuleHandler.Delete)

	priceChangeOrderGroup := v1.Group("/pricechangeorder")
	priceChangeOrderGroup.GET("", priceChangeOrderHandler.List)
	priceChangeOrderGroup.GET("/find", priceChangeOrderHandler.Find)
	priceChangeOrderGroup.POST("", priceChangeOrderHandler.Create)
	priceChangeOrderGroup.POST("/execute", priceChangeOrderHandler.Execute)

	priceChangeExecutionGroup := v1.Group("/pricechangeexecution")
	priceChangeExecutionGroup.GET("", priceChangeExecutionHandler.List)
	priceChangeExecutionGroup.GET("/find", priceChangeExecutionHandler.Find)
	priceChangeExecutionGroup.POST("/execute", priceChangeExecutionHandler.Execute)

	return errors.Wrap(mmServer.Run(), "error on customerServer.Run()")
}
