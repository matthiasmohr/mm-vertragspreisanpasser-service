package main

import (
	logger "github.com/enercity/lib-logger/v3"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/app"
	"github.com/matthiasmohr/mm-vertragspreisanpasser-service/pkg/contextbuilder"
	"log"
	"os"
)

// Application metadata that is set at compile time. See drone.yml from project template for details.
// Convention: https://lynqtech.atlassian.net/wiki/spaces/EL/pages/398270644/Conventions#HTTP-GET-%2Fversion
var (
	version     = "/"
	buildDate   = "/"
	commitHash  = "/"
	commitDate  = "/"
	buildBranch = "/"
	// description should be hard coded here. It's optional, so you may remove it.
	description = "Customer Sample Service"
)

func main() {
	cfg, err := app.NewConfig(version, buildDate, commitHash, commitDate, buildBranch, description)
	if err != nil {
		log.Printf("could not load config: %s\n", err.Error())
		os.Exit(1)
	}

	lg := logger.NewLogger(
		logger.Config{
			Format: cfg.Logger.Format,
			Level:  cfg.Logger.Level,
			DefaultFields: logger.DefaultFields{
				Build:    cfg.Info.CommitHash,
				Version:  cfg.Info.Version,
				Service:  "sample_service",
				ClientID: "lynqtech",
			},
			ContextFields: []string{"request_id"},
			ContextKeyFunc: func(key interface{}) interface{} {
				return contextbuilder.ContextKey(key.(string))
			},
		},
	)

	lg.WithFields(
		map[string]interface{}{
			"version":   cfg.Info.Version,
			"buildDate": cfg.Info.BuildDate,
		},
	).Info(cfg.Info.Description)

	err = app.Run(cfg, lg)

	if err != nil {
		lg.WithError(err).Error("app.Run() exit with error")
		os.Exit(1)
	}
}
