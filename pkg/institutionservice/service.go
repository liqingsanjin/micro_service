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

	reply.Items = pbIns
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
	edit := true
	if in.Type == "main" {
		edit = false
	}

	db := common.DB

	if edit {
		query := new(insmodel.InstitutionInfo)
		if in.Item != nil {
			query.InsIdCd = in.Item.InsIdCd
			query.InsCompanyCd = in.Item.InsCompanyCd
			query.InsType = in.Item.InsType
			query.InsName = in.Item.InsName
			query.InsProvCd = in.Item.InsProvCd
			query.InsCityCd = in.Item.InsCityCd
			query.InsRegionCd = in.Item.InsRegionCd
			query.InsSta = in.Item.InsSta
			query.InsStlmTp = in.Item.InsStlmTp
			query.InsAloStlmCycle = in.Item.InsAloStlmCycle
			query.InsAloStlmMd = in.Item.InsAloStlmMd
			query.InsStlmCNm = in.Item.InsStlmCNm
			query.InsStlmCAcct = in.Item.InsStlmCAcct
			query.InsStlmCBkNo = in.Item.InsStlmCBkNo
			query.InsStlmCBkNm = in.Item.InsStlmCBkNm
			query.InsStlmDNm = in.Item.InsStlmDNm
			query.InsStlmDAcct = in.Item.InsStlmDAcct
			query.InsStlmDBkNo = in.Item.InsStlmDBkNo
			query.InsStlmDBkNm = in.Item.InsStlmDBkNm
			query.MsgResvFld1 = in.Item.MsgResvFld1
			query.MsgResvFld2 = in.Item.MsgResvFld2
			query.MsgResvFld3 = in.Item.MsgResvFld3
			query.MsgResvFld4 = in.Item.MsgResvFld4
			query.MsgResvFld5 = in.Item.MsgResvFld5
			query.MsgResvFld6 = in.Item.MsgResvFld6
			query.MsgResvFld7 = in.Item.MsgResvFld7
			query.MsgResvFld8 = in.Item.MsgResvFld8
			query.MsgResvFld9 = in.Item.MsgResvFld9
			query.MsgResvFld10 = in.Item.MsgResvFld10
			query.RecOprId = in.Item.RecOprId
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
			Items: pbIns,
			Count: count,
			Page:  in.Page,
			Size:  in.Size,
		}, nil
	} else {
		query := new(insmodel.InstitutionInfoMain)
		if in.Item != nil {
			query.InsIdCd = in.Item.InsIdCd
			query.InsCompanyCd = in.Item.InsCompanyCd
			query.InsType = in.Item.InsType
			query.InsName = in.Item.InsName
			query.InsProvCd = in.Item.InsProvCd
			query.InsCityCd = in.Item.InsCityCd
			query.InsRegionCd = in.Item.InsRegionCd
			query.InsSta = in.Item.InsSta
			query.InsStlmTp = in.Item.InsStlmTp
			query.InsAloStlmCycle = in.Item.InsAloStlmCycle
			query.InsAloStlmMd = in.Item.InsAloStlmMd
			query.InsStlmCNm = in.Item.InsStlmCNm
			query.InsStlmCAcct = in.Item.InsStlmCAcct
			query.InsStlmCBkNo = in.Item.InsStlmCBkNo
			query.InsStlmCBkNm = in.Item.InsStlmCBkNm
			query.InsStlmDNm = in.Item.InsStlmDNm
			query.InsStlmDAcct = in.Item.InsStlmDAcct
			query.InsStlmDBkNo = in.Item.InsStlmDBkNo
			query.InsStlmDBkNm = in.Item.InsStlmDBkNm
			query.MsgResvFld1 = in.Item.MsgResvFld1
			query.MsgResvFld2 = in.Item.MsgResvFld2
			query.MsgResvFld3 = in.Item.MsgResvFld3
			query.MsgResvFld4 = in.Item.MsgResvFld4
			query.MsgResvFld5 = in.Item.MsgResvFld5
			query.MsgResvFld6 = in.Item.MsgResvFld6
			query.MsgResvFld7 = in.Item.MsgResvFld7
			query.MsgResvFld8 = in.Item.MsgResvFld8
			query.MsgResvFld9 = in.Item.MsgResvFld9
			query.MsgResvFld10 = in.Item.MsgResvFld10
			query.RecOprId = in.Item.RecOprId
		}
		ins, count, err := insmodel.QueryInstitutionInfoMain(db, query, in.Page, in.Size)
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
			Items: pbIns,
			Count: count,
			Page:  in.Page,
			Size:  in.Size,
		}, nil
	}
}

