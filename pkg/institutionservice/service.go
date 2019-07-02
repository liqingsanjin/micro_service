package institutionservice

import (
	"fmt"
	"userService/pkg/camunda"
	camundapb "userService/pkg/camunda/pb"
	"userService/pkg/common"
	camundamodel "userService/pkg/model/camunda"
	cleartxnM "userService/pkg/model/cleartxn"
	insmodel "userService/pkg/model/institution"
	"userService/pkg/pb"

	"github.com/sirupsen/logrus"

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

func (s *setService) SaveInstitution(ctx context.Context, in *pb.SaveInstitutionRequest) (*pb.SaveInstitutionReply, error) {
	var reply pb.SaveInstitutionReply
	if in.Field == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "机构信息为空",
		}
		return &reply, nil
	}

	if in.Field.InsIdCd == "" {
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
		ins.InsIdCd = in.Field.InsIdCd
		ins.InsCompanyCd = in.Field.InsCompanyCd
		ins.InsType = in.Field.InsType
		ins.InsName = in.Field.InsName
		ins.InsProvCd = in.Field.InsProvCd
		ins.InsCityCd = in.Field.InsCityCd
		ins.InsRegionCd = in.Field.InsRegionCd
		ins.InsSta = in.Field.InsSta
		ins.InsStlmTp = in.Field.InsStlmTp
		ins.InsAloStlmCycle = in.Field.InsAloStlmCycle
		ins.InsAloStlmMd = in.Field.InsAloStlmMd
		ins.InsStlmCNm = in.Field.InsStlmCNm
		ins.InsStlmCAcct = in.Field.InsStlmCAcct
		ins.InsStlmCBkNo = in.Field.InsStlmCBkNo
		ins.InsStlmCBkNm = in.Field.InsStlmCBkNm
		ins.InsStlmDNm = in.Field.InsStlmDNm
		ins.InsStlmDAcct = in.Field.InsStlmDAcct
		ins.InsStlmDBkNo = in.Field.InsStlmDBkNo
		ins.InsStlmDBkNm = in.Field.InsStlmDBkNm
		ins.MsgResvFld1 = in.Field.MsgResvFld1
		ins.MsgResvFld2 = in.Field.MsgResvFld2
		ins.MsgResvFld3 = in.Field.MsgResvFld3
		ins.MsgResvFld4 = in.Field.MsgResvFld4
		ins.MsgResvFld5 = in.Field.MsgResvFld5
		ins.MsgResvFld6 = in.Field.MsgResvFld6
		ins.MsgResvFld7 = in.Field.MsgResvFld7
		ins.MsgResvFld8 = in.Field.MsgResvFld8
		ins.MsgResvFld9 = in.Field.MsgResvFld9
		ins.MsgResvFld10 = in.Field.MsgResvFld10
		ins.RecOprId = in.Field.RecOprId
	}
	err := insmodel.SaveInstitution(db, ins)
	if err != nil {
		return nil, err
	}

	// 查询是否存在工作流
	camundaService := camunda.Get()
	listProcessInstanceRes, err := camundaService.ProcessInstance.List(ctx, &camundapb.ProcessInstanceListReq{
		BusinessKey: "ins_add_1:" + ins.InsIdCd,
	})
	if err != nil {
		return nil, err
	}
	if listProcessInstanceRes.Err != nil {
		logrus.Error(listProcessInstanceRes.Err)
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "WorkFlowError",
			Description: "查询工作流错误: " + listProcessInstanceRes.Err.Message,
		}
		return &reply, nil
	}
	var instanceId string
	if len(listProcessInstanceRes.Items) != 0 {
		// 已开启工作流
		instanceId = listProcessInstanceRes.Items[0].Id
	} else {
		// 需要开启工作流
		processes, err := camundamodel.QueryProcessDefinition(db, &camundamodel.ProcessDefinition{
			OperateName: "ins_add_1",
		})
		if err != nil {
			return nil, err
		}
		if len(processes) == 0 {
			reply.Err = &pb.Error{
				Code:        http.StatusBadRequest,
				Message:     "WorkFlowError",
				Description: "没有工作流信息",
			}
			return &reply, nil
		}

		// 开启工作流
		res, err := camundaService.ProcessDefinition.Start(ctx, &camundapb.StartProcessDefinitionReq{
			Id: processes[0].Id,
			Body: &camundapb.StartProcessDefinitionReqBody{
				BusinessKey: "ins_add_1:" + ins.InsIdCd,
			},
		})
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		if res.Err != nil {
			logrus.Error(res.Err)
			reply.Err = &pb.Error{
				Code:        http.StatusBadRequest,
				Message:     "WorkFlowError",
				Description: "开启工作流错误: " + res.Err.Message,
			}
			return &reply, nil
		}
		logrus.Debugln("工作流返回", res)
		instanceId = res.Item.Id
	}
	// 获取编辑任务
	taskListResp, err := camundaService.Task.GetList(ctx, &camundapb.GetListTaskReq{
		ProcessInstanceId: instanceId,
	})
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	if taskListResp.Err != nil {
		logrus.Error(taskListResp.Err)
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "WorkFlowError",
			Description: "查询任务错误: " + taskListResp.Err.Message,
		}
		return &reply, nil
	}
	if len(taskListResp.Tasks) == 0 {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "WorkFlowError",
			Description: "没有编辑任务",
		}
		return &reply, nil
	}
	taskCompleteResp, err := camundaService.Task.Complete(ctx, &camundapb.CompleteTaskReq{
		Id: taskListResp.Tasks[0].Id,
		Body: &camundapb.CompleteTaskReqBody{
			Variables: map[string]*camundapb.Variable{},
		},
	})
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	if taskCompleteResp.Err != nil {
		logrus.Error(taskCompleteResp.Err)
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "WorkFlowError",
			Description: "完成任务错误: " + taskCompleteResp.Err.Message,
		}
		return &reply, nil
	}
	logrus.Infoln("完成编辑任务", "ins_add1:"+ins.InsIdCd)

	db.Commit()

	return &reply, nil
}

