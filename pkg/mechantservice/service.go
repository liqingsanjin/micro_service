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
