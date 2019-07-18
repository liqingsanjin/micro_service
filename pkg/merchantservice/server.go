package merchantservice

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type merchantServer struct {
	ListMerchantHandler            grpctransport.Handler
	ListGroupMerchantHandler       grpctransport.Handler
	SaveMerchantHandler            grpctransport.Handler
	SaveMerchantBankAccountHandler grpctransport.Handler
	SaveMerchantBizDealHandler     grpctransport.Handler
	SaveMerchantBizFeeHandler      grpctransport.Handler
	SaveMerchantBusinessHandler    grpctransport.Handler
	SaveMerchantPictureHandler     grpctransport.Handler
	GetMerchantBankAccountHandler  grpctransport.Handler
}

func New(tracer grpctransport.ServerOption) pb.MerchantServer {
	svr := &merchantServer{}
	service := &merchantService{}
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}
	{
		endpoint := MakeListMerchantEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListMerchantHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeListGroupMerchantEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListGroupMerchantHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeSaveMerchantEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.SaveMerchantHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeSaveMerchantBankAccountEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.SaveMerchantBankAccountHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeSaveMerchantBizDealEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.SaveMerchantBizDealHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeSaveMerchantBizFeeEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.SaveMerchantBizFeeHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeSaveMerchantBusinessEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.SaveMerchantBusinessHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeSaveMerchantPictureEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.SaveMerchantPictureHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeGetMerchantBankAccountEndpoint(service)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.GetMerchantBankAccountHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	return svr
}

func (m *merchantServer) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
	_, res, err := m.ListMerchantHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}
func (m *merchantServer) ListGroupMerchant(ctx context.Context, in *pb.ListGroupMerchantRequest) (*pb.ListGroupMerchantReply, error) {
	_, res, err := m.ListGroupMerchantHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListGroupMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}
func (m *merchantServer) SaveMerchant(ctx context.Context, in *pb.SaveMerchantRequest) (*pb.SaveMerchantReply, error) {
	_, res, err := m.SaveMerchantHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantReply), nil
}
func (m *merchantServer) SaveMerchantBankAccount(ctx context.Context, in *pb.SaveMerchantBankAccountRequest) (*pb.SaveMerchantBankAccountReply, error) {
	_, res, err := m.SaveMerchantBankAccountHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBankAccountReply), nil
}

func (m *merchantServer) SaveMerchantBizDeal(ctx context.Context, in *pb.SaveMerchantBizDealRequest) (*pb.SaveMerchantBizDealReply, error) {
	_, res, err := m.SaveMerchantBizDealHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBizDealReply), nil
}

func (m *merchantServer) SaveMerchantBizFee(ctx context.Context, in *pb.SaveMerchantBizFeeRequest) (*pb.SaveMerchantBizFeeReply, error) {
	_, res, err := m.SaveMerchantBizFeeHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBizFeeReply), nil
}

func (m *merchantServer) SaveMerchantBusiness(ctx context.Context, in *pb.SaveMerchantBusinessRequest) (*pb.SaveMerchantBusinessReply, error) {
	_, res, err := m.SaveMerchantBusinessHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBusinessReply), nil
}

func (m *merchantServer) SaveMerchantPicture(ctx context.Context, in *pb.SaveMerchantPictureRequest) (*pb.SaveMerchantPictureReply, error) {
	_, res, err := m.SaveMerchantPictureHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantPictureReply), nil
}

func (m *merchantServer) GetMerchantBankAccount(ctx context.Context, in *pb.GetMerchantBankAccountRequest) (*pb.GetMerchantBankAccountReply, error) {
	_, res, err := m.GetMerchantBankAccountHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetMerchantBankAccountReply), nil
}
