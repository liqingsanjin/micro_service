package gateway

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type StaticEndpoints struct {
	SyncDataEndpoint                  endpoint.Endpoint
	GetDictionaryItemEndpoint         endpoint.Endpoint
	GetDicByProdAndBizEndpoint        endpoint.Endpoint
	GetDicByInsCmpCdEndpoint          endpoint.Endpoint
	CheckValuesEndpoint               endpoint.Endpoint
	GetDictionaryLayerItemEndpoint    endpoint.Endpoint
	GetDictionaryItemByPkEndpoint     endpoint.Endpoint
	GetUnionPayBankListByCodeEndpoint endpoint.Endpoint
	FindUnionPayMccListEndpoint       endpoint.Endpoint
}

func (s *StaticEndpoints) FindUnionPayMccList(ctx context.Context, in *pb.FindUnionPayMccListRequest) (*pb.FindUnionPayMccListReply, error) {
	res, err := s.FindUnionPayMccListEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.FindUnionPayMccListReply), nil
}

func (s *StaticEndpoints) GetUnionPayBankListByCode(ctx context.Context, in *pb.GetUnionPayBankListByCodeRequest) (*pb.GetUnionPayBankListByCodeReply, error) {
	res, err := s.GetUnionPayBankListByCodeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetUnionPayBankListByCodeReply), nil
}

func (s *StaticEndpoints) SyncData(ctx context.Context, in *pb.StaticSyncDataReq) (*pb.StaticSyncDataResp, error) {
	res, err := s.SyncDataEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticSyncDataResp), nil
}

func (s *StaticEndpoints) GetDictionaryItem(ctx context.Context, in *pb.StaticGetDictionaryItemReq) (*pb.StaticGetDictionaryItemResp, error) {
	res, err := s.GetDictionaryItemEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.StaticGetDictionaryItemResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *StaticEndpoints) GetDicByProdAndBiz(ctx context.Context, in *pb.StaticGetDicByProdAndBizReq) (*pb.StaticGetDicByProdAndBizResp, error) {
	res, err := s.GetDicByProdAndBizEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.StaticGetDicByProdAndBizResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *StaticEndpoints) GetDicByInsCmpCd(ctx context.Context, in *pb.StaticGetDicByInsCmpCdReq) (*pb.StaticGetDicByInsCmpCdResp, error) {
	res, err := s.GetDicByInsCmpCdEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.StaticGetDicByInsCmpCdResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *StaticEndpoints) CheckValues(ctx context.Context, in *pb.StaticCheckValuesReq) (*pb.StaticCheckValuesResp, error) {
	res, err := s.CheckValuesEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.StaticCheckValuesResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *StaticEndpoints) GetDictionaryLayerItem(ctx context.Context, in *pb.GetDictionaryLayerItemReq) (*pb.GetDictionaryLayerItemResp, error) {
	res, err := s.GetDictionaryLayerItemEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.GetDictionaryLayerItemResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *StaticEndpoints) GetDictionaryItemByPk(ctx context.Context, in *pb.GetDictionaryItemByPkReq) (*pb.GetDictionaryItemByPkResp, error) {
	res, err := s.GetDictionaryItemByPkEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.GetDictionaryItemByPkResp)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func NewStaticServiceGRPCClient(conn *grpc.ClientConn, tracer kitgrpc.ClientOption) *StaticEndpoints {
	endpoints := new(StaticEndpoints)
	options := make([]kitgrpc.ClientOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}
	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"SyncData",
			encodeRequest,
			decodeResponse,
			pb.StaticSyncDataResp{},
			options...,
		).Endpoint()
		endpoints.SyncDataEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"GetDictionaryItem",
			encodeRequest,
			decodeResponse,
			pb.StaticGetDictionaryItemResp{},
			options...,
		).Endpoint()
		endpoints.GetDictionaryItemEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"GetDicByProdAndBiz",
			encodeRequest,
			decodeResponse,
			pb.StaticGetDicByProdAndBizResp{},
			options...,
		).Endpoint()
		endpoints.GetDicByProdAndBizEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"GetDicByInsCmpCd",
			encodeRequest,
			decodeResponse,
			pb.StaticGetDicByInsCmpCdResp{},
			options...,
		).Endpoint()
		endpoints.GetDicByInsCmpCdEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"CheckValues",
			encodeRequest,
			decodeResponse,
			pb.StaticCheckValuesResp{},
			options...,
		).Endpoint()
		endpoints.CheckValuesEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"GetDictionaryLayerItem",
			encodeRequest,
			decodeResponse,
			pb.GetDictionaryLayerItemResp{},
			options...,
		).Endpoint()
		endpoints.GetDictionaryLayerItemEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"GetDictionaryItemByPk",
			encodeRequest,
			decodeResponse,
			pb.GetDictionaryItemByPkResp{},
			options...,
		).Endpoint()
		endpoints.GetDictionaryItemByPkEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"GetUnionPayBankListByCode",
			encodeRequest,
			decodeResponse,
			pb.GetUnionPayBankListByCodeReply{},
			options...,
		).Endpoint()
		endpoints.GetUnionPayBankListByCodeEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Static",
			"FindUnionPayMccList",
			encodeRequest,
			decodeResponse,
			pb.FindUnionPayMccListReply{},
			options...,
		).Endpoint()
		endpoints.FindUnionPayMccListEndpoint = endpoint
	}

	return endpoints
}
