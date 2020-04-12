package third_server

import (
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/database/automation"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/gy1229/oa/test"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestInitAction(t *testing.T) {
	test.InitTestConfig2()
	ActionInit()
	TriggerInit()
	aId := int64(2)
	InitAction(aId)
}

func InitAction(actionId int64) {
	da := &database.Action{
		Id: actionId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if actionId % 2 == 0 {
		action := mod_base.ActionGroup[actionId]
		da.Name = action.GetActionName()
		da.ActionType = action.GetActionType()
		da.ImageId = action.GetActionImageId()
	}else {
		action := mod_base.ActionGroup[actionId]
		da.Name = action.GetActionName()
		da.ActionType = action.GetActionType()
		da.ImageId = action.GetActionImageId()
	}
	err := automation.CreateAction(da)
	if err != nil {
		logrus.Error("err ", err.Error())
	}
}
