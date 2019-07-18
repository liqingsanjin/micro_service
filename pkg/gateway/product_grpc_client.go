package gateway

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type ProductEndpoints struct {
	ListTransMapEndpoint endpoint.Endpoint
	ListFeeMapEndpoint   endpoint.Endpoint
}

func (p *ProductEndpoints) ListFeeMap(ctx context.Context, in *pb.ListFeeMapRequest) (*pb.ListFeeMapReply, error) {
	res, err := p.ListFeeMapEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListFeeMapReply), nil
}

func (p *ProductEndpoints) ListTransMap(ctx context.Context, in *pb.ListTransMapRequest) (*pb.ListTransMapReply, error) {
	res, err := p.ListTransMapEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListTransMapReply), nil
}

func NewProductServiceClient(conn *grpc.ClientConn, tracer kitgrpc.ClientOption) *ProductEndpoints {
	endpoints := new(ProductEndpoints)
	options := make([]kitgrpc.ClientOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		e := grpctransport.NewClient(
			conn,
			"pb.Product",
			"ListTransMap",
			encodeRequest,
			decodeResponse,
			pb.ListTransMapReply{},
			options...,
		).Endpoint()
		endpoints.ListTransMapEndpoint = e
	}

	{
		e := grpctransport.NewClient(
			conn,
			"pb.Product",
			"ListFeeMap",
			encodeRequest,
			decodeResponse,
			pb.ListFeeMapReply{},
			options...,
		).Endpoint()
		endpoints.ListFeeMapEndpoint = e
	}

	return endpoints
}
