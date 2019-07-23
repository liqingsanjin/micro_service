package merchantservice

import (
	"context"
	"net/http"
	"time"
	"userService/pkg/common"
	"userService/pkg/pb"
	"userService/pkg/util"

	merchantmodel "userService/pkg/model/merchant"
)

type merchantService struct{}

func (m *merchantService) GetMerchantById(ctx context.Context, in *pb.GetMerchantByIdRequest) (*pb.GetMerchantByIdReply, error) {
	reply := new(pb.GetMerchantByIdReply)
	if in.Id == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "id不能为空",
		}
		return reply, nil
	}
	edit := true
	if in.Type == "main" {
		edit = false
	}
	db := common.DB

	if edit {
		info, err := merchantmodel.FindMerchantInfoById(db, in.Id)
		if err != nil {
			return nil, err
		}
		if info != nil {
			item := new(pb.MerchantField)
			item.MchtCd = info.MchtCd
			item.Sn = info.Sn
			item.AipBranCd = info.AipBranCd
			item.GroupCd = info.GroupCd
			item.OriChnl = info.OriChnl
			item.OriChnlDesc = info.OriChnlDesc
			item.BankBelongCd = info.BankBelongCd
			item.DvpBy = info.DvpBy
			item.MccCd18 = info.MccCd18
			item.ApplDate = info.ApplDate
			item.UpBcCd = info.UpBcCd
			item.UpAcCd = info.UpAcCd
			item.UpMccCd = info.UpMccCd
			item.Name = info.Name
			item.NameBusi = info.NameBusi
			item.BusiLiceNo = info.BusiLiceNo
			item.BusiRang = info.BusiRang
			item.BusiMain = info.BusiMain
			item.Certif = info.Certif
			item.CertifType = info.CertifType
			item.CertifNo = info.CertifNo
			item.CityCd = info.CityCd
			item.AreaCd = info.AreaCd
			item.RegAddr = info.RegAddr
			item.ContactName = info.ContactName
			item.ContactPhoneNo = info.ContactPhoneNo
			item.IsGroup = info.IsGroup
			item.MoneyToGroup = info.MoneyToGroup
			item.StlmWay = info.StlmWay
			item.StlmWayDesc = info.StlmWayDesc
			item.StlmInsCircle = info.StlmInsCircle
			item.Status = info.Status
			item.UcBcCd32 = info.UcBcCd32
			item.K2WorkflowId = info.K2WorkflowId
			item.SystemFlag = info.SystemFlag
			item.ApprovalUsername = info.ApprovalUsername
			item.FinalApprovalUsername = info.FinalApprovalUsername
			item.IsUpStandard = info.IsUpStandard
			item.BillingType = info.BillingType
			item.BillingLevel = info.BillingLevel
			item.Slogan = info.Slogan
			item.Ext1 = info.Ext1
			item.Ext2 = info.Ext2
			item.Ext3 = info.Ext3
			item.Ext4 = info.Ext4
			item.AreaStandard = info.AreaStandard
			item.MchtCdAreaCd = info.MchtCdAreaCd
			item.UcBcCdArea = info.UcBcCdArea
			item.RecOprId = info.RecOprId
			item.RecUpdOpr = info.RecUpdOpr
			item.OperIn = info.OperIn
			item.OemOrgCode = info.OemOrgCode
			item.IsEleInvoice = info.IsEleInvoice
			item.DutyParagraph = info.DutyParagraph
			item.TaxMachineBrand = info.TaxMachineBrand
			item.Ext5 = info.Ext5
			item.Ext6 = info.Ext6
			item.Ext7 = info.Ext7
			item.Ext8 = info.Ext8
			item.Ext9 = info.Ext9
			item.BusiLiceSt = info.BusiLiceSt
			item.BusiLiceDt = info.BusiLiceDt
			item.CertifSt = info.CertifSt
			item.CertifDt = info.CertifDt

			if !info.CreatedAt.IsZero() {
				item.CreatedAt = info.CreatedAt.Format(util.TimePattern)
			}
			if !info.UpdatedAt.IsZero() {
				item.UpdatedAt = info.UpdatedAt.Format(util.TimePattern)
			}
			if info.RecApllyTs != nil {
				item.RecApllyTs = info.RecApllyTs.Format(util.TimePattern)
			}
			if info.ApprDate != nil {
				item.ApprDate = info.ApprDate.Format(util.TimePattern)
			}
			if info.DeleteDate != nil {
				item.DeleteDate = info.DeleteDate.Format(util.TimePattern)
			}
			reply.Item = item
		}
	} else {
		info, err := merchantmodel.FindMerchantInfoMainById(db, in.Id)
		if err != nil {
			return nil, err
		}
		if info != nil {
			item := new(pb.MerchantField)
			item.MchtCd = info.MchtCd
			item.Sn = info.Sn
			item.AipBranCd = info.AipBranCd
			item.GroupCd = info.GroupCd
			item.OriChnl = info.OriChnl
			item.OriChnlDesc = info.OriChnlDesc
			item.BankBelongCd = info.BankBelongCd
			item.DvpBy = info.DvpBy
			item.MccCd18 = info.MccCd18
			item.ApplDate = info.ApplDate
			item.UpBcCd = info.UpBcCd
			item.UpAcCd = info.UpAcCd
			item.UpMccCd = info.UpMccCd
			item.Name = info.Name
			item.NameBusi = info.NameBusi
			item.BusiLiceNo = info.BusiLiceNo
			item.BusiRang = info.BusiRang
			item.BusiMain = info.BusiMain
			item.Certif = info.Certif
			item.CertifType = info.CertifType
			item.CertifNo = info.CertifNo
			item.CityCd = info.CityCd
			item.AreaCd = info.AreaCd
			item.RegAddr = info.RegAddr
			item.ContactName = info.ContactName
			item.ContactPhoneNo = info.ContactPhoneNo
			item.IsGroup = info.IsGroup
			item.MoneyToGroup = info.MoneyToGroup
			item.StlmWay = info.StlmWay
			item.StlmWayDesc = info.StlmWayDesc
			item.StlmInsCircle = info.StlmInsCircle
			item.Status = info.Status
			item.UcBcCd32 = info.UcBcCd32
			item.K2WorkflowId = info.K2WorkflowId
			item.SystemFlag = info.SystemFlag
			item.ApprovalUsername = info.ApprovalUsername
			item.FinalApprovalUsername = info.FinalApprovalUsername
			item.IsUpStandard = info.IsUpStandard
			item.BillingType = info.BillingType
			item.BillingLevel = info.BillingLevel
			item.Slogan = info.Slogan
			item.Ext1 = info.Ext1
			item.Ext2 = info.Ext2
			item.Ext3 = info.Ext3
			item.Ext4 = info.Ext4
			item.AreaStandard = info.AreaStandard
			item.MchtCdAreaCd = info.MchtCdAreaCd
			item.UcBcCdArea = info.UcBcCdArea
			item.RecOprId = info.RecOprId
			item.RecUpdOpr = info.RecUpdOpr
			item.OperIn = info.OperIn
			item.OemOrgCode = info.OemOrgCode
			item.IsEleInvoice = info.IsEleInvoice
			item.DutyParagraph = info.DutyParagraph
			item.TaxMachineBrand = info.TaxMachineBrand
			item.Ext5 = info.Ext5
			item.Ext6 = info.Ext6
			item.Ext7 = info.Ext7
			item.Ext8 = info.Ext8
			item.Ext9 = info.Ext9
			item.BusiLiceSt = info.BusiLiceSt
			item.BusiLiceDt = info.BusiLiceDt
			item.CertifSt = info.CertifSt
			item.CertifDt = info.CertifDt

			if !info.CreatedAt.IsZero() {
				item.CreatedAt = info.CreatedAt.Format(util.TimePattern)
			}
			if !info.UpdatedAt.IsZero() {
				item.UpdatedAt = info.UpdatedAt.Format(util.TimePattern)
			}
			if info.RecApllyTs != nil {
				item.RecApllyTs = info.RecApllyTs.Format(util.TimePattern)
			}
			if info.ApprDate != nil {
				item.ApprDate = info.ApprDate.Format(util.TimePattern)
			}
			if info.DeleteDate != nil {
				item.DeleteDate = info.DeleteDate.Format(util.TimePattern)
			}
			reply.Item = item
		}
	}

	return reply, nil
}

