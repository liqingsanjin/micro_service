package institutionservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	tnxHisDownload grpc.Handler
}

func (g *grpcServer) TnxHisDownload(ctx context.Context, in *pb.InstitutionTnxHisDownloadReq) (*pb.InstitutionTnxHisDownloadResp, error) {
	_, res, err := g.tnxHisDownload.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.InstitutionTnxHisDownloadResp), nil
}

//NewGRPCServer .
func NewGRPCServer(setEndpoint *SetEndpoint) pb.InstitutionServer {
	return &grpcServer{
		tnxHisDownload: grpcNewServer(setEndpoint.TnxHisDownloadEndpoint),
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
