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
			InsIdCd:         ins[i].InsIdCd,
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
			RecOprId:        ins[i].RecOprId,
		}
		if !ins[i].CreatedAt.IsZero() {
			pbIns[i].CreatedAt = ins[i].CreatedAt.Format("2006-01-02 15:04:05")
		}

		if !ins[i].UpdatedAt.IsZero() {
			pbIns[i].UpdatedAt = ins[i].UpdatedAt.Format("2006-01-02 15:04:05")
		}
	}

	reply.Institutions = pbIns
	reply.Count = count
	reply.Size = in.Size
	reply.Page = in.Page
	return reply, nil
}

func (s *setService) ListInstitutions(ctx context.Context, in *pb.ListInstitutionsRequest) (*pb.ListInstitutionsReply, error) {
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 10
	}

	db := common.DB

	query := new(insmodel.InstitutionInfo)
	if in.Institution != nil {
		query.InsIdCd = in.Institution.InsIdCd
		query.InsCompanyCd = in.Institution.InsCompanyCd
		query.InsType = in.Institution.InsType
		query.InsName = in.Institution.InsName
		query.InsProvCd = in.Institution.InsProvCd
		query.InsCityCd = in.Institution.InsCityCd
		query.InsRegionCd = in.Institution.InsRegionCd
		query.InsSta = in.Institution.InsSta
		query.InsStlmTp = in.Institution.InsStlmTp
		query.InsAloStlmCycle = in.Institution.InsAloStlmCycle
		query.InsAloStlmMd = in.Institution.InsAloStlmMd
		query.InsStlmCNm = in.Institution.InsStlmCNm
		query.InsStlmCAcct = in.Institution.InsStlmCAcct
		query.InsStlmCBkNo = in.Institution.InsStlmCBkNo
		query.InsStlmCBkNm = in.Institution.InsStlmCBkNm
		query.InsStlmDNm = in.Institution.InsStlmDNm
		query.InsStlmDAcct = in.Institution.InsStlmDAcct
		query.InsStlmDBkNo = in.Institution.InsStlmDBkNo
		query.InsStlmDBkNm = in.Institution.InsStlmDBkNm
		query.MsgResvFld1 = in.Institution.MsgResvFld1
		query.MsgResvFld2 = in.Institution.MsgResvFld2
		query.MsgResvFld3 = in.Institution.MsgResvFld3
		query.MsgResvFld4 = in.Institution.MsgResvFld4
		query.MsgResvFld5 = in.Institution.MsgResvFld5
		query.MsgResvFld6 = in.Institution.MsgResvFld6
		query.MsgResvFld7 = in.Institution.MsgResvFld7
		query.MsgResvFld8 = in.Institution.MsgResvFld8
		query.MsgResvFld9 = in.Institution.MsgResvFld9
		query.MsgResvFld10 = in.Institution.MsgResvFld10
		query.RecOprId = in.Institution.RecOprId
	}
	ins, count, err := insmodel.QueryInstitutionInfo(db, query, in.Page, in.Size)
	if err != nil {
		return nil, err
	}

	pbIns := make([]*pb.InstitutionField, len(ins))

	for i := range ins {
		pbIns[i] = &pb.InstitutionField{
			InsIdCd:         ins[i].InsIdCd,
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
			RecOprId:        ins[i].RecOprId,
		}
		if !ins[i].CreatedAt.IsZero() {
			pbIns[i].CreatedAt = ins[i].CreatedAt.Format("2006-01-02 15:04:05")
		}

		if !ins[i].UpdatedAt.IsZero() {
			pbIns[i].UpdatedAt = ins[i].UpdatedAt.Format("2006-01-02 15:04:05")
		}

	}

	return &pb.ListInstitutionsReply{
		Institutions: pbIns,
		Count:        count,
		Page:         in.Page,
		Size:         in.Size,
	}, nil
}

