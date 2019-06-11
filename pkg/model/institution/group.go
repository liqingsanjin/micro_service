package institution

import (
	"time"

	"github.com/jinzhu/gorm"
)

type InstitutionGroup struct {
	InsGroup  string    `gorm:"column:INS_GROUP;primary_key"`
	InsIdCd   string    `gorm:"column:INS_ID_CD;primary_key"`
	CreatedAt time.Time `gorm:"column:REC_CRT_TS"`
	UpdatedAt time.Time `gorm:"column:REC_UPD_TS"`
}

func (i InstitutionGroup) TableName() string {
	return "TBL_INS_GROUP"
}

func ListGroups(db *gorm.DB) ([]*InstitutionGroup, error) {
	gs := make([]*InstitutionGroup, 0)
	err := db.Select("INS_GROUP").Group("INS_GROUP").Find(&gs).Error
	if err == gorm.ErrRecordNotFound {
		return gs, nil
	}
	return gs, err
}
