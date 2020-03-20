package json_struct

import (
	"time"
)

type GetFileListRequest struct {
	RepositoryId string `json:"repository_id"`
	UserId string `json:"user_id"`
}

type File struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CreatorName string `json:"creator_name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type GetFileListResponse struct {
	FileList []*File      `json:"file_list"`
	Base     *BaseResponse `json:"base"`
}

type GetFileContentRequest struct {
	FileId string `json:"file_id"`
	UserId string `json:"user_id"`
}

type TableCell struct {
	Id string `json:"id"`
	Content string `json:"content"`
	Row string `json:"row"`
	Line string `json:"line"`
}

type TableContent struct {
	RowLen string `json:"row_len"`
	LineLen string `json:"line_len"`
	TableCells []*TableCell `json:"table_cells"`
}

type GetFileContentResponse struct {
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	Content      string       `json:"content",omitempty`
	TableContent *TableContent `json:"table_content", omitempty`
	Base         *BaseResponse  `json:"base"`
}

type UpdateTextContentRequest struct {
	FileId string `json:"file_id"`
	Content string `json:"content"`
	UserId string `json:"user_id"`
	Name string `json:"name"`
}

type UpdateTextContentResponse struct {
	Base *BaseResponse `json:"base"`
}

type UpdateTableContentRequest struct {
	FileId string `json:"file_id"`
	UserId string `json:"user_id"`
	Name string `json:"name"`
	TableCells []*TableCell `json:"table_cells"`
}

type UpdateTableContentResponse struct {
	Base *BaseResponse `json:"base"`
}

type CreateNewFileRequest struct {
	RepositoryId string `json:"repository_id"`
	UserId string `json:"user_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Content string `json:"content",omitempty`
	TableContent *TableContent `json:"table_content", omitempty`
}

type CreateNewFileResponse struct {
	Base *BaseResponse `json:"base"`
}

type DelFileRequest struct {
	FileId string `json:"file_id"`
	UserId string `json:"user_id"`
}

type DelFileResponse struct {
	Base *BaseResponse `json:"base"`
}