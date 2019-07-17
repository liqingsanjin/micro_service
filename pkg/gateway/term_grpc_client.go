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
	SaveTermEndpoint     endpoint.Endpoint
	SaveTermRiskEndpoint endpoint.Endpoint
}

func (t *TermEndpoints) SaveTermRisk(ctx context.Context, in *pb.SaveTermRiskRequest) (*pb.SaveTermRiskReply, error) {
	res, err := t.SaveTermRiskEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveTermRiskReply), nil
}

func (t *TermEndpoints) SaveTerm(ctx context.Context, in *pb.SaveTermRequest) (*pb.SaveTermReply, error) {
	res, err := t.SaveTermEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveTermReply), nil
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

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Term",
			"SaveTerm",
			encodeRequest,
			decodeResponse,
			pb.SaveTermReply{},
			options...,
		).Endpoint()
		endpoints.SaveTermEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Term",
			"SaveTermRisk",
			encodeRequest,
			decodeResponse,
			pb.SaveTermRiskReply{},
			options...,
		).Endpoint()
		endpoints.SaveTermRiskEndpoint = endpoint
	}

	return endpoints
}
