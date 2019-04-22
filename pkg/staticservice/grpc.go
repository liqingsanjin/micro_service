package staticservice

import (
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	syncData           grpc.Handler
	getDictionaryItem  grpc.Handler
	getDicByProdAndBiz grpc.Handler
	getDicByInsCmpCd   grpc.Handler
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

//NewGRPCServer .
func NewGRPCServer() pb.StaticServer {

	insSetService := NewSetService()
	syncDataEndpoint := MakeSyncDataEndpoint(insSetService)
	getDictionaryItemEndpoint := MakeGetDictionaryItemEndpoint(insSetService)
	getDicByProdAndBizEndpoint := MakeGetDicByProdAndBizEndpoint(insSetService)
	getDicByInsCmpCdEndpoint := MakeGetDicByInsCmpCdEndpoint(insSetService)
	setEndpoint := SetEndpoint{
		SyncDataEndpoint:           syncDataEndpoint,
		GetDictionaryItemEndpoint:  getDictionaryItemEndpoint,
		GetDicByProdAndBizEndpoint: getDicByProdAndBizEndpoint,
		GetDicByInsCmpCdEndpoint:   getDicByInsCmpCdEndpoint,
	}

	return &grpcServer{
		syncData:           grpcNewServer(setEndpoint.SyncDataEndpoint),
		getDictionaryItem:  grpcNewServer(setEndpoint.GetDictionaryItemEndpoint),
		getDicByProdAndBiz: grpcNewServer(setEndpoint.GetDicByProdAndBizEndpoint),
		getDicByInsCmpCd:   grpcNewServer(setEndpoint.GetDicByInsCmpCdEndpoint),
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
