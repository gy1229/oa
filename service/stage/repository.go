package stage

import (
	"fmt"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database"
	data_user "github.com/gy1229/oa/database/user"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/util"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func CreateRepository(req *json_struct.CreateRepositoryRequest) (*json_struct.CreateRepositoryResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	authority, err := strconv.Atoi(req.Authority)
	if err != nil {
		return nil, err
	}
	id := util.Int64toP(util.GenId())
	repository := database.StageRepository{
		Id:         id,
		Name:       req.Name,
		CreatorId:  &userId,
		Authority:  &authority,
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
		Status:     util.Int2P(0),
	}
	if err := database.DB.Create(repository).Error; err != nil {
		logrus.Error("CreateRepository err ", err.Error())
		return nil, err
	}
	return &json_struct.CreateRepositoryResponse{
		Id: fmt.Sprintf("%d", *id),
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func GetRepositoryList(req *json_struct.GetRepositoryListRequest) (*json_struct.GetRepositoryListResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	stage := make([]*database.StageRepository, 0)
	if err := database.DB.Where("creator_id = ? AND status = 0", userId).Find(&stage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &json_struct.GetRepositoryListResponse{
				RepositoryList: nil,
				Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
			}, nil
		}
		logrus.Error("GetRepositoryList err ", err.Error())
		return nil, err
	}
	repList := make([]*json_struct.Repository, 0)
	for k, v := range stage {
		repList = append(repList, &json_struct.Repository{
			Id:   strconv.FormatInt(*v.Id,10),
			Name: v.Name,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		})
		repList[k].CreatorName, _ = data_user.GetUserNameById(v.CreatorId)
	}
	return &json_struct.GetRepositoryListResponse{
		RepositoryList: repList,
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func UpdateRepository(req *json_struct.UpdateRepositoryRequest) (*json_struct.UpdateRepositoryResponse, error) {
	repId, err := strconv.ParseInt(req.RepositoryId, 10, 64)
	if err != nil {
		return nil, err
	}
	authority, err := strconv.Atoi(req.Authority)
	if err != nil {
		return nil, err
	}
	rep := database.StageRepository{
		Name:req.Name,
		Authority:&authority,
	}
	if err := database.DB.Model(&rep).Where("id = ?", repId).Updates(rep).Error; err != nil {
		logrus.Error("UpdateRepository err ", err.Error())
		return nil, err
	}
	return &json_struct.UpdateRepositoryResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func DelRepository(req *json_struct.DelRepositoryRequest) (*json_struct.DelRepositoryResponse, error) {
	repId, err := strconv.ParseInt(req.RepositoryId, 10, 64)
	if err != nil {
		return nil, err
	}
	rep := database.StageRepository{
		Status: util.Int2P(1),
	}
	if err := database.DB.Model(&rep).Where("id = ?", repId).Updates(rep).Error; err != nil {
		logrus.Error("UpdateRepository err ", err.Error())
		return nil, err
	}
	return &json_struct.DelRepositoryResponse{
		Base: &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}