func (s *setService) SaveInstitution(ctx context.Context, in *pb.SaveInstitutionRequest) (*pb.SaveInstitutionReply, error) {
	var reply pb.SaveInstitutionReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "机构信息为空",
		}
		return &reply, nil
	}

	if in.Item.InsIdCd == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "id不能为空",
		}
		return &reply, nil
	}
	db := common.DB.Begin()
	defer db.Rollback()

	ins := new(insmodel.InstitutionInfo)
	{
		ins.InsIdCd = in.Item.InsIdCd
		ins.InsCompanyCd = in.Item.InsCompanyCd
		ins.InsType = in.Item.InsType
		ins.InsName = in.Item.InsName
		ins.InsProvCd = in.Item.InsProvCd
		ins.InsCityCd = in.Item.InsCityCd
		ins.InsRegionCd = in.Item.InsRegionCd
		ins.InsSta = in.Item.InsSta
		ins.InsStlmTp = in.Item.InsStlmTp
		ins.InsAloStlmCycle = in.Item.InsAloStlmCycle
		ins.InsAloStlmMd = in.Item.InsAloStlmMd
		ins.InsStlmCNm = in.Item.InsStlmCNm
		ins.InsStlmCAcct = in.Item.InsStlmCAcct
		ins.InsStlmCBkNo = in.Item.InsStlmCBkNo
		ins.InsStlmCBkNm = in.Item.InsStlmCBkNm
		ins.InsStlmDNm = in.Item.InsStlmDNm
		ins.InsStlmDAcct = in.Item.InsStlmDAcct
		ins.InsStlmDBkNo = in.Item.InsStlmDBkNo
		ins.InsStlmDBkNm = in.Item.InsStlmDBkNm
		ins.MsgResvFld1 = in.Item.MsgResvFld1
		ins.MsgResvFld2 = in.Item.MsgResvFld2
		ins.MsgResvFld3 = in.Item.MsgResvFld3
		ins.MsgResvFld4 = in.Item.MsgResvFld4
		ins.MsgResvFld5 = in.Item.MsgResvFld5
		ins.MsgResvFld6 = in.Item.MsgResvFld6
		ins.MsgResvFld7 = in.Item.MsgResvFld7
		ins.MsgResvFld8 = in.Item.MsgResvFld8
		ins.MsgResvFld9 = in.Item.MsgResvFld9
		ins.MsgResvFld10 = in.Item.MsgResvFld10
		ins.RecOprId = in.Item.RecOprId
	}
	err := insmodel.SaveInstitution(db, ins)
	if err != nil {
		return nil, err
	}

	db.Commit()

	return &reply, nil
}

func (s *setService) SaveInstitutionFee(ctx context.Context, in *pb.SaveInstitutionFeeRequest) (*pb.SaveInstitutionFeeReply, error) {
	var reply pb.SaveInstitutionFeeReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "支付方式为空",
		}
		return &reply, nil
	}
	db := common.DB

	ins := new(insmodel.Fee)
	{
		ins.InsIdCd = in.Item.InsIdCd
		ins.ProdCd = in.Item.ProdCd
		ins.BizCd = in.Item.BizCd
		ins.SubBizCd = in.Item.SubBizCd
		ins.InsFeeBizCd = in.Item.InsFeeBizCd
		ins.InsFeeCd = in.Item.InsFeeCd
		ins.InsFeeTp = in.Item.InsFeeTp
		ins.InsFeeParam = in.Item.InsFeeParam
		ins.InsFeePercent = in.Item.InsFeePercent
		ins.InsFeePct = in.Item.InsFeePct
		ins.InsFeePctMin = in.Item.InsFeePctMin
		ins.InsFeePctMax = in.Item.InsFeePctMax
		ins.InsAFeeSame = in.Item.InsAFeeSame
		ins.InsAFeeParam = in.Item.InsAFeeParam
		ins.InsAFeePercent = in.Item.InsAFeePercent
		ins.InsAFeePct = in.Item.InsAFeePct
		ins.InsAFeePctMin = in.Item.InsAFeePctMin
		ins.InsAFeePctMax = in.Item.InsAFeePctMax
		ins.MsgResvFld1 = in.Item.MsgResvFld1
		ins.MsgResvFld2 = in.Item.MsgResvFld2
		ins.MsgResvFld3 = in.Item.MsgResvFld3
		ins.MsgResvFld4 = in.Item.MsgResvFld4
		ins.MsgResvFld5 = in.Item.MsgResvFld5
		ins.MsgResvFld6 = in.Item.MsgResvFld6
		ins.MsgResvFld7 = in.Item.MsgResvFld7
		ins.MsgResvFld8 = in.Item.MsgResvFld8
		ins.MsgResvFld9 = in.Item.MsgResvFld9
		ins.MsgResvFld10 = in.Item.MsgResvFld10
		ins.RecOprId = in.Item.RecOprId
		ins.RecUpdOpr = in.Item.RecUpdOpr
	}
	err := insmodel.SaveInstitutionFee(db, ins)
	if err != nil {
		return nil, err
	}

	return &reply, nil
}

