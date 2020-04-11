package automation

import (
	"github.com/gy1229/oa/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func CreateAction(action *database.Action) error {
	if err := database.DB.Create(action).Error; err != nil {
		logrus.Error("[CreateAction] err ", err.Error())
		return err
	}
	return nil
}

func UpdateAction(action *database.Action) error {
	if err := database.DB.Model(&action).Where("id = ?", action.Id).Updates(action).Error; err != nil {
		logrus.Error("[UpdateAction] err ", err.Error())
		return err
	}
	return nil
}

func DeletAction(action *database.Action) error {
	if err := database.DB.Delete(action).Error; err != nil {
		logrus.Error("[DeletAction] err ", err.Error())
		return err
	}
	return nil
}

func FindAction(eduEx *database.Action) error {
	if err := database.DB.Where("id = ?", eduEx.Id).First(&eduEx).Error; err != nil {
		logrus.Error("[FindAction] err ", err.Error())
		return err
	}
	return nil
}

func FindAllAction() ([]*database.Action, error) {
	action := make([]*database.Action, 0)
	if err := database.DB.Model(&action).Where("status = 0").Find(&action).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return action, nil
		}
		logrus.Error("[FindAllAction] err msg", err.Error())
		return nil, err
	}
	return action, nil
}