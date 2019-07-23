package gateway

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type MerchantEndpoints struct {
	ListMerchantEndpoint            endpoint.Endpoint
	ListGroupMerchantEndpoint       endpoint.Endpoint
	SaveMerchantEndpoint            endpoint.Endpoint
	SaveMerchantBankAccountEndpoint endpoint.Endpoint
	SaveMerchantBizDealEndpoint     endpoint.Endpoint
	SaveMerchantBizFeeEndpoint      endpoint.Endpoint
	SaveMerchantBusinessEndpoint    endpoint.Endpoint
	SaveMerchantPictureEndpoint     endpoint.Endpoint
	GetMerchantBankAccountEndpoint  endpoint.Endpoint
	GetMerchantBizDealEndpoint      endpoint.Endpoint
	GetMerchantBizFeeEndpoint       endpoint.Endpoint
	GetMerchantBusinessEndpoint     endpoint.Endpoint
	GetMerchantPictureEndpoint      endpoint.Endpoint
	GetMerchantByIdEndpoint         endpoint.Endpoint
}

func (m *MerchantEndpoints) GetMerchantById(ctx context.Context, in *pb.GetMerchantByIdRequest) (*pb.GetMerchantByIdReply, error) {
	res, err := m.GetMerchantByIdEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetMerchantByIdReply), nil
}

func (m *MerchantEndpoints) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
	res, err := m.ListMerchantEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}

func (m *MerchantEndpoints) ListGroupMerchant(ctx context.Context, in *pb.ListGroupMerchantRequest) (*pb.ListGroupMerchantReply, error) {
	res, err := m.ListGroupMerchantEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListGroupMerchantReply)
	if !ok {
		return nil, kit.ErrReplyTypeInvalid
	}
	return reply, nil
}

func (m *MerchantEndpoints) SaveMerchant(ctx context.Context, in *pb.SaveMerchantRequest) (*pb.SaveMerchantReply, error) {
	res, err := m.SaveMerchantEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantReply), nil
}

func (m *MerchantEndpoints) SaveMerchantBankAccount(ctx context.Context, in *pb.SaveMerchantBankAccountRequest) (*pb.SaveMerchantBankAccountReply, error) {
	res, err := m.SaveMerchantBankAccountEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBankAccountReply), nil
}

func (m *MerchantEndpoints) SaveMerchantBizDeal(ctx context.Context, in *pb.SaveMerchantBizDealRequest) (*pb.SaveMerchantBizDealReply, error) {
	res, err := m.SaveMerchantBizDealEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBizDealReply), nil
}

func (m *MerchantEndpoints) SaveMerchantBizFee(ctx context.Context, in *pb.SaveMerchantBizFeeRequest) (*pb.SaveMerchantBizFeeReply, error) {
	res, err := m.SaveMerchantBizFeeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBizFeeReply), nil
}

func (m *MerchantEndpoints) SaveMerchantBusiness(ctx context.Context, in *pb.SaveMerchantBusinessRequest) (*pb.SaveMerchantBusinessReply, error) {
	res, err := m.SaveMerchantBusinessEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantBusinessReply), nil
}

func (m *MerchantEndpoints) SaveMerchantPicture(ctx context.Context, in *pb.SaveMerchantPictureRequest) (*pb.SaveMerchantPictureReply, error) {
	res, err := m.SaveMerchantPictureEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveMerchantPictureReply), nil
}

func (m *MerchantEndpoints) GetMerchantBankAccount(ctx context.Context, in *pb.GetMerchantBankAccountRequest) (*pb.GetMerchantBankAccountReply, error) {
	res, err := m.GetMerchantBankAccountEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetMerchantBankAccountReply), nil
}

func (m *MerchantEndpoints) GetMerchantBizDeal(ctx context.Context, in *pb.GetMerchantBizDealRequest) (*pb.GetMerchantBizDealReply, error) {
	res, err := m.GetMerchantBizDealEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetMerchantBizDealReply), nil
}

