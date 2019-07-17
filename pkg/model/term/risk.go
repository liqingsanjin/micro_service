package term

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Risk struct {
	MchtCd           string    `gorm:"column:MCHT_CD"`
	TermId           string    `gorm:"column:TERM_ID"`
	CardType         string    `gorm:"column:CARD_TYPE"`
	TotalLimitMoney  float64   `gorm:"column:TOTAL_LIMITMONEY"`
	AccpetStartTime  string    `gorm:"column:ACCPET_START_TIME"`
	AccpetStartDate  string    `gorm:"column:ACCPET_START_DATE"`
	AccpetEndTime    string    `gorm:"column:ACCPET_END_TIME"`
	AccpetEndDate    string    `gorm:"column:ACCPET_END_DATE"`
	SingleLimitMoney float64   `gorm:"column:SINGLE_LIMITMONEY"`
	ControlWay       string    `gorm:"column:CONTROL_WAY"`
	SingleMinMoney   float64   `gorm:"column:SINGLE_MIN_MONEY"`
	TotalPeriod      string    `gorm:"column:TOTAL_PERIOD"`
	RecOprId         string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr        string    `gorm:"column:REC_UPD_OPR"`
	OperIn           string    `gorm:"column:OPER_IN"`
	CreatedAt        time.Time `gorm:"column:REC_CRT_TS"`
	UpdatedAt        time.Time `gorm:"column:REC_UPD_TS"`
}

func (Risk) TableName() string {
	return "TBL_EDIT_TERM_RISK_CFG"
}

type RiskMain struct {
	Risk
}

func (RiskMain) TableName() string {
	return "TBL_TERM_RISK_CFG"
}

func SaveRisk(db *gorm.DB, data *Risk) error {
	return db.Create(data).Error
}
