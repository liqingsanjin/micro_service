package cleartxn

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"
)

const (
	clearTxnName = "TBL_CLEAR_TXN"
)

type ClearTxn struct {
	CompanyCd        string    `gorm:"column:COMPANY_CD"`
	InsIDCd          string    `gorm:"column:INS_ID_CD"`
	AcqInsIDCd       string    `gorm:"column:ACQ_INS_ID_CD"`
	FwdInsIDCd       string    `gorm:"column:FWD_INS_ID_CD"`
	MchtCd           string    `gorm:"column:MCHT_CD"`
	MchtName         string    `gorm:"column:MCHT_NAME"`
	MchtShortName    string    `gorm:"column:MCHT_SHORT_NAME"`
	MccCd            string    `gorm:"column:MCC_CD"`
	MccCd42          string    `gorm:"column:MCC_CD_42"`
	MccDesc          string    `gorm:"column:MCC_DESC"`
	TransDateTime    string    `gorm:"column:TRANS_DATE_TIME"`
	StlmDate         string    `gorm:"column:STLM_DATE"`
	TransKind        string    `gorm:"column:TRANS_KIND"`
	TxnDesc          string    `gorm:"column:TXN_DESC"`
	TransState       string    `gorm:"column:TRANS_STATE`
	StlmFlg          string    `gorm:"column:STLM_FLG`
	TransAmt         string    `gorm:"column:TRANS_AMT"`
	Creditcardlimit  string    `gorm:"column:CREDITCARDLIMIT"`
	CupSsn           string    `gorm:"column:CUP_SSN"`
	AuthIDResp       string    `gorm:"column:AUTHR_ID_RESP"`
	Pan              string    `gorm:"column:PAN"`
	CardKindDis      string    `gorm:"column:CARD_KIND_DIS"`
	BankCode         string    `gorm:"column:BANK_CODE"`
	BankName         string    `gorm:"column:BANK_NAME"`
	BranchCd         string    `gorm:"column:BRANCH_CD"`
	BranchNm         string    `gorm:"column:BRANCH_NM"`
	TermID           string    `gorm:"column:TERM_ID"`
	OrgTransDateTime string    `gorm:"column:ORG_TRANS_DATE_TIME"`
	OrgCupSsn        string    `gorm:"column:ORG_CUP_SSN"`
	PosEntryMode     string    `gorm:"column:POS_ENTRY_MODE"`
	RspCode          string    `gorm:"column:RSP_CODE"`
	TrueFeeMod       string    `gorm:"column:TRUE_FEE_MOD"`
	TrueFeeBi        string    `gorm:"column:TRUE_FEE_BI"`
	TrueFeeFd        string    `gorm:"column:TRUE_FEE_FD"`
	TrueFeeFfd       string    `gorm:"column:TRUE_FEE_FFD"`
	Var1             string    `gorm:"column:VAR_1"`
	Var2             string    `gorm:"column:VAR_2"`
	Var3             string    `gorm:"column:VAR_3"`
	Var4             string    `gorm:"column:VAR_4"`
	VirFeeMod        string    `gorm:"column:VIR_FEE_MOD"`
	VirFeeBi         string    `gorm:"column:VIR_FEE_BI"`
	VirFeeBd         string    `gorm:"column:VIR_FEE_BD"`
	VirFeeFd         string    `gorm:"column:VIR_FEE_FD"`
	MchtFee          string    `gorm:"column:MCHT_FEE"`
	Var5             string    `gorm:"column:VAR_5"`
	MchtVirFee       string    `gorm:"column:MCHT_VIR_FEE"`
	StandBankFee     string    `gorm:"column:STAND_BANK_FEE"`
	BankFee          string    `gorm:"column:BANK_FEE"`
	HzjgFee          string    `gorm:"column:HZJG_FEE"`
	Jgsy             string    `gorm:"column:JGSY"`
	AipFee           string    `gorm:"column:AIP_FEE"`
	MchtSetAmt       string    `gorm:"column:MCHT_SET_AMT"`
	Hzjgyfppfwf      string    `gorm:"column:HZJGYFPPFWF"`
	Jgyfppfwf        string    `gorm:"column:JGYFPPFWF"`
	Aipyfppfwf       string    `gorm:"column:AIPYFPPFWF"`
	ErrFeeIn         string    `gorm:"column:ERR_FEE_IN"`
	ErrFeeOut        string    `gorm:"column:ERR_FEE_OUT"`
	ErrCode          string    `gorm:"column:ERR_CODE"`
	JtMchtCd         string    `gorm:"column:JT_MCHT_CD"`
	ExpandOrgCd      string    `gorm:"column:EXPAND_ORG_CD"`
	SpeServInst      string    `gorm:"column:SPE_SERV_INST"`
	PropIns          string    `gorm:"column:PROP_INS"`
	ExpandOrgFee     string    `gorm:"column:EXPAND_ORG_FEE"`
	SpeServFee       string    `gorm:"column:SPE_SERV_FEE"`
	PropInsFee       string    `gorm:"column:PROP_INS_FEE"`
	ExpandOrgPp      string    `gorm:"column:EXPAND_ORG_PP"`
	SpeServPp        string    `gorm:"column:SPE_SERV_PP"`
	PropInsPp        string    `gorm:"column:PROP_INS_PP"`
	ExpandFeeIn      string    `gorm:"column:EXPAND_FEE_IN"`
	ExpandFeeOut     string    `gorm:"column:EXPAND_FEE_OUT"`
	CupIfinsideSign  string    `gorm:"column:CUP_IFINSIDE_SIGN"`
	SpChargType      string    `gorm:"column:SP_CHARG_TYPE"`
	SpChargLev       string    `gorm:"column:SP_CHARG_LEV"`
	TermSsn          string    `gorm:"column:TERM_SSN"`
	SnSsn            string    `gorm:"column:SN_SSN"`
	UpChlID          string    `gorm:"column:UP_CHL_ID"`
	ConvMchtCd       string    `gorm:"column:CONV_MCHT_CD"`
	ConvTermID       string    `gorm:"column:CONV_TERM_ID"`
	ChlTrueFee       string    `gorm:"column:CHL_TRUE_FEE"`
	ChlStdFee        string    `gorm:"column:CHL_STD_FEE"`
	ChlFeePreFlg     string    `gorm:"column:CHL_FEE_PRE_FLG"`
	SysSer           string    `gorm:"column:SYS_SER"`
	Var6             string    `gorm:"column:VAR_6"`
	QudaoFee         string    `gorm:"column:QUDAO_FEE"`
	QudaoFeeMin      string    `gorm:"column:QUDAO_FEE_MIN"`
	QudaoFeeMix      string    `gorm:"column:QUDAO_FEE_MIX"`
	QudaoFeeFd       string    `gorm:"column:QUDAO_FEE_FD"`
	InsFee           string    `gorm:"column:INS_FEE"`
	InsMyFee         string    `gorm:"column:INS_MY_FEE"`
	InsCostFee       string    `gorm:"column:INS_COST_FEE"`
	InsMyFeeAmt      string    `gorm:"column:INS_MY_FEE_AMT"`
	InsSplitFee      string    `gorm:"column:INS_SPLIT_FEE"`
	InsResFee        string    `gorm:"column:INS_RES_FEE"`
	PinpFee          string    `gorm:"column:PINP_FEE"`
	PinpFeeInf       string    `gorm:"column:PINP_FEE_INF"`
	PinpFeeTop       string    `gorm:"column:PINP_FEE_TOP"`
	PinpStat         string    `gorm:"column:PINP_STAT"`
	T0Stat           string    `gorm:"column:T0_STAT"`
	KeyRsp           *string   `gorm:"column:KEY_RSP; primary_key"`
	Remark           *string   `gorm:"column:REMARK"`
	Remark1          *string   `gorm:"column:REMARK1"`
	Remark2          *string   `gorm:"column:REMARK2"`
	Remark3          *string   `gorm:"column:REMARK3"`
	Remark4          *string   `gorm:"column:REMARK4"`
	Remark5          *string   `gorm:"column:REMARK5"`
	MaSettleDt       time.Time `gorm:"column:MA_SETTLE_DT"`
	RetriRefNo       string    `gorm:"column:RETRI_REF_NO"`
	IndustryAddnInf  string    `gorm:"column:INDUSTRY_ADDN_INF"`
	SrcQid           float64   `gorm:"column:SRC_QID"`
	ProdCd           string    `gorm:"column:PROD_CD"`
}

