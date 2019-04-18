package institutionservice

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	cleartxnM "userService/pkg/model/cleartxn"

	"github.com/satori/go.uuid"
	"github.com/tealeg/xlsx"
)

const filePath = "/tmp"
const defaultSheep = "Worksheet"
const fileZipPath = "/home/dj/zip"

//通过日期将数据分为不同的文件
//@param fileDir test
//@return fileDir 下载完成之后存放临时文件的文件夹
func DownloadFileWithDay(clearTxn []*cleartxnM.ClearTxn) (string, error) {

	if len(clearTxn) == 0 {
		return "", nil
	}

	uid := uuid.NewV4().String()

	fileDirPath := path.Join(filePath, uid)
	os.MkdirAll(fileDirPath, os.ModePerm)
	var dataIndex string

	for index := 0; index < len(clearTxn); index++ {
		dataIndex = clearTxn[index].StlmDate
		file := xlsx.NewFile()
		sheet, err := file.AddSheet(defaultSheep)
		if err != nil {
			return "", err
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

		filename := path.Join(fileDirPath, dataIndex+".xlsx")
		err = file.Save(filename)
		if err != nil {
			return "", err
		}
	}
	return uid, nil
}

//@param dis 压缩前文件名称
//@param src 压缩后文件名称
func Compress(dis, src string) error {

	srcFile, err := os.Create(path.Join(fileZipPath, src))
	if err != nil {
		return err
	}
	defer srcFile.Close()

	zw := zip.NewWriter(srcFile)
	defer zw.Close()

	disFileDir := path.Join(filePath, dis)
	filesInfo, err := ioutil.ReadDir(disFileDir)
	if err != nil {
		return errors.New("File not exist")
	}

	files := make([]string, 0)
	for _, f := range filesInfo {
		files = append(files, path.Join(disFileDir, f.Name()))
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
