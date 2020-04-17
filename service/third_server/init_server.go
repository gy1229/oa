package third_server

import "github.com/gy1229/oa/service/mod/mod_base"

func ActionInit() {
	mod_base.ActionGroup[2]= &DemoAction{}
	mod_base.ActionGroup[4]= &DemoAction2{}
}

func TriggerInit() {
	mod_base.TriggerGroup[1] = &DemoTrigger{}
}
