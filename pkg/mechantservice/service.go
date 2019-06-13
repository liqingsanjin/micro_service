package merchantservice

import (
	"context"
	"userService/pkg/common"
	"userService/pkg/pb"

	merchantmodel "userService/pkg/model/merchant"
)

type merchantService struct{}

func (m *merchantService) ListMerchant(ctx context.Context, in *pb.ListMerchantRequest) (*pb.ListMerchantReply, error) {
	if in.Size == 0 {
		in.Size = 10
	}
	if in.Page == 0 {
		in.Page = 1
	}

	db := common.DB

	query := new(merchantmodel.MerchantInfo)
	if in.Merchant != nil {
		query.MchtCd = in.Merchant.MchtCd
		query.Sn = in.Merchant.Sn
		query.AipBranCd = in.Merchant.AipBranCd
		query.GroupCd = in.Merchant.GroupCd
		query.OriChnl = in.Merchant.OriChnl
		query.OriChnlDesc = in.Merchant.OriChnlDesc
		query.BankBelongCd = in.Merchant.BankBelongCd
		query.DvpBy = in.Merchant.DvpBy
		query.MccCd18 = in.Merchant.MccCd18
		query.ApplDate = in.Merchant.ApplDate
		query.UpBcCd = in.Merchant.UpBcCd
		query.UpAcCd = in.Merchant.UpAcCd
		query.UpMccCd = in.Merchant.UpMccCd
		query.Name = in.Merchant.Name
		query.NameBusi = in.Merchant.NameBusi
		query.BusiLiceNo = in.Merchant.BusiLiceNo
		query.BusiRang = in.Merchant.BusiRang
		query.BusiMain = in.Merchant.BusiMain
		query.Certif = in.Merchant.Certif
		query.CertifType = in.Merchant.CertifType
		query.CertifNo = in.Merchant.CertifNo
		query.CityCd = in.Merchant.CityCd
		query.AreaCd = in.Merchant.AreaCd
		query.RegAddr = in.Merchant.RegAddr
		query.ContactName = in.Merchant.ContactName
		query.ContactPhoneNo = in.Merchant.ContactPhoneNo
		query.IsGroup = in.Merchant.IsGroup
		query.MoneyToGroup = in.Merchant.MoneyToGroup
		query.StlmWay = in.Merchant.StlmWay
		query.StlmWayDesc = in.Merchant.StlmWayDesc
		query.StlmInsCircle = in.Merchant.StlmInsCircle
		query.Status = in.Merchant.Status
		query.UcBcCd32 = in.Merchant.UcBcCd32
		query.K2WorkflowId = in.Merchant.K2WorkflowId
		query.SystemFlag = in.Merchant.SystemFlag
		query.ApprovalUsername = in.Merchant.ApprovalUsername
		query.FinalApprovalUsername = in.Merchant.FinalApprovalUsername
		query.IsUpStandard = in.Merchant.IsUpStandard
		query.BillingType = in.Merchant.BillingType
		query.BillingLevel = in.Merchant.BillingLevel
		query.Slogan = in.Merchant.Slogan
		query.Ext1 = in.Merchant.Ext1
		query.Ext2 = in.Merchant.Ext2
		query.Ext3 = in.Merchant.Ext3
		query.Ext4 = in.Merchant.Ext4
		query.AreaStandard = in.Merchant.AreaStandard
		query.MchtCdAreaCd = in.Merchant.MchtCdAreaCd
		query.UcBcCdArea = in.Merchant.UcBcCdArea
		query.RecOprId = in.Merchant.RecOprId
		query.RecUpdOpr = in.Merchant.RecUpdOpr
		query.OperIn = in.Merchant.OperIn
		query.OemOrgCode = in.Merchant.OemOrgCode
		query.IsEleInvoice = in.Merchant.IsEleInvoice
		query.DutyParagraph = in.Merchant.DutyParagraph
		query.TaxMachineBrand = in.Merchant.TaxMachineBrand
		query.Ext5 = in.Merchant.Ext5
		query.Ext6 = in.Merchant.Ext6
		query.Ext7 = in.Merchant.Ext7
		query.Ext8 = in.Merchant.Ext8
		query.Ext9 = in.Merchant.Ext9
		query.BusiLiceSt = in.Merchant.BusiLiceSt
		query.BusiLiceDt = in.Merchant.BusiLiceDt
		query.CertifSt = in.Merchant.CertifSt
		query.CertifDt = in.Merchant.CertifDt
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
		if !merchants[i].ApprDate.IsZero() {
			pbMerchants[i].ApprDate = merchants[i].ApprDate.Format("2006-01-02 15:04:05")
		}
		if !merchants[i].DeleteDate.IsZero() {
			pbMerchants[i].DeleteDate = merchants[i].DeleteDate.Format("2006-01-02 15:04:05")
		}
		if !merchants[i].RecCrtTs.IsZero() {
			pbMerchants[i].RecCrtTs = merchants[i].RecCrtTs.Format("2006-01-02 15:04:05")
		}
		if !merchants[i].RecUpdTs.IsZero() {
			pbMerchants[i].RecUpdTs = merchants[i].RecUpdTs.Format("2006-01-02 15:04:05")
		}
		if !merchants[i].RecApllyTs.IsZero() {
			pbMerchants[i].RecApllyTs = merchants[i].RecApllyTs.Format("2006-01-02 15:04:05")
		}
	}

	return &pb.ListMerchantReply{
		Merchants: pbMerchants,
		Count:     count,
		Page:      in.Page,
		Size:      in.Size,
	}, nil
}

func (m *merchantService) ListGroupMerchant(ctx context.Context, in *pb.ListGroupMerchantRequest) (*pb.ListGroupMerchantReply, error) {
	if in.Size == 0 {
		in.Size = 10
	}
	if in.Page == 0 {
		in.Page = 1
	}
	db := common.DB

	query := new(merchantmodel.Group)
	if in.Group != nil {
		query.JtMchtCd = in.Group.JtMchtCd
		query.JtMchtNm = in.Group.JtMchtNm
		query.JtArea = in.Group.JtArea
		query.MchtStlmCNm = in.Group.MchtStlmCNm
		query.MchtStlmCAcct = in.Group.MchtStlmCAcct
		query.ChtStlmInsIdCd = in.Group.ChtStlmInsIdCd
		query.MchtStlmInsNm = in.Group.MchtStlmInsNm
		query.MchtPaySysAcct = in.Group.MchtPaySysAcct
		query.ProvCd = in.Group.ProvCd
		query.CityCd = in.Group.CityCd
		query.AipBranCd = in.Group.AipBranCd
		query.SystemFlag = in.Group.SystemFlag
		query.Status = in.Group.Status
		query.RecOprId = in.Group.RecOprId
		query.RecUpdOpr = in.Group.RecUpdOpr
		query.JtAddr = in.Group.JtAddr
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
		Groups: pbGroups,
		Count:  count,
		Page:   in.Page,
		Size:   in.Size,
	}, nil
}
