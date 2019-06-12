package institutionservice

import (
	"fmt"
	"userService/pkg/common"
	cleartxnM "userService/pkg/model/cleartxn"
	insmodel "userService/pkg/model/institution"
	"userService/pkg/pb"

	"net/http"

	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
)

//setService .
type setService struct {
}

//NewSetService return institution service with grpc registry type.
func NewSetService() pb.InstitutionServer {
	return &setService{}
}

//Download .
func (s *setService) TnxHisDownload(ctx context.Context, in *pb.InstitutionTnxHisDownloadReq) (*pb.InstitutionTnxHisDownloadResp, error) {
	if in.Name == "" {
		return nil, ErrDownloadFileNameEmpty
	}

	clearTxnEnty := cleartxnM.ClearTxn{}
	Txns, err := clearTxnEnty.GetWithTime(common.DB, in.StartTime, in.EndTime)
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

//GetTfrTrnLogs .
func (s *setService) GetTfrTrnLogs(ctx context.Context, in *pb.GetTfrTrnLogsReq) (*pb.GetTfrTrnLogsResp, error) {
	var cond cleartxnM.TfrTrnLog

	cond.MchntCd = in.MchntCd
	cond.PriAcctNo = in.PriAcctNO
	cond.KeyRsp = in.KeyRsp
	cond.PriAcctNo = in.PriAcctNO
	cond.CardClass = in.CardClass
	cond.RoutInsIdCd = in.RoutIndustryInsIdCd
	cond.FwdInsIdCd = in.FwdInsIdCd
	cond.IssInsIdCd = in.IssInsIdCd
	cond.RespCd = in.RespCd
	cond.TermId = in.TermId
	cond.ProdCd = in.ProdCd
	cond.BizCd = in.BizCd
	cond.MaTransCd = in.MaTransCd
	cond.MsgResvFld2 = in.MsgResvFld2

	var accountRegion = ""
	if in.BeginAt != "" && in.EndAt != "" {
		accountRegion = fmt.Sprintf("'%s' < TRANS_DT AND TRANS_DT < '%s'", in.BeginAt, in.EndAt)
	} else if in.BeginAt != "" {
		accountRegion = fmt.Sprintf("'%s' < TRANS_DT", in.BeginAt)
	} else if in.EndAt != "" {
		accountRegion = fmt.Sprintf("TRANS_DT < '%s'", in.EndAt)
	}

	trfTrnLogsEnty := cleartxnM.TfrTrnLog{}
	results, count, total, err := trfTrnLogsEnty.GetWithLimit(common.DB, &cond, accountRegion, in.Limit, in.Page)
	if err != nil {
		return nil, err
	}

	var items []*pb.GetTfrTrnLogsItem
	for _, insTxn := range results {
		item := pb.GetTfrTrnLogsItem{
			KeyRsp:       insTxn.KeyRsp,
			MchntCd:      insTxn.MchntCd,
			CardAccptrNm: insTxn.CardAccptrNm,
			TransDt:      insTxn.TransDt,
			MaSettleDt:   insTxn.MaSettleDt,
			TransMt:      insTxn.TransMt,
			MaTransCd:    insTxn.MaTransCd,
			FwdInsIdCd:   insTxn.FwdInsIdCd,
			TransAt:      insTxn.TransAt,
			PriAcctNo:    insTxn.PriAcctNo,
			IssInsIdCd:   insTxn.IssInsIdCd,
			TermId:       insTxn.TermId,
			ProdCd:       insTxn.ProdCd,
			CardClass:    insTxn.CardClass,
			TransSt:      insTxn.TransSt,
			RespCd:       insTxn.RespCd,
		}
		items = append(items, &item)
	}

	return &pb.GetTfrTrnLogsResp{Items: items, Count: count, Total: total}, nil
}

//GetTfrTrnLog .
func (s *setService) GetTfrTrnLog(ctx context.Context, in *pb.GetTfrTrnLogReq) (*pb.GetTfrTrnLogResp, error) {
	trfTrnLogsEnty := cleartxnM.TfrTrnLog{}
	resp := new(pb.GetTfrTrnLogResp)
	trfTrnLog, err := trfTrnLogsEnty.GetByKeyRsp(common.DB, in.KeyRsp)
	if err != nil {
		return nil, err
	}
	if trfTrnLog == nil {
		resp.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "输入的参数错误",
		}
		return resp, nil
	}
	copier.Copy(resp, trfTrnLog)
	return resp, nil
}