func (s *setService) SaveInstitutionFee(ctx context.Context, in *pb.SaveInstitutionFeeRequest) (*pb.SaveInstitutionFeeReply, error) {
	var reply pb.SaveInstitutionFeeReply
	if in.Field == nil {
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
		ins.InsIdCd = in.Field.InsIdCd
		ins.ProdCd = in.Field.ProdCd
		ins.BizCd = in.Field.BizCd
		ins.SubBizCd = in.Field.SubBizCd
		ins.InsFeeBizCd = in.Field.InsFeeBizCd
		ins.InsFeeCd = in.Field.InsFeeCd
		ins.InsFeeTp = in.Field.InsFeeTp
		ins.InsFeeParam = in.Field.InsFeeParam
		ins.InsFeePercent = in.Field.InsFeePercent
		ins.InsFeePct = in.Field.InsFeePct
		ins.InsFeePctMin = in.Field.InsFeePctMin
		ins.InsFeePctMax = in.Field.InsFeePctMax
		ins.InsAFeeSame = in.Field.InsAFeeSame
		ins.InsAFeeParam = in.Field.InsAFeeParam
		ins.InsAFeePercent = in.Field.InsAFeePercent
		ins.InsAFeePct = in.Field.InsAFeePct
		ins.InsAFeePctMin = in.Field.InsAFeePctMin
		ins.InsAFeePctMax = in.Field.InsAFeePctMax
		ins.MsgResvFld1 = in.Field.MsgResvFld1
		ins.MsgResvFld2 = in.Field.MsgResvFld2
		ins.MsgResvFld3 = in.Field.MsgResvFld3
		ins.MsgResvFld4 = in.Field.MsgResvFld4
		ins.MsgResvFld5 = in.Field.MsgResvFld5
		ins.MsgResvFld6 = in.Field.MsgResvFld6
		ins.MsgResvFld7 = in.Field.MsgResvFld7
		ins.MsgResvFld8 = in.Field.MsgResvFld8
		ins.MsgResvFld9 = in.Field.MsgResvFld9
		ins.MsgResvFld10 = in.Field.MsgResvFld10
		ins.RecOprId = in.Field.RecOprId
		ins.RecUpdOpr = in.Field.RecUpdOpr
	}
	err := insmodel.SaveInstitutionFee(db, ins)
	if err != nil {
		return nil, err
	}

	return &reply, nil
}

