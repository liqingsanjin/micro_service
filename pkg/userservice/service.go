package userservice

import (
	"context"
	"time"
	"userService/pkg/common"
	usermodel "userService/pkg/model/user"
	"userService/pkg/pb"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.Internal, err.Error())
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
		err = status.Error(codes.Internal, err.Error())
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
		return nil, status.Error(codes.Internal, err.Error())
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
		return nil, status.Error(codes.Internal, err.Error())
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
		return nil, status.Error(codes.Internal, err.Error())
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

func (u *userService) AddPermissionForRole(ctx context.Context, in *pb.AddPermissionForRoleRequest) (*pb.AddPermissionForRoleReply, error) {
	if in.Role == "" || len(in.Permission) == 0 {
		return nil, ErrInvalidParams
	}
	db := common.DB

	role, err := usermodel.FindRole(db, in.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if role == nil {
		return nil, ErrRoleNotFound
	}

	permission, err := usermodel.FindPermissionByName(db, in.Permission)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	if common.Enforcer.AddRoleForUser(in.Role, in.Permission) {
		return &pb.AddPermissionForRoleReply{}, nil
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
		return nil, status.Error(codes.Internal, err.Error())
	}

	if r != nil {
		return nil, ErrRoleExists
	}

	err = usermodel.SaveRole(db, &usermodel.Role{Role: in.Role})
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateRoleReply{}, err
}

func (u *userService) AddRoleForUser(ctx context.Context, in *pb.AddRoleForUserRequest) (*pb.AddRoleForUserReply, error) {
	if in.Username == "" || in.Role == "" {
		return nil, ErrInvalidParams
	}

	db := common.DB

	role, err := usermodel.FindRole(db, in.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if role == nil {
		return nil, ErrRoleNotFound
	}

	user, err := usermodel.FindUserByUserName(db, in.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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

func (u *userService) AddRoutes(ctx context.Context, in *pb.AddRoutesRequest) (*pb.AddRoutesReply, error) {
	if len(in.Routes) == 0 {
		return nil, ErrInvalidParams
	}
	db := common.DB

	rs, err := usermodel.FindRoutesByNames(db, in.Routes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(rs) != 0 {
		return nil, ErrRouteExists
	}

	err = usermodel.SaveRoutes(db, in.Routes)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.AddRoutesReply{}, err
}

func (u *userService) ListRoutes(ctx context.Context, in *pb.ListRoutesRequest) (*pb.ListRoutesReply, error) {
	db := common.DB

	routes, err := usermodel.ListRoutes(db)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	names := make([]string, len(routes))
	for i := range routes {
		names[i] = routes[i].Name
	}

	return &pb.ListRoutesReply{
		Routes: names,
	}, nil
}

func (u *userService) CreatePermission(ctx context.Context, in *pb.CreatePermissionRequest) (*pb.CreatePermissionReply, error) {
	if in.Name == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	p, err := usermodel.FindPermissionByName(db, in.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if p != nil {
		return nil, ErrPermissionExists
	}
	err = usermodel.SavePermission(db, in.Name)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.CreatePermissionReply{}, err
}

func (u *userService) UpdatePermission(ctx context.Context, in *pb.UpdatePermissionRequest) (*pb.UpdatePermissionReply, error) {
	db := common.DB

	p, err := usermodel.FindPermissionByID(db, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if p == nil {
		return nil, ErrPermissionNotFound
	}
	err = usermodel.UpdatePermission(db, in.Id, &usermodel.Permission{Name: in.Name})
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.UpdatePermissionReply{}, err
}

func (u *userService) AddRouteForPermission(ctx context.Context, in *pb.AddRouteForPermissionRequest) (*pb.AddRouteForPermissionReply, error) {
	db := common.DB

	if in.Permission == "" || in.Route == "" {
		return nil, ErrInvalidParams
	}

	permission, err := usermodel.FindPermissionByName(db, in.Permission)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	route, err := usermodel.FindRouteByName(db, in.Route)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if route == nil {
		return nil, ErrRouteNotFound
	}

	if common.Enforcer.AddPermissionForUser(permission.Name, in.Route) {
		return &pb.AddRouteForPermissionReply{}, nil
	}
	return nil, ErrPolicyExists

}

func (u *userService) RemoveRouteForPermission(ctx context.Context, in *pb.RemoveRouteForPermissionRequest) (*pb.RemoveRouteForPermissionReply, error) {
	if in.Permission == "" || in.Route == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	permission, err := usermodel.FindPermissionByName(db, in.Permission)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	route, err := usermodel.FindRouteByName(db, in.Route)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if route == nil {
		return nil, ErrRouteNotFound
	}

	if common.Enforcer.DeletePermissionForUser(permission.Name, in.Route) {
		return &pb.RemoveRouteForPermissionReply{}, nil
	}
	return nil, ErrPolicyNotFound
}

func (u *userService) RemovePermission(ctx context.Context, in *pb.RemovePermissionRequest) (*pb.RemovePermissionReply, error) {
	if in.Permission == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	permission, err := usermodel.FindPermissionByName(db, in.Permission)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	db = db.Begin()
	defer db.Rollback()
	err = usermodel.DeletePermission(db, permission)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	common.Enforcer.DeleteRole(in.Permission)
	err = db.Commit().Error
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.RemovePermissionReply{}, err
}

func (u *userService) ListPermissions(ctx context.Context, in *pb.ListPermissionsRequest) (*pb.ListPermissionsReply, error) {
	db := common.DB

	ps, err := usermodel.ListPermissions(db)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	names := make([]string, len(ps))
	for i := range ps {
		names[i] = ps[i].Name
	}

	return &pb.ListPermissionsReply{
		Permissions: names,
	}, nil
}

func (u *userService) AddPermissionForPermission(ctx context.Context, in *pb.AddPermissionForPermissionRequest) (*pb.AddPermissionForPermissionReply, error) {
	if in.From == "" || in.To == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	from, err := usermodel.FindPermissionByName(db, in.From)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if from == nil {
		return nil, ErrPermissionNotFound
	}
	to, err := usermodel.FindPermissionByName(db, in.To)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if to == nil {
		return nil, ErrPermissionNotFound
	}

	if common.Enforcer.AddRoleForUser(to.Name, from.Name) {
		return &pb.AddPermissionForPermissionReply{}, nil
	}
	return nil, ErrPolicyExists
}

func (u *userService) RemovePermissionForPermission(ctx context.Context, in *pb.RemovePermissionForPermissionRequest) (*pb.RemovePermissionForPermissionReply, error) {
	if in.From == "" || in.Child == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	from, err := usermodel.FindPermissionByName(db, in.From)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if from == nil {
		return nil, ErrPermissionNotFound
	}
	child, err := usermodel.FindPermissionByName(db, in.Child)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if child == nil {
		return nil, ErrPermissionNotFound
	}

	if common.Enforcer.DeleteRoleForUser(from.Name, child.Name) {
		return &pb.RemovePermissionForPermissionReply{}, nil
	}
	return nil, ErrPolicyExists
}

func (u *userService) ListRole(ctx context.Context, in *pb.ListRoleRequest) (*pb.ListRoleReply, error) {
	db := common.DB

	roles, err := usermodel.ListRole(db)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	names := make([]string, len(roles))
	for i := range roles {
		names[i] = roles[i].Role
	}
	return &pb.ListRoleReply{Roles: names}, nil
}

func (u *userService) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest) (*pb.UpdateRoleReply, error) {
	db := common.DB

	p, err := usermodel.FindRoleByID(db, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if p == nil {
		return nil, ErrRouteNotFound
	}

	err = usermodel.UpdateRole(db, in.Id, &usermodel.Role{Role: in.Name})
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.UpdateRoleReply{}, err
}

func (u *userService) RemovePermissionForRole(ctx context.Context, in *pb.RemovePermissionForRoleRequest) (*pb.RemovePermissionForRoleReply, error) {
	if in.Role == "" || len(in.Permission) == 0 {
		return nil, ErrInvalidParams
	}
	db := common.DB

	role, err := usermodel.FindRole(db, in.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if role == nil {
		return nil, ErrRoleNotFound
	}

	permission, err := usermodel.FindPermissionByName(db, in.Permission)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	if common.Enforcer.DeleteRoleForUser(in.Role, in.Permission) {
		return &pb.RemovePermissionForRoleReply{}, nil
	}
	return nil, ErrPolicyNotFound
}

func (u *userService) AddRoleForRole(ctx context.Context, in *pb.AddRoleForRoleRequest) (*pb.AddRoleForRoleReply, error) {
	if in.From == "" || in.To == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	from, err := usermodel.FindRole(db, in.From)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if from == nil {
		return nil, ErrRoleNotFound
	}
	to, err := usermodel.FindRole(db, in.To)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if to == nil {
		return nil, ErrRoleNotFound
	}

	if common.Enforcer.AddRoleForUser(to.Role, from.Role) {
		return &pb.AddRoleForRoleReply{}, nil
	}
	return nil, ErrPolicyExists
}

func (u *userService) RemoveRoleForRole(ctx context.Context, in *pb.RemoveRoleForRoleRequest) (*pb.RemoveRoleForRoleReply, error) {
	if in.From == "" || in.Child == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	from, err := usermodel.FindRole(db, in.From)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if from == nil {
		return nil, ErrRoleNotFound
	}
	child, err := usermodel.FindRole(db, in.Child)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if child == nil {
		return nil, ErrRoleNotFound
	}

	if common.Enforcer.DeleteRoleForUser(from.Role, child.Role) {
		return &pb.RemoveRoleForRoleReply{}, nil
	}
	return nil, ErrPolicyExists
}

func (u *userService) RemoveRole(ctx context.Context, in *pb.RemoveRoleRequest) (*pb.RemoveRoleReply, error) {
	if in.Role == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	role, err := usermodel.FindRole(db, in.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if role == nil {
		return nil, ErrRoleNotFound
	}

	db = db.Begin()
	defer db.Rollback()
	err = usermodel.DeleteRole(db, role)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	common.Enforcer.DeleteRole(in.Role)
	err = db.Commit().Error
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}
	return &pb.RemoveRoleReply{}, err
}

func (u *userService) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersReply, error) {
	db := common.DB

	us, err := usermodel.ListUsers(db)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	users := make([]*pb.UserField, 0)

	for _, u := range us {
		users = append(users, &pb.UserField{
			Id:       u.UserID,
			Username: u.UserName,
		})
	}

	return &pb.ListUsersReply{
		Users: users,
	}, nil
}

func (u *userService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	db := common.DB

	user, err := usermodel.FindUserByID(db, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	err = usermodel.UpdateUser(db, in.Id, &usermodel.User{
		UserName: in.Username,
		Email:    &in.Email,
		UserType: in.UserType,
	})
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateUserReply{}, err
}

func (u *userService) AddPermissionForUser(ctx context.Context, in *pb.AddPermissionForUserRequest) (*pb.AddPermissionForUserReply, error) {
	if in.Username == "" || in.Permission == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	user, err := usermodel.FindUserByUserName(db, in.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	permission, err := usermodel.FindPermissionByName(db, in.Permission)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	if common.Enforcer.AddRoleForUser(in.Username, in.Permission) {
		return &pb.AddPermissionForUserReply{}, nil
	}
	return nil, ErrPolicyExists
}

func (u *userService) RemovePermissionForUser(ctx context.Context, in *pb.RemovePermissionForUserRequest) (*pb.RemovePermissionForUserReply, error) {
	if in.Username == "" || in.Permission == "" {
		return nil, ErrInvalidParams
	}
	db := common.DB

	user, err := usermodel.FindUserByUserName(db, in.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	permission, err := usermodel.FindPermissionByName(db, in.Permission)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if permission == nil {
		return nil, ErrPermissionNotFound
	}

	if common.Enforcer.DeleteRoleForUser(in.Username, in.Permission) {
		return &pb.RemovePermissionForUserReply{}, nil
	}
	return nil, ErrPolicyNotFound
}

func (u *userService) RemoveRoleForUser(ctx context.Context, in *pb.RemoveRoleForUserRequest) (*pb.RemoveRoleForUserReply, error) {
	if in.Username == "" || in.Role == "" {
		return nil, ErrInvalidParams
	}

	db := common.DB

	role, err := usermodel.FindRole(db, in.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if role == nil {
		return nil, ErrRoleNotFound
	}

	user, err := usermodel.FindUserByUserName(db, in.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	if common.Enforcer.DeleteRoleForUser(in.Username, in.Role) {
		return &pb.RemoveRoleForUserReply{}, nil
	}
	return nil, ErrPolicyNotFound

}
