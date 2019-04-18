package userservice

import (
	"context"
	"time"
	"userService/pkg/common"
	usermodel "userService/pkg/model/user"
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
	db := common.DB.New()
	if in.GetUsername() == "" || in.GetPassword() == "" {
		return nil, ErrUserNamePasswordEmpty
	}

	// 查询用户
	user, err := usermodel.FindUserByUserName(db, in.GetUsername())
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
		Id:       user.UserID,
		Username: user.UserName,
		UserType: user.UserType,
		Token:    tk,
	}, err
}

func (u *userService) GetPermissions(ctx context.Context, in *pb.GetPermissionsRequest) (*pb.GetPermissionsReply, error) {
	db := common.DB
	user := ctx.Value("userInfo").(*UserInfo)

	permissions := common.Enforcer.GetImplicitPermissionsForUser(user.UserName)

	itemNames := make([]string, 0)
	for _, permission := range permissions {
		for _, ps := range permission {
			itemNames = append(itemNames, ps)
		}
	}

	menus, err := usermodel.GetAuthMenu(db, itemNames)
	if err != nil {
		return nil, err
	}

	idMap := make(map[int64][]*usermodel.Menu)
	for _, menu := range menus {
		if menu.Parent == nil {
			// 第一级菜单
			ms, ok := idMap[-1]
			if !ok {
				ms = make([]*usermodel.Menu, 0)
			}
			ms = append(ms, menu)
			idMap[-1] = ms
		} else {
			// 第二级菜单
			ms, ok := idMap[*menu.Parent]
			if !ok {
				ms = make([]*usermodel.Menu, 0)
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

func (u *userService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	db := common.DB
	if in.Username == "" || in.Password == "" || in.UserType == "" || in.Email == "" || in.LeaguerNo == "" {
		return nil, ErrInvalidParams
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(in.Password), 4)
	if err != nil {
		return nil, err
	}

	user := &usermodel.User{
		UserName:     in.Username,
		UserType:     in.UserType,
		Email:        &in.Email,
		LeaguerNO:    in.LeaguerNo,
		PasswordHash: string(bs),
		UserStatus:   1,
	}

	newUser, err := usermodel.SaveUser(db, user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		Id:         newUser.UserID,
		LeaguerNo:  newUser.LeaguerNO,
		Username:   newUser.UserName,
		Email:      *newUser.Email,
		UserType:   newUser.UserType,
		UserStatus: newUser.UserStatus,
		CreatedAt:  newUser.CreatedAt.UnixNano() / int64(time.Millisecond),
	}, nil
}

func (u *userService) AddPermission(ctx context.Context, in *pb.AddPermissionRequest) (*pb.AddPermissionReply, error) {
	if in.Role == "" || len(in.Permission) == 0 {
		return nil, ErrInvalidParams
	}

	if common.Enforcer.AddPolicy(in.Role, in.Permission) {
		return &pb.AddPermissionReply{}, nil
	} else {
		return nil, ErrPolicyExists
	}
}

func (u *userService) AddRole(ctx context.Context, in *pb.AddRoleRequest) (*pb.AddRoleReply, error) {
	if in.Role == "" || in.On == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	role, err := usermodel.FindRole(db, in.Role)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, ErrRoleNotFound
	}
	on, err := usermodel.FindRole(db, in.On)
	if err != nil {
		return nil, err
	}
	if on == nil {
		return nil, ErrRoleNotFound
	}

	if common.Enforcer.AddRoleForUser(in.Role, in.On) {
		return &pb.AddRoleReply{}, nil
	} else {
		return nil, ErrPolicyExists
	}
}

func (u *userService) CreateRole(ctx context.Context, in *pb.CreateRoleRequest) (*pb.CreateRoleReply, error) {
	if in.Role == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	r, err := usermodel.FindRole(db, in.Role)
	if err != nil {
		return nil, err
	}

	if r != nil {
		return nil, ErrRoleExists
	}

	err = usermodel.SaveRole(db, &usermodel.Role{Role: in.Role})
	return &pb.CreateRoleReply{}, err
}

func (u *userService) AddRoleForUser(ctx context.Context, in *pb.AddRoleForUserRequest) (*pb.AddRoleForUserReply, error) {
	if in.Username == "" || in.Role == "" {
		return nil, ErrInvalidParams
	}

	db := common.DB

	role, err := usermodel.FindRole(db, in.Role)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, ErrRoleNotFound
	}

	user, err := usermodel.FindUserByUserName(db, in.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	if common.Enforcer.AddRoleForUser(in.Username, in.Role) {
		return &pb.AddRoleForUserReply{}, nil
	} else {
		return nil, ErrPolicyExists
	}
}
