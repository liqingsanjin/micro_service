package institutionservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	tnxHisDownload           grpctransport.Handler
	getTfrTrnLogs            grpctransport.Handler
	getTfrTrnLog             grpctransport.Handler
	downloadTfrTrnLogs       grpctransport.Handler
	listGroupsHandler        grpctransport.Handler
	listInstitutionsHandler  grpctransport.Handler
	addInstitutionHandler    grpctransport.Handler
	addInstitutionFeeHandler grpctransport.Handler
}

func (g *grpcServer) TnxHisDownload(ctx context.Context, in *pb.InstitutionTnxHisDownloadReq) (*pb.InstitutionTnxHisDownloadResp, error) {
	_, res, err := g.tnxHisDownload.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.InstitutionTnxHisDownloadResp), nil
}

func (g *grpcServer) GetTfrTrnLogs(ctx context.Context, in *pb.GetTfrTrnLogsReq) (*pb.GetTfrTrnLogsResp, error) {
	_, res, err := g.getTfrTrnLogs.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetTfrTrnLogsResp), nil
}

func (g *grpcServer) GetTfrTrnLog(ctx context.Context, in *pb.GetTfrTrnLogReq) (*pb.GetTfrTrnLogResp, error) {
	_, res, err := g.getTfrTrnLog.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetTfrTrnLogResp), nil
}

func (g *grpcServer) DownloadTfrTrnLogs(ctx context.Context, in *pb.DownloadTfrTrnLogsReq) (*pb.DownloadTfrTrnLogsResp, error) {
	_, res, err := g.downloadTfrTrnLogs.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.DownloadTfrTrnLogsResp), nil
}

//NewGRPCServer .
func NewGRPCServer() pb.InstitutionServer {
	options := make([]grpctransport.ServerOption, 0)

	service := NewSetService()
	downEndpoint := MakeTnxHisDownloadEndpoint(service)
	getTfrTrnlogsEndpoint := MakeGetTfrTrnLogsEndpoint(service)
	getTfrTrnlogEndpoint := MakeGetTfrTrnLogEndpoint(service)
	downloadTfrTrnLogsEndpoint := MakeDownloadTfrTrnLogsEndpoint(service)

	server := &grpcServer{
		tnxHisDownload:     grpcNewServer(downEndpoint),
		getTfrTrnLogs:      grpcNewServer(getTfrTrnlogsEndpoint),
		getTfrTrnLog:       grpcNewServer(getTfrTrnlogEndpoint),
		downloadTfrTrnLogs: grpcNewServer(downloadTfrTrnLogsEndpoint),
	}

	{
		endpoint := MakeListGroupsEndpoint(service)
		server.listGroupsHandler = grpctransport.NewServer(
			endpoint,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		endpoint := MakeListInstitutionsEndpoint(service)
		server.listInstitutionsHandler = grpctransport.NewServer(
			endpoint,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		endpoint := MakeAddInstitutionEndpoint(service)
		server.addInstitutionHandler = grpctransport.NewServer(
			endpoint,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		endpoint := MakeAddInstitutionFeeEndpoint(service)
		server.addInstitutionFeeHandler = grpctransport.NewServer(
			endpoint,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	return server
}

func grpcNewServer(endpoint endpoint.Endpoint) *grpctransport.Server {
	return grpctransport.NewServer(endpoint, grpcDecode, grpcEncode)
}

func grpcDecode(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func grpcEncode(_ context.Context, res interface{}) (interface{}, error) {
	return res, nil
}

func (g *grpcServer) ListGroups(ctx context.Context, in *pb.ListGroupsRequest) (*pb.ListInstitutionsReply, error) {
	_, res, err := g.listGroupsHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListInstitutionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (g *grpcServer) ListInstitutions(ctx context.Context, in *pb.ListInstitutionsRequest) (*pb.ListInstitutionsReply, error) {
	_, res, err := g.listInstitutionsHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListInstitutionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (g *grpcServer) AddInstitution(ctx context.Context, in *pb.AddInstitutionRequest) (*pb.AddInstitutionReply, error) {
	_, res, err := g.addInstitutionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddInstitutionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil

}

func (g *grpcServer) AddInstitutionFee(ctx context.Context, in *pb.AddInstitutionFeeRequest) (*pb.AddInstitutionFeeReply, error) {
	_, res, err := g.addInstitutionFeeHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddInstitutionFeeReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}
