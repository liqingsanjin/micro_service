package task

import (
	"fmt"
	"userService/pkg/camunda/pb"
	camundamodel "userService/pkg/model/camunda"
	"userService/pkg/model/institution"

	"github.com/jinzhu/gorm"
)

func institutionRegister(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询机构id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询机构信息
	info, err := institution.FindInstitutionInfoById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("institution %s not found", instance.DataId)
	}
	// 查询fee，cash，control
	fees, _, err := institution.FindInstitutionFee(db, &institution.Fee{
		InsIdCd: info.InsIdCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	cashes, _, err := institution.FindInstitutionCash(db, &institution.Cash{
		InsIdCd: info.InsIdCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	controls, _, err := institution.FindInstitutionControl(db, &institution.Control{
		InsIdCd: info.InsIdCd,
	}, 1, 9999)
	if err != nil {
		return err
	}

	// 修改机构状态
	err = institution.UpdateInstitution(db, &institution.InstitutionInfo{
		InsIdCd: info.InsIdCd,
	}, &institution.InstitutionInfo{
		InsSta: "01",
	})
	if err != nil {
		return err
	}

	// 入正式表 institution fee cash control
	info.InsSta = "01"

	err = institution.SaveInstitutionMain(db, &institution.InstitutionInfoMain{
		InstitutionInfo: *info,
	})
	if err != nil {
		return err
	}
	for _, fee := range fees {
		err = institution.SaveInstitutionFeeMain(db, &institution.FeeMain{
			Fee: *fee,
		})
		if err != nil {
			return err
		}
	}
	for _, cash := range cashes {
		err = institution.SaveInstitutionCashMain(db, &institution.CashMain{
			Cash: *cash,
		})
		if err != nil {
			return err
		}
	}
	for _, control := range controls {
		err = institution.SaveInstitutionControlMain(db, &institution.ControlMain{
			Control: *control,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteInstitution(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询机构id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}
	// 删除 institution fee cash control
	err = institution.DeleteInstitution(db, &institution.InstitutionInfo{
		InsIdCd: instance.DataId,
	})
	if err != nil {
		return err
	}
	err = institution.DeleteInstitutionFee(db, &institution.Fee{
		InsIdCd: instance.DataId,
	})
	if err != nil {
		return err
	}
	err = institution.DeleteInstitutionCash(db, &institution.Cash{
		InsIdCd: instance.DataId,
	})
	if err != nil {
		return err
	}
	err = institution.DeleteInstitutionControl(db, &institution.Control{
		InsIdCd: instance.DataId,
	})
	if err != nil {
		return err
	}
	return nil
}

func institutionUpdate(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询机构id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询机构信息
	info, err := institution.FindInstitutionInfoById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("institution %s not found", instance.DataId)
	}
	// 查询fee，cash，control
	fees, _, err := institution.FindInstitutionFee(db, &institution.Fee{
		InsIdCd: info.InsIdCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	cashes, _, err := institution.FindInstitutionCash(db, &institution.Cash{
		InsIdCd: info.InsIdCd,
	}, 1, 9999)
	if err != nil {
		return err
	}
	controls, _, err := institution.FindInstitutionControl(db, &institution.Control{
		InsIdCd: info.InsIdCd,
	}, 1, 9999)
	if err != nil {
		return err
	}

	// 入正式表 institution fee cash control
	err = institution.SaveInstitutionMain(db, &institution.InstitutionInfoMain{
		InstitutionInfo: *info,
	})
	if err != nil {
		return err
	}
	for _, fee := range fees {
		err = institution.SaveInstitutionFeeMain(db, &institution.FeeMain{
			Fee: *fee,
		})
		if err != nil {
			return err
		}
	}
	for _, cash := range cashes {
		err = institution.SaveInstitutionCashMain(db, &institution.CashMain{
			Cash: *cash,
		})
		if err != nil {
			return err
		}
	}
	for _, control := range controls {
		err = institution.SaveInstitutionControlMain(db, &institution.ControlMain{
			Control: *control,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func institutionUpdateCancel(db *gorm.DB, in *pb.FetchAndLockExternalTaskRespItem) error {
	// 查询机构id
	instance, err := camundamodel.FindProcessInstanceByCamundaInstanceId(db, in.ProcessInstanceId)
	if err != nil {
		return err
	}
	if instance == nil {
		return fmt.Errorf("process %s not found", in.ProcessInstanceId)
	}

	// 查询正式表机构信息
	info, err := institution.FindInstitutionInfoMainById(db, instance.DataId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("institution %s not found", instance.DataId)
	}
	// 查询正式表fee，cash，control
	fees, _, err := institution.FindInstitutionFeeMain(db, &institution.FeeMain{
		Fee: institution.Fee{
			InsIdCd: info.InsIdCd,
		},
	}, 1, 9999)
	if err != nil {
		return err
	}
	cashes, _, err := institution.FindInstitutionCashMain(db, &institution.CashMain{
		Cash: institution.Cash{
			InsIdCd: info.InsIdCd,
		},
	}, 1, 9999)
	if err != nil {
		return err
	}
	controls, _, err := institution.FindInstitutionControlMain(db, &institution.ControlMain{
		Control: institution.Control{
			InsIdCd: info.InsIdCd,
		},
	}, 1, 9999)
	if err != nil {
		return err
	}

	// 入编辑表 institution fee cash control
	err = institution.SaveInstitution(db, &info.InstitutionInfo)
	if err != nil {
		return err
	}
	for _, fee := range fees {
		err = institution.SaveInstitutionFee(db, &fee.Fee)
		if err != nil {
			return err
		}
	}
	for _, cash := range cashes {
		err = institution.SaveInstitutionCash(db, &cash.Cash)
		if err != nil {
			return err
		}
	}
	for _, control := range controls {
		err = institution.SaveInstitutionControl(db, &control.Control)
		if err != nil {
			return err
		}
	}
	return nil
}