func (s *setService) AddInstitution(ctx context.Context, in *pb.AddInstitutionRequest) (*pb.AddInstitutionReply, error) {
	var reply pb.AddInstitutionReply
	if in.Institution == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "机构信息为空",
		}
		return &reply, nil
	}
	db := common.DB.Begin()
	defer db.Rollback()

	ins, err := insmodel.FindInstitutionInfoById(db, in.Institution.InsIdCd)
	if err != nil {
		return nil, err
	}
	if ins != nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "机构代码已存在",
		}
		return &reply, nil
	}

	ins = new(insmodel.InstitutionInfo)
	{
		ins.InsIdCd = in.Institution.InsIdCd
		ins.InsCompanyCd = in.Institution.InsCompanyCd
		ins.InsType = in.Institution.InsType
		ins.InsName = in.Institution.InsName
		ins.InsProvCd = in.Institution.InsProvCd
		ins.InsCityCd = in.Institution.InsCityCd
		ins.InsRegionCd = in.Institution.InsRegionCd
		ins.InsSta = in.Institution.InsSta
		ins.InsStlmTp = in.Institution.InsStlmTp
		ins.InsAloStlmCycle = in.Institution.InsAloStlmCycle
		ins.InsAloStlmMd = in.Institution.InsAloStlmMd
		ins.InsStlmCNm = in.Institution.InsStlmCNm
		ins.InsStlmCAcct = in.Institution.InsStlmCAcct
		ins.InsStlmCBkNo = in.Institution.InsStlmCBkNo
		ins.InsStlmCBkNm = in.Institution.InsStlmCBkNm
		ins.InsStlmDNm = in.Institution.InsStlmDNm
		ins.InsStlmDAcct = in.Institution.InsStlmDAcct
		ins.InsStlmDBkNo = in.Institution.InsStlmDBkNo
		ins.InsStlmDBkNm = in.Institution.InsStlmDBkNm
		ins.MsgResvFld1 = in.Institution.MsgResvFld1
		ins.MsgResvFld2 = in.Institution.MsgResvFld2
		ins.MsgResvFld3 = in.Institution.MsgResvFld3
		ins.MsgResvFld4 = in.Institution.MsgResvFld4
		ins.MsgResvFld5 = in.Institution.MsgResvFld5
		ins.MsgResvFld6 = in.Institution.MsgResvFld6
		ins.MsgResvFld7 = in.Institution.MsgResvFld7
		ins.MsgResvFld8 = in.Institution.MsgResvFld8
		ins.MsgResvFld9 = in.Institution.MsgResvFld9
		ins.MsgResvFld10 = in.Institution.MsgResvFld10
		ins.RecOprId = in.Institution.RecOprId
	}
	err = insmodel.SaveInstitution(db, ins)
	if err != nil {
		return nil, err
	}

	//todo 写入工作流

	db.Commit()

	return &reply, nil
}

func (s *setService) AddInstitutionFee(ctx context.Context, in *pb.AddInstitutionFeeRequest) (*pb.AddInstitutionFeeReply, error) {
	var reply pb.AddInstitutionFeeReply
	if in.InstitutionFee == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "机构信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	ins, err := insmodel.FindInstitutionFeeByPrimaryKey(
		db,
		in.InstitutionFee.InsIdCd,
		in.InstitutionFee.ProdCd,
		in.InstitutionFee.BizCd,
		in.InstitutionFee.SubBizCd,
		in.InstitutionFee.InsFeeCd,
	)
	if err != nil {
		return nil, err
	}
	if ins != nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "支付方式已存在",
		}
		return &reply, nil
	}

	ins = new(insmodel.Fee)
	{
		ins.InsIdCd = in.InstitutionFee.InsIdCd
		ins.ProdCd = in.InstitutionFee.ProdCd
		ins.BizCd = in.InstitutionFee.BizCd
		ins.SubBizCd = in.InstitutionFee.SubBizCd
		ins.InsFeeBizCd = in.InstitutionFee.InsFeeBizCd
		ins.InsFeeCd = in.InstitutionFee.InsFeeCd
		ins.InsFeeTp = in.InstitutionFee.InsFeeTp
		ins.InsFeeParam = in.InstitutionFee.InsFeeParam
		ins.InsFeePercent = in.InstitutionFee.InsFeePercent
		ins.InsFeePct = in.InstitutionFee.InsFeePct
		ins.InsFeePctMin = in.InstitutionFee.InsFeePctMin
		ins.InsFeePctMax = in.InstitutionFee.InsFeePctMax
		ins.InsAFeeSame = in.InstitutionFee.InsAFeeSame
		ins.InsAFeeParam = in.InstitutionFee.InsAFeeParam
		ins.InsAFeePercent = in.InstitutionFee.InsAFeePercent
		ins.InsAFeePct = in.InstitutionFee.InsAFeePct
		ins.InsAFeePctMin = in.InstitutionFee.InsAFeePctMin
		ins.InsAFeePctMax = in.InstitutionFee.InsAFeePctMax
		ins.MsgResvFld1 = in.InstitutionFee.MsgResvFld1
		ins.MsgResvFld2 = in.InstitutionFee.MsgResvFld2
		ins.MsgResvFld3 = in.InstitutionFee.MsgResvFld3
		ins.MsgResvFld4 = in.InstitutionFee.MsgResvFld4
		ins.MsgResvFld5 = in.InstitutionFee.MsgResvFld5
		ins.MsgResvFld6 = in.InstitutionFee.MsgResvFld6
		ins.MsgResvFld7 = in.InstitutionFee.MsgResvFld7
		ins.MsgResvFld8 = in.InstitutionFee.MsgResvFld8
		ins.MsgResvFld9 = in.InstitutionFee.MsgResvFld9
		ins.MsgResvFld10 = in.InstitutionFee.MsgResvFld10
		ins.RecOprId = in.InstitutionFee.RecOprId
		ins.RecUpdOpr = in.InstitutionFee.RecUpdOpr
	}
	err = insmodel.SaveInstitutionFee(db, ins)
	if err != nil {
		return nil, err
	}

	return &reply, nil
}
