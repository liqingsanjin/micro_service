package merchantservice

import (
	"context"

	"userService/pkg/kit"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
)

func MakeListMerchantEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListMerchantRequest)
		if !ok {
			return nil, kit.ErrRequestTypeInvalid
		}
		return service.ListMerchant(ctx, req)
	}
}

func MakeListGroupMerchantEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListGroupMerchantRequest)
		if !ok {
			return nil, kit.ErrRequestTypeInvalid
		}
		return service.ListGroupMerchant(ctx, req)
	}
}
