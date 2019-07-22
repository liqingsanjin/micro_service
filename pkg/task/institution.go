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
	fees, err := institution.FindInstitutionFee(db, &institution.Fee{
		InsIdCd: info.InsIdCd,
	})
	if err != nil {
		return err
	}
	cashes, err := institution.FindInstitutionCash(db, &institution.Cash{
		InsIdCd: info.InsIdCd,
	})
	if err != nil {
		return err
	}
	controls, err := institution.FindInstitutionControl(db, &institution.Control{
		InsIdCd: info.InsIdCd,
	})
	if err != nil {
		return err
	}

	// 修改机构状态
	err = institution.UpdateInstitution(db, &institution.InstitutionInfo{
		InsIdCd: info.InsIdCd,
	}, &institution.InstitutionInfo{
		InsSta: "1",
	})
	if err != nil {
		return err
	}

	// 入正式表 institution fee cash control
	info.InsSta = "1"

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
