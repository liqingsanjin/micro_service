package institutionservice

import (
	"userService/pkg/common"
	cleartxnM "userService/pkg/model/cleartxn"
	"userService/pkg/pb"

	"golang.org/x/net/context"
)

//SetService .
type SetService interface {
	TnxHisDownload(ctx context.Context, in *pb.InstitutionTnxHisDownloadReq) (*pb.InstitutionTnxHisDownloadResp, error)
}

//setService .
type setService struct {
}

//NewSetService return institution service with grpc registry type.
func NewSetService() SetService {
	return &setService{}
}

//Download .
func (s *setService) TnxHisDownload(ctx context.Context, in *pb.InstitutionTnxHisDownloadReq) (*pb.InstitutionTnxHisDownloadResp, error) {
	Txns, err := cleartxnM.DownloadInstitutionFile(common.DB, in.StartTime, in.EndTime)
	if err != nil {
		return nil, err
	}

	fileDir, err := DownloadFileWithDay(Txns)
	if err != nil {
		return nil, err
	}

	err = Compress(fileDir, in.Name)
	if err != nil {
		return nil, err
	}

	return &pb.InstitutionTnxHisDownloadResp{Result: true}, nil
}
