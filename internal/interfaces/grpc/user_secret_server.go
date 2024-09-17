package grpc

import (
	"context"

	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/interfaces/grpc/proto"
	"google.golang.org/grpc"
)

type UserSecretServer struct {
	proto.UnimplementedUserSecretsServer

	secretService interfaces.UserSecretService
}

func NewUserSecretServer(s *grpc.Server, secretService interfaces.UserSecretService) *UserSecretServer {
	server := &UserSecretServer{
		secretService: secretService,
	}

	proto.RegisterUserSecretsServer(s, server)

	return server
}

func (s UserSecretServer) CreateUserSecret(ctx context.Context, request *proto.CreateUserSecretRequest) (*proto.CreateUserSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserSecretServer) DeleteUserSecret(ctx context.Context, request *proto.DeleteUserSecretRequest) (*proto.DeleteUserSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserSecretServer) GetUserSecrets(ctx context.Context, request *proto.GetUserSecretsRequest) (*proto.GetUserSecretsResponse, error) {
	//TODO implement me
	panic("implement me")
}
