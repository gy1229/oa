package third_server

import (
	"encoding/json"
	"fmt"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database/automation"
	"github.com/gy1229/oa/gredis"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/gy1229/oa/util"
	"strconv"
)

const (
	WebhookUrl           = "webhook_url"
	WebhookUrlTitle      = "WebHook链接"
	WebhookBody 		= "webhook_body"
	WebhookBodyTitle = "WebHook内容"
	WebhookRowkey = "webhook_rowkey"
)

const (
	WebHookRedisKey = "webhook_redis"
)

type WebhookTrigger struct {
	mod_base.BaseTrigger
	Param map[string]string
}

type WebhookRedisStruct struct {
	Rowkey int64
	FlowDefinationId int64 `json:"flow_defination_id"`
}

func (b *WebhookTrigger) GetTriggerName() string {
	return "WebhookTrigger"
}

func (b *WebhookTrigger) PreInitAction() {
	return
}

func (b *WebhookTrigger) GetFrontStruct(userId int64) []*mod_base.FormData {
	return []*mod_base.FormData{
		{
			Title:    WebhookUrlTitle,
			Key:      constant.ItemText,
			Value:    fmt.Sprintf("http://49.235.180.218:19999/webhook/%d", util.GenId()),
			Position: "0",
			//Options:  GenFlieList(userId),
		},
		//{
		//	Title:    WebhookBodyTitle,
		//	Key:      constant.ItemMutiText,
		//	Value:    "",
		//	Position: "1",
		//},
	}
}

func (b *WebhookTrigger) SetRedisTrigger(fId int64, fd []*mod_base.FormData) error {
	var webhookUrl string
	//var webhookBody string
	for _, v := range fd {
		if v.Title == WebhookUrlTitle {
			webhookUrl = v.Value
		}
	}
	var rowkey int64
	if WebhookUrl == "" {
		return nil
	}
	rowkey, _ = strconv.ParseInt(webhookUrl[36:], 10, 64)
	err := gredis.LPush(WebHookRedisKey, WebhookRedisStruct{
		Rowkey: rowkey,
		FlowDefinationId: fId,
	})
	return err
}

func WebhookTriggerExec(userId int64, param map[string]interface{}) {
	rs, err := gredis.LRange(WebHookRedisKey)
	if err != nil {
		return
	}
	rowkey := param[WebhookRowkey].(int64)
	body := param[WebhookBody].(string)
	for _, v := range rs {
		wrs := &WebhookRedisStruct{}
		err = json.Unmarshal([]byte(v), &wrs)
		if err != nil {
			return
		}
		if wrs.Rowkey == rowkey {
			go webhookExec(userId, wrs.FlowDefinationId, rowkey, body)
		}

	}
}

func webhookExec(userId, flowId, rowkey int64, body string) {
	actionDef, err := automation.FindActionDefinationByFDefId2(flowId)
	if err != nil {
		return
	}
	action := mod_base.ActionGroup[actionDef.ActionId]
	param := make(map[string]interface{})
	param[mod_base.UserId] = userId
	param[WebhookRowkey] = strconv.FormatInt(rowkey, 10)
	param[WebhookBody] = body
	action.PreExecAction(actionDef.Id, param)
	action.ExecAction()
}