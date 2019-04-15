package userservice

import (
	"context"
	"time"

	"userService/pkg/common"
	"userService/pkg/model"
	"userService/pkg/pb"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

type UserClaims struct {
	ID int64
	jwt.StandardClaims
}

func (u *userService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	if in.GetUsername() == "" || in.GetPassword() == "" {
		return nil, ErrUserNamePasswordEmpty
	}

	// 查询用户
	user, err := model.FindUserByUserName(common.DB, in.GetUsername())
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrWrongUserNameOrPassword
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.GetPassword()))
	if err != nil {
		return nil, ErrWrongUserNameOrPassword
	}

	// 查询用户权限
	permission, err := model.GetPermissionsByUserID(common.DB, user.UserID)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	// 给用户添加权限
	// todo
	_ = permission

	// 生成token
	expiredAt := time.Now().Add(time.Hour * 72)
	tk, err := u.genToken(user.UserID, expiredAt)
	if err != nil {
		logrus.Errorln(err)
	}

	return &pb.LoginReply{
		Token: tk,
	}, err
}

func New() pb.UserServer {
	return &userService{}
}

func (u *userService) genToken(id int64, exTime time.Time) (string, error) {
	claims := UserClaims{
		ID: id,
	}
	claims.ExpiresAt = exTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SigningString()
}
