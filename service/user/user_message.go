package user

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/util"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

func LoadUserMessage(req *json_struct.LoadUserMessageRequest) (*json_struct.LoadUserMessageResponse, error) {
	user := database.OaUser{}
	if err := database.DB.Where("id = ?", req.UserId).First(&user).Error; err != nil {
		logrus.Error("LoadUserMessage err ", err.Error())
		return nil, err
	}
	return &json_struct.LoadUserMessageResponse{
		Account:  user.Account,
		UserName: user.UserName,
		Base: &json_struct.BaseResponse{Body:constant.SUCCESS},
	}, nil
}

func UpdateUserMessage(req *json_struct.UpdateUserRequest) (*json_struct.UpdateUserResponse, error) {
	user := database.OaUser{
		Password: req.Password,
		UserName: req.UserName,
	}
	if err := database.DB.Model(&user).Where("account = ?", req.Account).Updates(user).Error; err != nil {
		logrus.Error("UpdateUserMessage err ", err.Error())
		return nil, err
	}
	return &json_struct.UpdateUserResponse{
		Base: &json_struct.BaseResponse{Body:constant.SUCCESS},
	}, nil
}

func InsertUserMessage(req *json_struct.RegisterUserRequest) (*json_struct.RegisterUserResponse, error) {
	user := database.OaUser{
		Id:        util.Int64toP(util.GenId()),
		Account:    req.UserBase.Account,
		UserName:   req.UserName,
		Password:   req.UserBase.Password,
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
		ThirdId:    util.Int64toP(-1),
		ImageId:    util.Int64toP(-1),
	}
	if err := database.DB.Create(user).Error; err != nil {
		logrus.Error("InsertUserMessage err ", err.Error())
		return nil, err
	}
	return &json_struct.RegisterUserResponse{
		Base: &json_struct.BaseResponse{Body:constant.SUCCESS},
	}, nil
}

func LoginUser(req *json_struct.LoginUserRequest) (*json_struct.LoginUserResponse, error) {
	user := database.OaUser{
		Account: req.UserBase.Account,
	}
	if err := database.DB.Model(&user).Where("account = ?", user.Account).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &json_struct.LoginUserResponse{
				Base: &json_struct.BaseResponse{Body:constant.LoginFailAccount},
			}, nil
		}
		logrus.Error("UpdateUserMessage err ", err.Error())
		return nil, err
	}
	if user.Password == req.UserBase.Password {
		return &json_struct.LoginUserResponse{
			UserId: user.Id,
			Base: &json_struct.BaseResponse{Body:constant.SUCCESS},
		}, nil
	}
	return &json_struct.LoginUserResponse{
		Base: &json_struct.BaseResponse{Body:constant.LoginFailPassword},
	}, nil
}

func CertainAccount(req *json_struct.CertainAccountRequest) (*json_struct.CertainAccountResponse, error) {
	user := database.OaUser{
		Account:    req.Account,
	}
	if err := database.DB.Model(&user).Where("account = ?", user.Account).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &json_struct.CertainAccountResponse{
				Base: &json_struct.BaseResponse{Body:constant.SUCCESS},
			}, nil
		}
		logrus.Error("UpdateUserMessage err ", err.Error())
		return nil, err
	}
	return &json_struct.CertainAccountResponse{
		Base: &json_struct.BaseResponse{Body:constant.RegisterAccountExit},
	}, nil
}