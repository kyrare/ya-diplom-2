package grpc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/app/services"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
	"github.com/kyrare/ya-diplom-2/internal/interfaces/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserSecretServer struct {
	proto.UnimplementedUserSecretsServer

	logger        *services.Logger
	secretService interfaces.UserSecretService
	authService   interfaces.AuthService
}

func NewUserSecretServer(s *grpc.Server, secretService interfaces.UserSecretService, authService interfaces.AuthService, logger *services.Logger) *UserSecretServer {
	server := &UserSecretServer{
		secretService: secretService,
		authService:   authService,
		logger:        logger,
	}

	proto.RegisterUserSecretsServer(s, server)

	return server
}

func (s UserSecretServer) CreateUserSecret(ctx context.Context, request *proto.CreateUserSecretRequest) (*proto.CreateUserSecretResponse, error) {

	user, err := s.authService.GetUserByToken(request.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	if user == nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	t, err := entities.GetSecretTypeByString(request.Secret.Type)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid secret type")
	}

	d, err := s.makeData(t, request.Secret.Data)
	if err != nil {
		s.logger.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid secret data")
	}

	_, err = s.secretService.Create(ctx, &command.CreateUserSecretCommand{
		User:       user,
		SecretType: t,
		SecretName: request.Secret.Name,
		SecretData: &d,
	})

	return &proto.CreateUserSecretResponse{}, nil

}

func (s UserSecretServer) DeleteUserSecret(ctx context.Context, request *proto.DeleteUserSecretRequest) (*proto.DeleteUserSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserSecretServer) GetUserSecrets(ctx context.Context, request *proto.GetUserSecretsRequest) (*proto.GetUserSecretsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserSecretServer) makeData(secretType entities.UserSecretType, data []byte) (entities.UserSecretData, error) {

	switch secretType {
	case entities.UserSecretPasswordType:
		secretData := new(entities.UserSecretDataPassword)
		err := json.Unmarshal(data, secretData)
		if err != nil {
			return nil, err
		}
		return secretData, nil
	case entities.UserSecretBankCardType:
		secretData := new(entities.UserSecretDataBankCard)
		err := json.Unmarshal(data, secretData)
		if err != nil {
			return nil, err
		}
		return secretData, nil
	case entities.UserSecretTextType:
		secretData := entities.NewUserSecretText(string(data))

		return secretData, nil
	case entities.UserSecretFileType:
		panic("implement me")

		return nil, nil
	}

	return nil, errors.New("invalid secret type")
}
