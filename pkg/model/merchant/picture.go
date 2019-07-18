package merchant

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Picture struct {
	FileId     string    `gorm:"column:FILE_ID"`
	MchtCd     string    `gorm:"column:MCHT_CD"`
	DocType    string    `gorm:"column:DOC_TYPE"`
	FileType   string    `gorm:"column:FILE_TYPE"`
	FileName   string    `gorm:"column:FILE_NAME"`
	PIndex     int64     `gorm:"column:PINDEX"`
	PCode      string    `gorm:"column:PCODE"`
	Url        string    `gorm:"column:URL"`
	SystemFlag string    `gorm:"column:SYSTEMFLAG"`
	Status     string    `gorm:"column:STATUS"`
	RecOprId   string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr  string    `gorm:"column:REC_UPD_OPR"`
	CreatedAt  time.Time `gorm:"column:REC_CRT_TS"`
	UpdatedAt  time.Time `gorm:"column:REC_UPD_TS"`
}

type PictureMain struct {
	Picture
}

func (Picture) TableName() string {
	return "TBL_EDIT_MCHT_PICTURE"
}

func (PictureMain) TableName() string {
	return "TBL_MCHT_PICTURE"
}

func SavePicture(db *gorm.DB, data *Picture) error {
	return db.Create(data).Error
}

func QueryPicture(db *gorm.DB, query *Picture) ([]*Picture, error) {
	out := make([]*Picture, 0)
	err := db.Where(query).Find(&out).Error
	return out, err
}

func QueryPicTureMain(db *gorm.DB, query *PictureMain) ([]*PictureMain, error) {
	out := make([]*PictureMain, 0)
	err := db.Where(query).Find(&out).Error
	return out, err
}
