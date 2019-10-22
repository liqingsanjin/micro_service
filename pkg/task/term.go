package task

import (
	"fmt"
	"strings"
	camundamodel "userService/pkg/model/camunda"
	"userService/pkg/model/merchant"
	"userService/pkg/model/term"

	"github.com/jinzhu/gorm"
)

func termAdd(db *gorm.DB, instance *camundamodel.ProcessInstance) error {
	strs := strings.Split(instance.DataId, ",")
	if len(strs) == 2 {
		mchtCd := strs[0]
		termId := strs[1]

		// 查询终端信息
		info, err := term.FindTermByPk(db, mchtCd, termId)
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("term edit %s not found", instance.DataId)
		}

		err = term.UpdateTerm(db, &term.Info{MchtCd: mchtCd, TermId: termId}, &term.Info{Status: "01", SystemFlag: "01"})
		if err != nil {
			return err
		}

		info.Status = "01"
		info.SystemFlag = "01"
		err = term.SaveTermInfoMain(db, &term.InfoMain{Info: *info})
		if err != nil {
			return err
		}

		err = merchant.UpdateMerchant(db, &merchant.MerchantInfo{MchtCd: mchtCd}, &merchant.MerchantInfo{SystemFlag: "01"})

		return err
	} else {
		return fmt.Errorf("dataId 错误: %s", instance.DataId)
	}
}

func termDelete(db *gorm.DB, instance *camundamodel.ProcessInstance) error {
	strs := strings.Split(instance.DataId, ",")
	if len(strs) == 2 {
		mchtCd := strs[0]
		termId := strs[1]
		info, err := term.FindTermByPk(db, mchtCd, termId)
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("term edit %s not found", instance.DataId)
		}

		err = term.DeleteTerm(db, &term.Info{MchtCd: info.MchtCd, TermId: info.TermId})
		if err != nil {
			return err
		}

		err = merchant.UpdateMerchant(db, &merchant.MerchantInfo{MchtCd: mchtCd}, &merchant.MerchantInfo{SystemFlag: "01"})

		return err
	} else {
		return fmt.Errorf("dataId 错误: %s", instance.DataId)
	}
}
