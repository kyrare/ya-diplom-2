package grpc

import (
	"context"

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
	user, err := s.authService.GetUserByToken(ctx, request.Token)
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

	d, err := entities.MakeUserSecretData(t, request.Secret.Data)
	if err != nil {
		s.logger.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid secret data")
	}

	s.logger.Infof("Start create new secret, type: %s, user: %s, data: %s", t, user.Id, string(request.Secret.Data))

	_, err = s.secretService.Create(ctx, &command.CreateUserSecretCommand{
		User:       user,
		SecretType: t,
		SecretName: request.Secret.Name,
		SecretData: &d,
	})

	resp := &proto.CreateUserSecretResponse{}

	if err != nil {
		s.logger.Error(err)
		resp.Error = err.Error()
	}

	return resp, nil
}

func (s UserSecretServer) DeleteUserSecret(ctx context.Context, request *proto.DeleteUserSecretRequest) (*proto.DeleteUserSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserSecretServer) GetUserSecrets(ctx context.Context, request *proto.GetUserSecretsRequest) (*proto.GetUserSecretsResponse, error) {
	user, err := s.authService.GetUserByToken(ctx, request.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	if user == nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	resp := &proto.GetUserSecretsResponse{}

	s.logger.Infof("Получение секретов пользователя %s", user.Id.String())

	secrets, err := s.secretService.GetAllForUser(ctx, user.Id)
	if err != nil {
		s.logger.Error(err)
		resp.Error = err.Error()
	}

	result := make([]*proto.UserSecret, 0, len(secrets))
	for _, secret := range secrets {
		d, err := (*secret.Data).GetData()
		if err != nil {
			return nil, err
		}

		result = append(result, &proto.UserSecret{
			Name: secret.Name,
			Type: string(secret.Type),
			Data: d,
		})
	}

	resp.Secrets = result

	return resp, nil
}