func (m *merchantService) GetMerchantPicture(ctx context.Context, in *pb.GetMerchantPictureRequest) (*pb.GetMerchantPictureReply, error) {
	reply := new(pb.GetMerchantPictureReply)
	if in.Item == nil || in.Item.MchtCd == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "mchtCd不能为空",
		}
		return reply, nil
	}

	edit := true
	if in.Type == "main" {
		edit = false
	}
	db := common.DB

	if edit {
		query := new(merchantmodel.Picture)
		{
			query.MchtCd = in.Item.MchtCd
		}
		items, err := merchantmodel.QueryPicture(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantPictureField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantPictureField{
				FileId:     items[i].FileId,
				MchtCd:     items[i].MchtCd,
				DocType:    items[i].DocType,
				FileType:   items[i].FileType,
				FileName:   items[i].FileName,
				PIndex:     items[i].PIndex,
				PCode:      items[i].PCode,
				Url:        items[i].Url,
				SystemFlag: items[i].SystemFlag,
				Status:     items[i].Status,
				RecOprId:   items[i].RecOprId,
				RecUpdOpr:  items[i].RecUpdOpr,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems

	} else {
		query := new(merchantmodel.PictureMain)
		{
			query.MchtCd = in.Item.MchtCd
		}
		items, err := merchantmodel.QueryPicTureMain(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantPictureField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantPictureField{
				FileId:     items[i].FileId,
				MchtCd:     items[i].MchtCd,
				DocType:    items[i].DocType,
				FileType:   items[i].FileType,
				FileName:   items[i].FileName,
				PIndex:     items[i].PIndex,
				PCode:      items[i].PCode,
				Url:        items[i].Url,
				SystemFlag: items[i].SystemFlag,
				Status:     items[i].Status,
				RecOprId:   items[i].RecOprId,
				RecUpdOpr:  items[i].RecUpdOpr,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems

	}
	return reply, nil
}

func (m *merchantService) GetMerchantBusiness(ctx context.Context, in *pb.GetMerchantBusinessRequest) (*pb.GetMerchantBusinessReply, error) {
	reply := new(pb.GetMerchantBusinessReply)
	if in.Item == nil || in.Item.MchtCd == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "mchtCd不能为空",
		}
		return reply, nil
	}

	edit := true
	if in.Type == "main" {
		edit = false
	}
	db := common.DB

	if edit {
		query := new(merchantmodel.Business)
		{
			query.MchtCd = in.Item.MchtCd
		}
		items, err := merchantmodel.QueryBusiness(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantBusinessField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantBusinessField{
				MchtCd:                 items[i].MchtCd,
				ProdCd:                 items[i].ProdCd,
				ProdCdText:             items[i].ProdCdText,
				FeeMoneyCd:             items[i].FeeMoneyCd,
				FeeModeType:            items[i].FeeModeType,
				FeeSettlementType:      items[i].FeeSettlementType,
				FeeHoliday:             items[i].FeeHoliday,
				ServiceFeeType:         items[i].ServiceFeeType,
				ServiceFeeStaticAmount: items[i].ServiceFeeStaticAmount,
				ServiceFeeLevelCount:   items[i].ServiceFeeLevelCount,
				ServiceFeeMode:         items[i].ServiceFeeMode,
				ServiceFeeUnit:         items[i].ServiceFeeUnit,
				ServiceFeeTerm:         items[i].ServiceFeeTerm,
				ServiceFeeSumto:        items[i].ServiceFeeSumto,
				ServiceFeeCircle:       items[i].ServiceFeeCircle,
				ServiceFeeOthers:       items[i].ServiceFeeOthers,
				ServiceFeeStart:        items[i].ServiceFeeStart,
				ServiceFeeClct:         items[i].ServiceFeeClct,
				ServiceFeeClctOthers:   items[i].ServiceFeeClctOthers,
				SystemFlag:             items[i].SystemFlag,
				Ext1:                   items[i].Ext1,
				Ext2:                   items[i].Ext2,
				Ext3:                   items[i].Ext3,
				Ext4:                   items[i].Ext4,
				ServiceFeeYesNo:        items[i].ServiceFeeYesNo,
				RecOprId:               items[i].RecOprId,
				RecUpdOpr:              items[i].RecUpdOpr,
				OperIn:                 items[i].OperIn,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems
	} else {
		query := new(merchantmodel.BusinessMain)
		{
			query.MchtCd = in.Item.MchtCd
		}
		items, err := merchantmodel.QueryBusinessMain(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantBusinessField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantBusinessField{
				MchtCd:                 items[i].MchtCd,
				ProdCd:                 items[i].ProdCd,
				ProdCdText:             items[i].ProdCdText,
				FeeMoneyCd:             items[i].FeeMoneyCd,
				FeeModeType:            items[i].FeeModeType,
				FeeSettlementType:      items[i].FeeSettlementType,
				FeeHoliday:             items[i].FeeHoliday,
				ServiceFeeType:         items[i].ServiceFeeType,
				ServiceFeeStaticAmount: items[i].ServiceFeeStaticAmount,
				ServiceFeeLevelCount:   items[i].ServiceFeeLevelCount,
				ServiceFeeMode:         items[i].ServiceFeeMode,
				ServiceFeeUnit:         items[i].ServiceFeeUnit,
				ServiceFeeTerm:         items[i].ServiceFeeTerm,
				ServiceFeeSumto:        items[i].ServiceFeeSumto,
				ServiceFeeCircle:       items[i].ServiceFeeCircle,
				ServiceFeeOthers:       items[i].ServiceFeeOthers,
				ServiceFeeStart:        items[i].ServiceFeeStart,
				ServiceFeeClct:         items[i].ServiceFeeClct,
				ServiceFeeClctOthers:   items[i].ServiceFeeClctOthers,
				SystemFlag:             items[i].SystemFlag,
				Ext1:                   items[i].Ext1,
				Ext2:                   items[i].Ext2,
				Ext3:                   items[i].Ext3,
				Ext4:                   items[i].Ext4,
				ServiceFeeYesNo:        items[i].ServiceFeeYesNo,
				RecOprId:               items[i].RecOprId,
				RecUpdOpr:              items[i].RecUpdOpr,
				OperIn:                 items[i].OperIn,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems
	}

	return reply, nil
}

func (m *merchantService) GetMerchantBizFee(ctx context.Context, in *pb.GetMerchantBizFeeRequest) (*pb.GetMerchantBizFeeReply, error) {
	reply := new(pb.GetMerchantBizFeeReply)
	if in.Item == nil || in.Item.MchtCd == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "mchtCd不能为空",
		}
		return reply, nil
	}

	edit := true
	if in.Type == "main" {
		edit = false
	}
	db := common.DB
	if edit {
		query := new(merchantmodel.BizFee)
		{
			query.MchtCd = in.Item.MchtCd
		}
		items, err := merchantmodel.QueryBizFee(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantBizFeeField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantBizFeeField{
				MchtCd:          items[i].MchtCd,
				ProdCd:          items[i].ProdCd,
				BizCd:           items[i].BizCd,
				SubBizCd:        items[i].SubBizCd,
				MchtFeeMd:       items[i].MchtFeeMd,
				MchtFeePercent:  items[i].MchtFeePercent,
				MchtFeePctMin:   items[i].MchtFeePctMin,
				MchtFeePctMax:   items[i].MchtFeePctMax,
				MchtFeeSingle:   items[i].MchtFeeSingle,
				MchtAFeeSame:    items[i].MchtAFeeSame,
				MchtAFeeMd:      items[i].MchtAFeeMd,
				MchtAFeePercent: items[i].MchtAFeePercent,
				MchtAFeePctMin:  items[i].MchtAFeePctMin,
				MchtAFeePctMax:  items[i].MchtAFeePctMax,
				MchtAFeeSingle:  items[i].MchtAFeeSingle,
				OperIn:          items[i].OperIn,
				RecOprId:        items[i].RecOprId,
				RecUpdOpr:       items[i].RecUpdOpr,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems
	} else {
		query := new(merchantmodel.BizFeeMain)
		{
			query.MchtCd = in.Item.MchtCd
		}
		items, err := merchantmodel.QueryBizFeeMain(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantBizFeeField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantBizFeeField{
				MchtCd:          items[i].MchtCd,
				ProdCd:          items[i].ProdCd,
				BizCd:           items[i].BizCd,
				SubBizCd:        items[i].SubBizCd,
				MchtFeeMd:       items[i].MchtFeeMd,
				MchtFeePercent:  items[i].MchtFeePercent,
				MchtFeePctMin:   items[i].MchtFeePctMin,
				MchtFeePctMax:   items[i].MchtFeePctMax,
				MchtFeeSingle:   items[i].MchtFeeSingle,
				MchtAFeeSame:    items[i].MchtAFeeSame,
				MchtAFeeMd:      items[i].MchtAFeeMd,
				MchtAFeePercent: items[i].MchtAFeePercent,
				MchtAFeePctMin:  items[i].MchtAFeePctMin,
				MchtAFeePctMax:  items[i].MchtAFeePctMax,
				MchtAFeeSingle:  items[i].MchtAFeeSingle,
				OperIn:          items[i].OperIn,
				RecOprId:        items[i].RecOprId,
				RecUpdOpr:       items[i].RecUpdOpr,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems
	}
	return reply, nil
}

func (m *merchantService) GetMerchantBizDeal(ctx context.Context, in *pb.GetMerchantBizDealRequest) (*pb.GetMerchantBizDealReply, error) {
	reply := new(pb.GetMerchantBizDealReply)
	if in.Item == nil || in.Item.MchtCd == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "mchtCd不能为空",
		}
		return reply, nil
	}

	edit := true
	if in.Type == "main" {
		edit = false
	}
	db := common.DB

	if edit {
		query := new(merchantmodel.BizDeal)
		{
			query.MchtCd = in.Item.MchtCd
		}
		items, err := merchantmodel.QueryBizDeal(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantBizDealField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantBizDealField{
				MchtCd:    items[i].MchtCd,
				ProdCd:    items[i].ProdCd,
				BizCd:     items[i].BizCd,
				TransCd:   items[i].TransCd,
				OperIn:    items[i].OperIn,
				RecOprId:  items[i].RecOprId,
				RecUpdOpr: items[i].RecUpdOpr,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems

	} else {
		query := new(merchantmodel.BizDealMain)
		{
			query.MchtCd = in.Item.MchtCd
		}
		items, err := merchantmodel.QueryBizDealMain(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantBizDealField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantBizDealField{
				MchtCd:    items[i].MchtCd,
				ProdCd:    items[i].ProdCd,
				BizCd:     items[i].BizCd,
				TransCd:   items[i].TransCd,
				OperIn:    items[i].OperIn,
				RecOprId:  items[i].RecOprId,
				RecUpdOpr: items[i].RecUpdOpr,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems
	}

	return reply, nil
}

func (m *merchantService) GetMerchantBankAccount(ctx context.Context, in *pb.GetMerchantBankAccountRequest) (*pb.GetMerchantBankAccountReply, error) {
	reply := new(pb.GetMerchantBankAccountReply)
	if in.Item == nil || in.Item.OwnerCd == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "ownerCd不能为空",
		}
		return reply, nil
	}

	edit := true
	if in.Type == "main" {
		edit = false
	}
	db := common.DB

	if edit {
		query := new(merchantmodel.BankAccount)
		{
			query.OwnerCd = in.Item.OwnerCd
		}
		items, err := merchantmodel.QueryBankAccount(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantBankAccountField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantBankAccountField{
				OwnerCd:      items[i].OwnerCd,
				AccountType:  items[i].AccountType,
				Name:         items[i].Name,
				Account:      items[i].Account,
				UcBcCd:       items[i].UcBcCd,
				Province:     items[i].Province,
				City:         items[i].City,
				BankCode:     items[i].BankCode,
				BankName:     items[i].BankName,
				OperIn:       items[i].OperIn,
				RecOprId:     items[i].RecOprId,
				RecUpdOpr:    items[i].RecUpdOpr,
				MsgResvFld1:  items[i].MsgResvFld1,
				MsgResvFld2:  items[i].MsgResvFld2,
				MsgResvFld3:  items[i].MsgResvFld3,
				MsgResvFld4:  items[i].MsgResvFld4,
				MsgResvFld5:  items[i].MsgResvFld5,
				MsgResvFld6:  items[i].MsgResvFld6,
				MsgResvFld7:  items[i].MsgResvFld7,
				MsgResvFld8:  items[i].MsgResvFld8,
				MsgResvFld9:  items[i].MsgResvFld9,
				MsgResvFld10: items[i].MsgResvFld10,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems

	} else {
		query := new(merchantmodel.BankAccountMain)
		{
			query.OwnerCd = in.Item.OwnerCd
		}
		items, err := merchantmodel.QueryBankAccountMain(db, query)
		if err != nil {
			return nil, err
		}
		pbItems := make([]*pb.MerchantBankAccountField, len(items))

		for i := range items {
			pbItems[i] = &pb.MerchantBankAccountField{
				OwnerCd:      items[i].OwnerCd,
				AccountType:  items[i].AccountType,
				Name:         items[i].Name,
				Account:      items[i].Account,
				UcBcCd:       items[i].UcBcCd,
				Province:     items[i].Province,
				City:         items[i].City,
				BankCode:     items[i].BankCode,
				BankName:     items[i].BankName,
				OperIn:       items[i].OperIn,
				RecOprId:     items[i].RecOprId,
				RecUpdOpr:    items[i].RecUpdOpr,
				MsgResvFld1:  items[i].MsgResvFld1,
				MsgResvFld2:  items[i].MsgResvFld2,
				MsgResvFld3:  items[i].MsgResvFld3,
				MsgResvFld4:  items[i].MsgResvFld4,
				MsgResvFld5:  items[i].MsgResvFld5,
				MsgResvFld6:  items[i].MsgResvFld6,
				MsgResvFld7:  items[i].MsgResvFld7,
				MsgResvFld8:  items[i].MsgResvFld8,
				MsgResvFld9:  items[i].MsgResvFld9,
				MsgResvFld10: items[i].MsgResvFld10,
			}
			if !items[i].CreatedAt.IsZero() {
				pbItems[i].CreatedAt = items[i].CreatedAt.Format(util.TimePattern)
			}
			if !items[i].UpdatedAt.IsZero() {
				pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Items = pbItems
	}

	return reply, nil
}

func (m *merchantService) SaveMerchant(ctx context.Context, in *pb.SaveMerchantRequest) (*pb.SaveMerchantReply, error) {
	var reply pb.SaveMerchantReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "保存信息为空",
		}
		return &reply, nil
	}

	if in.Item.MchtCd == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "id不能为空",
		}
		return &reply, nil
	}
	db := common.DB

	mtch := new(merchantmodel.MerchantInfo)
	{
		mtch.MchtCd = in.Item.MchtCd
		mtch.Sn = in.Item.Sn
		mtch.AipBranCd = in.Item.AipBranCd
		mtch.GroupCd = in.Item.GroupCd
		mtch.OriChnl = in.Item.OriChnl
		mtch.OriChnlDesc = in.Item.OriChnlDesc
		mtch.BankBelongCd = in.Item.BankBelongCd
		mtch.DvpBy = in.Item.DvpBy
		mtch.MccCd18 = in.Item.MccCd18
		mtch.ApplDate = in.Item.ApplDate
		mtch.UpBcCd = in.Item.UpBcCd
		mtch.UpAcCd = in.Item.UpAcCd
		mtch.UpMccCd = in.Item.UpMccCd
		mtch.Name = in.Item.Name
		mtch.NameBusi = in.Item.NameBusi
		mtch.BusiLiceNo = in.Item.BusiLiceNo
		mtch.BusiRang = in.Item.BusiRang
		mtch.BusiMain = in.Item.BusiMain
		mtch.Certif = in.Item.Certif
		mtch.CertifType = in.Item.CertifType
		mtch.CertifNo = in.Item.CertifNo
		mtch.CityCd = in.Item.CityCd
		mtch.AreaCd = in.Item.AreaCd
		mtch.RegAddr = in.Item.RegAddr
		mtch.ContactName = in.Item.ContactName
		mtch.ContactPhoneNo = in.Item.ContactPhoneNo
		mtch.IsGroup = in.Item.IsGroup
		mtch.MoneyToGroup = in.Item.MoneyToGroup
		mtch.StlmWay = in.Item.StlmWay
		mtch.StlmWayDesc = in.Item.StlmWayDesc
		mtch.StlmInsCircle = in.Item.StlmInsCircle
		mtch.Status = in.Item.Status
		mtch.UcBcCd32 = in.Item.UcBcCd32
		mtch.K2WorkflowId = in.Item.K2WorkflowId
		mtch.SystemFlag = in.Item.SystemFlag
		mtch.ApprovalUsername = in.Item.ApprovalUsername
		mtch.FinalApprovalUsername = in.Item.FinalApprovalUsername
		mtch.IsUpStandard = in.Item.IsUpStandard
		mtch.BillingType = in.Item.BillingType
		mtch.BillingLevel = in.Item.BillingLevel
		mtch.Slogan = in.Item.Slogan
		mtch.Ext1 = in.Item.Ext1
		mtch.Ext2 = in.Item.Ext2
		mtch.Ext3 = in.Item.Ext3
		mtch.Ext4 = in.Item.Ext4
		mtch.AreaStandard = in.Item.AreaStandard
		mtch.MchtCdAreaCd = in.Item.MchtCdAreaCd
		mtch.UcBcCdArea = in.Item.UcBcCdArea
		mtch.RecOprId = in.Item.RecOprId
		mtch.RecUpdOpr = in.Item.RecUpdOpr
		mtch.OperIn = in.Item.OperIn
		mtch.IsEleInvoice = in.Item.IsEleInvoice
		mtch.DutyParagraph = in.Item.DutyParagraph
		mtch.TaxMachineBrand = in.Item.TaxMachineBrand
		mtch.Ext5 = in.Item.Ext5
		mtch.Ext6 = in.Item.Ext6
		mtch.Ext7 = in.Item.Ext7
		mtch.Ext8 = in.Item.Ext8
		mtch.Ext9 = in.Item.Ext9
		mtch.BusiLiceSt = in.Item.BusiLiceSt
		mtch.BusiLiceDt = in.Item.BusiLiceDt
		mtch.CertifSt = in.Item.CertifSt
		mtch.CertifDt = in.Item.CertifDt
		mtch.OemOrgCode = in.Item.OemOrgCode
		if in.Item.ApprDate != "" {
			t, _ := time.Parse("2006-01-02 15:04:05", in.Item.ApprDate)
			mtch.ApprDate = &t
		}
		if in.Item.DeleteDate != "" {
			t, _ := time.Parse("2006-01-02 15:04:05", in.Item.DeleteDate)
			mtch.DeleteDate = &t
		}
		if in.Item.RecApllyTs != "" {
			t, _ := time.Parse("2006-01-02 15:04:05", in.Item.RecApllyTs)
			mtch.RecApllyTs = &t
		}
	}
	err := merchantmodel.SaveMerchant(db, mtch)

	return &reply, err
}

func (m *merchantService) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
	if in.Size == 0 {
		in.Size = 10
	}
	if in.Page == 0 {
		in.Page = 1
	}

	db := common.DB

	edit := true
	if in.Type == "main" {
		edit = false
	}

	if edit {
		query := new(merchantmodel.MerchantInfo)
		if in.Item != nil {
			query.MchtCd = in.Item.MchtCd
			query.Sn = in.Item.Sn
			query.AipBranCd = in.Item.AipBranCd
			query.GroupCd = in.Item.GroupCd
			query.OriChnl = in.Item.OriChnl
			query.OriChnlDesc = in.Item.OriChnlDesc
			query.BankBelongCd = in.Item.BankBelongCd
			query.DvpBy = in.Item.DvpBy
			query.MccCd18 = in.Item.MccCd18
			query.ApplDate = in.Item.ApplDate
			query.UpBcCd = in.Item.UpBcCd
			query.UpAcCd = in.Item.UpAcCd
			query.UpMccCd = in.Item.UpMccCd
			query.Name = in.Item.Name
			query.NameBusi = in.Item.NameBusi
			query.BusiLiceNo = in.Item.BusiLiceNo
			query.BusiRang = in.Item.BusiRang
			query.BusiMain = in.Item.BusiMain
			query.Certif = in.Item.Certif
			query.CertifType = in.Item.CertifType
			query.CertifNo = in.Item.CertifNo
			query.CityCd = in.Item.CityCd
			query.AreaCd = in.Item.AreaCd
			query.RegAddr = in.Item.RegAddr
			query.ContactName = in.Item.ContactName
			query.ContactPhoneNo = in.Item.ContactPhoneNo
			query.IsGroup = in.Item.IsGroup
			query.MoneyToGroup = in.Item.MoneyToGroup
			query.StlmWay = in.Item.StlmWay
			query.StlmWayDesc = in.Item.StlmWayDesc
			query.StlmInsCircle = in.Item.StlmInsCircle
			query.Status = in.Item.Status
			query.UcBcCd32 = in.Item.UcBcCd32
			query.K2WorkflowId = in.Item.K2WorkflowId
			query.SystemFlag = in.Item.SystemFlag
			query.ApprovalUsername = in.Item.ApprovalUsername
			query.FinalApprovalUsername = in.Item.FinalApprovalUsername
			query.IsUpStandard = in.Item.IsUpStandard
			query.BillingType = in.Item.BillingType
			query.BillingLevel = in.Item.BillingLevel
			query.Slogan = in.Item.Slogan
			query.Ext1 = in.Item.Ext1
			query.Ext2 = in.Item.Ext2
			query.Ext3 = in.Item.Ext3
			query.Ext4 = in.Item.Ext4
			query.AreaStandard = in.Item.AreaStandard
			query.MchtCdAreaCd = in.Item.MchtCdAreaCd
			query.UcBcCdArea = in.Item.UcBcCdArea
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
			query.OperIn = in.Item.OperIn
			query.OemOrgCode = in.Item.OemOrgCode
			query.IsEleInvoice = in.Item.IsEleInvoice
			query.DutyParagraph = in.Item.DutyParagraph
			query.TaxMachineBrand = in.Item.TaxMachineBrand
			query.Ext5 = in.Item.Ext5
			query.Ext6 = in.Item.Ext6
			query.Ext7 = in.Item.Ext7
			query.Ext8 = in.Item.Ext8
			query.Ext9 = in.Item.Ext9
			query.BusiLiceSt = in.Item.BusiLiceSt
			query.BusiLiceDt = in.Item.BusiLiceDt
			query.CertifSt = in.Item.CertifSt
			query.CertifDt = in.Item.CertifDt
		}

		merchants, count, err := merchantmodel.QueryMerchantInfos(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbMerchants := make([]*pb.MerchantField, len(merchants))
		for i := range merchants {
			pbMerchants[i] = &pb.MerchantField{
				MchtCd:                merchants[i].MchtCd,
				Sn:                    merchants[i].Sn,
				AipBranCd:             merchants[i].AipBranCd,
				GroupCd:               merchants[i].GroupCd,
				OriChnl:               merchants[i].OriChnl,
				OriChnlDesc:           merchants[i].OriChnlDesc,
				BankBelongCd:          merchants[i].BankBelongCd,
				DvpBy:                 merchants[i].DvpBy,
				MccCd18:               merchants[i].MccCd18,
				ApplDate:              merchants[i].ApplDate,
				UpBcCd:                merchants[i].UpBcCd,
				UpAcCd:                merchants[i].UpAcCd,
				UpMccCd:               merchants[i].UpMccCd,
				Name:                  merchants[i].Name,
				NameBusi:              merchants[i].NameBusi,
				BusiLiceNo:            merchants[i].BusiLiceNo,
				BusiRang:              merchants[i].BusiRang,
				BusiMain:              merchants[i].BusiMain,
				Certif:                merchants[i].Certif,
				CertifType:            merchants[i].CertifType,
				CertifNo:              merchants[i].CertifNo,
				CityCd:                merchants[i].CityCd,
				AreaCd:                merchants[i].AreaCd,
				RegAddr:               merchants[i].RegAddr,
				ContactName:           merchants[i].ContactName,
				ContactPhoneNo:        merchants[i].ContactPhoneNo,
				IsGroup:               merchants[i].IsGroup,
				MoneyToGroup:          merchants[i].MoneyToGroup,
				StlmWay:               merchants[i].StlmWay,
				StlmWayDesc:           merchants[i].StlmWayDesc,
				StlmInsCircle:         merchants[i].StlmInsCircle,
				Status:                merchants[i].Status,
				UcBcCd32:              merchants[i].UcBcCd32,
				K2WorkflowId:          merchants[i].K2WorkflowId,
				SystemFlag:            merchants[i].SystemFlag,
				ApprovalUsername:      merchants[i].ApprovalUsername,
				FinalApprovalUsername: merchants[i].FinalApprovalUsername,
				IsUpStandard:          merchants[i].IsUpStandard,
				BillingType:           merchants[i].BillingType,
				BillingLevel:          merchants[i].BillingLevel,
				Slogan:                merchants[i].Slogan,
				Ext1:                  merchants[i].Ext1,
				Ext2:                  merchants[i].Ext2,
				Ext3:                  merchants[i].Ext3,
				Ext4:                  merchants[i].Ext4,
				AreaStandard:          merchants[i].AreaStandard,
				MchtCdAreaCd:          merchants[i].MchtCdAreaCd,
				UcBcCdArea:            merchants[i].UcBcCdArea,
				RecOprId:              merchants[i].RecOprId,
				RecUpdOpr:             merchants[i].RecUpdOpr,
				OperIn:                merchants[i].OperIn,
				OemOrgCode:            merchants[i].OemOrgCode,
				IsEleInvoice:          merchants[i].IsEleInvoice,
				DutyParagraph:         merchants[i].DutyParagraph,
				TaxMachineBrand:       merchants[i].TaxMachineBrand,
				Ext5:                  merchants[i].Ext5,
				Ext6:                  merchants[i].Ext6,
				Ext7:                  merchants[i].Ext7,
				Ext8:                  merchants[i].Ext8,
				Ext9:                  merchants[i].Ext9,
				BusiLiceSt:            merchants[i].BusiLiceSt,
				BusiLiceDt:            merchants[i].BusiLiceDt,
				CertifSt:              merchants[i].CertifSt,
				CertifDt:              merchants[i].CertifDt,
			}
			if merchants[i] != nil {
				pbMerchants[i].ApprDate = merchants[i].ApprDate.Format("2006-01-02 15:04:05")
			}
			if merchants[i] != nil {
				pbMerchants[i].DeleteDate = merchants[i].DeleteDate.Format("2006-01-02 15:04:05")
			}
			if !merchants[i].CreatedAt.IsZero() {
				pbMerchants[i].CreatedAt = merchants[i].CreatedAt.Format("2006-01-02 15:04:05")
			}
			if merchants[i].UpdatedAt.IsZero() {
				pbMerchants[i].UpdatedAt = merchants[i].UpdatedAt.Format("2006-01-02 15:04:05")
			}
			if merchants[i] != nil {
				pbMerchants[i].RecApllyTs = merchants[i].RecApllyTs.Format("2006-01-02 15:04:05")
			}
		}

		return &pb.ListMerchantReply{
			Items: pbMerchants,
			Count: count,
			Page:  in.Page,
			Size:  in.Size,
		}, nil
	} else {
		query := new(merchantmodel.MerchantInfoMain)
		if in.Item != nil {
			query.MchtCd = in.Item.MchtCd
			query.Sn = in.Item.Sn
			query.AipBranCd = in.Item.AipBranCd
			query.GroupCd = in.Item.GroupCd
			query.OriChnl = in.Item.OriChnl
			query.OriChnlDesc = in.Item.OriChnlDesc
			query.BankBelongCd = in.Item.BankBelongCd
			query.DvpBy = in.Item.DvpBy
			query.MccCd18 = in.Item.MccCd18
			query.ApplDate = in.Item.ApplDate
			query.UpBcCd = in.Item.UpBcCd
			query.UpAcCd = in.Item.UpAcCd
			query.UpMccCd = in.Item.UpMccCd
			query.Name = in.Item.Name
			query.NameBusi = in.Item.NameBusi
			query.BusiLiceNo = in.Item.BusiLiceNo
			query.BusiRang = in.Item.BusiRang
			query.BusiMain = in.Item.BusiMain
			query.Certif = in.Item.Certif
			query.CertifType = in.Item.CertifType
			query.CertifNo = in.Item.CertifNo
			query.CityCd = in.Item.CityCd
			query.AreaCd = in.Item.AreaCd
			query.RegAddr = in.Item.RegAddr
			query.ContactName = in.Item.ContactName
			query.ContactPhoneNo = in.Item.ContactPhoneNo
			query.IsGroup = in.Item.IsGroup
			query.MoneyToGroup = in.Item.MoneyToGroup
			query.StlmWay = in.Item.StlmWay
			query.StlmWayDesc = in.Item.StlmWayDesc
			query.StlmInsCircle = in.Item.StlmInsCircle
			query.Status = in.Item.Status
			query.UcBcCd32 = in.Item.UcBcCd32
			query.K2WorkflowId = in.Item.K2WorkflowId
			query.SystemFlag = in.Item.SystemFlag
			query.ApprovalUsername = in.Item.ApprovalUsername
			query.FinalApprovalUsername = in.Item.FinalApprovalUsername
			query.IsUpStandard = in.Item.IsUpStandard
			query.BillingType = in.Item.BillingType
			query.BillingLevel = in.Item.BillingLevel
			query.Slogan = in.Item.Slogan
			query.Ext1 = in.Item.Ext1
			query.Ext2 = in.Item.Ext2
			query.Ext3 = in.Item.Ext3
			query.Ext4 = in.Item.Ext4
			query.AreaStandard = in.Item.AreaStandard
			query.MchtCdAreaCd = in.Item.MchtCdAreaCd
			query.UcBcCdArea = in.Item.UcBcCdArea
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
			query.OperIn = in.Item.OperIn
			query.OemOrgCode = in.Item.OemOrgCode
			query.IsEleInvoice = in.Item.IsEleInvoice
			query.DutyParagraph = in.Item.DutyParagraph
			query.TaxMachineBrand = in.Item.TaxMachineBrand
			query.Ext5 = in.Item.Ext5
			query.Ext6 = in.Item.Ext6
			query.Ext7 = in.Item.Ext7
			query.Ext8 = in.Item.Ext8
			query.Ext9 = in.Item.Ext9
			query.BusiLiceSt = in.Item.BusiLiceSt
			query.BusiLiceDt = in.Item.BusiLiceDt
			query.CertifSt = in.Item.CertifSt
			query.CertifDt = in.Item.CertifDt
		}

		merchants, count, err := merchantmodel.QueryMerchantInfosMain(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbMerchants := make([]*pb.MerchantField, len(merchants))
		for i := range merchants {
			pbMerchants[i] = &pb.MerchantField{
				MchtCd:                merchants[i].MchtCd,
				Sn:                    merchants[i].Sn,
				AipBranCd:             merchants[i].AipBranCd,
				GroupCd:               merchants[i].GroupCd,
				OriChnl:               merchants[i].OriChnl,
				OriChnlDesc:           merchants[i].OriChnlDesc,
				BankBelongCd:          merchants[i].BankBelongCd,
				DvpBy:                 merchants[i].DvpBy,
				MccCd18:               merchants[i].MccCd18,
				ApplDate:              merchants[i].ApplDate,
				UpBcCd:                merchants[i].UpBcCd,
				UpAcCd:                merchants[i].UpAcCd,
				UpMccCd:               merchants[i].UpMccCd,
				Name:                  merchants[i].Name,
				NameBusi:              merchants[i].NameBusi,
				BusiLiceNo:            merchants[i].BusiLiceNo,
				BusiRang:              merchants[i].BusiRang,
				BusiMain:              merchants[i].BusiMain,
				Certif:                merchants[i].Certif,
				CertifType:            merchants[i].CertifType,
				CertifNo:              merchants[i].CertifNo,
				CityCd:                merchants[i].CityCd,
				AreaCd:                merchants[i].AreaCd,
				RegAddr:               merchants[i].RegAddr,
				ContactName:           merchants[i].ContactName,
				ContactPhoneNo:        merchants[i].ContactPhoneNo,
				IsGroup:               merchants[i].IsGroup,
				MoneyToGroup:          merchants[i].MoneyToGroup,
				StlmWay:               merchants[i].StlmWay,
				StlmWayDesc:           merchants[i].StlmWayDesc,
				StlmInsCircle:         merchants[i].StlmInsCircle,
				Status:                merchants[i].Status,
				UcBcCd32:              merchants[i].UcBcCd32,
				K2WorkflowId:          merchants[i].K2WorkflowId,
				SystemFlag:            merchants[i].SystemFlag,
				ApprovalUsername:      merchants[i].ApprovalUsername,
				FinalApprovalUsername: merchants[i].FinalApprovalUsername,
				IsUpStandard:          merchants[i].IsUpStandard,
				BillingType:           merchants[i].BillingType,
				BillingLevel:          merchants[i].BillingLevel,
				Slogan:                merchants[i].Slogan,
				Ext1:                  merchants[i].Ext1,
				Ext2:                  merchants[i].Ext2,
				Ext3:                  merchants[i].Ext3,
				Ext4:                  merchants[i].Ext4,
				AreaStandard:          merchants[i].AreaStandard,
				MchtCdAreaCd:          merchants[i].MchtCdAreaCd,
				UcBcCdArea:            merchants[i].UcBcCdArea,
				RecOprId:              merchants[i].RecOprId,
				RecUpdOpr:             merchants[i].RecUpdOpr,
				OperIn:                merchants[i].OperIn,
				OemOrgCode:            merchants[i].OemOrgCode,
				IsEleInvoice:          merchants[i].IsEleInvoice,
				DutyParagraph:         merchants[i].DutyParagraph,
				TaxMachineBrand:       merchants[i].TaxMachineBrand,
				Ext5:                  merchants[i].Ext5,
				Ext6:                  merchants[i].Ext6,
				Ext7:                  merchants[i].Ext7,
				Ext8:                  merchants[i].Ext8,
				Ext9:                  merchants[i].Ext9,
				BusiLiceSt:            merchants[i].BusiLiceSt,
				BusiLiceDt:            merchants[i].BusiLiceDt,
				CertifSt:              merchants[i].CertifSt,
				CertifDt:              merchants[i].CertifDt,
			}
			if merchants[i].ApprDate != nil {
				pbMerchants[i].ApprDate = merchants[i].ApprDate.Format("2006-01-02 15:04:05")
			}
			if merchants[i].DeleteDate != nil {
				pbMerchants[i].DeleteDate = merchants[i].DeleteDate.Format("2006-01-02 15:04:05")
			}
			if !merchants[i].CreatedAt.IsZero() {
				pbMerchants[i].CreatedAt = merchants[i].CreatedAt.Format("2006-01-02 15:04:05")
			}
			if !merchants[i].UpdatedAt.IsZero() {
				pbMerchants[i].UpdatedAt = merchants[i].UpdatedAt.Format("2006-01-02 15:04:05")
			}
			if merchants[i].RecApllyTs != nil {
				pbMerchants[i].RecApllyTs = merchants[i].RecApllyTs.Format("2006-01-02 15:04:05")
			}
		}

		return &pb.ListMerchantReply{
			Items: pbMerchants,
			Count: count,
			Page:  in.Page,
			Size:  in.Size,
		}, nil

	}
}

func (m *merchantService) ListGroupMerchant(ctx context.Context, in *pb.ListGroupMerchantRequest) (*pb.ListGroupMerchantReply, error) {
	if in.Size == 0 {
		in.Size = 10
	}
	if in.Page == 0 {
		in.Page = 1
	}
	db := common.DB

	edit := true
	if in.Type == "main" {
		edit = false
	}

	if edit {
		query := new(merchantmodel.Group)
		if in.Item != nil {
			query.JtMchtCd = in.Item.JtMchtCd
			query.JtMchtNm = in.Item.JtMchtNm
			query.JtArea = in.Item.JtArea
			query.MchtStlmCNm = in.Item.MchtStlmCNm
			query.MchtStlmCAcct = in.Item.MchtStlmCAcct
			query.ChtStlmInsIdCd = in.Item.ChtStlmInsIdCd
			query.MchtStlmInsNm = in.Item.MchtStlmInsNm
			query.MchtPaySysAcct = in.Item.MchtPaySysAcct
			query.ProvCd = in.Item.ProvCd
			query.CityCd = in.Item.CityCd
			query.AipBranCd = in.Item.AipBranCd
			query.SystemFlag = in.Item.SystemFlag
			query.Status = in.Item.Status
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
			query.JtAddr = in.Item.JtAddr
		}

		groups, count, err := merchantmodel.QueryMerchantGroups(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbGroups := make([]*pb.MerchantGroupField, len(groups))

		for i := range groups {
			pbGroups[i] = &pb.MerchantGroupField{
				JtMchtCd:       groups[i].JtMchtCd,
				JtMchtNm:       groups[i].JtMchtNm,
				JtArea:         groups[i].JtArea,
				MchtStlmCNm:    groups[i].MchtStlmCNm,
				MchtStlmCAcct:  groups[i].MchtStlmCAcct,
				ChtStlmInsIdCd: groups[i].ChtStlmInsIdCd,
				MchtStlmInsNm:  groups[i].MchtStlmInsNm,
				MchtPaySysAcct: groups[i].MchtPaySysAcct,
				ProvCd:         groups[i].ProvCd,
				CityCd:         groups[i].CityCd,
				AipBranCd:      groups[i].AipBranCd,
				SystemFlag:     groups[i].SystemFlag,
				Status:         groups[i].Status,
				RecOprId:       groups[i].RecOprId,
				RecUpdOpr:      groups[i].RecUpdOpr,
				JtAddr:         groups[i].JtAddr,
			}
			if !groups[i].CreatedAt.IsZero() {
				pbGroups[i].CreatedAt = groups[i].CreatedAt.Format("2006-01-02 15:04:05")
			}
			if !groups[i].UpdatedAt.IsZero() {
				pbGroups[i].UpdatedAt = groups[i].UpdatedAt.Format("2006-01-02 15:04:05")
			}
		}
		return &pb.ListGroupMerchantReply{
			Items: pbGroups,
			Count: count,
			Page:  in.Page,
			Size:  in.Size,
		}, nil
	} else {
		query := new(merchantmodel.GroupMain)
		if in.Item != nil {
			query.JtMchtCd = in.Item.JtMchtCd
			query.JtMchtNm = in.Item.JtMchtNm
			query.JtArea = in.Item.JtArea
			query.MchtStlmCNm = in.Item.MchtStlmCNm
			query.MchtStlmCAcct = in.Item.MchtStlmCAcct
			query.ChtStlmInsIdCd = in.Item.ChtStlmInsIdCd
			query.MchtStlmInsNm = in.Item.MchtStlmInsNm
			query.MchtPaySysAcct = in.Item.MchtPaySysAcct
			query.ProvCd = in.Item.ProvCd
			query.CityCd = in.Item.CityCd
			query.AipBranCd = in.Item.AipBranCd
			query.SystemFlag = in.Item.SystemFlag
			query.Status = in.Item.Status
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
			query.JtAddr = in.Item.JtAddr
		}

		groups, count, err := merchantmodel.QueryMerchantGroupsMain(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbGroups := make([]*pb.MerchantGroupField, len(groups))

		for i := range groups {
			pbGroups[i] = &pb.MerchantGroupField{
				JtMchtCd:       groups[i].JtMchtCd,
				JtMchtNm:       groups[i].JtMchtNm,
				JtArea:         groups[i].JtArea,
				MchtStlmCNm:    groups[i].MchtStlmCNm,
				MchtStlmCAcct:  groups[i].MchtStlmCAcct,
				ChtStlmInsIdCd: groups[i].ChtStlmInsIdCd,
				MchtStlmInsNm:  groups[i].MchtStlmInsNm,
				MchtPaySysAcct: groups[i].MchtPaySysAcct,
				ProvCd:         groups[i].ProvCd,
				CityCd:         groups[i].CityCd,
				AipBranCd:      groups[i].AipBranCd,
				SystemFlag:     groups[i].SystemFlag,
				Status:         groups[i].Status,
				RecOprId:       groups[i].RecOprId,
				RecUpdOpr:      groups[i].RecUpdOpr,
				JtAddr:         groups[i].JtAddr,
			}
			if !groups[i].CreatedAt.IsZero() {
				pbGroups[i].CreatedAt = groups[i].CreatedAt.Format("2006-01-02 15:04:05")
			}
			if !groups[i].UpdatedAt.IsZero() {
				pbGroups[i].UpdatedAt = groups[i].UpdatedAt.Format("2006-01-02 15:04:05")
			}
		}
		return &pb.ListGroupMerchantReply{
			Items: pbGroups,
			Count: count,
			Page:  in.Page,
			Size:  in.Size,
		}, nil
	}
}

func (m *merchantService) SaveMerchantBankAccount(ctx context.Context, in *pb.SaveMerchantBankAccountRequest) (*pb.SaveMerchantBankAccountReply, error) {
	var reply pb.SaveMerchantBankAccountReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "保存信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	data := new(merchantmodel.BankAccount)
	{
		data.OwnerCd = in.Item.OwnerCd
		data.AccountType = in.Item.AccountType
		data.Name = in.Item.Name
		data.Account = in.Item.Account
		data.UcBcCd = in.Item.UcBcCd
		data.Province = in.Item.Province
		data.City = in.Item.City
		data.BankCode = in.Item.BankCode
		data.BankName = in.Item.BankName
		data.OperIn = in.Item.OperIn
		data.RecOprId = in.Item.RecOprId
		data.RecUpdOpr = in.Item.RecUpdOpr
		data.MsgResvFld1 = in.Item.MsgResvFld1
		data.MsgResvFld2 = in.Item.MsgResvFld2
		data.MsgResvFld3 = in.Item.MsgResvFld3
		data.MsgResvFld4 = in.Item.MsgResvFld4
		data.MsgResvFld5 = in.Item.MsgResvFld5
		data.MsgResvFld6 = in.Item.MsgResvFld6
		data.MsgResvFld7 = in.Item.MsgResvFld7
		data.MsgResvFld8 = in.Item.MsgResvFld8
		data.MsgResvFld9 = in.Item.MsgResvFld9
		data.MsgResvFld10 = in.Item.MsgResvFld10
	}
	err := merchantmodel.SaveBankAccount(db, data)
	return &reply, err
}

func (m *merchantService) SaveMerchantBizDeal(ctx context.Context, in *pb.SaveMerchantBizDealRequest) (*pb.SaveMerchantBizDealReply, error) {
	var reply pb.SaveMerchantBizDealReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "保存信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	data := new(merchantmodel.BizDeal)
	{
		data.MchtCd = in.Item.MchtCd
		data.ProdCd = in.Item.ProdCd
		data.BizCd = in.Item.BizCd
		data.TransCd = in.Item.TransCd
		data.OperIn = in.Item.OperIn
		data.RecOprId = in.Item.RecOprId
		data.RecUpdOpr = in.Item.RecUpdOpr
	}
	err := merchantmodel.SaveBizDeal(db, data)
	return &reply, err
}

func (m *merchantService) SaveMerchantBizFee(ctx context.Context, in *pb.SaveMerchantBizFeeRequest) (*pb.SaveMerchantBizFeeReply, error) {
	var reply pb.SaveMerchantBizFeeReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "保存信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	data := new(merchantmodel.BizFee)
	{
		data.MchtCd = in.Item.MchtCd
		data.ProdCd = in.Item.ProdCd
		data.BizCd = in.Item.BizCd
		data.SubBizCd = in.Item.SubBizCd
		data.MchtFeeMd = in.Item.MchtFeeMd
		data.MchtFeePercent = in.Item.MchtFeePercent
		data.MchtFeePctMin = in.Item.MchtFeePctMin
		data.MchtFeePctMax = in.Item.MchtFeePctMax
		data.MchtFeeSingle = in.Item.MchtFeeSingle
		data.MchtAFeeSame = in.Item.MchtAFeeSame
		data.MchtAFeeMd = in.Item.MchtAFeeMd
		data.MchtAFeePercent = in.Item.MchtAFeePercent
		data.MchtAFeePctMin = in.Item.MchtAFeePctMin
		data.MchtAFeePctMax = in.Item.MchtAFeePctMax
		data.MchtAFeeSingle = in.Item.MchtAFeeSingle
		data.OperIn = in.Item.OperIn
		data.RecOprId = in.Item.RecOprId
		data.RecUpdOpr = in.Item.RecUpdOpr
	}
	err := merchantmodel.SaveBizFee(db, data)
	return &reply, err
}

func (m *merchantService) SaveMerchantBusiness(ctx context.Context, in *pb.SaveMerchantBusinessRequest) (*pb.SaveMerchantBusinessReply, error) {
	var reply pb.SaveMerchantBusinessReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "保存信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	data := new(merchantmodel.Business)
	{
		data.MchtCd = in.Item.MchtCd
		data.ProdCd = in.Item.ProdCd
		data.ProdCdText = in.Item.ProdCdText
		data.FeeMoneyCd = in.Item.FeeMoneyCd
		data.FeeModeType = in.Item.FeeModeType
		data.FeeSettlementType = in.Item.FeeSettlementType
		data.FeeHoliday = in.Item.FeeHoliday
		data.ServiceFeeType = in.Item.ServiceFeeType
		data.ServiceFeeStaticAmount = in.Item.ServiceFeeStaticAmount
		data.ServiceFeeLevelCount = in.Item.ServiceFeeLevelCount
		data.ServiceFeeMode = in.Item.ServiceFeeMode
		data.ServiceFeeUnit = in.Item.ServiceFeeUnit
		data.ServiceFeeTerm = in.Item.ServiceFeeTerm
		data.ServiceFeeSumto = in.Item.ServiceFeeSumto
		data.ServiceFeeCircle = in.Item.ServiceFeeCircle
		data.ServiceFeeOthers = in.Item.ServiceFeeOthers
		data.ServiceFeeStart = in.Item.ServiceFeeStart
		data.ServiceFeeClct = in.Item.ServiceFeeClct
		data.ServiceFeeClctOthers = in.Item.ServiceFeeClctOthers
		data.SystemFlag = in.Item.SystemFlag
		data.Ext1 = in.Item.Ext1
		data.Ext2 = in.Item.Ext2
		data.Ext3 = in.Item.Ext3
		data.Ext4 = in.Item.Ext4
		data.ServiceFeeYesNo = in.Item.ServiceFeeYesNo
		data.RecOprId = in.Item.RecOprId
		data.RecUpdOpr = in.Item.RecUpdOpr
		data.OperIn = in.Item.OperIn
	}
	err := merchantmodel.SaveBusiness(db, data)
	return &reply, err
}

func (m *merchantService) SaveMerchantPicture(ctx context.Context, in *pb.SaveMerchantPictureRequest) (*pb.SaveMerchantPictureReply, error) {
	var reply pb.SaveMerchantPictureReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "保存信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	data := new(merchantmodel.Picture)
	{
		data.FileId = in.Item.FileId
		data.MchtCd = in.Item.MchtCd
		data.DocType = in.Item.DocType
		data.FileType = in.Item.FileType
		data.FileName = in.Item.FileName
		data.PIndex = in.Item.PIndex
		data.PCode = in.Item.PCode
		data.Url = in.Item.Url
		data.SystemFlag = in.Item.SystemFlag
		data.Status = in.Item.Status
		data.RecOprId = in.Item.RecOprId
		data.RecUpdOpr = in.Item.RecUpdOpr
	}
	err := merchantmodel.SavePicture(db, data)
	return &reply, err
}
