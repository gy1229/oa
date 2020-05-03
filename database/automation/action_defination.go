package automation

import (
	"errors"
	"github.com/gy1229/oa/database"
	"github.com/sirupsen/logrus"
)

func CreateActionDefination(ad *database.ActionDefination) error {
	if err := database.DB.Create(ad).Error; err != nil {
		logrus.Error("CreateActionDefination err ", err.Error())
		return err
	}
	return nil
}

func UpdateActionDefination(ad *database.ActionDefination) error {
	if err := database.DB.Model(&ad).Where("id = ?", ad.Id).Updates(ad).Error; err != nil {
		logrus.Error("UpdateActionDefination err ", err.Error())
		return err
	}
	return nil
}

func DeletActionDefination(ad *database.ActionDefination) error {
	if err := database.DB.Delete(ad).Error; err != nil {
		logrus.Error("DeletActionDefination err ", err.Error())
		return err
	}
	return nil
}

func FindActionDefination(ad *database.ActionDefination) error {
	if err := database.DB.Where("id = ?", ad.Id).First(&ad).Error; err != nil {
		logrus.Error("FindActionDefination err ", err.Error())
		return err
	}
	return nil
}

func FindActionDefinationByFDefId(fdefId int64) ([]*database.ActionDefination, error) {
	aDef := make([]*database.ActionDefination, 0)
	if err := database.DB.Where("flow_defination_id = ? && status = 0", fdefId).Find(&aDef).Error; err != nil {
		logrus.Error("[FindActionDefinationByFDefId] err ", err.Error())
		return nil, err
	}
	return aDef, nil
}

func DeleteActionDefinationById(id int64) error {
	a := database.ActionDefination{
		Status: 1,
	}
	if err := database.DB.Model(&a).Where("id = ?", id).Updates(a).Error; err != nil {
		logrus.Error("[UpdateFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func FindActionDefinationByFDefId2(fdefId int64) (*database.ActionDefination, error) {
	aDef := make([]*database.ActionDefination, 0)
	if err := database.DB.Where("flow_defination_id = ? && status = 0 && action_type = 2", fdefId).Find(&aDef).Error; err != nil {
		logrus.Error("[FindActionDefinationByFDefId] err ", err.Error())
		return nil, err
	}
	if len(aDef) == 0 {
		return nil, errors.New("cant find adef")
	}
	return aDef[0], nil
}

func FindActionDefinationByFDefId1(fdefId int64) (*database.ActionDefination, error) {
	aDef := make([]*database.ActionDefination, 0)
	if err := database.DB.Where("flow_defination_id = ? && status = 0 && action_type = 1", fdefId).Find(&aDef).Error; err != nil {
		logrus.Error("[FindActionDefinationByFDefId] err ", err.Error())
		return nil, err
	}
	return aDef[0], nil
}
