package server

import (
	"context"
	"database/sql"

	"github.com/kyrare/ya-diplom-2/internal/infrastructure/connection"
	"github.com/kyrare/ya-diplom-2/internal/service/logger"
)

type App struct {
	config *Config
	logger *logger.Logger
	db     *sql.DB
}

func NewApp(
	config *Config,
	logger *logger.Logger,
) *App {
	return &App{
		config: config,
		logger: logger,
	}
}

func (app *App) Configure(ctx context.Context) error {
	db, err := connection.NewPostgresql(
		app.config.DB.Name,
		app.config.DB.Host,
		app.config.DB.Port,
		app.config.DB.User,
		app.config.DB.Password,
		app.logger,
	)

	if err != nil {
		return err
	}
	app.db = db

	return nil
}

func (app *App) Run(ctx context.Context) error {

	return nil
}
