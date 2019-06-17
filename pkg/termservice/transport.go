package termservice

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
)

func MakeListTermInfoEndpoint(service pb.TermServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListTermInfoRequest)
		if !ok {
			return nil, kit.ErrRequestTypeInvalid
		}
		return service.ListTermInfo(ctx, req)
	}
}