func (s *setService) SaveInstitutionControl(ctx context.Context, in *pb.SaveInstitutionControlRequest) (*pb.SaveInstitutionControlReply, error) {
	var reply pb.SaveInstitutionControlReply
	if in.Field == nil {
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
		ins.InsIdCd = in.Field.InsIdCd
		ins.InsCompanyCd = in.Field.InsCompanyCd
		ins.ProdCd = in.Field.ProdCd
		ins.BizCd = in.Field.BizCd
		ins.CtrlSta = in.Field.CtrlSta
		ins.InsBegTm = in.Field.InsBegTm
		ins.InsEndTm = in.Field.InsEndTm
		ins.MsgResvFld1 = in.Field.MsgResvFld1
		ins.MsgResvFld2 = in.Field.MsgResvFld2
		ins.MsgResvFld3 = in.Field.MsgResvFld3
		ins.MsgResvFld4 = in.Field.MsgResvFld4
		ins.MsgResvFld5 = in.Field.MsgResvFld5
		ins.MsgResvFld6 = in.Field.MsgResvFld6
		ins.MsgResvFld7 = in.Field.MsgResvFld7
		ins.MsgResvFld8 = in.Field.MsgResvFld8
		ins.MsgResvFld9 = in.Field.MsgResvFld9
		ins.MsgResvFld10 = in.Field.MsgResvFld10
		ins.RecOprId = in.Field.RecOprId
		ins.RecUpdOpr = in.Field.RecUpdOpr
	}
	err := insmodel.SaveInstitutionControl(db, ins)
	if err != nil {
		return nil, err
	}

	return &reply, nil
}

func (s *setService) SaveInstitutionCash(ctx context.Context, in *pb.SaveInstitutionCashRequest) (*pb.SaveInstitutionCashReply, error) {
	var reply pb.SaveInstitutionCashReply
	if in.Field == nil {
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
		ins.InsIdCd = in.Field.InsIdCd
		ins.ProdCd = in.Field.ProdCd
		ins.InsDefaultFlag = in.Field.InsDefaultFlag
		ins.InsDefaultCash = in.Field.InsDefaultCash
		ins.InsCurrentCash = in.Field.InsCurrentCash
		ins.MsgResvFld1 = in.Field.MsgResvFld1
		ins.MsgResvFld2 = in.Field.MsgResvFld2
		ins.MsgResvFld3 = in.Field.MsgResvFld3
		ins.MsgResvFld4 = in.Field.MsgResvFld4
		ins.MsgResvFld5 = in.Field.MsgResvFld5
		ins.MsgResvFld6 = in.Field.MsgResvFld6
		ins.MsgResvFld7 = in.Field.MsgResvFld7
		ins.MsgResvFld8 = in.Field.MsgResvFld8
		ins.MsgResvFld9 = in.Field.MsgResvFld9
		ins.MsgResvFld10 = in.Field.MsgResvFld10
		ins.RecOprid = in.Field.RecOprid
		ins.RecUpdopr = in.Field.RecUpdopr
	}
	err := insmodel.SaveInstitutionCash(db, ins)
	if err != nil {
		return nil, err
	}

	return &reply, nil
}
