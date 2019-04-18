package userservice

import (
	"context"
	"userService/pkg/pb"

	stdjwt "github.com/dgrijalva/jwt-go"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type userServer struct {
	loginHandler           grpctransport.Handler
	getPermissionsHandler  grpctransport.Handler
	checkPermissionHandler grpctransport.Handler
	registerHandler        grpctransport.Handler
	addPermissionHandler   grpctransport.Handler
	addRoleHandler         grpctransport.Handler
	createRoleHandler      grpctransport.Handler
	addRoleForUserHandler  grpctransport.Handler
}

func New() pb.UserServer {
	svr := &userServer{}
	userService := &userService{}

	{
		loginEndpoint := makeLoginEndpoint(userService)
		loginEndpoint = logginMiddleware(loginEndpoint)
		svr.loginHandler = grpctransport.NewServer(
			loginEndpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		getPermissionEndpoint := makeGetPermissionsEndpoint(userService)
		getPermissionEndpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(getPermissionEndpoint)
		getPermissionEndpoint = logginMiddleware(getPermissionEndpoint)
		svr.getPermissionsHandler = grpctransport.NewServer(
			getPermissionEndpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		checkPermissionEndpoint := makeCheckPermissionEndpoint(userService)
		checkPermissionEndpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(checkPermissionEndpoint)
		checkPermissionEndpoint = logginMiddleware(checkPermissionEndpoint)
		svr.checkPermissionHandler = grpctransport.NewServer(
			checkPermissionEndpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		registerEndpoint := makeRegisterEndpoint(userService)
		registerEndpoint = logginMiddleware(registerEndpoint)
		svr.registerHandler = grpctransport.NewServer(
			registerEndpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		addPermissionEndpoint := makeAddPermissionEndpoint(userService)
		addPermissionEndpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(addPermissionEndpoint)
		addPermissionEndpoint = logginMiddleware(addPermissionEndpoint)
		svr.addPermissionHandler = grpctransport.NewServer(
			addPermissionEndpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		addRoleEndpoint := makeAddRoleEndpoint(userService)
		addRoleEndpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(addRoleEndpoint)
		addRoleEndpoint = logginMiddleware(addRoleEndpoint)
		svr.addRoleHandler = grpctransport.NewServer(
			addRoleEndpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		createRoleEndpoint := makeCreateRoleEndpoint(userService)
		createRoleEndpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(createRoleEndpoint)
		createRoleEndpoint = logginMiddleware(createRoleEndpoint)
		svr.createRoleHandler = grpctransport.NewServer(
			createRoleEndpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		addRoleForUserEndpoint := makeAddRoleForUserEndpoint(userService)
		addRoleForUserEndpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(addRoleForUserEndpoint)
		addRoleForUserEndpoint = logginMiddleware(addRoleForUserEndpoint)
		svr.addRoleForUserHandler = grpctransport.NewServer(
			addRoleForUserEndpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	return svr
}

func (u *userServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	_, res, err := u.loginHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.LoginReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) GetPermissions(ctx context.Context, in *pb.GetPermissionsRequest) (*pb.GetPermissionsReply, error) {
	_, res, err := u.getPermissionsHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.GetPermissionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) CheckPermission(ctx context.Context, in *pb.CheckPermissionRequest) (*pb.CheckPermissionReply, error) {
	_, res, err := u.checkPermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.CheckPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	_, res, err := u.registerHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RegisterReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) AddPermission(ctx context.Context, in *pb.AddPermissionRequest) (*pb.AddPermissionReply, error) {
	_, res, err := u.addPermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) AddRole(ctx context.Context, in *pb.AddRoleRequest) (*pb.AddRoleReply, error) {
	_, res, err := u.addRoleHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) CreateRole(ctx context.Context, in *pb.CreateRoleRequest) (*pb.CreateRoleReply, error) {
	_, res, err := u.createRoleHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.CreateRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) AddRoleForUser(ctx context.Context, in *pb.AddRoleForUserRequest) (*pb.AddRoleForUserReply, error) {
	_, res, err := u.addRoleForUserHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddRoleForUserReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}
