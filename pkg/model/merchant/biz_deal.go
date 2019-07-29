package merchant

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BizDeal struct {
	MchtCd    string    `gorm:"column:MCHT_CD;primary_key"`
	ProdCd    string    `gorm:"column:PROD_CD;primary_key"`
	BizCd     string    `gorm:"column:BIZ_CD;primary_key"`
	TransCd   string    `gorm:"column:TRANS_CD;primary_key"`
	OperIn    string    `gorm:"column:OPER_IN"`
	RecOprId  string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr string    `gorm:"column:REC_UPD_OPR"`
	CreatedAt time.Time `gorm:"column:REC_CRT_TS"`
	UpdatedAt time.Time `gorm:"column:REC_UPD_TS"`
}

type BizDealMain struct {
	BizDeal
}

func (BizDeal) TableName() string {
	return "TBL_EDIT_MCHT_BIZ_DEAL"
}

func (BizDealMain) TableName() string {
	return "TBL_MCHT_BIZ_DEAL"
}

func SaveBizDeal(db *gorm.DB, data *BizDeal) error {
	return db.Save(data).Error
}

func SaveBizDealMain(db *gorm.DB, data *BizDealMain) error {
	return db.Save(data).Error
}

func QueryBizDeal(db *gorm.DB, query *BizDeal) ([]*BizDeal, error) {
	out := make([]*BizDeal, 0)
	err := db.Where(query).Find(&out).Error
	return out, err
}

func QueryBizDealMain(db *gorm.DB, query *BizDealMain) ([]*BizDealMain, error) {
	out := make([]*BizDealMain, 0)
	err := db.Where(query).Find(&out).Error
	return out, err
}

func DeleteBizDeal(db *gorm.DB, query *BizDeal) error {
	return db.Where(query).Delete(&BizDeal{}).Error
}
