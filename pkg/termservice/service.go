package termservice

import (
	"context"
	"time"
	"userService/pkg/common"
	termmodel "userService/pkg/model/term"
	"userService/pkg/pb"
	"userService/pkg/util"
)

type service struct{}

func (s *service) ListTermActivationState(ctx context.Context, in *pb.ListTermActivationStateRequest) (*pb.ListTermActivationStateReply, error) {
	reply := new(pb.ListTermActivationStateReply)
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
		query := new(termmodel.ActivationState)
		if in.Item != nil {
			query.ActiveCode = in.Item.ActiveCode
			query.ActiveType = in.Item.ActiveType
			query.MchtCd = in.Item.MchtCd
			query.TermId = in.Item.TermId
			query.NewKsn = in.Item.NewKsn
			query.OldKsn = in.Item.OldKsn
			query.IsActive = in.Item.IsActive
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
		}
		infos, count, err := termmodel.QueryActivationState(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbInfos := make([]*pb.TermActivationStateField, len(infos))
		for i := range infos {
			pbInfos[i] = &pb.TermActivationStateField{
				ActiveCode: infos[i].ActiveCode,
				ActiveType: infos[i].ActiveType,
				MchtCd:     infos[i].MchtCd,
				TermId:     infos[i].TermId,
				NewKsn:     infos[i].NewKsn,
				OldKsn:     infos[i].OldKsn,
				IsActive:   infos[i].IsActive,
				RecOprId:   infos[i].RecOprId,
				RecUpdOpr:  infos[i].RecUpdOpr,
			}
			if !infos[i].CreatedAt.IsZero() {
				pbInfos[i].CreatedAt = infos[i].CreatedAt.Format(util.TimePattern)
			}
			if !infos[i].UpdatedAt.IsZero() {
				pbInfos[i].UpdatedAt = infos[i].UpdatedAt.Format(util.TimePattern)
			}
			if infos[i].ActiveDate.Valid {
				pbInfos[i].ActiveDate = infos[i].ActiveDate.Time.Format(util.TimePattern)
			}
			if infos[i].CreateDate.Valid {
				pbInfos[i].CreateDate = infos[i].CreateDate.Time.Format(util.TimePattern)
			}
		}

		reply.Count = count
		reply.Items = pbInfos

	} else {
		query := new(termmodel.ActivationStateMain)
		if in.Item != nil {
			query.ActiveCode = in.Item.ActiveCode
			query.ActiveType = in.Item.ActiveType
			query.MchtCd = in.Item.MchtCd
			query.TermId = in.Item.TermId
			query.NewKsn = in.Item.NewKsn
			query.OldKsn = in.Item.OldKsn
			query.IsActive = in.Item.IsActive
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
		}
		infos, count, err := termmodel.QueryActivationStateMain(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbInfos := make([]*pb.TermActivationStateField, len(infos))
		for i := range infos {
			pbInfos[i] = &pb.TermActivationStateField{
				ActiveCode: infos[i].ActiveCode,
				ActiveType: infos[i].ActiveType,
				MchtCd:     infos[i].MchtCd,
				TermId:     infos[i].TermId,
				NewKsn:     infos[i].NewKsn,
				OldKsn:     infos[i].OldKsn,
				IsActive:   infos[i].IsActive,
				RecOprId:   infos[i].RecOprId,
				RecUpdOpr:  infos[i].RecUpdOpr,
			}
			if !infos[i].CreatedAt.IsZero() {
				pbInfos[i].CreatedAt = infos[i].CreatedAt.Format(util.TimePattern)
			}
			if !infos[i].UpdatedAt.IsZero() {
				pbInfos[i].UpdatedAt = infos[i].UpdatedAt.Format(util.TimePattern)
			}
			if infos[i].ActiveDate.Valid {
				pbInfos[i].ActiveDate = infos[i].ActiveDate.Time.Format(util.TimePattern)
			}
			if infos[i].CreateDate.Valid {
				pbInfos[i].CreateDate = infos[i].CreateDate.Time.Format(util.TimePattern)
			}
		}

		reply.Count = count
		reply.Items = pbInfos
	}

	reply.Size = in.Size
	reply.Page = in.Page
	return reply, nil
}

func (s *service) ListTermRisk(ctx context.Context, in *pb.ListTermRiskRequest) (*pb.ListTermRiskReply, error) {
	reply := new(pb.ListTermRiskReply)
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
		query := new(termmodel.Risk)
		if in.Item != nil {
			query.MchtCd = in.Item.MchtCd
			query.TermId = in.Item.TermId
			query.CardType = in.Item.CardType
			query.TotalLimitMoney = in.Item.TotalLimitMoney
			query.AccpetStartTime = in.Item.AccpetStartTime
			query.AccpetStartDate = in.Item.AccpetStartDate
			query.AccpetEndTime = in.Item.AccpetEndTime
			query.AccpetEndDate = in.Item.AccpetEndDate
			query.SingleLimitMoney = in.Item.SingleLimitMoney
			query.ControlWay = in.Item.ControlWay
			query.SingleMinMoney = in.Item.SingleMinMoney
			query.TotalPeriod = in.Item.TotalPeriod
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
			query.OperIn = in.Item.OperIn
		}
		infos, count, err := termmodel.QueryTermRisk(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbInfos := make([]*pb.TermRiskField, len(infos))
		for i := range infos {
			pbInfos[i] = &pb.TermRiskField{
				MchtCd:           infos[i].MchtCd,
				TermId:           infos[i].TermId,
				CardType:         infos[i].CardType,
				TotalLimitMoney:  infos[i].TotalLimitMoney,
				AccpetStartTime:  infos[i].AccpetStartTime,
				AccpetStartDate:  infos[i].AccpetStartDate,
				AccpetEndTime:    infos[i].AccpetEndTime,
				AccpetEndDate:    infos[i].AccpetEndDate,
				SingleLimitMoney: infos[i].SingleLimitMoney,
				ControlWay:       infos[i].ControlWay,
				SingleMinMoney:   infos[i].SingleMinMoney,
				TotalPeriod:      infos[i].TotalPeriod,
				RecOprId:         infos[i].RecOprId,
				RecUpdOpr:        infos[i].RecUpdOpr,
				OperIn:           infos[i].OperIn,
			}
			if !infos[i].CreatedAt.IsZero() {
				pbInfos[i].CreatedAt = infos[i].CreatedAt.Format(util.TimePattern)
			}
			if !infos[i].UpdatedAt.IsZero() {
				pbInfos[i].UpdatedAt = infos[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Count = count
		reply.Items = pbInfos

	} else {

		query := new(termmodel.RiskMain)
		if in.Item != nil {
			query.MchtCd = in.Item.MchtCd
			query.TermId = in.Item.TermId
			query.CardType = in.Item.CardType
			query.TotalLimitMoney = in.Item.TotalLimitMoney
			query.AccpetStartTime = in.Item.AccpetStartTime
			query.AccpetStartDate = in.Item.AccpetStartDate
			query.AccpetEndTime = in.Item.AccpetEndTime
			query.AccpetEndDate = in.Item.AccpetEndDate
			query.SingleLimitMoney = in.Item.SingleLimitMoney
			query.ControlWay = in.Item.ControlWay
			query.SingleMinMoney = in.Item.SingleMinMoney
			query.TotalPeriod = in.Item.TotalPeriod
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
			query.OperIn = in.Item.OperIn
		}
		infos, count, err := termmodel.QueryTermRiskMain(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbInfos := make([]*pb.TermRiskField, len(infos))
		for i := range infos {
			pbInfos[i] = &pb.TermRiskField{
				MchtCd:           infos[i].MchtCd,
				TermId:           infos[i].TermId,
				CardType:         infos[i].CardType,
				TotalLimitMoney:  infos[i].TotalLimitMoney,
				AccpetStartTime:  infos[i].AccpetStartTime,
				AccpetStartDate:  infos[i].AccpetStartDate,
				AccpetEndTime:    infos[i].AccpetEndTime,
				AccpetEndDate:    infos[i].AccpetEndDate,
				SingleLimitMoney: infos[i].SingleLimitMoney,
				ControlWay:       infos[i].ControlWay,
				SingleMinMoney:   infos[i].SingleMinMoney,
				TotalPeriod:      infos[i].TotalPeriod,
				RecOprId:         infos[i].RecOprId,
				RecUpdOpr:        infos[i].RecUpdOpr,
				OperIn:           infos[i].OperIn,
			}
			if !infos[i].CreatedAt.IsZero() {
				pbInfos[i].CreatedAt = infos[i].CreatedAt.Format(util.TimePattern)
			}
			if !infos[i].UpdatedAt.IsZero() {
				pbInfos[i].UpdatedAt = infos[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Count = count
		reply.Items = pbInfos
	}

	reply.Size = in.Size
	reply.Page = in.Page
	return reply, nil
}

func (s *service) ListTermInfo(ctx context.Context, in *pb.ListTermInfoRequest) (*pb.ListTermInfoReply, error) {
	reply := new(pb.ListTermInfoReply)
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
		query := new(termmodel.Info)
		if in.Item != nil {
			query.MchtCd = in.Item.MchtCd
			query.TermId = in.Item.TermId
			query.TermTp = in.Item.TermTp
			query.Belong = in.Item.Belong
			query.BelongSub = in.Item.BelongSub
			query.TmnlMoneyIntype = in.Item.TmnlMoneyIntype
			query.TmnlMoney = in.Item.TmnlMoney
			query.TmnlBrand = in.Item.TmnlBrand
			query.TmnlModelNo = in.Item.TmnlModelNo
			query.TmnlBarcode = in.Item.TmnlBarcode
			query.DeviceCd = in.Item.DeviceCd
			query.InstallLocation = in.Item.InstallLocation
			query.TmnlIntype = in.Item.TmnlIntype
			query.DialOut = in.Item.DialOut
			query.DealTypes = in.Item.DealTypes
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
			query.AppCd = in.Item.AppCd
			query.SystemFlag = in.Item.SystemFlag
			query.Status = in.Item.Status
			query.ActiveCode = in.Item.ActiveCode
			query.NoFlag = in.Item.NoFlag
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
		}
		infos, count, err := termmodel.QueryTermInfo(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbInfos := make([]*pb.TermInfoField, len(infos))
		for i := range infos {
			pbInfos[i] = &pb.TermInfoField{
				MchtCd:          infos[i].MchtCd,
				TermId:          infos[i].TermId,
				TermTp:          infos[i].TermTp,
				Belong:          infos[i].Belong,
				BelongSub:       infos[i].BelongSub,
				TmnlMoneyIntype: infos[i].TmnlMoneyIntype,
				TmnlMoney:       infos[i].TmnlMoney,
				TmnlBrand:       infos[i].TmnlBrand,
				TmnlModelNo:     infos[i].TmnlModelNo,
				TmnlBarcode:     infos[i].TmnlBarcode,
				DeviceCd:        infos[i].DeviceCd,
				InstallLocation: infos[i].InstallLocation,
				TmnlIntype:      infos[i].TmnlIntype,
				DialOut:         infos[i].DialOut,
				DealTypes:       infos[i].DealTypes,
				RecOprId:        infos[i].RecOprId,
				RecUpdOpr:       infos[i].RecUpdOpr,
				AppCd:           infos[i].AppCd,
				SystemFlag:      infos[i].SystemFlag,
				Status:          infos[i].Status,
				ActiveCode:      infos[i].ActiveCode,
				NoFlag:          infos[i].NoFlag,
				MsgResvFld1:     infos[i].MsgResvFld1,
				MsgResvFld2:     infos[i].MsgResvFld2,
				MsgResvFld3:     infos[i].MsgResvFld3,
				MsgResvFld4:     infos[i].MsgResvFld4,
				MsgResvFld5:     infos[i].MsgResvFld5,
				MsgResvFld6:     infos[i].MsgResvFld6,
				MsgResvFld7:     infos[i].MsgResvFld7,
				MsgResvFld8:     infos[i].MsgResvFld8,
				MsgResvFld9:     infos[i].MsgResvFld9,
				MsgResvFld10:    infos[i].MsgResvFld10,
			}
			if !infos[i].CreatedAt.IsZero() {
				pbInfos[i].CreatedAt = infos[i].CreatedAt.Format(util.TimePattern)
			}
			if !infos[i].UpdatedAt.IsZero() {
				pbInfos[i].UpdatedAt = infos[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Count = count
		reply.Items = pbInfos
	} else {
		query := new(termmodel.InfoMain)
		if in.Item != nil {
			query.MchtCd = in.Item.MchtCd
			query.TermId = in.Item.TermId
			query.TermTp = in.Item.TermTp
			query.Belong = in.Item.Belong
			query.BelongSub = in.Item.BelongSub
			query.TmnlMoneyIntype = in.Item.TmnlMoneyIntype
			query.TmnlMoney = in.Item.TmnlMoney
			query.TmnlBrand = in.Item.TmnlBrand
			query.TmnlModelNo = in.Item.TmnlModelNo
			query.TmnlBarcode = in.Item.TmnlBarcode
			query.DeviceCd = in.Item.DeviceCd
			query.InstallLocation = in.Item.InstallLocation
			query.TmnlIntype = in.Item.TmnlIntype
			query.DialOut = in.Item.DialOut
			query.DealTypes = in.Item.DealTypes
			query.RecOprId = in.Item.RecOprId
			query.RecUpdOpr = in.Item.RecUpdOpr
			query.AppCd = in.Item.AppCd
			query.SystemFlag = in.Item.SystemFlag
			query.Status = in.Item.Status
			query.ActiveCode = in.Item.ActiveCode
			query.NoFlag = in.Item.NoFlag
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
		}
		infos, count, err := termmodel.QueryTermInfoMain(db, query, in.Page, in.Size)
		if err != nil {
			return nil, err
		}

		pbInfos := make([]*pb.TermInfoField, len(infos))
		for i := range infos {
			pbInfos[i] = &pb.TermInfoField{
				MchtCd:          infos[i].MchtCd,
				TermId:          infos[i].TermId,
				TermTp:          infos[i].TermTp,
				Belong:          infos[i].Belong,
				BelongSub:       infos[i].BelongSub,
				TmnlMoneyIntype: infos[i].TmnlMoneyIntype,
				TmnlMoney:       infos[i].TmnlMoney,
				TmnlBrand:       infos[i].TmnlBrand,
				TmnlModelNo:     infos[i].TmnlModelNo,
				TmnlBarcode:     infos[i].TmnlBarcode,
				DeviceCd:        infos[i].DeviceCd,
				InstallLocation: infos[i].InstallLocation,
				TmnlIntype:      infos[i].TmnlIntype,
				DialOut:         infos[i].DialOut,
				DealTypes:       infos[i].DealTypes,
				RecOprId:        infos[i].RecOprId,
				RecUpdOpr:       infos[i].RecUpdOpr,
				AppCd:           infos[i].AppCd,
				SystemFlag:      infos[i].SystemFlag,
				Status:          infos[i].Status,
				ActiveCode:      infos[i].ActiveCode,
				NoFlag:          infos[i].NoFlag,
				MsgResvFld1:     infos[i].MsgResvFld1,
				MsgResvFld2:     infos[i].MsgResvFld2,
				MsgResvFld3:     infos[i].MsgResvFld3,
				MsgResvFld4:     infos[i].MsgResvFld4,
				MsgResvFld5:     infos[i].MsgResvFld5,
				MsgResvFld6:     infos[i].MsgResvFld6,
				MsgResvFld7:     infos[i].MsgResvFld7,
				MsgResvFld8:     infos[i].MsgResvFld8,
				MsgResvFld9:     infos[i].MsgResvFld9,
				MsgResvFld10:    infos[i].MsgResvFld10,
			}
			if !infos[i].CreatedAt.IsZero() {
				pbInfos[i].CreatedAt = infos[i].CreatedAt.Format(util.TimePattern)
			}
			if !infos[i].UpdatedAt.IsZero() {
				pbInfos[i].UpdatedAt = infos[i].UpdatedAt.Format(util.TimePattern)
			}
		}

		reply.Count = count
		reply.Items = pbInfos
	}
	reply.Size = in.Size
	reply.Page = in.Page

	return reply, nil
}

func (s *service) SaveTerm(ctx context.Context, in *pb.SaveTermRequest) (*pb.SaveTermReply, error) {
	var reply pb.SaveTermReply
	var err error
	db := common.DB
	db = db.Begin()
	defer db.Rollback()

	if in.Term != nil {
		data := new(termmodel.Info)
		data.MchtCd = in.Term.MchtCd
		data.TermId = in.Term.TermId
		data.TermTp = in.Term.TermTp
		data.Belong = in.Term.Belong
		data.BelongSub = in.Term.BelongSub
		data.TmnlMoneyIntype = in.Term.TmnlMoneyIntype
		data.TmnlMoney = in.Term.TmnlMoney
		data.TmnlBrand = in.Term.TmnlBrand
		data.TmnlModelNo = in.Term.TmnlModelNo
		data.TmnlBarcode = in.Term.TmnlBarcode
		data.DeviceCd = in.Term.DeviceCd
		data.InstallLocation = in.Term.InstallLocation
		data.TmnlIntype = in.Term.TmnlIntype
		data.DialOut = in.Term.DialOut
		data.DealTypes = in.Term.DealTypes
		data.RecOprId = in.Term.RecOprId
		data.RecUpdOpr = in.Term.RecUpdOpr
		data.AppCd = in.Term.AppCd
		data.SystemFlag = in.Term.SystemFlag
		data.Status = in.Term.Status
		data.ActiveCode = in.Term.ActiveCode
		data.NoFlag = in.Term.NoFlag
		data.MsgResvFld1 = in.Term.MsgResvFld1
		data.MsgResvFld2 = in.Term.MsgResvFld2
		data.MsgResvFld3 = in.Term.MsgResvFld3
		data.MsgResvFld4 = in.Term.MsgResvFld4
		data.MsgResvFld5 = in.Term.MsgResvFld5
		data.MsgResvFld6 = in.Term.MsgResvFld6
		data.MsgResvFld7 = in.Term.MsgResvFld7
		data.MsgResvFld8 = in.Term.MsgResvFld8
		data.MsgResvFld9 = in.Term.MsgResvFld9
		data.MsgResvFld10 = in.Term.MsgResvFld10
		err = termmodel.SaveTermInfo(db, data)
		if err != nil {
			return nil, err
		}
	}

	if in.Risk != nil {
		data := new(termmodel.Risk)
		data.MchtCd = in.Risk.MchtCd
		data.TermId = in.Risk.TermId
		data.CardType = in.Risk.CardType
		data.TotalLimitMoney = in.Risk.TotalLimitMoney
		data.AccpetStartTime = in.Risk.AccpetStartTime
		data.AccpetStartDate = in.Risk.AccpetStartDate
		data.AccpetEndTime = in.Risk.AccpetEndTime
		data.AccpetEndDate = in.Risk.AccpetEndDate
		data.SingleLimitMoney = in.Risk.SingleLimitMoney
		data.ControlWay = in.Risk.ControlWay
		data.SingleMinMoney = in.Risk.SingleMinMoney
		data.TotalPeriod = in.Risk.TotalPeriod
		data.RecOprId = in.Risk.RecOprId
		data.RecUpdOpr = in.Risk.RecUpdOpr
		data.OperIn = in.Risk.OperIn
		err = termmodel.SaveRisk(db, data)
		if err != nil {
			return nil, err
		}
	}
	if in.Activation != nil {
		data := new(termmodel.ActivationState)
		data.ActiveCode = in.Activation.ActiveCode
		data.ActiveType = in.Activation.ActiveType
		data.MchtCd = in.Activation.MchtCd
		data.TermId = in.Activation.TermId
		data.NewKsn = in.Activation.NewKsn
		data.OldKsn = in.Activation.OldKsn
		data.IsActive = in.Activation.IsActive
		data.RecOprId = in.Activation.RecOprId
		data.RecUpdOpr = in.Activation.RecUpdOpr
		if in.Activation.ActiveDate != "" {
			data.ActiveDate.Time, err = time.Parse(util.TimePattern, in.Activation.ActiveDate)
			if err == nil {
				data.ActiveDate.Valid = true
			}
		}
		if in.Activation.CreateDate != "" {
			data.CreateDate.Time, err = time.Parse(util.TimePattern, in.Activation.CreateDate)
			if err == nil {
				data.CreateDate.Valid = true
			}
		}
		err = termmodel.SaveActivationState(db, data)
		if err != nil {
			return nil, err
		}
	}
	db.Commit()
	return &reply, err
}
