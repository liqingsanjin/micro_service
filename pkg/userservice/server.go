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
}

func New() pb.UserServer {
	svr := &userServer{}
	userService := &userService{}
	svr.loginHandler = grpctransport.NewServer(
		makeLoginEndpoint(userService),
		decodeRequest,
		encodeResponse,
	)

	getPermissionEndpoint := makeGetPermissionsEndpoint(userService)
	getPermissionEndpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(getPermissionEndpoint)
	svr.getPermissionsHandler = grpctransport.NewServer(
		getPermissionEndpoint,
		decodeRequest,
		encodeResponse,
	)

	checkPermissionEndpoint := makeCheckPermissionEndpoint(userService)
	checkPermissionEndpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(checkPermissionEndpoint)
	svr.checkPermissionHandler = grpctransport.NewServer(
		checkPermissionEndpoint,
		decodeRequest,
		encodeResponse,
	)

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
