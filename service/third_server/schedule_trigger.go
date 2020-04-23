package third_server

import (
	"encoding/json"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database/automation"
	"github.com/gy1229/oa/gredis"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/gy1229/oa/util"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	ScheduleTimeSelect        = "schedule_time_select"
	ScheduleTimeSelectTitle   = "开始时间"
	ScheduleExecFunction      = "schedule_exec_function"
	ScheduleExecFunctionTitle = "触发方式"

	NowTimeStamp = "now_time_stamp"

	PerHour  = "每时"
	PerDay   = "每天"
	PerWeek  = "每周"
	PerMonth = "每月"

	PerHourInt  = 60 * 60
	PerDayInt   = PerHourInt * 24
	PerWeekInt  = PerDayInt * 7
	PerMonthInt = PerDayInt * 30
)

const (
	ScheduleRedisKey = "schedule_redis"
)

type ScheduleTrigger struct {
	mod_base.BaseTrigger
	Param map[string]string
}

type ScheduleRedisStruct struct {
	FlowDefinationId int64
	StartTime        int64
	ExecInterval     int64
}

func (b *ScheduleTrigger) GetTriggerName() string {
	return "ScheduleTrigger"
}

func (b *ScheduleTrigger) PreInitAction() {
	return
}

func (b *ScheduleTrigger) GetFrontStruct(userId int64) []*mod_base.FormData {
	return []*mod_base.FormData{
		{
			Title:    ScheduleTimeSelectTitle,
			Key:      constant.ItemSingle,
			Value:    "",
			Position: "0",
			Options:  GenFlieList(userId),
		},
		{
			Title:    ScheduleExecFunctionTitle,
			Key:      constant.ItemSingle,
			Value:    "",
			Position: "1",
			Options: []Option{
				{
					Id:    "0",
					Value: PerHour,
				},
				{Id: "1",
					Value: PerDay,
				}, {
					Id:    "2",
					Value: PerWeek,
				}, {
					Id:    "3",
					Value: PerMonth,
				},
			},
		},
	}
}

func (b *ScheduleTrigger) SetRedisTrigger(fId int64, fd []*mod_base.FormData) error {
	var startTime int64
	var startInterval int64
	var err error
	for _, v := range fd {
		switch v.Title {
		case ScheduleTimeSelectTitle:
			timeStamp := util.TranUnixString2TimeStamp(v.Value)
			startTime = timeStamp - (timeStamp % 60)
		case ScheduleExecFunctionTitle:
			switch v.Value {
			case "0":
				startInterval = PerHourInt
			case "1":
				startInterval = PerDayInt
			case "2":
				startInterval = PerWeekInt
			case "3":
				startInterval = PerMonthInt
			default:
				startInterval = 100000
			}
		default:
			logrus.Error("[SetRedisTrigger] cant find title", v.Title)
			return nil
		}
	}
	err = gredis.LPush(ScheduleRedisKey, ScheduleRedisStruct{
		StartTime:        startTime,
		ExecInterval:     startInterval,
		FlowDefinationId: fId,
	})
	return err
}

func ScheduleTriggerExec(userId int64, param map[string]interface{}) {
	rs, err := gredis.LRange(ScheduleRedisKey)
	if err != nil {
		return
	}
	timeStamp := param[NowTimeStamp].(int64)
	if err != nil {
		return
	}
	for _, v := range rs {
		srs := &ScheduleRedisStruct{}
		err = json.Unmarshal([]byte(v), &srs)
		if err != nil {
			return
		}
		if timeStamp-srs.StartTime%srs.ExecInterval == 0 {
			go scheduleExec(timeStamp, userId, srs.FlowDefinationId)
		}
	}
}

func scheduleExec(timeStamp, userId, flowId int64) {
	actionDef, err := automation.FindActionDefinationByFDefId2(flowId)
	if err != nil {
		return
	}
	action := mod_base.ActionGroup[actionDef.ActionId]
	param := make(map[string]interface{})
	param[mod_base.UserId] = userId
	param[NowTimeStamp] = timeStamp
	action.PreExecAction(actionDef.ActionId, param)
	action.ExecAction()
}

func ScheduleTaskStart() {

	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				timeStamp := time.Now().Unix()
				param := make(map[string]interface{})
				param[NowTimeStamp] = timeStamp - (timeStamp % 60)
				go ScheduleTriggerExec(0, param)
			}
		}
	}()
}
