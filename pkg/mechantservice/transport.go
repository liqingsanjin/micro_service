package merchantservice

import (
	"context"

	"userService/pkg/kit"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
)

func MakeMerchantQueryEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.MerchantQueryRequest)
		if !ok {
			return nil, kit.ErrRequestTypeInvalid
		}
		return service.MerchantQuery(ctx, req)
	}
}
