package gateway

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type InstitutionEndpoints struct {
	TnxHisDownloadEndpoint     endpoint.Endpoint
	GetTfrTrnLogsEndpoint      endpoint.Endpoint
	GetTfrTrnLogEndpoint       endpoint.Endpoint
	DownloadTfrTrnLogsEndpoint endpoint.Endpoint
}

func NewInstitutionServiceGRPCClient(conn *grpc.ClientConn, tracer kitgrpc.ClientOption) *InstitutionEndpoints {
	endpoints := new(InstitutionEndpoints)
	options := make([]kitgrpc.ClientOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}
	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"TnxHisDownload",
			encodeRequest,
			decodeResponse,
			pb.InstitutionTnxHisDownloadResp{},
			options...,
		).Endpoint()
		endpoints.TnxHisDownloadEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"GetTfrTrnLogs",
			encodeRequest,
			decodeResponse,
			pb.GetTfrTrnLogsResp{},
			options...,
		).Endpoint()
		endpoints.GetTfrTrnLogsEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"GetTfrTrnLog",
			encodeRequest,
			decodeResponse,
			pb.GetTfrTrnLogResp{},
			options...,
		).Endpoint()
		endpoints.GetTfrTrnLogEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"DownloadTfrTrnLogs",
			encodeRequest,
			decodeResponse,
			pb.DownloadTfrTrnLogsResp{},
			options...,
		).Endpoint()
		endpoints.DownloadTfrTrnLogsEndpoint = endpoint
	}
	return endpoints
}
func (s *InstitutionEndpoints) TnxHisDownload(ctx context.Context, in *pb.InstitutionTnxHisDownloadReq) (*pb.InstitutionTnxHisDownloadResp, error) {
	res, err := s.TnxHisDownloadEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.InstitutionTnxHisDownloadResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *InstitutionEndpoints) GetTfrTrnLogs(ctx context.Context, in *pb.GetTfrTrnLogsReq) (*pb.GetTfrTrnLogsResp, error) {
	res, err := s.GetTfrTrnLogsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.GetTfrTrnLogsResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *InstitutionEndpoints) GetTfrTrnLog(ctx context.Context, in *pb.GetTfrTrnLogReq) (*pb.GetTfrTrnLogResp, error) {
	res, err := s.GetTfrTrnLogEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.GetTfrTrnLogResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *InstitutionEndpoints) DownloadTfrTrnLogs(ctx context.Context, in *pb.DownloadTfrTrnLogsReq) (*pb.DownloadTfrTrnLogsResp, error) {
	res, err := s.DownloadTfrTrnLogsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.DownloadTfrTrnLogsResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}
