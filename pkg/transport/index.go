package transport

import (
	"context"
	"userService/pkg/camunda/pb"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetProcessDefinitionEndpoint(service pb.ProcessDefinitionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.GetProcessDefinitionReq)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.Get(ctx, req)
	}
}

func MakeStartProcessDefinitionEndpoint(service pb.ProcessDefinitionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.StartProcessDefinitionReq)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.Start(ctx, req)
	}
}

func MakeListProcessInstanceEndpoint(service pb.ProcessInstanceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ProcessInstanceListReq)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.List(ctx, req)
	}
}
