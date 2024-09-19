package grpc

import (
	"context"

	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/interfaces/grpc/proto"
	"google.golang.org/grpc"
)

type AuthServer struct {
	proto.UnimplementedAuthServer

	userService interfaces.UserService
}

func NewAuthServer(s *grpc.Server, userService interfaces.UserService) *AuthServer {
	server := &AuthServer{
		userService: userService,
	}

	proto.RegisterAuthServer(s, server)

	return server
}

func (s AuthServer) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	var response proto.RegisterResponse

	_, err := s.userService.Create(&command.CreateUserCommand{
		Login:    request.Login,
		Password: request.Password,
	})
	if err != nil {
		response.Error = err.Error()
	}

	return &response, nil
}

func (s AuthServer) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s AuthServer) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}
