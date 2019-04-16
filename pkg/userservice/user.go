package userservice

import (
	"context"
	"time"

	"userService/pkg/common"
	"userService/pkg/model"
	"userService/pkg/pb"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

type UserInfo struct {
	ID       int64
	UserName string
}

func (u *userService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	db := common.DB
	if in.GetUsername() == "" || in.GetPassword() == "" {
		return nil, ErrUserNamePasswordEmpty
	}

	// 查询用户
	user, err := model.FindUserByUserName(db, in.GetUsername())
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

	// 生成token
	expiredAt := time.Now().Add(time.Hour * 72)
	tk, err := genToken(&UserInfo{
		ID:       user.UserID,
		UserName: user.UserName,
	}, expiredAt)
	if err != nil {
		logrus.Errorln(err)
	}

	return &pb.LoginReply{
		Token: tk,
	}, err
}

func (u *userService) GetPermissions(ctx context.Context, in *pb.GetPermissionsRequest) (*pb.GetPermissionsReply, error) {
	user := ctx.Value("userInfo").(*UserInfo)

	permissions := common.Enforcer.GetImplicitPermissionsForUser(user.UserName)

	itemNames := make([]string, 0)
	for _, permission := range permissions {
		for _, ps := range permission {
			itemNames = append(itemNames, ps)
		}
	}

	menus, err := model.GetAuthMenu(common.DB, itemNames)
	if err != nil {
		return nil, err
	}

	idMap := make(map[int64][]*model.Menu)
	for _, menu := range menus {
		if menu.Parent == nil {
			// 第一级菜单
			ms, ok := idMap[-1]
			if !ok {
				ms = make([]*model.Menu, 0)
			}
			ms = append(ms, menu)
			idMap[-1] = ms
		} else {
			// 第二级菜单
			ms, ok := idMap[*menu.Parent]
			if !ok {
				ms = make([]*model.Menu, 0)
			}
			ms = append(ms, menu)
			idMap[*menu.Parent] = ms
		}
	}

	replyMenus := make([]*pb.Menu, 0)
	rootMenus := idMap[-1]
	for _, rootMenu := range rootMenus {
		children := make([]*pb.Menu, 0, len(idMap[rootMenu.ID]))
		for _, child := range idMap[rootMenu.ID] {
			children = append(children, &pb.Menu{
				Menu:  child.Name,
				Route: child.MenuRoute,
				Data:  child.MenuData,
			})
		}
		replyMenus = append(replyMenus, &pb.Menu{
			Menu:     rootMenu.Name,
			Route:    rootMenu.MenuRoute,
			Data:     rootMenu.MenuData,
			Children: children,
		})
	}

	return &pb.GetPermissionsReply{
		Menus: replyMenus,
	}, nil
}

func (u *userService) CheckPermission(ctx context.Context, in *pb.CheckPermissionRequest) (*pb.CheckPermissionReply, error) {
	user := ctx.Value("userInfo").(*UserInfo)

	ok := common.Enforcer.Enforce(user.UserName, in.GetRoute())
	return &pb.CheckPermissionReply{
		Result: ok,
	}, nil
}