func (s *setService) DownloadTfrTrnLogs(ctx context.Context, in *pb.DownloadTfrTrnLogsReq) (*pb.DownloadTfrTrnLogsResp, error) {
	if in.Name == "" {
		return nil, ErrDownloadFileNameEmpty
	}

	var cond cleartxnM.TfrTrnLog

	cond.MchntCd = in.MchntCd
	cond.PriAcctNo = in.PriAcctNO
	cond.KeyRsp = in.KeyRsp
	cond.PriAcctNo = in.PriAcctNO
	cond.CardClass = in.CardClass
	cond.RoutIndustryInsIdCd = in.RoutIndustryInsIdCd
	cond.FwdInsIdCd = in.FwdInsIdCd
	cond.IssInsIdCd = in.IssInsIdCd
	cond.RespCd = in.RespCd
	cond.TermId = in.TermId
	cond.ProdCd = in.ProdCd
	cond.BizCd = in.BizCd
	cond.MaTransCd = in.MaTransCd
	cond.MsgResvFld2 = in.MsgResvFld2

	var accountRegion = ""
	if in.BeginAt != "" && in.EndAt != "" {
		accountRegion = fmt.Sprintf("'%s' < TRANS_DT AND TRANS_DT < '%s'", in.BeginAt, in.EndAt)
	} else if in.BeginAt != "" {
		accountRegion = fmt.Sprintf("'%s' < TRANS_DT", in.BeginAt)
	} else if in.EndAt != "" {
		accountRegion = fmt.Sprintf("TRANS_DT < '%s'", in.EndAt)
	}

	trfTrnLogsEnty := cleartxnM.TfrTrnLog{}
	results, err := trfTrnLogsEnty.Get(common.DB, &cond, accountRegion)
	if err != nil {
		return nil, err
	}
	uid, err := DownloadTfrTrnLogs(results)
	if err != nil {
		return nil, err
	}

	err = Compress(uid, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.DownloadTfrTrnLogsResp{Code: true}, nil
}

func (s *setService) ListGroups(ctx context.Context, in *pb.ListGroupsRequest) (*pb.ListInstitutionsReply, error) {
	reply := new(pb.ListInstitutionsReply)
	db := common.DB
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 10
	}

	groups, count, err := insmodel.ListGroups(db, in.Page, in.Size)
	if err != nil {
		return nil, err
	}

	list := make([]string, len(groups))
	for i := range groups {
		list[i] = groups[i].InsGroup
	}
	ins, err := insmodel.FindInstitutionInfosByIdList(db, list)
	if err != nil {
		return nil, err
	}

	pbIns := make([]*pb.InstitutionField, len(ins))
	for i := range ins {
		pbIns[i] = &pb.InstitutionField{
			InsIdCd:         ins[i].InsIDCd,
			InsCompanyCd:    ins[i].InsCompanyCd,
			InsType:         ins[i].InsType,
			InsName:         ins[i].InsName,
			InsProvCd:       ins[i].InsProvCd,
			InsCityCd:       ins[i].InsCityCd,
			InsRegionCd:     ins[i].InsRegionCd,
			InsSta:          ins[i].InsSta,
			InsStlmTp:       ins[i].InsStlmTp,
			InsAloStlmCycle: ins[i].InsAloStlmCycle,
			InsAloStlmMd:    ins[i].InsAloStlmMd,
			InsStlmCNm:      ins[i].InsStlmCNm,
			InsStlmCAcct:    ins[i].InsStlmCAcct,
			InsStlmCBkNo:    ins[i].InsStlmCBkNo,
			InsStlmCBkNm:    ins[i].InsStlmCBkNm,
			InsStlmDNm:      ins[i].InsStlmDNm,
			InsStlmDAcct:    ins[i].InsStlmDAcct,
			InsStlmDBkNo:    ins[i].InsStlmDBkNo,
			InsStlmDBkNm:    ins[i].InsStlmDBkNm,
			MsgResvFld1:     ins[i].MsgResvFld1,
			MsgResvFld2:     ins[i].MsgResvFld2,
			MsgResvFld3:     ins[i].MsgResvFld3,
			MsgResvFld4:     ins[i].MsgResvFld4,
			MsgResvFld5:     ins[i].MsgResvFld5,
			MsgResvFld6:     ins[i].MsgResvFld6,
			MsgResvFld7:     ins[i].MsgResvFld7,
			MsgResvFld8:     ins[i].MsgResvFld8,
			MsgResvFld9:     ins[i].MsgResvFld9,
			MsgResvFld10:    ins[i].MsgResvFld10,
			RecOprId:        ins[i].RecOprID,
			RecCrtTs:        ins[i].CreatedAt.Unix(),
			RecUpdTs:        ins[i].UpdatedAt.Unix(),
		}
	}

	reply.Institutions = pbIns
	reply.Count = count
	reply.Size = in.Size
	reply.Page = in.Page
	return reply, nil
}
