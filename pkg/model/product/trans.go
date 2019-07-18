package product

import "github.com/jinzhu/gorm"

type Trans struct {
	ProdCd      string `gorm:"cloumn:PROD_CD"`
	BizCd       string `gorm:"cloumn:BIZ_CD"`
	TransCd     string `gorm:"cloumn:TRANS_CD"`
	UpdateDate  string `gorm:"cloumn:UPDATE_DATE"`
	Description string `gorm:"cloumn:DESCRIPTION"`
	ResvFld1    string `gorm:"cloumn:RESV_FLD1"`
	ResvFld2    string `gorm:"cloumn:RESV_FLD2"`
	ResvFld3    string `gorm:"cloumn:RESV_FLD3"`
}

func (Trans) TableName() string {
	return "TBL_PROD_BIZ_TRANS_MAP"
}

func ListTrans(db *gorm.DB, query *Trans) ([]*Trans, error) {
	out := make([]*Trans, 0)
	err := db.Where(query).Find(&out).Error
	return out, err
}
