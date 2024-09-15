package server

import (
	"context"
	"database/sql"

	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/app/services"
	"github.com/kyrare/ya-diplom-2/internal/infrastructure/db/postgres"
)

type App struct {
	config            *Config
	logger            *services.Logger
	db                *sql.DB
	userService       interfaces.UserService
	userSecretService interfaces.UserSecretService
}

func NewApp(
	config *Config,
	logger *services.Logger,
) *App {
	return &App{
		config: config,
		logger: logger,
	}
}

func (app *App) Configure(ctx context.Context) error {
	db, err := postgres.NewPostgresql(
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

	userRepository := postgres.NewPostgresUserRepository(db)

	app.userService = services.NewUserService(userRepository)

	userSecretRepository := postgres.NewPostgresUserSecretRepository(db, userRepository)

	app.userSecretService = services.NewUserSecretService(userSecretRepository)

	return nil
}

func (app *App) Run(ctx context.Context) error {

	return nil
}
