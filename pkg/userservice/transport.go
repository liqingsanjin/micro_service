package userservice

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
)

func makeLoginEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.LoginRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.Login(ctx, req)
	}
}

func makeGetPermissionsEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.GetPermissionsRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.GetPermissions(ctx, req)
	}
}

func makeCheckPermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.CheckPermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.CheckPermission(ctx, req)
	}
}

func makeRegisterEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.RegisterRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.Register(ctx, req)
	}
}

func makeAddPermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddPermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddPermission(ctx, req)
	}
}

func makeAddRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddRole(ctx, req)
	}
}

func makeCreateRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.CreateRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.CreateRole(ctx, req)
	}
}

func makeAddRoleForUserEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddRoleForUserRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddRoleForUser(ctx, req)
	}
}

func makeAddRoutesEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddRoutesRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddRoutes(ctx, req)
	}
}

func makeListRoutesEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListRoutesRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.ListRoutes(ctx, req)
	}
}

func makeCreatePermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.CreatePermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.CreatePermission(ctx, req)
	}
}

func makeUpdatePermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.UpdatePermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.UpdatePermission(ctx, req)
	}
}

func decodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func encodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
