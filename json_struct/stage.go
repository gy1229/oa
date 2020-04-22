package json_struct

import "time"

type UploadFileResponse struct {
	Base *BaseResponse `json:"base"`
}

type CreateRepositoryRequest struct {
	Name      string `json:"name"`
	Authority string `json:"authority"`
	UserId    string `json:"user_id"`
}

type CreateRepositoryResponse struct {
	Id   string        `json:"id"`
	Base *BaseResponse `json:"base"`
}

type Repository struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	CreatorName string    `json:"creator_name"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	Authority   *int      `json:"authority"`
}

type GetRepositoryListRequest struct {
	UserId string `json:"user_id"`
}

type GetRepositoryListResponse struct {
	RepositoryList []*Repository `json:"repository_list"`
	Base           *BaseResponse `json:"base"`
}

type UpdateRepositoryRequest struct {
	RepositoryId string `json:"repository_id"`
	Name         string `json:"name"`
	Authority    string `json:"authority"`
	UserId       string `json:"user_id"`
}

type UpdateRepositoryResponse struct {
	Base *BaseResponse `json:"base"`
}

type DelRepositoryRequest struct {
	RepositoryId string `json:"repository_id"`
	UserId       string `json:"user_id"`
}

type DelRepositoryResponse struct {
	Base *BaseResponse `json:"base"`
}

type GetRepositoryByIdResponse struct {
}
