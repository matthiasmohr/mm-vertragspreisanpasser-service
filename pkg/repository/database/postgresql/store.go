package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/enercity/be-service-sample/pkg/repository"
	logger "github.com/enercity/lib-logger/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB

	customer repository.Customer
}

func New(cfg Config, lg logger.Logger) (repository.Store, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DatabaseName,
	)), &gorm.Config{
		Logger: gormLogger(cfg, lg),
	})
	if err != nil {
		return nil, fmt.Errorf("create new repository: %w", err)
	}

	return &Repository{
		db: db,

		customer: newCustomer(db),
	}, nil
}

type dbLogger struct {
	logger.Logger
}

func (dl dbLogger) Printf(msg string, args ...interface{}) {
	dl.Debug(msg, args)
}

func gormLogger(cfg Config, lg logger.Logger) glogger.Interface {
	level := glogger.Silent

	if cfg.LogMode {
		switch cfg.Level {
		case "trace", "debug", "info":
			level = glogger.Info
		case "warning":
			level = glogger.Warn
		case "error", "fatal", "panic":
			level = glogger.Error
		}
	}

	return glogger.New(
		dbLogger{lg},
		glogger.Config{
			SlowThreshold: time.Second,
			LogLevel:      level,
			Colorful:      false,
		},
	)
}

func newWithDB(db *gorm.DB) *Repository {
	return &Repository{
		db: db,

		customer: newCustomer(db),
	}
}

// BeginTransaction starts the new database transaction.
func (r *Repository) BeginTransaction() (repository.Store, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return newWithDB(tx), nil
}

// Commit commits all the changes from the current transaction to the database.
func (r *Repository) Commit() error {
	return r.db.Commit().Error
}

// Rollback rolls back all the changes made in the current database transaction.
func (r *Repository) Rollback() error {
	if err := r.db.Rollback().Error; !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	return nil
}

// Case returns the case repository.
func (r *Repository) Customer() repository.Customer {
	return r.customer
}
