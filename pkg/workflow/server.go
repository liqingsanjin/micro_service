package workflow

import (
	"context"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	ListTaskHandler   grpctransport.Handler
	HandleTaskHandler grpctransport.Handler
	StartHandler      grpctransport.Handler
}

func (s *server) ListTask(context.Context, *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	panic("implement me")
}

func (s *server) HandleTask(context.Context, *pb.HandleTaskRequest) (*pb.HandleTaskReply, error) {
	panic("implement me")
}

func (s *server) Start(ctx context.Context, in *pb.StartWorkflowRequest) (*pb.StartWorkflowReply, error) {
	_, res, err := s.StartHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StartWorkflowReply), nil
}

func New(tracer grpctransport.ServerOption) pb.WorkflowServer {
	svr := &server{}
	svc := &service{}
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		endpoint := MakeListTaskEndpoint(svc)
		svr.ListTaskHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeHandleTaskEndpoint(svc)
		svr.HandleTaskHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeStartWorkflowEndpoint(svc)
		svr.StartHandler = grpctransport.NewServer(
			endpoint,
			decodeRequest,
			encodeResponse,
			options...,
		)
	}

	return svr
}
