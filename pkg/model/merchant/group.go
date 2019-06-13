package merchant

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Group struct {
	JtMchtCd       string    `gorm:"column:JT_MCHT_CD"`
	JtMchtNm       string    `gorm:"column:JT_MCHT_NM"`
	JtArea         string    `gorm:"column:JT_AREA"`
	MchtStlmCNm    string    `gorm:"column:MCHT_STLM_C_NM"`
	MchtStlmCAcct  string    `gorm:"column:MCHT_STLM_C_ACCT"`
	ChtStlmInsIdCd string    `gorm:"column:CHT_STLM_INS_ID_CD"`
	MchtStlmInsNm  string    `gorm:"column:MCHT_STLM_INS_NM"`
	MchtPaySysAcct string    `gorm:"column:MCHT_PAY_SYS_ACCT"`
	ProvCd         string    `gorm:"column:PROV_CD"`
	CityCd         string    `gorm:"column:CITY_CD"`
	AipBranCd      string    `gorm:"column:AIP_BRAN_CD"`
	SystemFlag     string    `gorm:"column:SYSTEMFLAG"`
	Status         string    `gorm:"column:STATUS"`
	UpdatedAt      time.Time `gorm:"column:REC_UPD_TS"`
	CreatedAt      time.Time `gorm:"column:REC_CRT_TS"`
	RecOprId       string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr      string    `gorm:"column:REC_UPD_OPR"`
	JtAddr         string    `gorm:"column:JT_ADDR"`
}

func (g Group) TableName() string {
	return "TBL_EDIT_GROUP_MCHT_INF"
}

func QueryMerchantGroups(db *gorm.DB, query *Group, page int32, size int32) ([]*Group, int32, error) {
	out := make([]*Group, 0)
	var count int32
	db.Model(&Group{}).Where(query).Count(&count)
	err := db.Where(query).Offset((page - 1) * size).Limit(size).Find(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, count, nil
	}
	return out, count, err
}
