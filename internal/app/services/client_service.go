package services

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
	"github.com/kyrare/ya-diplom-2/internal/interfaces/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientService struct {
	conn     *grpc.ClientConn
	jwtToken string
	logger   *Logger
}

func NewClientService(serverAddr string, logger *Logger) *ClientService {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return &ClientService{
		conn:   conn,
		logger: logger,
	}
}

func (s *ClientService) Login(ctx context.Context, command *command.LoginCommand) error {

	// todo правильно ли при каждом вызове создавать клиент, или лучше его создать в конструкторе?
	client := proto.NewAuthClient(s.conn)

	resp, err := client.Login(ctx, &proto.LoginRequest{
		Login:    command.Login,
		Password: command.Password,
	})

	if err != nil {
		return err
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	s.jwtToken = resp.JwtToken

	return nil
}

func (s *ClientService) Register(ctx context.Context, command *command.RegisterCommand) error {
	client := proto.NewAuthClient(s.conn)

	resp, err := client.Register(ctx, &proto.RegisterRequest{
		Login:    command.Login,
		Password: command.Password,
	})

	if err != nil {
		return err
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	s.jwtToken = resp.JwtToken

	return nil
}

func (s *ClientService) GetUserSecrets(ctx context.Context) ([]*entities.UserSecret, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClientService) GetUserSecretById(ctx context.Context, id uuid.UUID) (*entities.UserSecret, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ClientService) CreateUserSecret(ctx context.Context, command *command.ClientCreateUserSecretCommand) error {
	if !s.isAuth() {
		return errors.New("not authorized")
	}

	data, err := command.SecretData.GetData()
	if err != nil {
		return err
	}

	client := proto.NewUserSecretsClient(s.conn)

	resp, err := client.CreateUserSecret(ctx, &proto.CreateUserSecretRequest{
		Token: s.jwtToken,
		Secret: &proto.UserSecret{
			Name: command.SecretName,
			Type: string(command.SecretType),
			Data: data,
		},
	})

	if err != nil {
		return err
	}
	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return nil
}

func (s *ClientService) isAuth() bool {
	return s.jwtToken != ""
}
