package staticservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	syncData               grpc.Handler
	getDictionaryItem      grpc.Handler
	getDicByProdAndBiz     grpc.Handler
	getDicByInsCmpCd       grpc.Handler
	checkValues            grpc.Handler
	getDictionaryLayerItem grpc.Handler
	getDictionaryItemByPk  grpc.Handler
}

func (g *grpcServer) GetDictionaryItemByPk(ctx context.Context, in *pb.GetDictionaryItemByPkReq) (*pb.GetDictionaryItemByPkResp, error) {
	_, res, err := g.getDictionaryItemByPk.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetDictionaryItemByPkResp), nil
}

func (g *grpcServer) GetDictionaryLayerItem(ctx context.Context, in *pb.GetDictionaryLayerItemReq) (*pb.GetDictionaryLayerItemResp, error) {
	_, res, err := g.getDictionaryLayerItem.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetDictionaryLayerItemResp), nil
}

func (g *grpcServer) SyncData(ctx context.Context, in *pb.StaticSyncDataReq) (*pb.StaticSyncDataResp, error) {
	_, res, err := g.syncData.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticSyncDataResp), nil
}

func (g *grpcServer) GetDictionaryItem(ctx context.Context, in *pb.StaticGetDictionaryItemReq) (*pb.StaticGetDictionaryItemResp, error) {
	_, res, err := g.getDictionaryItem.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticGetDictionaryItemResp), nil
}

func (g *grpcServer) GetDicByProdAndBiz(ctx context.Context, in *pb.StaticGetDicByProdAndBizReq) (*pb.StaticGetDicByProdAndBizResp, error) {
	_, res, err := g.getDicByProdAndBiz.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticGetDicByProdAndBizResp), nil
}

func (g *grpcServer) GetDicByInsCmpCd(ctx context.Context, in *pb.StaticGetDicByInsCmpCdReq) (*pb.StaticGetDicByInsCmpCdResp, error) {
	_, res, err := g.getDicByInsCmpCd.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticGetDicByInsCmpCdResp), nil
}

func (g *grpcServer) CheckValues(ctx context.Context, in *pb.StaticCheckValuesReq) (*pb.StaticCheckValuesResp, error) {
	_, res, err := g.checkValues.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StaticCheckValuesResp), nil
}

//NewGRPCServer .
func NewGRPCServer() pb.StaticServer {
	svr := &grpcServer{}

	svc := &setService{}

	{
		e := MakeSyncDataEndpoint(svc)
		svr.syncData = grpcNewServer(e)
	}

	{
		e := MakeGetDictionaryItemEndpoint(svc)
		svr.getDictionaryItem = grpcNewServer(e)
	}

	{
		e := MakeGetDicByProdAndBizEndpoint(svc)
		svr.getDicByProdAndBiz = grpcNewServer(e)
	}

	{
		e := MakeGetDicByInsCmpCdEndpoint(svc)
		svr.getDicByInsCmpCd = grpcNewServer(e)
	}

	{
		e := MakeCheckValuesEndpoint(svc)
		svr.checkValues = grpcNewServer(e)
	}

	{
		e := MakeGetDictionaryLayerItemEndpoint(svc)
		svr.getDictionaryLayerItem = grpcNewServer(e)
	}

	{
		e := MakeGetDictionaryItemByPkEndpoint(svc)
		svr.getDictionaryItemByPk = grpcNewServer(e)
	}

	return svr
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
