package automation

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database"
	"github.com/sirupsen/logrus"
)

func CreateFlowInstance(fi *database.FlowInstance) error {
	if err := database.DB.Create(fi).Error; err != nil {
		logrus.Error("[CreateFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func UpdateFlowInstance(fi *database.FlowInstance) error {
	if err := database.DB.Model(&fi).Where("id = ?", fi.Id).Updates(fi).Error; err != nil {
		logrus.Error("[UpdateFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func DeletFlowInstance(fi *database.FlowInstance) error {
	if err := database.DB.Delete(fi).Error; err != nil {
		logrus.Error("[DeletFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func FindFlowInstance(fi *database.FlowInstance) error {
	if err := database.DB.Where("id = ?", fi.Id).First(&fi).Error; err != nil {
		logrus.Error("[FindFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func CreateFlowInstanceByDetail(id, fdId int64) error {
	fi := &database.FlowInstance{
		Id:               id,
		FlowDefinationId: fdId,
		RunStatus:        constant.FlowInstanceStart,
	}
	if err := database.DB.Create(fi).Error; err != nil {
		logrus.Error("[CreateFlowInstanceByDetail] err msg", err.Error())
		return err
	}
	return nil
}

func UpdateFlowInstanceStatus(id int64, status int) error {
	fi := &database.FlowInstance{
		RunStatus:        status,
	}
	if err := database.DB.Model(&fi).Where("id = ?", id).Updates(fi).Error; err != nil {
		logrus.Error("[UpdateFlowDefination] err msg ", err.Error())
		return err
	}
	return nil
}
