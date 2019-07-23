package gateway

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type InstitutionEndpoints struct {
	TnxHisDownloadEndpoint         endpoint.Endpoint
	GetTfrTrnLogsEndpoint          endpoint.Endpoint
	GetTfrTrnLogEndpoint           endpoint.Endpoint
	DownloadTfrTrnLogsEndpoint     endpoint.Endpoint
	ListGroupsEndpoint             endpoint.Endpoint
	ListInstitutionsEndpoint       endpoint.Endpoint
	SaveInstitutionEndpoint        endpoint.Endpoint
	SaveInstitutionFeeEndpoint     endpoint.Endpoint
	SaveInstitutionControlEndpoint endpoint.Endpoint
	SaveInstitutionCashEndpoint    endpoint.Endpoint
	GetInstitutionByIdEndpoint     endpoint.Endpoint
}

func (s *InstitutionEndpoints) GetInstitutionById(ctx context.Context, in *pb.GetInstitutionByIdRequest) (*pb.GetInstitutionByIdReply, error) {
	res, err := s.GetInstitutionByIdEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetInstitutionByIdReply), nil
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

func (s *InstitutionEndpoints) ListGroups(ctx context.Context, in *pb.ListGroupsRequest) (*pb.ListInstitutionsReply, error) {
	res, err := s.ListGroupsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListInstitutionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *InstitutionEndpoints) ListInstitutions(ctx context.Context, in *pb.ListInstitutionsRequest) (*pb.ListInstitutionsReply, error) {
	res, err := s.ListInstitutionsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListInstitutionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *InstitutionEndpoints) SaveInstitution(ctx context.Context, in *pb.SaveInstitutionRequest) (*pb.SaveInstitutionReply, error) {
	res, err := s.SaveInstitutionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *InstitutionEndpoints) SaveInstitutionFee(ctx context.Context, in *pb.SaveInstitutionFeeRequest) (*pb.SaveInstitutionFeeReply, error) {
	res, err := s.SaveInstitutionFeeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionFeeReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *InstitutionEndpoints) SaveInstitutionControl(ctx context.Context, in *pb.SaveInstitutionControlRequest) (*pb.SaveInstitutionControlReply, error) {
	res, err := s.SaveInstitutionControlEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionControlReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *InstitutionEndpoints) SaveInstitutionCash(ctx context.Context, in *pb.SaveInstitutionCashRequest) (*pb.SaveInstitutionCashReply, error) {
	res, err := s.SaveInstitutionCashEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionCashReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
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

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"ListGroups",
			encodeRequest,
			decodeResponse,
			pb.ListInstitutionsReply{},
			options...,
		).Endpoint()
		endpoints.ListGroupsEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"ListInstitutions",
			encodeRequest,
			decodeResponse,
			pb.ListInstitutionsReply{},
			options...,
		).Endpoint()
		endpoints.ListInstitutionsEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"SaveInstitution",
			encodeRequest,
			decodeResponse,
			pb.SaveInstitutionReply{},
			options...,
		).Endpoint()
		endpoints.SaveInstitutionEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"SaveInstitutionFee",
			encodeRequest,
			decodeResponse,
			pb.SaveInstitutionFeeReply{},
			options...,
		).Endpoint()
		endpoints.SaveInstitutionFeeEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"SaveInstitutionControl",
			encodeRequest,
			decodeResponse,
			pb.SaveInstitutionControlReply{},
			options...,
		).Endpoint()
		endpoints.SaveInstitutionControlEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"SaveInstitutionCash",
			encodeRequest,
			decodeResponse,
			pb.SaveInstitutionCashReply{},
			options...,
		).Endpoint()
		endpoints.SaveInstitutionCashEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Institution",
			"GetInstitutionById",
			encodeRequest,
			decodeResponse,
			pb.GetInstitutionByIdReply{},
			options...,
		).Endpoint()
		endpoints.GetInstitutionByIdEndpoint = endpoint
	}

	return endpoints
}
