package stage

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database/stage"
	data_user "github.com/gy1229/oa/database/user"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/util"
	"github.com/sirupsen/logrus"
	"strconv"
)

func UploadFile(c *gin.Context) {

}

func GetFileList(req *json_struct.GetFileListRequest) (*json_struct.GetFileListResponse, error) {
	respId, err := strconv.ParseInt(req.RepositoryId, 10, 64)
	if err != nil {
		return nil, err
	}
	dfiles, err := stage.DGetFileListByRepId(respId)
	if err != nil {
		return nil, err
	}
	files := make([]*json_struct.File, 0)
	for k, v := range dfiles {
		files = append(files, &json_struct.File{
			Id:         strconv.FormatInt(v.Id, 10),
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		})
		files[k].CreatorName, _ = data_user.GetUserNameById(v.CreatorId)
		files[k].Name, _ = stage.DGetFileName(v.Id, v.Type)
	}
	return &json_struct.GetFileListResponse{
		FileList: files,
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil

}

func GetFileContent(req *json_struct.GetFileContentRequest) (*json_struct.GetFileContentResponse, error) {
	fileId, err := strconv.ParseInt(req.FileId, 10, 64)
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	logrus.Info("[GetFileContent]user id is ", userId)
	fileDetail, err := stage.DGetFileDetalByFileId(fileId)
	if err != nil {
		return nil, err
	}
	switch fileDetail.Type {
	case constant.TextFile:
		fileText, err := stage.DGetTextFileByFileId(fileId)
		if err != nil {
			return nil, err
		}
		return &json_struct.GetFileContentResponse{
			Name:    fileText.Name,
			Type:    strconv.Itoa(constant.TextFile),
			Content: &fileText.Content,
			Base:    &json_struct.BaseResponse{Body: constant.SUCCESS},
		}, nil
	case constant.TableFile:
		resp, err := getTableFileContent(fileId)
		return resp, err
	default:
		return nil, errors.New("cont match file type")
	}

}

func getTableFileContent(fileId int64) (*json_struct.GetFileContentResponse, error) {
	fileTable, err := stage.DGetTableFileByFileId(fileId)
	if err != nil {
		return nil, err
	}
	tableCells, err := stage.DGetTableCellsByFileId(fileTable.Id)
	if err != nil {
		return nil, err
	}
	tCells := make([]*json_struct.TableCell, 0)
	for _, v := range tableCells {
		tCells = append(tCells, &json_struct.TableCell{
			Id:      strconv.FormatInt(v.Id, 10),
			Content: v.Content,
			Row:     strconv.FormatInt(v.Row, 10),
			Line:    strconv.FormatInt(v.Line, 10),
		})
	}
	content := &json_struct.TableContent{
		RowLen:     strconv.FormatInt(fileTable.RowLen, 10),
		LineLen:    strconv.FormatInt(fileTable.LineLen, 10),
		TableCells: tCells,
	}
	return &json_struct.GetFileContentResponse{
		Name:         fileTable.Name,
		Type:         strconv.Itoa(constant.TextFile),
		TableContent: content,
		Base:         &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func UpdateTextContent(req *json_struct.UpdateTextContentRequest) (*json_struct.UpdateTextContentResponse, error) {
	fileId, err := strconv.ParseInt(req.FileId, 10, 64)
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	logrus.Info("[UpdateTextContent] userId is ", userId)
	err = stage.DUpdateTextContent(fileId, req.Content, req.Name)
	if err != nil {
		return nil, err
	}
	return &json_struct.UpdateTextContentResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func UpdateTableContent(req *json_struct.UpdateTableContentRequest) (*json_struct.UpdateTableContentResponse, error) {
	fileId, err := strconv.ParseInt(req.FileId, 10, 64)
	if err != nil {
		logrus.Error("[UpdateTableContent] fileId ParseInt")
		return nil, err
	}
	logrus.Info("[UpdateTextContent] fileId is ", fileId)
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[UpdateTableContent] userId ParseInt")
		return nil, err
	}
	logrus.Info("[UpdateTextContent] userId is ", userId)
	for _, v := range req.TableCells {
		cellId, err := strconv.ParseInt(v.Id, 10, 64)
		if err != nil {
			logrus.Error("[UpdateTableContent] cellId ParseInt")
			return nil, err
		}
		if err := stage.DUpdateTableContent(cellId, v.Content); err != nil {
			logrus.Error("[UpdateTableContent] stage.DUpdateTableContent ", err.Error())
			return nil, err
		}
	}
	return &json_struct.UpdateTableContentResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func CreateNewFile(req *json_struct.CreateNewFileRequest) (*json_struct.CreateNewFileResponse, error) {
	repositoryId, err := strconv.ParseInt(req.RepositoryId, 10, 64)
	if err != nil {
		logrus.Error("[CreateNewFile] repositoryId ParseInt")
		return nil, err
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[CreateNewFile] userId ParseInt")
		return nil, err
	}
	ttype, err := strconv.Atoi(req.Type)
	if err != nil {
		logrus.Error("[CreateNewFile] ttype ParseInt")
		return nil, err
	}
	id := util.GenId()
	err = stage.DCreateFileDeatil(req.Name, userId, repositoryId, id, ttype)
	if err != nil {
		logrus.Error("[CreateNewFile] err ", err.Error())
		return nil, err
	}
	switch req.Type {
	case constant.TextFileString:
		err = stage.DCreateTextFile(req.Name, req.Content, id)
		if err != nil {
			return nil, err
		}
	case constant.TableFileString:
		CreateTableFile(req.Name, req.TableContent, id)

	}
	return &json_struct.CreateNewFileResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func CreateTableFile(name string, tableContent *json_struct.TableContent, fileId int64) error {
	id := util.GenId()
	rowLen, err := strconv.ParseInt(tableContent.RowLen, 10, 64)
	if err != nil {
		logrus.Error("[CreateTableFile] rowLen ParseInt ")
		return err
	}
	lineLen, err := strconv.ParseInt(tableContent.LineLen, 10, 64)
	if err != nil {
		logrus.Error("[CreateTableFile] lineLen ParseInt ")
		return err
	}
	err = stage.DCreateTableFile(id, fileId, rowLen, lineLen, name)
	if err != nil {
		return err
	}
	for _, v := range tableContent.TableCells {
		row, err := strconv.ParseInt(v.Row, 10, 64)
		if err != nil {
			return err
		}
		line, err := strconv.ParseInt(v.Line, 10, 64)
		if err != nil {
			return err
		}
		err = stage.DCreateTableCell(id, row, line, v.Content)
		if err != nil {
			logrus.Error("[CreateTableFile]  err ", err.Error())
			return err
		}
	}
	return nil
}

func DelFile(req *json_struct.DelFileRequest) (*json_struct.DelFileResponse, error) {
	fileId, err := strconv.ParseInt(req.FileId, 10, 64)
	if err != nil {
		logrus.Error("[DelFile] fileId ParseInt")
		return nil, err
	}
	err = stage.DDelTableDetailById(fileId)
	if err != nil {
		return nil, err
	}
	return &json_struct.DelFileResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}
