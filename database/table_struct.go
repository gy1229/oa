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
	Id             int64     `json:"id"`
	CreatorId      int64     `json:"creator_id"`
	FileName       string    `json:"file_name"`
	StageRespId    int64     `json:"stage_resp_id"`
	Type           int       `json:"type"`
	UpdateTime     time.Time `json:"update_time"`
	CreateTime     time.Time `json:"create_time"`
	Status         int       `json:"status"`
}

type FileText struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Status     int    `json:"status"`
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
}
