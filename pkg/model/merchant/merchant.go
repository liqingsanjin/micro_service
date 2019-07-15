package merchant

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	TableMerchantInfo        = "TBL_EDIT_MCHT_INF"
	TableMerchantBankAccount = "TBL_EDIT_MCHT_BANKACCOUNT"
)

type MerchantInfo struct {
	MchtCd                string     `gorm:"column:MCHT_CD;primary_key"`
	Sn                    string     `gorm:"column:SN"`
	AipBranCd             string     `gorm:"column:AIP_BRAN_CD;index"`
	GroupCd               string     `gorm:"column:GROUP_CD"`
	OriChnl               string     `gorm:"column:ORI_CHNL"`
	OriChnlDesc           string     `gorm:"column:ORI_CHNL_DESC"`
	BankBelongCd          string     `gorm:"column:BANK_BELONG_CD"`
	DvpBy                 string     `gorm:"column:DVP_BY"`
	MccCd18               string     `gorm:"column:MCC_CD_18"`
	ApplDate              string     `gorm:"column:APPL_DATE"`
	UpBcCd                string     `gorm:"column:UP_BC_CD"`
	UpAcCd                string     `gorm:"column:UP_AC_CD"`
	UpMccCd               string     `gorm:"column:UP_MCC_CD"`
	Name                  string     `gorm:"column:NAME"`
	NameBusi              string     `gorm:"column:NAME_BUSI"`
	BusiLiceNo            string     `gorm:"column:BUSI_LICE_NO"`
	BusiRang              string     `gorm:"column:BUSI_RANG"`
	BusiMain              string     `gorm:"column:BUSI_MAIN"`
	Certif                string     `gorm:"column:CERTIF"`
	CertifType            string     `gorm:"column:CERTIF_TYPE"`
	CertifNo              string     `gorm:"column:CERTIF_NO"`
	CityCd                string     `gorm:"column:CITY_CD"`
	AreaCd                string     `gorm:"column:AREA_CD"`
	RegAddr               string     `gorm:"column:REG_ADDR"`
	ContactName           string     `gorm:"column:CONTACT_NAME"`
	ContactPhoneNo        string     `gorm:"column:CONTACT_PHONENO"`
	IsGroup               string     `gorm:"column:ISGROUP"`
	MoneyToGroup          string     `gorm:"column:MONEYTOGROUP"`
	StlmWay               string     `gorm:"column:STLM_WAY"`
	StlmWayDesc           string     `gorm:"column:STLM_WAY_DESC"`
	StlmInsCircle         string     `gorm:"column:STLM_INS_CIRCLE"`
	ApprDate              *time.Time `gorm:"column:APPR_DATE"`
	Status                string     `gorm:"column:STATUS"`
	DeleteDate            *time.Time `gorm:"column:DELETE_DATE"`
	UcBcCd32              string     `gorm:"column:UC_BC_CD_32"`
	K2WorkflowId          string     `gorm:"column:K2WORKFLOWID"`
	SystemFlag            string     `gorm:"column:SYSTEMFLAG"`
	ApprovalUsername      string     `gorm:"column:APPROVALUSERNAME"`
	FinalApprovalUsername string     `gorm:"column:FINALARRPOVALUSERNAME"`
	IsUpStandard          string     `gorm:"column:IS_UP_STANDARD"`
	BillingType           string     `gorm:"column:BILLINGTYPE"`
	BillingLevel          string     `gorm:"column:BILLINGLEVEL"`
	Slogan                string     `gorm:"column:SLOGAN"`
	Ext1                  string     `gorm:"column:EXT1"`
	Ext2                  string     `gorm:"column:EXT2"`
	Ext3                  string     `gorm:"column:EXT3"`
	Ext4                  string     `gorm:"column:EXT4"`
	AreaStandard          string     `gorm:"column:AREA_STANDARD"`
	MchtCdAreaCd          string     `gorm:"column:MCHTCD_AREA_CD"`
	UcBcCdArea            string     `gorm:"column:UC_BC_CD_AREA"`
	RecOprId              string     `gorm:"column:REC_OPR_ID"`
	RecUpdOpr             string     `gorm:"column:REC_UPD_OPR"`
	CreatedAt             time.Time  `gorm:"column:REC_CRT_TS"`
	UpdatedAt             time.Time  `gorm:"column:REC_UPD_TS"`
	OperIn                string     `gorm:"column:OPER_IN"`
	RecApllyTs            *time.Time `gorm:"column:REC_APLLY_TS"`
	OemOrgCode            string     `gorm:"column:OEM_ORG_CODE"`
	IsEleInvoice          string     `gorm:"column:IS_ELE_INVOICE"`
	DutyParagraph         string     `gorm:"column:DUTY_PARAGRAPH"`
	TaxMachineBrand       string     `gorm:"column:TAX_MACHINE_BRAND"`
	Ext5                  string     `gorm:"column:EXT5"`
	Ext6                  string     `gorm:"column:EXT6"`
	Ext7                  string     `gorm:"column:EXT7"`
	Ext8                  string     `gorm:"column:EXT8"`
	Ext9                  string     `gorm:"column:EXT9"`
	BusiLiceSt            string     `gorm:"column:BUSI_LICE_ST"`
	BusiLiceDt            string     `gorm:"column:BUSI_LICE_DT"`
	CertifSt              string     `gorm:"column:CERTIF_ST"`
	CertifDt              string     `gorm:"column:CERTIF_DT"`
}

