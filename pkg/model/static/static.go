package static

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	tableDictItem            = "TBL_DICTIONARYITEM"
	tableInsProdBizFeeMapInf = "TBL_INS_PROD_BIZ_FEE_MAP_INF"
	tableProdBizTransMap     = "TBL_PROD_BIZ_TRANS_MAP"
	tableInsInf              = "TBL_INS_INF"
)

type DictionaryItem struct {
	DicType    string    `gorm:"column:DIC_TYPE" json:"DicType"`
	DicCode    string    `gorm:"column:DIC_CODE" json:"DicCode"`
	DicName    string    `gorm:"column:DIC_NAME" json:"DicName"`
	DispOrder  string    `gorm:"column:DISP_ORDER" json:"DispOrder"`
	UpdateTime time.Time `gorm:"column:UPDATE_TIME" json:"UpdateTime"`
	Memo       string    `gorm:"column:MEMO" json:"Memo"`
}

func (d DictionaryItem) TableName() string {
	return tableDictItem
}

type InsProdBizFeeMapInf struct {
	ProdCd       string    `gorm:"column:PROD_CD"`
	BizCd        string    `gorm:"column:BIZ_CD"`
	MccMTp       string    `gorm:"column:MCC_M_TP"`
	MccSTp       string    `gorm:"column:MCC_S_TP"`
	InsFeeBizCd  string    `gorm:"column:INS_FEE_BIZ_CD"`
	InsFeeBizNm  string    `gorm:"column:INS_FEE_BIZ_NM"`
	MsgResvFld1  string    `gorm:"column:MSG_RESV_FLD1"`
	MsgResvFld2  string    `gorm:"column:MSG_RESV_FLD2"`
	MsgResvFld3  string    `gorm:"column:MSG_RESV_FLD3"`
	MsgResvFld4  string    `gorm:"column:MSG_RESV_FLD4"`
	MsgResvFld5  string    `gorm:"column:MSG_RESV_FLD5"`
	MsgResvFld6  string    `gorm:"column:MSG_RESV_FLD6"`
	MsgResvFld7  string    `gorm:"column:MSG_RESV_FLD7"`
	MsgResvFld8  string    `gorm:"column:MSG_RESV_FLD8"`
	MsgResvFld9  string    `gorm:"column:MSG_RESV_FLD9"`
	MsgResvFld10 string    `gorm:"column:MSG_RESV_FLD10"`
	RecOprID     string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr    string    `gorm:"column:REC_UPD_OPR"`
	RecCrtTs     time.Time `gorm:"column:REC_CRT_TS"`
	RecUpdTs     time.Time `gorm:"column:REC_UPD_TS"`
}

func (i InsProdBizFeeMapInf) TableName() string {
	return tableInsProdBizFeeMapInf
}

type ProdBizTransMap struct {
	ProdCd      string `gorm:"column:PROD_CD"`
	BizCd       string `gorm:"column:BIZ_CD"`
	TransCd     string `gorm:"column:TRANS_CD"`
	UpdateDate  string `gorm:"column:UPDATE_DATE"`
	Description string `gorm:"column:DESCRIPTION"`
	ResvFld1    string `gorm:"column:RESV_FLD1"`
	ResvFld2    string `gorm:"column:RESV_FLD2"`
	ResvFld3    string `gorm:"column:RESV_FLD3"`
}

func (p ProdBizTransMap) TableName() string {
	return tableProdBizTransMap
}

type InsInf struct {
	InsIDCd         string    `gorm:"column:INS_ID_CD"`
	InsCompanyCd    string    `gorm:"column:INS_COMPANY_CD"`
	InsType         string    `gorm:"column:INS_TYPE"`
	InsName         string    `gorm:"column:INS_NAME"`
	InsProvCd       string    `gorm:"column:INS_PROV_CD"`
	InsCityCd       string    `gorm:"column:INS_CITY_CD"`
	InsRegionCd     string    `gorm:"column:INS_REGION_CD"`
	InsSta          string    `gorm:"column:INS_STA"`
	InsStlmTp       string    `gorm:"column:INS_STLM_TP"`
	InsAloStlmCycle string    `gorm:"column:INS_ALO_STLM_CYCLE"`
	InsAloStlmMd    string    `gorm:"column:INS_ALO_STLM_MD"`
	InsStlmCNm      string    `gorm:"column:INS_STLM_C_NM"`
	InsStlmCAcct    string    `gorm:"column:INS_STLM_C_ACCT"`
	InsStlmCBkNo    string    `gorm:"column:INS_STLM_C_BK_NO"`
	InsStlmCBkNm    string    `gorm:"column:INS_STLM_C_BK_NM"`
	InsStlmDNm      string    `gorm:"column:INS_STLM_D_NM"`
	InsStlmDAcct    string    `gorm:"column:INS_STLM_D_ACCT"`
	InsStlmDBkNo    string    `gorm:"column:INS_STLM_D_BK_NO"`
	InsStlmDBkNm    string    `gorm:"column:INS_STLM_D_BK_NM"`
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
	RecOprID        string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr       string    `gorm:"column:REC_UPD_OPR"`
	RecCrtTs        time.Time `gorm:"column:REC_CRT_TS"`
	RecUpdTs        time.Time `gorm:"column:REC_UPD_TS"`
}

func (p InsInf) TableName() string {
	return tableInsInf
}

func GetDictionaryItem(db *gorm.DB) []*DictionaryItem {
	dictionaryItem := make([]*DictionaryItem, 0)
	db.Debug().Find(&dictionaryItem)
	return dictionaryItem
}

func GetInsProdBizFeeMapInf(db *gorm.DB) []*InsProdBizFeeMapInf {
	insProdBizFeeMapInf := make([]*InsProdBizFeeMapInf, 0)
	db.Debug().Find(&insProdBizFeeMapInf)
	return insProdBizFeeMapInf
}

func GetProdBizTransMap(db *gorm.DB) []*ProdBizTransMap {
	prodBizTransMap := make([]*ProdBizTransMap, 0)
	db.Debug().Find(&prodBizTransMap)
	return prodBizTransMap
}

func GetInsInf(db *gorm.DB) []*InsInf {
	insInf := make([]*InsInf, 0)
	db.Debug().Find(&insInf)
	return insInf
}
