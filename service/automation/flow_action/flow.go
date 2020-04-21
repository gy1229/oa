package flow_action

import (
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/gy1229/oa/util"
	"github.com/sirupsen/logrus"
	"strconv"
)

func SetOfficeFlow(fd int64, ActionList []*json_struct.ActionDetail) {
	triggerId, err := strconv.ParseInt(ActionList[0].ActionId, 10, 64)
	if err != nil {
		logrus.Error("[SetOfficeFlow] ParseInt triggerId err ", err.Error())
		return
	}
	logrus.Info("[SetOfficeFlow] trigger id ", triggerId, "flow defination id", fd)
	trigger := mod_base.TriggerGroup[triggerId]
	formData := make([]*mod_base.FormData, 0)
	util.TranHttpStruct2Database(ActionList[0].BehaviorInstanceList, &formData)
	err = trigger.SetRedisTrigger(fd, formData)
	if err != nil {
		logrus.Error("[SetOfficeFlow] SetRedisTrigger err ", err.Error())
		return
	}
}
