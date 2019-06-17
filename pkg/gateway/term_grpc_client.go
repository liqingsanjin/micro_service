package gateway

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type TermEndpoints struct {
	ListTermInfoEndpoint endpoint.Endpoint
}

func NewTermServiceClient(conn *grpc.ClientConn, tracer kitgrpc.ClientOption) *TermEndpoints {
	endpoints := new(TermEndpoints)
	options := make([]kitgrpc.ClientOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Term",
			"ListTermInfo",
			encodeRequest,
			decodeResponse,
			pb.ListTermInfoReply{},
			options...,
		).Endpoint()
		endpoints.ListTermInfoEndpoint = endpoint
	}

	return endpoints
}

func (t *TermEndpoints) ListTermInfo(ctx context.Context, in *pb.ListTermInfoRequest) (*pb.ListTermInfoReply, error) {
	res, err := t.ListTermInfoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListTermInfoReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}
