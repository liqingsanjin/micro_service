package static

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BankList struct {
	Id        int64     `gorm:"column:ID;primary_key;auto_increment"`
	Code      string    `gorm:"column:CODE;index"`
	Name      string    `gorm:"column:NAME"`
	UpdatedAt time.Time `gorm:"column:UPDATE_DATE"`
}

func (BankList) TableName() string {
	return "TBL_UNIONPAY_BANKLIST"
}

type Mcc struct {
	Id           int64     `gorm:"column:ID;primary_key;auto_increment"`
	Code         string    `gorm:"column:CODE"`
	Name         string    `gorm:"column:NAME"`
	Category     string    `gorm:"column:CATEGORY"`
	CategoryType string    `gorm:"column:CATEGORY_TYPE"`
	Industry     string    `gorm:"column:INDUSTRY"`
	Status       string    `gorm:"column:STATUS"`
	UpdatedAt    time.Time `gorm:"column:UPDATE_TIME"`
}

func (Mcc) TableName() string {
	return "TBL_UNIONPAY_MCC"
}

func FindBankListByCode(db *gorm.DB, code string, page int32, size int32) ([]*BankList, int32, error) {
	out := make([]*BankList, 0)
	var count int32
	db.Model(&BankList{}).Where("CODE like ?", code+"%").Count(&count)
	err := db.Where("CODE like ?", code+"%").Offset((page - 1) * size).Limit(size).Find(&out).Error
	return out, count, err
}

func QueryMcc(db *gorm.DB, query *Mcc, page int32, size int32) ([]*Mcc, int32, error) {
	out := make([]*Mcc, 0)
	var count int32
	db.Model(&Mcc{}).Where(query).Count(&count)
	err := db.Where(query).Offset((page - 1) * size).Limit(size).Find(&out).Error
	return out, count, err
}
