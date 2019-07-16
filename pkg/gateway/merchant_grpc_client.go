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

type MerchantEndpoints struct {
	ListMerchantEndpoint            endpoint.Endpoint
	ListGroupMerchantEndpoint       endpoint.Endpoint
	SaveMerchantEndpoint            endpoint.Endpoint
	SaveMerchantBankAccountEndpoint endpoint.Endpoint
}

func NewMerchantServiceClient(conn *grpc.ClientConn, tracer kitgrpc.ClientOption) *MerchantEndpoints {
	endpoints := new(MerchantEndpoints)
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

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"SaveMerchant",
			encodeRequest,
			decodeResponse,
			pb.SaveMerchantReply{},
			options...,
		).Endpoint()
		endpoints.SaveMerchantEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"SaveMerchantBankAccount",
			encodeRequest,
			decodeResponse,
			pb.SaveMerchantBankAccountReply{},
			options...,
		).Endpoint()
		endpoints.SaveMerchantBankAccountEndpoint = endpoint
	}

	return endpoints
}

func (m *MerchantEndpoints) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
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

func (m *MerchantEndpoints) ListGroupMerchant(ctx context.Context, in *pb.ListGroupMerchantRequest) (*pb.ListGroupMerchantReply, error) {
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

func (m *MerchantEndpoints) SaveMerchant(ctx context.Context, in *pb.SaveMerchantRequest) (*pb.SaveMerchantReply, error) {
	res, err := m.SaveMerchantEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantReply), nil
}

func (m *MerchantEndpoints) SaveMerchantBankAccount(ctx context.Context, in *pb.SaveMerchantBankAccountRequest) (*pb.SaveMerchantBankAccountReply, error) {
	res, err := m.SaveMerchantBankAccountEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBankAccountReply), nil
}
