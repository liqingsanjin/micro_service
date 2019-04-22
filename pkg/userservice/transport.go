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

func makeAddPermissionForRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddPermissionForRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddPermissionForRole(ctx, req)
	}
}

func makeAddRoleForRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddRoleForRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddRoleForRole(ctx, req)
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

func makeAddRouteForPermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddRouteForPermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddRouteForPermission(ctx, req)
	}
}

func makeRemoveRouteForPermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.RemoveRouteForPermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.RemoveRouteForPermission(ctx, req)
	}
}

func makeRemovePermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.RemovePermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.RemovePermission(ctx, req)
	}
}

func makeListPermissionsEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListPermissionsRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.ListPermissions(ctx, req)
	}
}

func makeAddPermissionForPermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddPermissionForPermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddPermissionForPermission(ctx, req)
	}
}

func makeRemovePermissionForPermissionEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.RemovePermissionForPermissionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.RemovePermissionForPermission(ctx, req)
	}
}

func makeListRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.ListRole(ctx, req)
	}
}

func makeUpdateRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.UpdateRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.UpdateRole(ctx, req)
	}
}

func makeRemovePermissionForRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.RemovePermissionForRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.RemovePermissionForRole(ctx, req)
	}
}

func makeRemoveRoleForRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.RemoveRoleForRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.RemoveRoleForRole(ctx, req)
	}
}

func makeRemoveRoleEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.RemoveRoleRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.RemoveRole(ctx, req)
	}
}

func makeListUsersEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListUsersRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.ListUsers(ctx, req)
	}
}

func makeUpdateUserEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.UpdateUserRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.UpdateUser(ctx, req)
	}
}

func makeAddPermissionForUserEndpoint(service pb.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddPermissionForUserRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddPermissionForUser(ctx, req)
	}
}

func decodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func encodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
