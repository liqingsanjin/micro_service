package institutionservice

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrDownloadFileNameEmpty = status.Error(codes.NotFound, "empty file name")
)
