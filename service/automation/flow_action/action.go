package flow_action

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/database/automation"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/gy1229/oa/util"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func GetActionList(req *json_struct.GetActionListRequest) (*json_struct.GetActionListResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[GetActionList] userId ParseInt")
		return nil, err
	}
	logrus.Info("[GetActionList] userid : ", userId)
	action, err := automation.FindAllAction()
	if err != nil {
		logrus.Error("[GetActionList] FindAllAction err", err.Error())
		return nil, err
	}
	actionArr := make([]*json_struct.Action, 0)
	for _, v := range action {
		actionArr = append(actionArr, &json_struct.Action{
			ActionId:   strconv.FormatInt(v.Id, 10),
			ActionIcon: strconv.FormatInt(v.ImageId, 10),
			ActionName: v.Name,
		})
	}
	return &json_struct.GetActionListResponse{
		ActionList: actionArr,
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func GetActionDefination(req *json_struct.GetActionDefinationRequest) (*json_struct.GetActionDefinationResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[GetActionDefination] userId ParseInt")
		return nil, err
	}
	logrus.Info("[GetActionDefination] userid : ", userId)
	actionId, err := strconv.ParseInt(req.ActionId, 10, 64)
	if err != nil {
		logrus.Error("[GetActionDefination] userId ParseInt")
		return nil, err
	}
	logrus.Info("[GetActionDefination] actionId : ", actionId)
	formData := make([]*mod_base.FormData, 0)
	switch actionId%2 {
	case 1:
		trigger := mod_base.TriggerGroup[actionId]
		formData = trigger.GetFrontStruct()
	default:
		action := mod_base.ActionGroup[actionId]
		formData = action.GetFrontStruct()
	}
	rFormData := make([]*json_struct.FormData, 0 )
	util.TranHttpStruct2Database(formData, &rFormData)
	return &json_struct.GetActionDefinationResponse{
		BehaviorDefinationList: rFormData,
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil

}

func GetFlowDefinationDetail(req *json_struct.GetFlowDefinationDetailRequest) (*json_struct.GetFlowDefinationDetailResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[GetFlowDefinationDetail] userId ParseInt")
		return nil, err
	}
	logrus.Info("[GetFlowDefinationDetail] userid : ", userId)
	flowDefinationId, err := strconv.ParseInt(req.FlowDefinationId, 10, 64)
	if err != nil {
		logrus.Error("[GetFlowDefinationDetail] userId ParseInt")
		return nil, err
	}
	logrus.Info("[GetFlowDefinationDetail] userid : ", flowDefinationId)
	actionDetail, err := GetActionDetailsByFlowDefid(flowDefinationId)
	if err != nil {
		logrus.Error("[GetFlowDefinationDetail] GetActionDetailsByFlowDefid ParseInt")
		return nil, err
	}
	return &json_struct.GetFlowDefinationDetailResponse{
		ActionList: actionDetail,
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func CreateFlowDefination(req *json_struct.CreateFlowDefinationRequest) (*json_struct.CreateFlowDefinationResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[CreateFlowDefination] userId ParseInt")
		return nil, err
	}
	logrus.Info("[CreateFlowDefination] userid : ", userId)
	flowdefId := util.GenId()
	if err := automation.CreateFlowDefinationByArgs(flowdefId, userId, req.FlowDefinationName); err != nil {
		logrus.Error("[CreateFlowDefination] CreateFlowDefinationByArgs err", err.Error())
		return nil, err
	}
	err = CreateActionDetail(req.ActionList, flowdefId)
	if err != nil {
		logrus.Error("[CreateFlowDefination] CreateActionDetail err", err.Error())
		return nil, err
	}
	return &json_struct.CreateFlowDefinationResponse{
		FlowDefinationId: strconv.FormatInt(flowdefId, 10),
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func UpdateFlowDefination(req *json_struct.UpdateFlowDefinationRequest) (*json_struct.UpdateFlowDefinationResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[UpdateFlowDefination] userId ParseInt")
		return nil, err
	}
	logrus.Info("[UpdateFlowDefination] userid : ", userId)
	flowdefId, err := strconv.ParseInt(req.FlowDefinationId, 10, 64)
	if err != nil {
		logrus.Error("[UpdateFlowDefination] flowdefId ParseInt")
		return nil, err
	}
	logrus.Info("[UpdateFlowDefination] userid : ", flowdefId)
	err = automation.UpdateFlowDefinationNameById(flowdefId, req.FlowDefinationName)
	if err != nil {
		logrus.Error("[UpdateFlowDefination] UpdateFlowDefinationNameById err", err.Error())
		return nil, err
	}
	actionDetial, err := GetActionDetailsByFlowDefid(flowdefId)
	if err != nil {
		logrus.Error("[UpdateFlowDefination] GetActionDetailsByFlowDefid err", err.Error())
		return nil, err
	}
	err = DeleteByActionDetail(actionDetial)
	if err != nil {
		logrus.Error("[UpdateFlowDefination] DeleteByActionDetail err", err.Error())
		return nil, err
	}
	err = CreateActionDetail(req.ActionList, flowdefId)
	if err != nil {
		logrus.Error("[CreateFlowDefination] CreateActionDetail err", err.Error())
		return nil, err
	}
	return &json_struct.UpdateFlowDefinationResponse{
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func GetFlowDefinationList(req *json_struct.GetFlowDefinationListRequest) (*json_struct.GetFlowDefinationListResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[GetFlowDefinationList] userId ParseInt")
		return nil, err
	}
	logrus.Info("[GetFlowDefinationList] userid : ", userId)
	fds, err := automation.FindFlowDefinationByUserId(userId)
	if err != nil {
		logrus.Error("[GetFlowDefinationList] FindFlowDefinationByUserId err ", err.Error())
		return nil, err
	}
	fdDetails := make([]*json_struct.FlowDefination, 0)
	for _, v := range fds {
		fdDetails = append(fdDetails, &json_struct.FlowDefination{
			FlowDefinationId:   strconv.FormatInt(v.Id, 10),
			FlowDefinationName: v.Name,
		})
	}
	return &json_struct.GetFlowDefinationListResponse{
		FlowDefinationList: fdDetails,
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func DeleteFlowDeination(req *json_struct.DeleteFlowDeinationRequest) (*json_struct.DeleteFlowDeinationResponse, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		logrus.Error("[DeleteFlowDeination] userId ParseInt")
		return nil, err
	}
	logrus.Info("[DeleteFlowDeination] userid : ", userId)
	flowdefId, err := strconv.ParseInt(req.FlowDefinationId, 10, 64)
	if err != nil {
		logrus.Error("[DeleteFlowDeination] flowdefId ParseInt")
		return nil, err
	}
	logrus.Info("[DeleteFlowDeination] userid : ", flowdefId)
	if err := automation.DelFlowDefinationById(flowdefId); err != nil {
		logrus.Error("[DeleteFlowDeination] DelFlowDefinationById err", err.Error())
		return nil, err
	}
	actionDetial, err := GetActionDetailsByFlowDefid(flowdefId)
	if err != nil {
		logrus.Error("[DeleteFlowDeination] GetActionDetailsByFlowDefid err", err.Error())
		return nil, err
	}
	err = DeleteByActionDetail(actionDetial)
	if err != nil {
		logrus.Error("[DeleteFlowDeination] DeleteByActionDetail err", err.Error())
		return nil, err
	}
	return &json_struct.DeleteFlowDeinationResponse{
		Base:     &json_struct.BaseResponse{Body: constant.SUCCESS},
	}, nil
}

func GetActionDetailsByFlowDefid(flowdefId int64) ([]*json_struct.ActionDetail, error) {
	actionList, err :=automation.FindActionDefinationByFDefId(flowdefId)
	if err != nil {
		logrus.Error("[GetActionDetailsByFlowDefid] FindActionDefinationByFDefId  err:", err.Error())
		return nil, err
	}
	actionDetail := make([]*json_struct.ActionDetail, 0)
	for _, v := range actionList {
		fdatas, err := automation.FindFormDataByADefId(v.Id)
		if err != nil {
			logrus.Error("[GetActionDetailsByFlowDefid] FindFormDataByADefId  err:", err.Error())
			return nil, err
		}
		jf := make([]*json_struct.FormData, 0)
		for _, f := range fdatas {
			jf = append(jf, &json_struct.FormData{
				Id: strconv.FormatInt(f.Id, 10),
				Key:      f.Key,
				Value:    f.Value,
				Title:   f.Title,
				Position: strconv.Itoa(f.Position),
			})
		}
		actionDetail = append(actionDetail, &json_struct.ActionDetail{
			ActionId:             strconv.FormatInt(v.Id, 10),
			ActionName:           "test",
			ActionIcon:           "icon",
			ActionPosition:       strconv.Itoa(v.Position),
			ActionType:           v.ActionType,
			BehaviorInstanceList: jf,
		})
	}
	return actionDetail, nil
}

func DeleteByActionDetail(actionDetail []*json_struct.ActionDetail) error {
	for _, v := range actionDetail {
		actionId, _ := strconv.ParseInt(v.ActionId, 10, 64)
		if err := automation.DeleteActionDefinationById(actionId); err != nil {
			return err
		}
		for _, f := range v.BehaviorInstanceList {
			formDataId, _ := strconv.ParseInt(f.Id, 10, 64)
			if err := automation.DeleteFormDataById(formDataId); err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateActionDetail(ActionList []*json_struct.ActionDetail, flowDefId int64) error {
	for _, action := range ActionList {
		actionId := util.GenId()
		position, err := strconv.Atoi(action.ActionPosition)
		if err != nil {
			logrus.Error("[CreateFlowDefination] Parse ActionPosition err", err.Error())
			return err
		}
		ad := database.ActionDefination{
			Id:               actionId,
			FlowDefinationId: flowDefId,
			Position:         position,
			ActionType:       action.ActionType,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		if err := automation.CreateActionDefination(&ad); err != nil {
			logrus.Error("[CreateFlowDefination] CreateActionDefination err", err.Error())
			return err
		}
		for _, bh := range action.BehaviorInstanceList {
			bPosition, err := strconv.Atoi(bh.Position)
			if err != nil {
				logrus.Error("[CreateFlowDefination] Parse BHPosition err", err.Error())
				return err
			}
			fd := database.FormData{
				Id:                 util.GenId(),
				ActionDefinationId: actionId,
				Key:                bh.Key,
				Value:              bh.Value,
				Position:           bPosition,
				Title: 				bh.Title,
			}
			if err := automation.CreateFormData(&fd); err != nil {
				logrus.Error("[CreateFlowDefination] CreateFormData err", err.Error())
				return err
			}
		}
	}
	return nil
}