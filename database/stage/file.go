package stage

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/util"
	"github.com/sirupsen/logrus"
	"time"
)

// 数据库相关D开头
func DGetFileListByRepId(repId int64) ([]*database.FileDetail, error) {
	files := make([]*database.FileDetail, 0)
	if err := database.DB.Where("stage_resp_id = ? AND status=0", repId).Find(&files).Error; err != nil {
		logrus.Error("DGetFileListByRepId err ", err.Error())
		return nil, err
	}
	return files, nil
}

func DGetFileDetalByFileId(fileId int64) (*database.FileDetail, error) {
	file := &database.FileDetail{}
	if err := database.DB.Where("id = ?", fileId).Find(&file).Error; err != nil {
		logrus.Error("DGetFileDetalByFileId err ", err.Error())
		return nil, err
	}
	return file, nil
}

func DGetTextFileByFileId(fileId int64) (*database.FileText, error) {
	fileText := &database.FileText{}
	if err := database.DB.Where("file_id = ?", fileId).Find(&fileText).Error; err != nil {
		logrus.Error("DGetTextFileByFileId err ", err.Error())
		return nil, err
	}
	return fileText, nil
}

func DGetTableFileByFileId(fileId int64) (*database.FileTable, error) {
	fileTable := &database.FileTable{}
	if err := database.DB.Where("file_id = ?", fileId).Find(&fileTable).Error; err != nil {
		logrus.Error("DGetTableFileByFileId err ", err.Error())
		return nil, err
	}
	return fileTable, nil
}

func DGetTableCellsByFileId(fileId int64) ([]*database.TableCell, error) {
	var tableCells []*database.TableCell
	if err := database.DB.Where("file_table_id = ? AND status = 0", fileId).Find(&tableCells).Error; err != nil {
		logrus.Error("DGetTextFileContent err ", err.Error())
		return nil, err
	}
	return tableCells, nil
}

func DUpdateTextContent(fileId int64, content string, name string) error {
	fileText := database.FileText{
		Content: content,
		Name:    name,
	}
	if err := database.DB.Model(&fileText).Where("file_id = ?", fileId).Updates(&fileText).Error; err != nil {
		logrus.Error("[DUpdateTextContent] err ", err.Error())
		return err
	}
	return nil
}

func DUpdateTableContent(cellId int64, content string) error {
	cell := database.TableCell{
		Content: content,
	}
	if err := database.DB.Model(&cell).Where("id = ?", cellId).Updates(&cell).Error; err != nil {
		logrus.Error("[DUpdateTableContent] err ", err.Error())
		return err
	}
	return nil
}

func DCreateTextFile(name, content string, fileId int64) error {
	fileText := database.FileText{
		Id:         util.GenId(),
		Name:       name,
		Content:    content,
		Status:     0,
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
		FileId:     fileId,
	}
	if err := database.DB.Create(fileText).Error; err != nil {
		logrus.Error("InsertUserMessage err ", err.Error())
		return err
	}
	return nil
}

func DCreateFileDeatil(name string, userId, stageRepId, id int64, ttype int) error {
	fileDetail := database.FileDetail{
		Id:          id,
		CreatorId:   userId,
		StageRespId: stageRepId,
		Type:        ttype,
		UpdateTime:  time.Now(),
		CreateTime:  time.Now(),
		Status:      0,
	}
	if err := database.DB.Create(fileDetail).Error; err != nil {
		logrus.Error("[DCreateFileDeatil] err ", err.Error())
		return err
	}
	return nil
}

func DCreateTableFile(id, fileId, rowLen, lineLen int64, name string) error {
	tableFile := database.FileTable{
		Id:         id,
		Name:       name,
		FileId:     fileId,
		Status:     "0",
		RowLen:     rowLen,
		LineLen:    lineLen,
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
	}
	if err := database.DB.Create(tableFile).Error; err != nil {
		logrus.Error("[DCreateTableFile] err ", err.Error())
		return err
	}
	return nil
}

func DCreateTableCell(tableFileId, row, line int64, content string) error {
	tableCell := database.TableCell{
		Id:      util.GenId(),
		FileTableId:  tableFileId,
		Content: content,
		Row:     row,
		Line:    line,
	}
	if err := database.DB.Create(tableCell).Error; err != nil {
		logrus.Error("[DCreateTableCell] err ", err.Error())
		return err
	}
	return nil
}

func DDelTableDetailById(fileId int64) error {
	fileDetail := database.FileDetail{
		Status: 1,
	}
	if err := database.DB.Model(&fileDetail).Where("id = ?", fileId).Updates(&fileDetail).Error; err != nil {
		logrus.Error("[DUpdateTableContent] err ", err.Error())
		return err
	}
	return nil
}

func DGetFileName(file_id int64, ttype int) (string, error) {
	switch ttype {
	case constant.TableFile:
		fileTable := &database.FileTable{}
		if err := database.DB.Where("file_id = ?", file_id).Find(&fileTable).Error; err != nil {
			logrus.Error("DGetFileListByRepId err ", err.Error())
			return "", err
		}
		return fileTable.Name, nil
	default:
		fileText := &database.FileText{}
		if err := database.DB.Where("file_id = ?", file_id).Find(&fileText).Error; err != nil {
			logrus.Error("DGetFileListByRepId err ", err.Error())
			return "", err
		}
		return fileText.Name, nil
	}

}

func DUpdateTableFileByFileId(table *database.FileTable) error{
	if err := database.DB.Model(&table).Where("file_id = ?", table.FileId).Updates(table).Error; err != nil {
		logrus.Error("[UpdateFlowDefination] err msg ", err.Error())
		return err
	}
	return nil
}

func DDeleteTableCellById(cellId int64) error {
	cell := database.TableCell{
		Status: 1,
	}
	if err := database.DB.Model(&cell).Where("id = ?", cellId).Updates(&cell).Error; err != nil {
		logrus.Error("[DDeleteTableCellById] err ", err.Error())
		return err
	}
	return nil
}