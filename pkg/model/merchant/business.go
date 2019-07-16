package merchant

import "github.com/jinzhu/gorm"

type Business struct {
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
