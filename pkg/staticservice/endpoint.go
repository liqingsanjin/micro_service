package staticservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

//SetEndpoint .
type SetEndpoint struct {
	SyncDataEndpoint           endpoint.Endpoint
	GetDictionaryItemEndpoint  endpoint.Endpoint
	GetDicByProdAndBizEndpoint endpoint.Endpoint
	GetDicByInsCmpCdEndpoint   endpoint.Endpoint
	CheckValuesEndpoint        endpoint.Endpoint
}

//MakeSyncDataEndpoint .
func MakeSyncDataEndpoint(s pb.StaticServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.StaticSyncDataReq)
		return s.SyncData(ctx, req)
	}
}

//MakeGetDictionaryItemEndpoint .
func MakeGetDictionaryItemEndpoint(s pb.StaticServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.StaticGetDictionaryItemReq)
		return s.GetDictionaryItem(ctx, req)
	}
}

//MakeGetDicByProdAndBizEndpoint .
func MakeGetDicByProdAndBizEndpoint(s pb.StaticServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.StaticGetDicByProdAndBizReq)
		return s.GetDicByProdAndBiz(ctx, req)
	}
}

//MakeGetDicByInsCmpCdEndpoint .
func MakeGetDicByInsCmpCdEndpoint(s pb.StaticServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.StaticGetDicByInsCmpCdReq)
		return s.GetDicByInsCmpCd(ctx, req)
	}
}

func MakeCheckValuesEndpoint(s pb.StaticServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.StaticCheckValuesReq)
		return s.CheckValues(ctx, req)
	}
}
