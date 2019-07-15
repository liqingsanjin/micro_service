package merchantservice

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type merchantServer struct {
	ListMerchantHandler      grpctransport.Handler
	ListGroupMerchantHandler grpctransport.Handler
	SaveMerchantHandler      grpctransport.Handler
}

func New(tracer grpctransport.ServerOption) pb.MerchantServer {
	svr := &merchantServer{}
	service := &merchantService{}
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}
	{
		endpoint := MakeListMerchantEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListMerchantHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeListGroupMerchantEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListGroupMerchantHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeSaveMerchantEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.SaveMerchantHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	return svr
}

func (m *merchantServer) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
	_, res, err := m.ListMerchantHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}
func (m *merchantServer) ListGroupMerchant(ctx context.Context, in *pb.ListGroupMerchantRequest) (*pb.ListGroupMerchantReply, error) {
	_, res, err := m.ListGroupMerchantHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListGroupMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}
func (m *merchantServer) SaveMerchant(ctx context.Context, in *pb.SaveMerchantRequest) (*pb.SaveMerchantReply, error) {
	_, res, err := m.SaveMerchantHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantReply), nil
}
