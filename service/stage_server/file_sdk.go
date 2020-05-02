package stage_server

import (
	"bufio"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database/stage"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"path"
	"strconv"
)

func UploadFile2Stage(fileHeader *multipart.FileHeader, userIdS, repIdS string) error {
	//b, err := ioutil.ReadAll(file)
	//if err != nil {
	//	return err
	//}
	userId, err := strconv.ParseInt(userIdS, 10, 64)
	if err != nil {
		return nil
	}
	repId, err := strconv.ParseInt(repIdS, 10, 64)
	if err != nil {
		return nil
	}
	fileId := util.GenId()
	fileType := 2

	fNameLast := path.Ext(fileHeader.Filename)
	if fNameLast == ".txt" {
		Read2StageTextFile(fileHeader, fileId)
		fileType = constant.TextFile
	} else if fNameLast == ".xlsx" {
		Read2StageTableFile(fileHeader, fileId)
		fileType = constant.TableFile
	} else {
		return errors.New("cant find type")
	}
	err = stage.DCreateFileDeatil(fileHeader.Filename, userId, repId, fileId, fileType)
	if err != nil {
		return err
	}
	return nil
}

func Read2StageTextFile(fileHeader *multipart.FileHeader, fileId int64) error {
	file, _ := fileHeader.Open()
	bufSize := viper.GetInt("server.BufSize")
	buf := make([]byte, bufSize) //一次读取多少个字节
	bfRd := bufio.NewReader(file)
	textByte := make([]byte, 1)
	for {
		n, err := bfRd.Read(buf)
		textByte = util.BytesCombine(textByte, buf[:n])
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			return err
		}
	}
	err := stage.DCreateTextFile(fileHeader.Filename, string(textByte), fileId)
	return err
}

func Read2StageTableFile(fHeader *multipart.FileHeader, fileId int64) error {
	f, _ := fHeader.Open()
	file, err := excelize.OpenReader(f)
	if err != nil {
		logrus.Error("[Read2StageTableFile] err msg ", err.Error())
		return err
	}
	rows := file.GetRows("Sheet1")
	rlen := len(rows)
	if rlen == 0 {
		return errors.New("cant find row")
	}
	llen := len(rows[0])
	tableCells := make([]*json_struct.TableCell, rlen*llen)
	for r, row := range rows {
		for l, cell := range row {
			tableCells[llen*r+l] = &json_struct.TableCell{
				Content: cell,
				Row:     strconv.Itoa(r),
				Line:    strconv.Itoa(l),
			}
		}
	}
	tableContent := &json_struct.TableContent{
		RowLen:     strconv.Itoa(rlen),
		LineLen:    strconv.Itoa(llen),
		TableCells: tableCells,
	}
	err = CreateTableFile(fHeader.Filename, tableContent, fileId)
	if err != nil {
		return err
	}
	return nil
}
