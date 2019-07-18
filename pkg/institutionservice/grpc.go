package institutionservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type server struct {
	tnxHisDownload               grpctransport.Handler
	getTfrTrnLogs                grpctransport.Handler
	getTfrTrnLog                 grpctransport.Handler
	downloadTfrTrnLogs           grpctransport.Handler
	listGroupsHandler            grpctransport.Handler
	listInstitutionsHandler      grpctransport.Handler
	addInstitutionHandler        grpctransport.Handler
	addInstitutionFeeHandler     grpctransport.Handler
	addInstitutionControlHandler grpctransport.Handler
	addInstitutionCashHandler    grpctransport.Handler
}

func (g *server) TnxHisDownload(ctx context.Context, in *pb.InstitutionTnxHisDownloadReq) (*pb.InstitutionTnxHisDownloadResp, error) {
	_, res, err := g.tnxHisDownload.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.InstitutionTnxHisDownloadResp), nil
}

func (g *server) GetTfrTrnLogs(ctx context.Context, in *pb.GetTfrTrnLogsReq) (*pb.GetTfrTrnLogsResp, error) {
	_, res, err := g.getTfrTrnLogs.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetTfrTrnLogsResp), nil
}

func (g *server) GetTfrTrnLog(ctx context.Context, in *pb.GetTfrTrnLogReq) (*pb.GetTfrTrnLogResp, error) {
	_, res, err := g.getTfrTrnLog.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetTfrTrnLogResp), nil
}

func (g *server) DownloadTfrTrnLogs(ctx context.Context, in *pb.DownloadTfrTrnLogsReq) (*pb.DownloadTfrTrnLogsResp, error) {
	_, res, err := g.downloadTfrTrnLogs.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.DownloadTfrTrnLogsResp), nil
}

//Newserver .
func New(tracer grpctransport.ServerOption) pb.InstitutionServer {
	svc := new(service)
	svr := new(server)
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		e := MakeTnxHisDownloadEndpoint(svc)
		svr.tnxHisDownload = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}
	{
		e := MakeGetTfrTrnLogsEndpoint(svc)
		svr.getTfrTrnLogs = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}
	{
		e := MakeGetTfrTrnLogEndpoint(svc)
		svr.getTfrTrnLog = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}
	{
		e := MakeDownloadTfrTrnLogsEndpoint(svc)
		svr.downloadTfrTrnLogs = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeListGroupsEndpoint(svc)
		svr.listGroupsHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeListInstitutionsEndpoint(svc)
		svr.listInstitutionsHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeSaveInstitutionEndpoint(svc)
		svr.addInstitutionHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeSaveInstitutionFeeEndpoint(svc)
		svr.addInstitutionFeeHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeSaveInstitutionControlEndpoint(svc)
		svr.addInstitutionControlHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeSaveInstitutionCashEndpoint(svc)
		svr.addInstitutionCashHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	return svr
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

func (g *server) ListGroups(ctx context.Context, in *pb.ListGroupsRequest) (*pb.ListInstitutionsReply, error) {
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

func (g *server) ListInstitutions(ctx context.Context, in *pb.ListInstitutionsRequest) (*pb.ListInstitutionsReply, error) {
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

func (g *server) SaveInstitution(ctx context.Context, in *pb.SaveInstitutionRequest) (*pb.SaveInstitutionReply, error) {
	_, res, err := g.addInstitutionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil

}

func (g *server) SaveInstitutionFee(ctx context.Context, in *pb.SaveInstitutionFeeRequest) (*pb.SaveInstitutionFeeReply, error) {
	_, res, err := g.addInstitutionFeeHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionFeeReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (g *server) SaveInstitutionControl(ctx context.Context, in *pb.SaveInstitutionControlRequest) (*pb.SaveInstitutionControlReply, error) {
	_, res, err := g.addInstitutionControlHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionControlReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}
func (g *server) SaveInstitutionCash(ctx context.Context, in *pb.SaveInstitutionCashRequest) (*pb.SaveInstitutionCashReply, error) {
	_, res, err := g.addInstitutionCashHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionCashReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}
