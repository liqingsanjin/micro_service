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

type MerchantEndpoint struct {
	ListMerchantEndpoint      endpoint.Endpoint
	ListGroupMerchantEndpoint endpoint.Endpoint
}

func NewMerchantServiceClient(conn *grpc.ClientConn, tracer kitgrpc.ClientOption) *MerchantEndpoint {
	endpoints := new(MerchantEndpoint)
	options := make([]kitgrpc.ClientOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"ListMerchant",
			encodeRequest,
			decodeResponse,
			pb.ListMerchantReply{},
			options...,
		).Endpoint()
		endpoints.ListMerchantEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"ListGroupMerchant",
			encodeRequest,
			decodeResponse,
			pb.ListGroupMerchantReply{},
			options...,
		).Endpoint()
		endpoints.ListGroupMerchantEndpoint = endpoint
	}

	return endpoints
}

func (m *MerchantEndpoint) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
	res, err := m.ListMerchantEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}

func (m *MerchantEndpoint) ListGroupMerchant(ctx context.Context, in *pb.ListGroupMerchantRequest) (*pb.ListGroupMerchantReply, error) {
	res, err := m.ListGroupMerchantEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListGroupMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}
