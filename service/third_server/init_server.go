package third_server

import "github.com/gy1229/oa/service/mod/mod_base"

func ActionInit() {
	mod_base.ActionGroup[2]= &DemoAction{}
}

func TriggerInit() {

}