type MerchantInfoMain struct {
	MerchantInfo
}

func (m MerchantInfoMain) TableName() string {
	return "TBL_MCHT_INF"
}

func (m MerchantInfo) TableName() string {
	return TableMerchantInfo
}

type MerchantBankAccount struct {
	OwnerCd      string    `gorm:"column:OWNER_CD;primary_key"`
	AccountType  string    `gorm:"column:ACCOUNTTYPE;primary_key"`
	Name         string    `gorm:"column:NAME"`
	Account      string    `gorm:"column:ACCOUNT"`
	UcBcCd       string    `gorm:"column:UC_BC_CD"`
	Province     string    `gorm:"column:PROVINCE"`
	City         string    `gorm:"column:CITY"`
	BankCode     string    `gorm:"column:BANK_CODE"`
	BankName     string    `gorm:"column:BANK_NAME"`
	OperIn       string    `gorm:"column:OPER_IN"`
	RecOprId     string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr    string    `gorm:"column:REC_UPD_OPR"`
	RecCrtTs     time.Time `gorm:"column:REC_CRT_TS"`
	RecUpdTs     time.Time `gorm:"column:REC_UPD_TS"`
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
}

func (m MerchantBankAccount) TableName() string {
	return TableMerchantBankAccount
}

func FindMerchantInfoById(db *gorm.DB, id string) (*MerchantInfo, error) {
	out := new(MerchantInfo)
	err := db.Where("MCHT_CD = ?", id).First(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return out, err
}

func QueryMerchantInfos(db *gorm.DB, query *MerchantInfo, page int32, size int32) ([]*MerchantInfo, int32, error) {
	out := make([]*MerchantInfo, 0)
	var count int32
	db.Model(&MerchantInfo{}).Where(query).Count(&count)
	err := db.Where(query).Offset((page - 1) * size).Limit(size).Find(&out).Error
	return out, count, err
}

func QueryMerchantInfosMain(db *gorm.DB, query *MerchantInfoMain, page int32, size int32) ([]*MerchantInfoMain, int32, error) {
	out := make([]*MerchantInfoMain, 0)
	var count int32
	db.Model(&MerchantInfoMain{}).Where(query).Count(&count)
	err := db.Where(query).Offset((page - 1) * size).Limit(size).Find(&out).Error
	return out, count, err
}

func SaveMerchant(db *gorm.DB, info *MerchantInfo) error {
	return db.Create(info).Error
}
