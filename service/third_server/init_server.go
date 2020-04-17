package third_server

import "github.com/gy1229/oa/service/mod/mod_base"

func ActionInit() {
	mod_base.ActionGroup[2]= &DemoAction{}
	mod_base.ActionGroup[4]= &DemoAction2{}
}

func TriggerInit() {
	mod_base.TriggerGroup[1] = &DemoTrigger{}
}

func GetActionName(actionId int64) string {
	if actionId % 2 == 0 {
		return mod_base.ActionGroup[actionId].GetActionName()
	}else {
		return mod_base.TriggerGroup[actionId].GetTriggerName()
	}
}

func GetActionIcon(actionId int64) int64 {
	if actionId % 2 == 0 {
		return mod_base.ActionGroup[actionId].GetActionImageId()
	}else {
		return mod_base.TriggerGroup[actionId].GetTriggerImageId()
	}
}