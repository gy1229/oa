package automation

import (
	"github.com/gy1229/oa/database"
	"github.com/sirupsen/logrus"
)

func CreateFormData(fi *database.FormData) error {
	if err := database.DB.Create(fi).Error; err != nil {
		logrus.Error("[CreateFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func UpdateFormData(fi *database.FormData) error {
	if err := database.DB.Model(&fi).Where("id = ?", fi.Id).Updates(fi).Error; err != nil {
		logrus.Error("[UpdateFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func DeleteFormData(fi *database.FormData) error {
	if err := database.DB.Delete(fi).Error; err != nil {
		logrus.Error("[DeletFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func FindFormData(fi *database.FormData) error {
	if err := database.DB.Where("id = ?", fi.Id).First(&fi).Error; err != nil {
		logrus.Error("[FindFlowInstance] err ", err.Error())
		return err
	}
	return nil
}

func FindFormDataByADefId(aDefId int64) ([]*database.FormData, error) {
	fData := make([]*database.FormData, 0)
	if err := database.DB.Model(&database.FormData{}).Where("action_defination_id = ? AND status = 0", aDefId).Find(&fData).Error; err != nil {
		logrus.Error("[FindActionDefinationByFDefId] err ", err.Error())
		return nil, err
	}
	return fData, nil
}

func DeleteFormDataById(id int64) error {
	f := database.FormData{
		Status: 1,
	}
	if err := database.DB.Model(&f).Where("id = ?", id).Updates(f).Error; err != nil {
		logrus.Error("[UpdateFlowInstance] err ", err.Error())
		return err
	}
	return nil
}
