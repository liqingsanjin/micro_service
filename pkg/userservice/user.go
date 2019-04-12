package userservice

import (
	"context"
	"userService/pb"
)

type userService struct{}

func (u *userService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	return nil, nil
}

func New() pb.UserServer {
	return &userService{}
}