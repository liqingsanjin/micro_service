package merchantservice

import (
	"context"

	"userService/pkg/kit"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
)

func MakeListMerchantEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListMerchantRequest)
		if !ok {
			return nil, kit.ErrRequestTypeInvalid
		}
		return service.ListMerchant(ctx, req)
	}
}

func MakeListGroupMerchantEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*pb.ListGroupMerchantRequest)
		if !ok {
			return nil, kit.ErrRequestTypeInvalid
		}
		return service.ListGroupMerchant(ctx, req)
	}
}

func MakeSaveMerchantEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveMerchant(ctx, request.(*pb.SaveMerchantRequest))
	}
}

func MakeSaveMerchantBankAccountEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveMerchantBankAccount(ctx, request.(*pb.SaveMerchantBankAccountRequest))
	}
}

func MakeSaveMerchantBizDealEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveMerchantBizDeal(ctx, request.(*pb.SaveMerchantBizDealRequest))
	}
}

func MakeSaveMerchantBizFeeEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveMerchantBizFee(ctx, request.(*pb.SaveMerchantBizFeeRequest))
	}
}

func MakeSaveMerchantBusinessEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveMerchantBusiness(ctx, request.(*pb.SaveMerchantBusinessRequest))
	}
}

func MakeSaveMerchantPictureEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveMerchantPicture(ctx, request.(*pb.SaveMerchantPictureRequest))
	}
}

func MakeGetMerchantBankAccountEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.GetMerchantBankAccount(ctx, request.(*pb.GetMerchantBankAccountRequest))
	}
}

func MakeGetMerchantBizDealEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.GetMerchantBizDeal(ctx, request.(*pb.GetMerchantBizDealRequest))
	}
}

func MakeGetMerchantBizFeeEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.GetMerchantBizFee(ctx, request.(*pb.GetMerchantBizFeeRequest))
	}
}

func MakeGetMerchantBusinessEndpoint(service pb.MerchantServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.GetMerchantBusiness(ctx, request.(*pb.GetMerchantBusinessRequest))
	}
}
