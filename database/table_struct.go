package database

import "time"

type OaUser struct {
	Id         int64
	Account    string    `json:"account"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
	ThirdId    int64     `json:"third_id"`
	ImageId    int64     `json:"image_id"`
}

type Third struct {
	Id int64
	Account string `json:"account"`
	Password   string    `json:"password"`
	UserId int64
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
	App string `json:"app"`
}
type StageRepository struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	CreatorId  int64     `json:"creator_id"`
	Authority  int       `json:"authority"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
	Status     int       `json:"status"`
}

type FileDetail struct {
	Id          int64     `json:"id"`
	CreatorId   int64     `json:"creator_id"`
	StageRespId int64     `json:"stage_resp_id"`
	Type        int       `json:"type"`
	UpdateTime  time.Time `json:"update_time"`
	CreateTime  time.Time `json:"create_time"`
	Status      int       `json:"status"`
	Name string `json:"name"`
}

type FileText struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
	FileId     int64     `json:"file_id"`
}

type FileTable struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	FileId     int64     `json:"file_id"`
	Status     string    `json:"status"`
	RowLen     int64     `json:"row_len"`
	LineLen    int64     `json:"line_len"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
}

type TableCell struct {
	Id          int64  `json:"id"`
	FileTableId int64  `json:"file_table_id"`
	Content     string `json:"content"`
	Row         int64  `json:"row"`
	Line        int64  `json:"line"`
	Status      int `json:"status"`
}

type ImageFile struct {
	Id         int64     `json:"id"`
	ImageFile  string    `json:"image_file"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	ImageSize  int64     `json:"image_size"`
}

type ActionDefination struct {
	Id         int64     `json:"id"`
	ActionId   int64 `json:"action_id"`
	FlowDefinationId int64 `json:"flow_defination_id"`
	Position int `json:"postion"`
	ActionType string `json:"action_type"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status int `json:"status"`
}

type Action struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ImageId int64 `json:"image_id"`
	Status int `json:"status"`
	ActionType string `json:"action_type"` // 1 是Trigger 2 是Action
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type FlowDefination struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	CreatorId int64 `json:"creator_id"`
	Status int `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type FlowInstance struct {
	Id int64 `json:"id"`
	FlowDefinationId int64 `json:"flow_defination_id"`
	CreateTime time.Time `json:"create_time"`
	RunStatus int `json:"run_status"`
}

type FormData struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	ActionDefinationId int64 `json:"action_defination_id"`
	Key string `json:"key"`
	Value string `json:"value"`
	Position int `json:"position"`
	Status int `json:"status"`
}

type Test struct {
	Id int `json:"id"`
	Name string `json:"name"`
}