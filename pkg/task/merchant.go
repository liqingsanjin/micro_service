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
	accounts, _, err := merchant.QueryBankAccount(db, &merchant.BankAccount{
		OwnerCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	fees, _, err := merchant.QueryBizFee(db, &merchant.BizFee{
		MchtCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	deals, _, err := merchant.QueryBizDeal(db, &merchant.BizDeal{
		MchtCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	business, _, err := merchant.QueryBusiness(db, &merchant.Business{
		MchtCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	pictures, _, err := merchant.QueryPicture(db, &merchant.Picture{
		MchtCd: info.MchtCd,
	}, 1, 9999)
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

	// 入库
	err = merchant.UpdateMerchant(db, &merchant.MerchantInfo{
		MchtCd: info.MchtCd,
	}, &merchant.MerchantInfo{
		Status: "01",
	})
	if err != nil {
		return err
	}

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
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 删除商户相关信息
	err = merchant.DeleteMerchant(db, &merchant.MerchantInfo{
		MchtCd: instance.DataId,
	})
	if err != nil {
		return err
	}

	err = merchant.DeleteBankAccount(db, &merchant.BankAccount{
		OwnerCd: instance.DataId,
	})
	if err != nil {
		return err
	}

	err = merchant.DeleteBizDeal(db, &merchant.BizDeal{
		MchtCd: instance.DataId,
	})
	if err != nil {
		return err
	}

	err = merchant.DeleteBizFee(db, &merchant.BizFee{
		MchtCd: instance.DataId,
	})
	if err != nil {
		return err
	}

	err = merchant.DeleteBusiness(db, &merchant.Business{
		MchtCd: instance.DataId,
	})
	if err != nil {
		return err
	}

	err = merchant.DeletePicture(db, &merchant.Picture{
		MchtCd: instance.DataId,
	})
	if err != nil {
		return err
	}

	err = term.DeleteTerm(db, &term.Info{
		MchtCd: instance.DataId,
	})
	if err != nil {
		return err
	}

	err = term.DeleteRisk(db, &term.Risk{
		MchtCd: instance.DataId,
	})
	if err != nil {
		return err
	}

	return nil
}

func merchantUpdate(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
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
	accounts, _, err := merchant.QueryBankAccount(db, &merchant.BankAccount{
		OwnerCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	fees, _, err := merchant.QueryBizFee(db, &merchant.BizFee{
		MchtCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	deals, _, err := merchant.QueryBizDeal(db, &merchant.BizDeal{
		MchtCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	business, _, err := merchant.QueryBusiness(db, &merchant.Business{
		MchtCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	pictures, _, err := merchant.QueryPicture(db, &merchant.Picture{
		MchtCd: info.MchtCd,
	}, 1, 9999)
	if err != nil {
		return err
	}

	// 入库
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

	return nil
}

func merchantUpdateCancel(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询正式表商户信息
	info, err := merchant.FindMerchantInfoMainById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("merchant %s not found", instance.DataId)
	}

	// 查询正式表term, bank_account, biz_fee, biz_deal, business, picture, term_risk_cfg
	accounts, _, err := merchant.QueryBankAccountMain(db, &merchant.BankAccountMain{
		BankAccount: merchant.BankAccount{
			OwnerCd: info.MchtCd,
		},
	}, 1, 9999)
	if err != nil {
		return err
	}
	fees, _, err := merchant.QueryBizFeeMain(db, &merchant.BizFeeMain{
		BizFee: merchant.BizFee{
			MchtCd: info.MchtCd,
		},
	}, 1, 9999)
	if err != nil {
		return err
	}
	deals, _, err := merchant.QueryBizDealMain(db, &merchant.BizDealMain{
		BizDeal: merchant.BizDeal{
			MchtCd: info.MchtCd,
		},
	}, 1, 9999)
	if err != nil {
		return err
	}
	business, _, err := merchant.QueryBusinessMain(db, &merchant.BusinessMain{
		Business: merchant.Business{
			MchtCd: info.MchtCd,
		},
	}, 1, 9999)
	if err != nil {
		return err
	}
	pictures, _, err := merchant.QueryPictureMain(db, &merchant.PictureMain{
		Picture: merchant.Picture{
			MchtCd: info.MchtCd,
		},
	}, 1, 9999)
	if err != nil {
		return err
	}

	// 入库编辑表
	err = merchant.SaveMerchant(db, &info.MerchantInfo)
	if err != nil {
		return err
	}

	for i := range accounts {
		err = merchant.SaveBankAccount(db, &accounts[i].BankAccount)
		if err != nil {
			return err
		}
	}
	for i := range fees {
		err = merchant.SaveBizFee(db, &fees[i].BizFee)
		if err != nil {
			return err
		}
	}
	for i := range deals {
		err = merchant.SaveBizDeal(db, &deals[i].BizDeal)
		if err != nil {
			return err
		}
	}
	for i := range business {
		err = merchant.SaveBusiness(db, &business[i].Business)
		if err != nil {
			return err
		}
	}
	for i := range pictures {
		err = merchant.SavePicture(db, &pictures[i].Picture)
		if err != nil {
			return err
		}
	}

	return nil
}

func merchantUnFreeze(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询编辑表商户信息
	info, err := merchant.FindMerchantInfoById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("merchant %s not found", instance.DataId)
	}

	err = merchant.SaveMerchantMain(db, &merchant.MerchantInfoMain{
		MerchantInfo: *info,
	})
	return err
}

func cancelMerchantUnFreeze(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询主表商户信息
	info, err := merchant.FindMerchantInfoMainById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("merchant %s not found", instance.DataId)
	}

	err = merchant.SaveMerchant(db, &info.MerchantInfo)
	return err
}

func merchantFreeze(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询编辑表商户信息
	info, err := merchant.FindMerchantInfoById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("merchant %s not found", instance.DataId)
	}

	err = merchant.SaveMerchantMain(db, &merchant.MerchantInfoMain{
		MerchantInfo: *info,
	})
	return err
}

func cancelMerchantFreeze(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询主表商户信息
	info, err := merchant.FindMerchantInfoMainById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("merchant %s not found", instance.DataId)
	}

	err = merchant.SaveMerchant(db, &info.MerchantInfo)
	return err
}

func merchantUnregister(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询编辑表商户信息
	info, err := merchant.FindMerchantInfoById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("merchant %s not found", instance.DataId)
	}

	err = merchant.SaveMerchantMain(db, &merchant.MerchantInfoMain{
		MerchantInfo: *info,
	})
	return err
}

func cancelMerchantUnregister(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询商户id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询主表商户信息
	info, err := merchant.FindMerchantInfoMainById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("merchant %s not found", instance.DataId)
	}

	err = merchant.SaveMerchant(db, &info.MerchantInfo)
	return err
}