func (m *MerchantEndpoints) GetMerchantBizFee(ctx context.Context, in *pb.GetMerchantBizFeeRequest) (*pb.GetMerchantBizFeeReply, error) {
	res, err := m.GetMerchantBizFeeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetMerchantBizFeeReply), nil
}

func (m *MerchantEndpoints) GetMerchantBusiness(ctx context.Context, in *pb.GetMerchantBusinessRequest) (*pb.GetMerchantBusinessReply, error) {
	res, err := m.GetMerchantBusinessEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetMerchantBusinessReply), nil
}

func (m *MerchantEndpoints) GetMerchantPicture(ctx context.Context, in *pb.GetMerchantPictureRequest) (*pb.GetMerchantPictureReply, error) {
	res, err := m.GetMerchantPictureEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetMerchantPictureReply), nil
}

func NewMerchantServiceClient(conn *grpc.ClientConn, tracer kitgrpc.ClientOption) *MerchantEndpoints {
	endpoints := new(MerchantEndpoints)
	options := make([]kitgrpc.ClientOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"ListMerchant",
			encodeRequest,
			decodeResponse,
			pb.ListMerchantReply{},
			options...,
		).Endpoint()
		endpoints.ListMerchantEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"ListGroupMerchant",
			encodeRequest,
			decodeResponse,
			pb.ListGroupMerchantReply{},
			options...,
		).Endpoint()
		endpoints.ListGroupMerchantEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"SaveMerchant",
			encodeRequest,
			decodeResponse,
			pb.SaveMerchantReply{},
			options...,
		).Endpoint()
		endpoints.SaveMerchantEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"SaveMerchantBankAccount",
			encodeRequest,
			decodeResponse,
			pb.SaveMerchantBankAccountReply{},
			options...,
		).Endpoint()
		endpoints.SaveMerchantBankAccountEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"SaveMerchantBizDeal",
			encodeRequest,
			decodeResponse,
			pb.SaveMerchantBizDealReply{},
			options...,
		).Endpoint()
		endpoints.SaveMerchantBizDealEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"SaveMerchantBizFee",
			encodeRequest,
			decodeResponse,
			pb.SaveMerchantBizFeeReply{},
			options...,
		).Endpoint()
		endpoints.SaveMerchantBizFeeEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"SaveMerchantBusiness",
			encodeRequest,
			decodeResponse,
			pb.SaveMerchantBusinessReply{},
			options...,
		).Endpoint()
		endpoints.SaveMerchantBusinessEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"SaveMerchantPicture",
			encodeRequest,
			decodeResponse,
			pb.SaveMerchantPictureReply{},
			options...,
		).Endpoint()
		endpoints.SaveMerchantPictureEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"GetMerchantBankAccount",
			encodeRequest,
			decodeResponse,
			pb.GetMerchantBankAccountReply{},
			options...,
		).Endpoint()
		endpoints.GetMerchantBankAccountEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"GetMerchantBizDeal",
			encodeRequest,
			decodeResponse,
			pb.GetMerchantBizDealReply{},
			options...,
		).Endpoint()
		endpoints.GetMerchantBizDealEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"GetMerchantBizFee",
			encodeRequest,
			decodeResponse,
			pb.GetMerchantBizFeeReply{},
			options...,
		).Endpoint()
		endpoints.GetMerchantBizFeeEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"GetMerchantBusiness",
			encodeRequest,
			decodeResponse,
			pb.GetMerchantBusinessReply{},
			options...,
		).Endpoint()
		endpoints.GetMerchantBusinessEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"GetMerchantPicture",
			encodeRequest,
			decodeResponse,
			pb.GetMerchantPictureReply{},
			options...,
		).Endpoint()
		endpoints.GetMerchantPictureEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.Merchant",
			"GetMerchantById",
			encodeRequest,
			decodeResponse,
			pb.GetMerchantByIdReply{},
			options...,
		).Endpoint()
		endpoints.GetMerchantByIdEndpoint = endpoint
	}

	return endpoints
}
