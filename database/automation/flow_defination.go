package automation

import (
	"github.com/gy1229/oa/database"
	"github.com/sirupsen/logrus"
)

func CreateFlowDefination(id, creatorId int64, name string) error {
	fd := &database.FlowDefination{
		Id:         id,
		Name:       name,
		CreatorId:  creatorId,
	}
	if err := database.DB.Create(fd).Error; err != nil {
		logrus.Error("[CreateFlowDefination] err msg", err.Error())
		return err
	}
	return nil
}

func GetFlowDefination(id int64) (*database.FlowDefination, error) {
	fd := &database.FlowDefination{
		Id: id,
	}
	if err := database.DB.Where("id = ?", id).Find(&fd).Error; err != nil {
		logrus.Error("[GetFlowDefination] err ", err.Error())
		return nil, err
	}
	return fd, nil
}

func DelFlowDefination(id int64) error {
	fd := &database.FlowDefination{
		Status:     0,
	}
	if err := database.DB.Model(&fd).Where("id = ?", id).Updates(fd).Error; err != nil {
		logrus.Error("[DelFlowDefination] err msg ", err.Error())
		return err
	}
	return nil
}

func UpdateFlowDefination(id int64, name string) error {
	fd := &database.FlowDefination{
		Name:       name,
	}
	if err := database.DB.Model(&fd).Where("id = ?", id).Updates(fd).Error; err != nil {
		logrus.Error("[UpdateFlowDefination] err msg ", err.Error())
		return err
	}
	return nil
}