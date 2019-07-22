package staticservice

import (
	"context"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	syncData                         grpctransport.Handler
	getDictionaryItem                grpctransport.Handler
	getDicByProdAndBiz               grpctransport.Handler
	getDicByInsCmpCd                 grpctransport.Handler
	checkValues                      grpctransport.Handler
	getDictionaryLayerItem           grpctransport.Handler
	getDictionaryItemByPk            grpctransport.Handler
	GetUnionPayBankListByCodeHandler grpctransport.Handler
}

func (g *server) GetUnionPayBankListByCode(ctx context.Context, in *pb.GetUnionPayBankListByCodeRequest) (*pb.GetUnionPayBankListByCodeReply, error) {
	_, res, err := g.GetUnionPayBankListByCodeHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetUnionPayBankListByCodeReply), nil
}

func (g *server) GetDictionaryItemByPk(ctx context.Context, in *pb.GetDictionaryItemByPkReq) (*pb.GetDictionaryItemByPkResp, error) {
	_, res, err := g.getDictionaryItemByPk.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetDictionaryItemByPkResp), nil
}

func (g *server) GetDictionaryLayerItem(ctx context.Context, in *pb.GetDictionaryLayerItemReq) (*pb.GetDictionaryLayerItemResp, error) {
	_, res, err := g.getDictionaryLayerItem.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetDictionaryLayerItemResp), nil
}

func (g *server) SyncData(ctx context.Context, in *pb.StaticSyncDataReq) (*pb.StaticSyncDataResp, error) {
	_, res, err := g.syncData.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticSyncDataResp), nil
}

func (g *server) GetDictionaryItem(ctx context.Context, in *pb.StaticGetDictionaryItemReq) (*pb.StaticGetDictionaryItemResp, error) {
	_, res, err := g.getDictionaryItem.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticGetDictionaryItemResp), nil
}

func (g *server) GetDicByProdAndBiz(ctx context.Context, in *pb.StaticGetDicByProdAndBizReq) (*pb.StaticGetDicByProdAndBizResp, error) {
	_, res, err := g.getDicByProdAndBiz.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticGetDicByProdAndBizResp), nil
}

func (g *server) GetDicByInsCmpCd(ctx context.Context, in *pb.StaticGetDicByInsCmpCdReq) (*pb.StaticGetDicByInsCmpCdResp, error) {
	_, res, err := g.getDicByInsCmpCd.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticGetDicByInsCmpCdResp), nil
}

func (g *server) CheckValues(ctx context.Context, in *pb.StaticCheckValuesReq) (*pb.StaticCheckValuesResp, error) {
	_, res, err := g.checkValues.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticCheckValuesResp), nil
}

//Newserver .
func New(tracer grpctransport.ServerOption) pb.StaticServer {
	svr := new(server)
	svc := new(service)
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		e := MakeSyncDataEndpoint(svc)
		svr.syncData = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeGetDictionaryItemEndpoint(svc)
		svr.getDictionaryItem = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeGetDicByProdAndBizEndpoint(svc)
		svr.getDicByProdAndBiz = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeGetDicByInsCmpCdEndpoint(svc)
		svr.getDicByInsCmpCd = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeCheckValuesEndpoint(svc)
		svr.checkValues = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeGetDictionaryLayerItemEndpoint(svc)
		svr.getDictionaryLayerItem = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeGetDictionaryItemByPkEndpoint(svc)
		svr.getDictionaryItemByPk = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeGetUnionPayBankListByCodeEndpoint(svc)
		svr.GetUnionPayBankListByCodeHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	return svr
}

func grpcDecode(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func grpcEncode(_ context.Context, res interface{}) (interface{}, error) {
	return res, nil
}
