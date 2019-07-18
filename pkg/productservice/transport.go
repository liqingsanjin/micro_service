package productservice

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
)

func MakeListTransMapEndpoint(service pb.ProductServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.ListTransMap(ctx, request.(*pb.ListTransMapRequest))
	}
}
