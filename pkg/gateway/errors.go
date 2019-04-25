package gateway

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrReplyTypeInvalid   = status.Error(codes.Internal, "reply type invalid")
	ErrRequestTypeInvalid = status.Error(codes.Internal, "request type invalid")
)
