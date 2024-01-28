package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/enercity/be-service-sample/pkg/repository/database/postgresql"
	"github.com/enercity/be-service-sample/pkg/server"
	logger "github.com/enercity/lib-logger/v3"
	"github.com/spf13/viper"
)

// Config holds the general application configuration.
type Config struct {
	Logger   logger.Config
	Info     InfoConfig
	Server   server.Config
	Database postgresql.Config
}

// InfoConfig configures application information.
type InfoConfig struct {
	Version     string
	BuildDate   string
	Description string
	CommitHash  string
	CommitDate  string
	BuildBranch string
}

// NewConfig returns a *Config.
func NewConfig(version, buildDate, commitHash, commitDate, buildBranch, description string) (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.SetEnvPrefix("VERTRAGSPREISANPASSER")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		var configNotFoundErr viper.ConfigFileNotFoundError
		if errors.As(err, &configNotFoundErr) {
			return nil, configNotFoundErr
		}
		// if it is not a viper.ConfigFileNotFoundError
		return nil, fmt.Errorf("unknown error reading config: %w", err)
	}

	return &Config{
		Database: postgresql.Config{
			Host:         viper.GetString("database.host"),
			Port:         viper.GetInt("database.port"),
			DatabaseName: viper.GetString("database.name"),
			Username:     viper.GetString("database.user"),
			Password:     viper.GetString("database.password"),
			LogMode:      viper.GetBool("database.log_mode"),
		},
		Info: InfoConfig{
			Version:     version,
			BuildDate:   buildDate,
			Description: description,
			CommitHash:  commitHash,
			CommitDate:  commitDate,
			BuildBranch: buildBranch,
		},
		Logger: logger.Config{
			Level:  viper.GetString("logger.level"),
			Format: viper.GetString("logger.format"),
		},
		Server: server.Config{
			Address:      viper.GetString("server.address"),
			ReadTimeout:  viper.GetDuration("server.read_timeout"),
			WriteTimeout: viper.GetDuration("server.write_timeout"),
			Debug:        viper.GetBool("server.debug"),
			Version:      version,
			CORS: server.CORSConfig{
				AllowCredentials: viper.GetBool("server.cors.allow_credentials"),
				Headers:          viper.GetStringSlice("server.cors.headers"),
				Methods:          viper.GetStringSlice("server.cors.methods"),
				Origins:          viper.GetStringSlice("server.cors.origins"),
			},
		},
	}, nil
}
