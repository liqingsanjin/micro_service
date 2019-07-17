package termservice

import (
	"context"
	"net/http"
	"userService/pkg/common"
	termmodel "userService/pkg/model/term"
	"userService/pkg/pb"
)

type service struct{}

func (s *service) SaveTermRisk(ctx context.Context, in *pb.SaveTermRiskRequest) (*pb.SaveTermRiskReply, error) {
	var reply pb.SaveTermRiskReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "保存信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	data := new(termmodel.Risk)
	{
		data.MchtCd = in.Item.MchtCd
		data.TermId = in.Item.TermId
		data.CardType = in.Item.CardType
		data.TotalLimitMoney = in.Item.TotalLimitMoney
		data.AccpetStartTime = in.Item.AccpetStartTime
		data.AccpetStartDate = in.Item.AccpetStartDate
		data.AccpetEndTime = in.Item.AccpetEndTime
		data.AccpetEndDate = in.Item.AccpetEndDate
		data.SingleLimitMoney = in.Item.SingleLimitMoney
		data.ControlWay = in.Item.ControlWay
		data.SingleMinMoney = in.Item.SingleMinMoney
		data.TotalPeriod = in.Item.TotalPeriod
		data.RecOprId = in.Item.RecOprId
		data.RecUpdOpr = in.Item.RecUpdOpr
		data.OperIn = in.Item.OperIn
	}
	err := termmodel.SaveRisk(db, data)
	return &reply, err
}

func (s *service) ListTermInfo(ctx context.Context, in *pb.ListTermInfoRequest) (*pb.ListTermInfoReply, error) {
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 10
	}

	db := common.DB

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
			pbInfos[i].CreatedAt = infos[i].CreatedAt.Format("2006-01-02 15:04:05")
		}
		if !infos[i].UpdatedAt.IsZero() {
			pbInfos[i].UpdatedAt = infos[i].UpdatedAt.Format("2006-01-02 15:04:05")
		}
	}
	return &pb.ListTermInfoReply{
		Page:  in.Page,
		Size:  in.Size,
		Count: count,
		Items: pbInfos,
	}, nil
}

func (s *service) SaveTerm(ctx context.Context, in *pb.SaveTermRequest) (*pb.SaveTermReply, error) {
	var reply pb.SaveTermReply
	if in.Item == nil {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     "InvalidParamsError",
			Description: "保存信息为空",
		}
		return &reply, nil
	}
	db := common.DB

	data := new(termmodel.Info)
	{
		data.MchtCd = in.Item.MchtCd
		data.TermId = in.Item.TermId
		data.TermTp = in.Item.TermTp
		data.Belong = in.Item.Belong
		data.BelongSub = in.Item.BelongSub
		data.TmnlMoneyIntype = in.Item.TmnlMoneyIntype
		data.TmnlMoney = in.Item.TmnlMoney
		data.TmnlBrand = in.Item.TmnlBrand
		data.TmnlModelNo = in.Item.TmnlModelNo
		data.TmnlBarcode = in.Item.TmnlBarcode
		data.DeviceCd = in.Item.DeviceCd
		data.InstallLocation = in.Item.InstallLocation
		data.TmnlIntype = in.Item.TmnlIntype
		data.DialOut = in.Item.DialOut
		data.DealTypes = in.Item.DealTypes
		data.RecOprId = in.Item.RecOprId
		data.RecUpdOpr = in.Item.RecUpdOpr
		data.AppCd = in.Item.AppCd
		data.SystemFlag = in.Item.SystemFlag
		data.Status = in.Item.Status
		data.ActiveCode = in.Item.ActiveCode
		data.NoFlag = in.Item.NoFlag
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
	err := termmodel.SaveTermInfo(db, data)
	return &reply, err
}
