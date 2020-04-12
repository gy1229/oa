package automation

import (
	"github.com/gy1229/oa/database"
	"github.com/sirupsen/logrus"
	"time"
)

func CreateFlowDefination(fd *database.FlowDefination) error {
	if err := database.DB.Create(fd).Error; err != nil {
		logrus.Error("[CreateFlowDefination] err ", err.Error())
		return err
	}
	return nil
}

func UpdateFlowDefination(fd *database.FlowDefination) error {
	if err := database.DB.Model(&fd).Where("id = ?", fd.Id).Updates(fd).Error; err != nil {
		logrus.Error("[UpdateFlowDefination] err ", err.Error())
		return err
	}
	return nil
}

func DeletFlowDefination(fd *database.FlowDefination) error {
	if err := database.DB.Delete(fd).Error; err != nil {
		logrus.Error("[DeletFlowDefination] err ", err.Error())
		return err
	}
	return nil
}

func FindFlowDefination(fd *database.FlowDefination) error {
	if err := database.DB.Where("id = ?", fd.Id).First(&fd).Error; err != nil {
		logrus.Error("[FindFlowDefination] err ", err.Error())
		return err
	}
	return nil
}

func CreateFlowDefinationByArgs(id, creatorId int64, name string) error {
	fd := &database.FlowDefination{
		Id:         id,
		Name:       name,
		CreatorId:  creatorId,
		Status: 0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := database.DB.Create(fd).Error; err != nil {
		logrus.Error("[CreateFlowDefinationByArgs] err msg", err.Error())
		return err
	}
	return nil
}

func GetFlowDefinationById(id int64) (*database.FlowDefination, error) {
	fd := &database.FlowDefination{
		Id: id,
	}
	if err := database.DB.Where("id = ?", id).Find(&fd).Error; err != nil {
		logrus.Error("[GetFlowDefinationById] err ", err.Error())
		return nil, err
	}
	return fd, nil
}

func DelFlowDefinationById(id int64) error {
	fd := &database.FlowDefination{
		Status:     1,
	}
	if err := database.DB.Model(&fd).Where("id = ?", id).Updates(fd).Error; err != nil {
		logrus.Error("[DelFlowDefinationById] err msg ", err.Error())
		return err
	}
	return nil
}

func UpdateFlowDefinationNameById(id int64, name string) error {
	fd := &database.FlowDefination{
		Name:       name,
	}
	if err := database.DB.Model(&fd).Where("id = ?", id).Updates(fd).Error; err != nil {
		logrus.Error("[UpdateFlowDefinationNameById] err msg ", err.Error())
		return err
	}
	return nil
}

func FindFlowDefinationByUserId(userId int64) ([]*database.FlowDefination, error) {
	fds := make([]*database.FlowDefination, 0)
	if err := database.DB.Where("creator_id = ? AND status = 0", userId).Find(&fds).Error; err != nil {
		logrus.Error("[FindFlowDefinationByUserId] err ", err.Error())
		return nil, err
	}
	return fds, nil
}