package database_user

import (
	"github.com/gy1229/oa/database"
	"github.com/sirupsen/logrus"
)

func GetUserNameById(id int64) (string, error) {
	user := &database.OaUser{
		Id: id,
	}
	if err := database.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		logrus.Error("[GetUserNameById] err ", err.Error())
		return "", err
	}
	return user.UserName, nil
}

func FindUserIdByAccount(account string) (int64, error) {
	user := &database.OaUser{}
	if err := database.DB.Where("account = ?", account).Find(&user).Error; err != nil {
		logrus.Error("[FindUserIdByAccount] err ", err.Error())
		return 0, err
	}
	return user.Id, nil
}
func UpdateUserMessage(account, userName, userPassword string) error {
	user := database.OaUser{
		Password: userPassword,
		UserName: userName,
	}
	if err := database.DB.Model(&user).Where("account = ?", account).Updates(&user).Error; err != nil {
		logrus.Error("[UpdateUserMessage] err msg", err.Error())
		return err
	}
	return nil
}

func UpdateUserThirdMessage(userId int64, emailAddr, emailPass string) error{
	third := database.Third{
		Account: emailAddr,
		Password: emailPass,
	}
	if err := database.DB.Model(&third).Where("user_id = ?", userId).Updates(&third).Error; err != nil {
		logrus.Error("[UpdateUserThirdMessage] err msg", err.Error())
		return err
	}
	return nil
}

func FindThirdAccountByUserId(userId int64) (*database.Third, error) {
	third := &database.Third{}
	if err := database.DB.Model(&third).Where("user_id = ?", userId).Find(&third).Error; err != nil {
		logrus.Error("[FindThirdAccountByUserId] err msg", err.Error())
		return nil, err
	}
	return third, nil
}