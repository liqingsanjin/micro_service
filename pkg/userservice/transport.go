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
func decodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func encodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
