package static

import "github.com/jinzhu/gorm"

const (
	orgDictionaryItem   = "TBL_ORG_DICTIONARYITEM"
	dictionaryLayerItem = "TBL_DICTIONARYLAYERITEM"
)

type Area struct {
	Name    string `gorm:"column:NAME"`
	DicCode string `gorm:"column:DIC_CODE"`
}

func FindProvince(db *gorm.DB, orgCode string) ([]*Area, error) {
	out := make([]*Area, 0)
	err := db.Table(dictionaryLayerItem).
		Select("NAME, DIC_CODE").
		Joins("left join TBL_ORG_DICTIONARYITEM b on b.ITEM_CODE = DIC_CODE").
		Where(
			"ORG_CODE = ? and TYPE_CODE = ? and DIC_TYPE = ? and DIC_LEVEL = ?",
			orgCode,
			"UnionPay_Bank_32Area",
			"UnionPay_Area",
			"1",
		).Order("DIC_CODE", true).Scan(&out).Error
	return out, err
}

func FindCity(db *gorm.DB, dicCode string, level string) ([]*Area, error) {
	out := make([]*Area, 0)
	err := db.Table(dictionaryLayerItem).
		Select("NAME, DIC_CODE").
		Where("DIC_PCODE = ? and DIC_LEVEL = ? and dic_type = ?",
			dicCode,
			level,
			"UnionPay_Area",
		).
		Scan(&out).Error
	return out, err
}
