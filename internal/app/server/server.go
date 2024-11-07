package server

import (
	"context"
	"database/sql"
	"net"

	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/app/services"
	"github.com/kyrare/ya-diplom-2/internal/infrastructure/db/postgres"
	"github.com/kyrare/ya-diplom-2/internal/infrastructure/s3/minio"
	igrpc "github.com/kyrare/ya-diplom-2/internal/interfaces/grpc"
	"google.golang.org/grpc"
)

type App struct {
	config            *services.Config
	logger            *services.Logger
	db                *sql.DB
	authService       interfaces.AuthService
	userService       interfaces.UserService
	userSecretService interfaces.UserSecretService
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

	app.authService = services.NewAuthService(userRepository, app.config.Jwt.Secret, app.config.Jwt.Duration, app.logger)

	app.userService = services.NewUserService(userRepository, app.authService, app.logger)

	userSecretRepository := postgres.NewPostgresUserSecretRepository(db, userRepository)

	minioClient, err := minio.NewClient(
		app.config.Minio.Endpoint,
		app.config.Minio.AccessKey,
		app.config.Minio.SecretKey,
		app.config.Minio.UseSSL,
		app.logger,
	)
	if err != nil {
		return err
	}
	userSecretFileRepository := minio.NewMinioUserSecretFileRepository("user-secrets", minioClient)

	app.userSecretService = services.NewUserSecretService(userSecretRepository, userSecretFileRepository)

	return nil
}

func (app *App) Run(ctx context.Context) error {
	listen, err := net.Listen("tcp", app.config.GRPC.Address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	igrpc.NewAuthServer(server, app.userService, app.authService)
	igrpc.NewUserSecretServer(server, app.userSecretService, app.authService, app.logger)

	app.logger.Infof("start gRPC server on %s", listen.Addr())
	// получаем запрос gRPC
	if err := server.Serve(listen); err != nil {
		return err
	}

	return nil
}
