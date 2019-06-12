package merchantservice

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type merchantServer struct {
	MerchantQueryHandler grpctransport.Handler
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
		svr.MerchantQueryHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	return svr
}

func (m *merchantServer) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
	_, res, err := m.MerchantQueryHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}
