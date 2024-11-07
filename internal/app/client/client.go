package client

import (
	"context"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/app/services"
	"github.com/kyrare/ya-diplom-2/internal/interfaces/tui/bubbletea"
)

type App struct {
	config        *services.Config
	logger        *services.Logger
	clientService interfaces.ClientService
}

func NewApp(
	config *services.Config,
	logger *services.Logger,
) *App {
	return &App{
		config: config,
		logger: logger,
	}
}

func (app *App) Configure(ctx context.Context) error {

	app.clientService = services.NewClientService(app.config.GRPC.Address, app.logger)

	//db, err := postgres.NewPostgresql(
	//	app.config.DB.Name,
	//	app.config.DB.Host,
	//	app.config.DB.Port,
	//	app.config.DB.User,
	//	app.config.DB.Password,
	//	app.logger,
	//)
	//
	//if err != nil {
	//	return err
	//}
	//
	//userRepository := postgres.NewPostgresUserRepository(db)
	//
	//authService := services.NewAuthService(userRepository, app.config.Jwt.Secret, app.config.Jwt.Duration, app.logger)
	//
	//userService := services.NewUserService(userRepository, authService, app.logger)
	//
	//userService.Create(&command.CreateUserCommand{
	//	Login:    "test",
	//	Password: "123123",
	//})

	return nil
}

func (app *App) Run(ctx context.Context) error {
	p := tea.NewProgram(bubbletea.New(app.clientService))
	if _, err := p.Run(); err != nil {
		app.logger.Error(err)
		os.Exit(1)
	}

	return nil
}
