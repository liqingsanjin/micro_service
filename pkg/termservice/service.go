package termservice

import (
	"context"
	"userService/pkg/common"
	termmodel "userService/pkg/model/term"
	"userService/pkg/pb"
)

type service struct{}

func (s *service) ListTermInfo(ctx context.Context, in *pb.ListTermInfoRequest) (*pb.ListTermInfoReply, error) {
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 10
	}

	db := common.DB

	query := new(termmodel.Info)
	if in.Term != nil {
		query.MchtCd = in.Term.MchtCd
		query.TermId = in.Term.TermId
		query.TermTp = in.Term.TermTp
		query.Belong = in.Term.Belong
		query.BelongSub = in.Term.BelongSub
		query.TmnlMoneyIntype = in.Term.TmnlMoneyIntype
		query.TmnlMoney = in.Term.TmnlMoney
		query.TmnlBrand = in.Term.TmnlBrand
		query.TmnlModelNo = in.Term.TmnlModelNo
		query.TmnlBarcode = in.Term.TmnlBarcode
		query.DeviceCd = in.Term.DeviceCd
		query.InstallLocation = in.Term.InstallLocation
		query.TmnlIntype = in.Term.TmnlIntype
		query.DialOut = in.Term.DialOut
		query.DealTypes = in.Term.DealTypes
		query.RecOprId = in.Term.RecOprId
		query.RecUpdOpr = in.Term.RecUpdOpr
		query.AppCd = in.Term.AppCd
		query.SystemFlag = in.Term.SystemFlag
		query.Status = in.Term.Status
		query.ActiveCode = in.Term.ActiveCode
		query.NoFlag = in.Term.NoFlag
		query.MsgResvFld1 = in.Term.MsgResvFld1
		query.MsgResvFld2 = in.Term.MsgResvFld2
		query.MsgResvFld3 = in.Term.MsgResvFld3
		query.MsgResvFld4 = in.Term.MsgResvFld4
		query.MsgResvFld5 = in.Term.MsgResvFld5
		query.MsgResvFld6 = in.Term.MsgResvFld6
		query.MsgResvFld7 = in.Term.MsgResvFld7
		query.MsgResvFld8 = in.Term.MsgResvFld8
		query.MsgResvFld9 = in.Term.MsgResvFld9
		query.MsgResvFld10 = in.Term.MsgResvFld10
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
		Terms: pbInfos,
	}, nil
}
