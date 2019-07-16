package merchant

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Business struct {
	MchtCd    string    `gorm:"column:MCHT_CD"`
	ProdCd    string    `gorm:"column:PROD_CD"`
	BizCd     string    `gorm:"column:BIZ_CD"`
	TransCd   string    `gorm:"column:TRANS_CD"`
	OperIn    string    `gorm:"column:OPER_IN"`
	RecOprId  string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr string    `gorm:"column:REC_UPD_OPR"`
	CreatedAt time.Time `gorm:"column:REC_CRT_TS"`
	UpdatedAt time.Time `gorm:"column:REC_UPD_TS"`
}

type BusinessMain struct {
	Business
}

func (Business) TableName() string {
	return "TBL_EDIT_MCHT_BUSINESS"
}

func (BusinessMain) TableName() string {
	return "TBL_MCHT_BUSINESS"
}

func SaveBusiness(db *gorm.DB, data *Business) error {
	return db.Create(data).Error
}
