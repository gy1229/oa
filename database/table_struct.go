package database

import "time"

type OaUser struct {
	Id         *int64
	Account    string    `json:"account"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
	ThirdId    *int64    `json:"third_id"`
	ImageId    *int64    `json:"image_id"`
}

type StageRepository struct {
	Id         *int64    `json:"id"`
	Name       string    `json:"name"`
	CreatorId  *int64    `json:"creator_id"`
	Authority  *int      `json:"authority"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
	Status     *int      `json:"status"`
}
