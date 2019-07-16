package merchant

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BizFee struct {
	MchtCd    string    `gorm:"column:MCHT_CD"`
	ProdCd    string    `gorm:"column:PROD_CD"`
	BizCd     string    `gorm:"column:BIZ_CD"`
	TransCd   string    `gorm:"column:TRANS_CD"`
	OperIn    string    `gorm:"column:OPER_IN"`
	RecOprId  string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr string    `gorm:"column:REC_UPD_OPR"`
	CreatedAt time.Time `gorm:"column:REC_CRT_TS"`
	UpdateAt  time.Time `gorm:"column:REC_UPD_TS"`
}

type BizFeeMain struct {
	BizFee
}

func (BizFee) TableName() string {
	return "TBL_EDIT_MCHT_BIZ_FEE"
}

func (BizFeeMain) TableName() string {
	return "TBL_MCHT_BIZ_FEE"
}

func SaveBizFee(db *gorm.DB, data *BizFee) error {
	return db.Create(data).Error
}
