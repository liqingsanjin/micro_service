package merchantservice

import (
	"context"
	"userService/pkg/pb"
)

type merchantService struct{}

func (m *merchantService) MerchantQuery(ctx context.Context, in *pb.MerchantQueryRequest) (*pb.MerchantQueryReply, error) {
	reply := &pb.MerchantQueryReply{}
	return reply, nil
}
