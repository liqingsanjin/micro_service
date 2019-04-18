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
func MakeDownloadEndpoint(s SetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.InstitutionTnxHisDownloadReq)
		return s.TnxHisDownload(ctx, req)
	}
}
