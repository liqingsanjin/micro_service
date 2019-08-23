package termservice

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
)

func MakeListTermInfoEndpoint(service pb.TermServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.ListTermInfo(ctx, request.(*pb.ListTermInfoRequest))
	}
}

func MakeSaveTermEndpoint(service pb.TermServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveTerm(ctx, request.(*pb.SaveTermRequest))
	}
}

func MakeSaveTermRiskEndpoint(service pb.TermServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveTermRisk(ctx, request.(*pb.SaveTermRiskRequest))
	}
}

func MakeListTermRiskEndpoint(service pb.TermServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.ListTermRisk(ctx, request.(*pb.ListTermRiskRequest))
	}
}

func MakeSaveTermActivationStateEndpoint(service pb.TermServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.SaveTermActivationState(ctx, request.(*pb.SaveTermActivationStateRequest))
	}
}

func MakeListTermActivationStateEndpoint(service pb.TermServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.ListTermActivationState(ctx, request.(*pb.ListTermActivationStateRequest))
	}
}
