package userservice

import (
	"context"
	"userService/pkg/pb"

	stdjwt "github.com/dgrijalva/jwt-go"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type userServer struct {
	loginHandler                         grpctransport.Handler
	getPermissionsHandler                grpctransport.Handler
	checkPermissionHandler               grpctransport.Handler
	registerHandler                      grpctransport.Handler
	addPermissionForRoleHandler          grpctransport.Handler
	addRoleHandler                       grpctransport.Handler
	createRoleHandler                    grpctransport.Handler
	addRoleForUserHandler                grpctransport.Handler
	addRoutesHandler                     grpctransport.Handler
	listRouteHandler                     grpctransport.Handler
	createPermissionHandler              grpctransport.Handler
	updatePermissionHandler              grpctransport.Handler
	addRouteForPermissionHandler         grpctransport.Handler
	removeRouteForPermissionHandler      grpctransport.Handler
	removePermissionHandler              grpctransport.Handler
	listPermissionsHandler               grpctransport.Handler
	addPermissionForPermissionHandler    grpctransport.Handler
	removePermissionForPermissionHandler grpctransport.Handler
	listRoleHandler                      grpctransport.Handler
	updateRoleHandler                    grpctransport.Handler
	removePermissionForRoleHandler       grpctransport.Handler
}

func New() pb.UserServer {
	svr := &userServer{}
	userService := &userService{}

	{
		endpoint := makeLoginEndpoint(userService)
		endpoint = logginMiddleware(endpoint)
		svr.loginHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeGetPermissionsEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.getPermissionsHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeCheckPermissionEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.checkPermissionHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeRegisterEndpoint(userService)
		endpoint = logginMiddleware(endpoint)
		svr.registerHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeAddPermissionForRoleEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.addPermissionForRoleHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeAddRoleEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.addRoleHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeCreateRoleEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.createRoleHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeAddRoleForUserEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.addRoleForUserHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeAddRoutesEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.addRoutesHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeListRoutesEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.listRouteHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeCreatePermissionEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.createPermissionHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeUpdatePermissionEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.updatePermissionHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeAddRouteForPermissionEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.addRouteForPermissionHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeRemoveRouteForPermissionEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.removeRouteForPermissionHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeRemovePermissionEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.removePermissionHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeListPermissionsEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.listPermissionsHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeAddPermissionForPermissionEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.addPermissionForPermissionHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeRemovePermissionForPermissionEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.removePermissionForPermissionHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeListRoleEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.listRoleHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeUpdateRoleEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.updateRoleHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
		)
	}

	{
		endpoint := makeRemovePermissionForRoleEndpoint(userService)
		endpoint = jwtParser(keyFunc, stdjwt.SigningMethodHS256, UserClaimFactory)(endpoint)
		endpoint = logginMiddleware(endpoint)
		svr.removePermissionForRoleHandler = grpctransport.NewServer(
			endpoint,
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

func (u *userServer) AddPermissionForRole(ctx context.Context, in *pb.AddPermissionForRoleRequest) (*pb.AddPermissionForRoleReply, error) {
	_, res, err := u.addPermissionForRoleHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddPermissionForRoleReply)
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

func (u *userServer) AddPolicy(ctx context.Context, in *pb.AddPolicyRequest) (*pb.AddPolicyReply, error) {
	return nil, nil
}

func (u *userServer) AddRoutes(ctx context.Context, in *pb.AddRoutesRequest) (*pb.AddRoutesReply, error) {
	_, res, err := u.addRoutesHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddRoutesReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) ListRoutes(ctx context.Context, in *pb.ListRoutesRequest) (*pb.ListRoutesReply, error) {
	_, res, err := u.listRouteHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListRoutesReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) CreatePermission(ctx context.Context, in *pb.CreatePermissionRequest) (*pb.CreatePermissionReply, error) {
	_, res, err := u.createPermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.CreatePermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) UpdatePermission(ctx context.Context, in *pb.UpdatePermissionRequest) (*pb.UpdatePermissionReply, error) {
	_, res, err := u.updatePermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.UpdatePermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) AddRouteForPermission(ctx context.Context, in *pb.AddRouteForPermissionRequest) (*pb.AddRouteForPermissionReply, error) {
	_, res, err := u.addRouteForPermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddRouteForPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) RemoveRouteForPermission(ctx context.Context, in *pb.RemoveRouteForPermissionRequest) (*pb.RemoveRouteForPermissionReply, error) {
	_, res, err := u.removeRouteForPermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemoveRouteForPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) RemovePermission(ctx context.Context, in *pb.RemovePermissionRequest) (*pb.RemovePermissionReply, error) {
	_, res, err := u.removePermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemovePermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) ListPermissions(ctx context.Context, in *pb.ListPermissionsRequest) (*pb.ListPermissionsReply, error) {
	_, res, err := u.listPermissionsHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListPermissionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) AddPermissionForPermission(ctx context.Context, in *pb.AddPermissionForPermissionRequest) (*pb.AddPermissionForPermissionReply, error) {
	_, res, err := u.addPermissionForPermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddPermissionForPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) RemovePermissionForPermission(ctx context.Context, in *pb.RemovePermissionForPermissionRequest) (*pb.RemovePermissionForPermissionReply, error) {
	_, res, err := u.removePermissionForPermissionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemovePermissionForPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) ListRole(ctx context.Context, in *pb.ListRoleRequest) (*pb.ListRoleReply, error) {
	_, res, err := u.listRoleHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest) (*pb.UpdateRoleReply, error) {
	_, res, err := u.updateRoleHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.UpdateRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *userServer) RemovePermissionForRole(ctx context.Context, in *pb.RemovePermissionForRoleRequest) (*pb.RemovePermissionForRoleReply, error) {
	_, res, err := u.removePermissionForRoleHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemovePermissionForRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}
