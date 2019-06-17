package termservice

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	ListTermInfoHandler grpctransport.Handler
}

func New(tracer grpctransport.ServerOption) pb.TermServer {
	svr := &server{}
	svc := &service{}
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}
	{
		endpoint := MakeListTermInfoEndpoint(svc)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListTermInfoHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	return svr
}

func (s *server) ListTermInfo(ctx context.Context, in *pb.ListTermInfoRequest) (*pb.ListTermInfoReply, error) {
	_, res, err := s.ListTermInfoHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListTermInfoReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}