func (c ClearTxn) TableName() string {
	return clearTxnName
}

//DownloadInstitutionFile 根据开始时间跟结束时间查找交易流水
//@return Institution： 符合的实例
func DownloadInstitutionFile(db *gorm.DB, startTime, endTime string) ([]*ClearTxn, error) {
	clearTxn := []*ClearTxn{}
	err := db.Debug().Where("STLM_DATE < ? AND STLM_DATE > ?", endTime, startTime).Find(&clearTxn).Error
	if err != nil {
		return nil, err
	}
	return clearTxn, err
}

//通过日期将数据分为不同的文件
//@param fileDir "/home/dj/test/"
func downloadFile(clearTxn []*ClearTxn, fileDir string) error {

	if len(clearTxn) == 0 {
		return errors.New("Not found correct record")
	}

	os.MkdirAll(fileDir, os.ModePerm)
	var dataIndex string

	for index := 0; index < len(clearTxn); index++ {
		dataIndex = clearTxn[index].StlmDate
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("Worksheet")
		if err != nil {
			println(err)
		}

		var cell *xlsx.Cell
		{
			row := sheet.AddRow()
			cell = row.AddCell()
			cell.Value = "商户编码"
			cell = row.AddCell()
			cell.Value = "交易日期"
			cell = row.AddCell()
			cell.Value = "交易时间"
			cell = row.AddCell()
			cell.Value = "清算日期"
			cell = row.AddCell()
			cell.Value = "终端编号"
			cell = row.AddCell()
			cell.Value = "交易类型"
			cell = row.AddCell()
			cell.Value = "交易流水号"
			cell = row.AddCell()
			cell.Value = "交易卡号"
			cell = row.AddCell()
			cell.Value = "卡类型"
			cell = row.AddCell()
			cell.Value = "交易本金"
			cell = row.AddCell()
			cell.Value = "交易手续费"
			cell = row.AddCell()
			cell.Value = "交易结算资金"
			cell = row.AddCell()
			cell.Value = "应收差错费用"
			cell = row.AddCell()
			cell.Value = "应付差费用"
			cell = row.AddCell()
			cell.Value = "系统流水号"
			cell = row.AddCell()
			cell.Value = "机构基准收入"
			cell = row.AddCell()
			cell.Value = "机构实际收入"
			cell = row.AddCell()
			cell.Value = "机构营销返佣"
			cell = row.AddCell()
			cell.Value = "代理编码"
			cell = row.AddCell()
			cell.Value = "会员号"
		}

		var j = index
		for ; j < len(clearTxn); j++ {
			if dataIndex == clearTxn[j].StlmDate {
				row := sheet.AddRow()
				cell = row.AddCell()
				cell.Value = clearTxn[j].MchtCd
				cell = row.AddCell()
				cell.Value = clearTxn[j].StlmDate
				cell = row.AddCell()
				cell.Value = clearTxn[j].TransDateTime
				cell = row.AddCell()
				cell.Value = "清算日期"
				cell = row.AddCell()
				cell.Value = "终端编号"
				cell = row.AddCell()
				cell.Value = clearTxn[j].TxnDesc
				cell = row.AddCell()
				cell.Value = "交易流水号"
				cell = row.AddCell()
				cell.Value = "交易卡号"
				cell = row.AddCell()
				cell.Value = clearTxn[j].CardKindDis
				cell = row.AddCell()
				cell.Value = "交易本金"
				cell = row.AddCell()
				cell.Value = "交易手续费"
				cell = row.AddCell()
				cell.Value = "交易结算资金"
				cell = row.AddCell()
				cell.Value = "应收差错费用"
				cell = row.AddCell()
				cell.Value = "应付差费用"
				cell = row.AddCell()
				cell.Value = "系统流水号"
				cell = row.AddCell()
				cell.Value = "机构基准收入"
				cell = row.AddCell()
				cell.Value = "机构实际收入"
				cell = row.AddCell()
				cell.Value = "机构营销返佣"
				cell = row.AddCell()
				cell.Value = "代理编码"
				cell = row.AddCell()
				cell.Value = "会员号"
			} else {
				break
			}

		}
		index = j - 1

		filename := fileDir + dataIndex + ".xlsx"
		err = file.Save(filename)
		if err != nil {
			println(err)
		}
	}
	compress(fileDir, "/home/dj/"+"result.zip")
	return nil
}

func compress(dis, src string) error {
	disFile, err := os.Create(src)
	if err != nil {
		return err
	}
	defer disFile.Close()

	zw := zip.NewWriter(disFile)
	defer zw.Close()

	filesInfo, err := ioutil.ReadDir(dis)
	if err != nil {
		errors.New("File not exist")
	}

	files := make([]string, 0)
	for _, f := range filesInfo {
		files = append(files, dis+f.Name())
	}

	fmt.Println(files)

	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Using FileInfoHeader() above only uses the basename of the file. If we want
		// to preserve the folder structure we can overwrite this with the full path.
		header.Name = file

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		if _, err = io.Copy(writer, zipfile); err != nil {
			return err
		}
	}
	return nil
}
