package automation

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database"
	"github.com/sirupsen/logrus"
)

func CreateFlowInstance(id, fdId int64) error {
	fi := &database.FlowInstance{
		Id:               id,
		FlowDefinationId: fdId,
		RunStatus:        constant.FlowInstanceStart,
	}
	if err := database.DB.Create(fi).Error; err != nil {
		logrus.Error("[CreateFlowInstance] err msg", err.Error())
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
