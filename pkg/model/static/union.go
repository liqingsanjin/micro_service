package static

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BankList struct {
	Id        int64     `gorm:"column:ID"`
	Code      string    `gorm:"column:CODE"`
	Name      string    `gorm:"column:NAME"`
	UpdatedAt time.Time `gorm:"column:UPDATE_DATE"`
}

func (BankList) TableName() string {
	return "TBL_UNIONPAY_BANKLIST"
}

func FindBankListByCode(db *gorm.DB, code string, page int32, size int32) ([]*BankList, int32, error) {
	out := make([]*BankList, 0)
	var count int32
	db.Model(&BankList{}).Where("CODE like ?", code+"%").Count(&count)
	err := db.Where("CODE like ?", code+"%").Offset((page - 1) * size).Limit(size).Find(&out).Error
	return out, count, err
}
