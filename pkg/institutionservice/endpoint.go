package institutionservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

//MakeDownloadEndpoint .
func MakeTnxHisDownloadEndpoint(s pb.InstitutionServer) endpoint.Endpoint {
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

func MakeDownloadTfrTrnLogsEndpoint(s pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.DownloadTfrTrnLogsReq)
		return s.DownloadTfrTrnLogs(ctx, req)
	}
}

func MakeListGroupsEndpoint(service pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListGroupsRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.ListGroups(ctx, req)
	}
}

func MakeListInstitutionsEndpoint(service pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListInstitutionsRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.ListInstitutions(ctx, req)
	}
}

func MakeAddInstitutionEndpoint(service pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddInstitutionRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddInstitution(ctx, req)
	}
}

func MakeAddInstitutionFeeEndpoint(service pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddInstitutionFeeRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddInstitutionFee(ctx, req)
	}
}

func MakeAddInstitutionControlEndpoint(service pb.InstitutionServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.AddInstitutionControlRequest)
		if !ok {
			return nil, ErrRequestTypeInvalid
		}
		return service.AddInstitutionControl(ctx, req)
	}
}
