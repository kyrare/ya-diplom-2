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

// type UserServiceServer interface {
//	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
//	AuthUser(context.Context, *AuthUserRequest) (*AuthUserResponse, error)
//	mustEmbedUnimplementedUserServiceServer()
//}

func (s UserServer) RegisterUser(ctx context.Context, request *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	var response proto.RegisterUserResponse

	_, err := s.userService.Create(&command.CreateUserCommand{
		Login:    request.Login,
		Password: request.Password,
	})
	if err != nil {
		response.Error = err.Error()
	}

	return &response, err
}

func (s UserServer) AuthUser(ctx context.Context, request *proto.AuthUserRequest) (*proto.AuthUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
