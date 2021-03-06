package third_server

import (
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/sirupsen/logrus"
)

func ActionInit() {
	ScheduleTaskStart()
	mod_base.ActionGroup[2] = &DemoAction{}
	mod_base.ActionGroup[4] = &MailAction{}
	mod_base.ActionGroup[6] = &TxtAction{}
	mod_base.ActionGroup[8] = &WebhookAction{}
}

func TriggerInit() {
	mod_base.TriggerGroup[1] = &DemoTrigger{}
	mod_base.TriggerGroup[3] = &TxtFileTrigger{}
	mod_base.TriggerGroup[5] = &ScheduleTrigger{}
	mod_base.TriggerGroup[7] = &WebhookTrigger{}
}

func GetActionName(actionId int64) string {
	logrus.Info("[GetActionName] actionId : ", actionId)
	if actionId%2 == 0 {
		return mod_base.ActionGroup[actionId].GetActionName()
	} else {
		return mod_base.TriggerGroup[actionId].GetTriggerName()
	}
}

func GetActionIcon(actionId int64) int64 {
	if actionId%2 == 0 {
		return mod_base.ActionGroup[actionId].GetActionImageId()
	} else {
		return mod_base.TriggerGroup[actionId].GetTriggerImageId()
	}
}
