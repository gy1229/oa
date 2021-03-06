package user_server

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database"
	database_user "github.com/gy1229/oa/database/user"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/util"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func LoadUserMessage(req *json_struct.LoadUserMessageRequest) (*json_struct.LoadUserMessageResponse, error) {
	user := database.OaUser{}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("LoadUserMessage err ", err.Error())
		return nil, err
	}
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		logrus.Error("[LoadUserMessage] err msg", err.Error())
		return nil, err
	}
	return &json_struct.LoadUserMessageResponse{
		Account:  user.Account,
		UserName: user.UserName,
		ImageId:  strconv.FormatInt(user.ImageId, 10),
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func UpdateUserMessage(req *json_struct.UpdateUserRequest) (*json_struct.UpdateUserResponse, error) {
	err := database_user.UpdateUserMessage(req.Account, req.UserName, req.Password)
	if err != nil {
		return nil, err
	}
	userId, err := database_user.FindUserIdByAccount(req.Account)
	if err != nil {
		return nil, err
	}
	if req.EmailAddr != "" {
		err := database_user.UpdateUserThirdMessage(userId, req.EmailAddr, req.EmailPass)
		if err != nil {
			return nil, err
		}
	}
	return &json_struct.UpdateUserResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func InsertUserMessage(req *json_struct.RegisterUserRequest) (*json_struct.RegisterUserResponse, error) {
	imageId, err := strconv.ParseInt(req.ImageId, 10, 64)
	if err != nil {
		return nil, err
	}
	user := database.OaUser{
		Id:         util.GenId(),
		Account:    req.UserBase.Account,
		UserName:   req.UserName,
		Password:   req.UserBase.Password,
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
		ThirdId:    -1,
		ImageId:    imageId,
	}
	if err := database.DB.Create(user).Error; err != nil {
		logrus.Error("[InsertUserMessage] err msg", err.Error())
		return nil, err
	}
	return &json_struct.RegisterUserResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func LoginUser(req *json_struct.LoginUserRequest) (*json_struct.LoginUserResponse, error) {
	user := database.OaUser{
		Account: req.UserBase.Account,
	}
	if err := database.DB.Model(&user).Where("account = ?", user.Account).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &json_struct.LoginUserResponse{
				Base: &json_struct.BaseResponse{Body: constant.LoginFailAccount},
			}, nil
		}
		logrus.Error("[LoginUser] err msg", err.Error())
		return nil, err
	}
	if user.Password == req.UserBase.Password {
		return &json_struct.LoginUserResponse{
			UserId:  strconv.FormatInt(user.Id, 10),
			ImageId: strconv.FormatInt(user.ImageId, 10),
			Base:    &json_struct.BaseResponse{Body: constant.SUCCESS},
		}, nil
	}
	return &json_struct.LoginUserResponse{
		Base: &json_struct.BaseResponse{Body: constant.LoginFailPassword},
	}, nil
}

func CertainAccount(req *json_struct.CertainAccountRequest) (*json_struct.CertainAccountResponse, error) {
	user := database.OaUser{
		Account: req.Account,
	}
	if err := database.DB.Model(&user).Where("account = ?", user.Account).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &json_struct.CertainAccountResponse{
				Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
			}, nil
		}
		logrus.Error("[CertainAccount] err msg", err.Error())
		return nil, err
	}
	return &json_struct.CertainAccountResponse{
		Base: &json_struct.BaseResponse{Body: constant.RegisterAccountExit},
	}, nil
}
