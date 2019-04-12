package userservice

import (
	"context"

	"userService/pb"
	"userService/pkg/common"
	"userService/pkg/model"

	"github.com/sirupsen/logrus"
)

type userService struct{}

func (u *userService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	logrus.Infoln("开始查询用户")
	if in.GetUsername() == "" || in.GetPassword() == "" {
		return nil, ErrUserNamePasswordEmpty
	}
	user, err := model.FindUserByUserName(common.DB, in.GetUsername())
	if err != nil {
		return nil, err
	}
	logrus.Infoln(user)
	return &pb.LoginReply{}, nil
}

func New() pb.UserServer {
	return &userService{}
}
