package productservice

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	ListTransMapHandler grpctransport.Handler
	ListFeeMapHandler   grpctransport.Handler
}

func (s *server) ListFeeMap(ctx context.Context, in *pb.ListFeeMapRequest) (*pb.ListFeeMapReply, error) {
	_, res, err := s.ListFeeMapHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListFeeMapReply), nil
}

func (s *server) ListTransMap(ctx context.Context, in *pb.ListTransMapRequest) (*pb.ListTransMapReply, error) {
	_, res, err := s.ListTransMapHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}

	return res.(*pb.ListTransMapReply), nil
}

func New(tracer grpctransport.ServerOption) pb.ProductServer {
	svc := new(service)
	svr := new(server)
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		endpoint := MakeListTransMapEndpoint(svc)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListTransMapHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeListFeeMapEndpoint(svc)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListFeeMapHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	return svr
}
