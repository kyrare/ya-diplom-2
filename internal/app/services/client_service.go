package services

import (
	"context"
	"errors"
	"fmt"
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
	if !s.isAuth() {
		return nil, errors.New("not authorized")
	}

	client := proto.NewUserSecretsClient(s.conn)

	resp, err := client.GetUserSecrets(ctx, &proto.GetUserSecretsRequest{
		Token: s.jwtToken,
	})

	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	secrets := make([]*entities.UserSecret, 0, len(resp.Secrets))

	for _, secret := range resp.Secrets {
		t := entities.UserSecretType(secret.Type)
		d, err := entities.MakeUserSecretData(t, secret.Data)
		if err != nil {
			return nil, err
		}

		id, _ := uuid.Parse(secret.Id)

		secrets = append(secrets, &entities.UserSecret{
			Id:   id,
			Type: entities.UserSecretType(secret.Type),
			Name: secret.Name,
			Data: &d,
		})
	}

	fmt.Println("secrets", secrets)
	fmt.Println("secrets len", len(secrets))

	return secrets, nil
}

func (s *ClientService) DeleteUserSecret(ctx context.Context, id uuid.UUID) error {
	if !s.isAuth() {
		return errors.New("not authorized")
	}

	client := proto.NewUserSecretsClient(s.conn)

	resp, err := client.DeleteUserSecret(ctx, &proto.DeleteUserSecretRequest{
		Token: s.jwtToken,
		Id:    id.String(),
	})

	if err != nil {
		return err
	}
	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return nil
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
