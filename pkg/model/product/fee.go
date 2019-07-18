package product

type Fee struct {
	ProdCd      string `gorm:"column:PROD_CD"`
	BizCd       string `gorm:"column:BIZ_CD"`
	SubBizCd    string `gorm:"column:SUB_BIZ_CD"`
	UpdateDate  string `gorm:"column:UPDATE_DATE"`
	Description string `gorm:"column:DESCRIPTION"`
	ResvFld1    string `gorm:"column:RESV_FLD1"`
	ResvFld2    string `gorm:"column:RESV_FLD2"`
	ResvFld3    string `gorm:"column:RESV_FLD3"`
}

func (Fee) TableName() string {
	return "TBL_PROD_BIZ_FEE_MAP"
}
