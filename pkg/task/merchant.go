package task

import (
	"fmt"
	"userService/pkg/camunda/pb"
	camundamodel "userService/pkg/model/camunda"
	"userService/pkg/model/merchant"
	"userService/pkg/model/term"

	"github.com/jinzhu/gorm"
)

func merchantRegister(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询商户信息
	info, err := merchant.FindMerchantInfoById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("merchant %s not found", instance.DataId)
	}

	// 查询term, bank_account, biz_fee, biz_deal, business, picture, term_risk_cfg
	accounts, err := merchant.QueryBankAccount(db, &merchant.BankAccount{
		OwnerCd: info.MchtCd,
	})
	if err != nil {
		return err
	}
	fees, err := merchant.QueryBizFee(db, &merchant.BizFee{
		MchtCd: info.MchtCd,
	})
	if err != nil {
		return err
	}
	deals, err := merchant.QueryBizDeal(db, &merchant.BizDeal{
		MchtCd: info.MchtCd,
	})
	if err != nil {
		return err
	}
	business, err := merchant.QueryBusiness(db, &merchant.Business{
		MchtCd: info.MchtCd,
	})
	if err != nil {
		return err
	}
	pictures, err := merchant.QueryPicture(db, &merchant.Picture{
		MchtCd: info.MchtCd,
	})
	if err != nil {
		return err
	}
	terms, _, err := term.QueryTermInfo(db, &term.Info{
		MchtCd: info.MchtCd,
	}, 1, 10000)
	if err != nil {
		return err
	}
	risks, _, err := term.QueryTermRisk(db, &term.Risk{
		MchtCd: info.MchtCd,
	}, 1, 10000)
	if err != nil {
		return err
	}

	// todo 入库
	info.Status = "01"
	err = merchant.SaveMerchantMain(db, &merchant.MerchantInfoMain{
		MerchantInfo: *info,
	})
	if err != nil {
		return err
	}

	for i := range accounts {
		err = merchant.SaveBankAccountMain(db, &merchant.BankAccountMain{
			BankAccount: *accounts[i],
		})
		if err != nil {
			return err
		}
	}
	for i := range fees {
		err = merchant.SaveBizFeeMain(db, &merchant.BizFeeMain{
			BizFee: *fees[i],
		})
		if err != nil {
			return err
		}
	}
	for i := range deals {
		err = merchant.SaveBizDealMain(db, &merchant.BizDealMain{
			BizDeal: *deals[i],
		})
		if err != nil {
			return err
		}
	}
	for i := range business {
		err = merchant.SaveBusinessMain(db, &merchant.BusinessMain{
			Business: *business[i],
		})
		if err != nil {
			return err
		}
	}
	for i := range pictures {
		err = merchant.SavePictureMain(db, &merchant.PictureMain{
			Picture: *pictures[i],
		})
		if err != nil {
			return err
		}
	}
	for i := range terms {
		err = term.SaveTermInfoMain(db, &term.InfoMain{
			Info: *terms[i],
		})
		if err != nil {
			return err
		}
	}
	for i := range risks {
		err = term.SaveRiskMain(db, &term.RiskMain{
			Risk: *risks[i],
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteMerchant(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	return nil
}
