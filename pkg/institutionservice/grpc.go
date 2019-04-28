package institutionservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	tnxHisDownload     grpc.Handler
	getTfrTrnLogs      grpc.Handler
	getTfrTrnLog       grpc.Handler
	downloadTfrTrnLogs grpc.Handler
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
	insSetService := NewSetService()
	downEndpoint := MakeDownloadEndpoint(insSetService)
	getTfrTrnlogsEndpoint := MakeGetTfrTrnLogsEndpoint(insSetService)
	getTfrTrnlogEndpoint := MakeGetTfrTrnLogEndpoint(insSetService)
	downloadTfrTrnLogsEndpoint := MakeDownloadTfrTrnLogsEndpoint(insSetService)

	return &grpcServer{
		tnxHisDownload:     grpcNewServer(downEndpoint),
		getTfrTrnLogs:      grpcNewServer(getTfrTrnlogsEndpoint),
		getTfrTrnLog:       grpcNewServer(getTfrTrnlogEndpoint),
		downloadTfrTrnLogs: grpcNewServer(downloadTfrTrnLogsEndpoint),
	}
}

func grpcNewServer(endpoint endpoint.Endpoint) *grpc.Server {
	return grpc.NewServer(endpoint, grpcDecode, grpcEncode)
}

func grpcDecode(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func grpcEncode(_ context.Context, res interface{}) (interface{}, error) {
	return res, nil
}
