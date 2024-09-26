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
	authService interfaces.AuthService
}

func NewAuthServer(s *grpc.Server, userService interfaces.UserService, authService interfaces.AuthService) *AuthServer {
	server := &AuthServer{
		userService: userService,
		authService: authService,
	}

	proto.RegisterAuthServer(s, server)

	return server
}

func (s AuthServer) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	resp := new(proto.RegisterResponse)

	_, err := s.userService.Create(&command.CreateUserCommand{
		Login:    request.Login,
		Password: request.Password,
	})
	if err != nil {
		resp.Error = err.Error()
		return resp, nil
	}

	login, err := s.authService.Login(&command.LoginCommand{
		Login:    request.Login,
		Password: request.Password,
	})
	if err != nil {
		resp.Error = err.Error()
		return resp, nil
	}

	resp.JwtToken = login.JwtToken

	return resp, nil
}

func (s AuthServer) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	resp := new(proto.LoginResponse)

	login, err := s.authService.Login(&command.LoginCommand{
		Login:    request.Login,
		Password: request.Password,
	})
	if err != nil {
		resp.Error = err.Error()
		return resp, nil
	}

	resp.JwtToken = login.JwtToken

	return resp, nil
}
