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
	}

	return endpoints
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
