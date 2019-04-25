package institutionservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

//SetEndpoint .
type SetEndpoint struct {
	TnxHisDownloadEndpoint endpoint.Endpoint
}

//MakeDownloadEndpoint .
func MakeDownloadEndpoint(s pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.InstitutionTnxHisDownloadReq)
		return s.TnxHisDownload(ctx, req)
	}
}

//MakeDownloadEndpoint .
func MakeGetTfrTrnLogsEndpoint(s pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetTfrTrnLogsReq)
		return s.GetTfrTrnLogs(ctx, req)
	}
}

func MakeGetTfrTrnLogEndpoint(s pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetTfrTrnLogReq)
		return s.GetTfrTrnLog(ctx, req)
	}
}