func (s *setService) SaveInstitutionControl(ctx context.Context, in *pb.SaveInstitutionControlRequest) (*pb.SaveInstitutionControlReply, error) {
	var reply pb.SaveInstitutionControlReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "机构权限为空",
		}
		return &reply, nil
	}
	db := common.DB

	ins := new(insmodel.Control)
	{
		ins.InsIdCd = in.Item.InsIdCd
		ins.InsCompanyCd = in.Item.InsCompanyCd
		ins.ProdCd = in.Item.ProdCd
		ins.BizCd = in.Item.BizCd
		ins.CtrlSta = in.Item.CtrlSta
		ins.InsBegTm = in.Item.InsBegTm
		ins.InsEndTm = in.Item.InsEndTm
		ins.MsgResvFld1 = in.Item.MsgResvFld1
		ins.MsgResvFld2 = in.Item.MsgResvFld2
		ins.MsgResvFld3 = in.Item.MsgResvFld3
		ins.MsgResvFld4 = in.Item.MsgResvFld4
		ins.MsgResvFld5 = in.Item.MsgResvFld5
		ins.MsgResvFld6 = in.Item.MsgResvFld6
		ins.MsgResvFld7 = in.Item.MsgResvFld7
		ins.MsgResvFld8 = in.Item.MsgResvFld8
		ins.MsgResvFld9 = in.Item.MsgResvFld9
		ins.MsgResvFld10 = in.Item.MsgResvFld10
		ins.RecOprId = in.Item.RecOprId
		ins.RecUpdOpr = in.Item.RecUpdOpr
	}
	err := insmodel.SaveInstitutionControl(db, ins)
	if err != nil {
		return nil, err
	}

	return &reply, nil
}

func (s *setService) SaveInstitutionCash(ctx context.Context, in *pb.SaveInstitutionCashRequest) (*pb.SaveInstitutionCashReply, error) {
	var reply pb.SaveInstitutionCashReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "头寸信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	ins := new(insmodel.Cash)
	{
		ins.InsIdCd = in.Item.InsIdCd
		ins.ProdCd = in.Item.ProdCd
		ins.InsDefaultFlag = in.Item.InsDefaultFlag
		ins.InsDefaultCash = in.Item.InsDefaultCash
		ins.InsCurrentCash = in.Item.InsCurrentCash
		ins.MsgResvFld1 = in.Item.MsgResvFld1
		ins.MsgResvFld2 = in.Item.MsgResvFld2
		ins.MsgResvFld3 = in.Item.MsgResvFld3
		ins.MsgResvFld4 = in.Item.MsgResvFld4
		ins.MsgResvFld5 = in.Item.MsgResvFld5
		ins.MsgResvFld6 = in.Item.MsgResvFld6
		ins.MsgResvFld7 = in.Item.MsgResvFld7
		ins.MsgResvFld8 = in.Item.MsgResvFld8
		ins.MsgResvFld9 = in.Item.MsgResvFld9
		ins.MsgResvFld10 = in.Item.MsgResvFld10
		ins.RecOprId = in.Item.RecOprId
		ins.RecUpdOpr = in.Item.RecUpdOpr
	}
	err := insmodel.SaveInstitutionCash(db, ins)
	if err != nil {
		return nil, err
	}

	return &reply, nil
}
