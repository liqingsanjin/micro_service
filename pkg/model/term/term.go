package term

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Info struct {
	MchtCd          string    `gorm:"column:MCHT_CD"`
	TermId          string    `gorm:"column:TERM_ID"`
	TermTp          string    `gorm:"column:TERM_TP"`
	Belong          string    `gorm:"column:BELONG"`
	BelongSub       string    `gorm:"column:BELONG_SUB"`
	TmnlMoneyIntype string    `gorm:"column:TMNL_MONEY_INTYPE"`
	TmnlMoney       float64   `gorm:"column:TMNL_MONEY"`
	TmnlBrand       string    `gorm:"column:TMNL_BRAND"`
	TmnlModelNo     string    `gorm:"column:TMNL_MODEL_NO"`
	TmnlBarcode     string    `gorm:"column:TMNL_BARCODE"`
	DeviceCd        string    `gorm:"column:DEVICE_CD"`
	InstallLocation string    `gorm:"column:INSTALLLOCATION"`
	TmnlIntype      string    `gorm:"column:TMNL_INTYPE"`
	DialOut         string    `gorm:"column:DIAL_OUT"`
	DealTypes       string    `gorm:"column:DEAL_TYPES"`
	RecOprId        string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr       string    `gorm:"column:REC_UPD_OPR"`
	AppCd           string    `gorm:"column:APP_CD"`
	SystemFlag      string    `gorm:"column:SYSTEMFLAG"`
	Status          string    `gorm:"column:STATUS"`
	ActiveCode      string    `gorm:"column:ACTIVE_CODE"`
	NoFlag          string    `gorm:"column:NO_FLAG"`
	MsgResvFld1     string    `gorm:"column:MSG_RESV_FLD1"`
	MsgResvFld2     string    `gorm:"column:MSG_RESV_FLD2"`
	MsgResvFld3     string    `gorm:"column:MSG_RESV_FLD3"`
	MsgResvFld4     string    `gorm:"column:MSG_RESV_FLD4"`
	MsgResvFld5     string    `gorm:"column:MSG_RESV_FLD5"`
	MsgResvFld6     string    `gorm:"column:MSG_RESV_FLD6"`
	MsgResvFld7     string    `gorm:"column:MSG_RESV_FLD7"`
	MsgResvFld8     string    `gorm:"column:MSG_RESV_FLD8"`
	MsgResvFld9     string    `gorm:"column:MSG_RESV_FLD9"`
	MsgResvFld10    string    `gorm:"column:MSG_RESV_FLD10"`
	CreatedAt       time.Time `gorm:"column:REC_CRT_TS"`
	UpdatedAt       time.Time `gorm:"column:REC_UPD_TS"`
}

func (Info) TableName() string {
	return "TBL_EDIT_TERM_INF"
}

type InfoMain struct {
	Info
}

func (InfoMain) TableName() string {
	return "TBL_TERM_INF"
}

func QueryTermInfo(db *gorm.DB, query *Info, page int32, size int32) ([]*Info, int32, error) {
	out := make([]*Info, 0)
	var count int32
	db.Model(&Info{}).Where(query).Count(&count)
	err := db.Where(query).Offset((page - 1) * size).Limit(size).Find(&out).Error
	return out, count, err
}

func SaveTermInfo(db *gorm.DB, data *Info) error {
	return db.Create(data).Error
}

func SaveTermInfoMain(db *gorm.DB, data *InfoMain) error {
	return db.Create(data).Error
}

func QueryTermInfoMain(db *gorm.DB, query *InfoMain, page int32, size int32) ([]*InfoMain, int32, error) {
	out := make([]*InfoMain, 0)
	var count int32
	db.Model(&InfoMain{}).Where(query).Count(&count)
	err := db.Where(query).Offset((page - 1) * size).Limit(size).Find(&out).Error
	return out, count, err
}
