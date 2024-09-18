package grpc

import (
	"context"

	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/interfaces/grpc/proto"
	"google.golang.org/grpc"
)

type UserServer struct {
	proto.UnimplementedUserServer

	userService interfaces.UserService
}

func NewUserServer(s *grpc.Server, userService interfaces.UserService) *UserServer {
	server := &UserServer{
		userService: userService,
	}

	proto.RegisterUserServer(s, server)

	return server
}

func (s UserServer) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {
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

func (s UserServer) Auth(ctx context.Context, request *proto.AuthRequest) (*proto.AuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServer) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}
