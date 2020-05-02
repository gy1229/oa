package third_server

import (
	"errors"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database/automation"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	WebhookActionAidUrl      = "webhook_action_aid_url"
	WebhookActionAidUrlTitle = "目标地址"
	WebhookActionHeader = "webhook_action_header"
	WebhookActionHeaderTitle = "请求头"
	WebhookActionBody = "webhook_action_body"
	WebhookActionBodyTitle = "请求体"
)

const (
	WebhookContentTypeJson = "content-type = application/json"
	WebhookContentTypeForm = "content-type = multipart/form-data"
)

type WebhookAction struct {
	mod_base.BaseAction
	Param map[string]interface{}
}

func (b *WebhookAction) PreExecAction(actionId int64, param map[string]interface{}) {
	b.Param = param
	formDatas, err := automation.FindFormDataByADefId(actionId)
	if err != nil {
		return
	}
	for _, v := range formDatas {
		b.Param[v.Key] = v.Value
	}

}

func (b *WebhookAction) GetActionName() string {
	return "WebhookAction"
}


func (b *WebhookAction) GetFrontStruct(userId int64) []*mod_base.FormData {
	//third, _ := data_user.FindThirdAccountByUserId(userId)
	return []*mod_base.FormData{
		{
			Title:    WebhookActionAidUrlTitle,
			Key:      constant.ItemText,
			Value:    "",
			Position: "0",
		},
		{
			Title:    WebhookActionHeaderTitle,
			Key:      constant.ItemMutiText,
			Value:    "",
			Position: "1",
			Options: []*mod_base.Option{
				{
					Id:"0",
					Value: WebhookContentTypeJson,
				},
				{
					Id:"1",
					Value:WebhookContentTypeForm,
				},
			},
		},
		{
			Title:    WebhookActionBodyTitle,
			Key:      constant.ItemMutiText,
			Value:    "",
			Position: "2",
		},
	}
}

func (b *WebhookAction) ExecAction() error {

	aidUrl, ok := b.Param[WebhookActionAidUrl].(string)
	if ok {
		logrus.Error("[WebhookActionAidUrl] cant get WebhookActionAidUrl")
		return errors.New("cant get sender")
	}
	body, ok := b.Param[WebhookActionBody].(string)
	if ok {
		logrus.Error("[WebhookActionBody] cant get EmailContent")
		return errors.New("cant get EmailContent")
	}
	return SenHttpRequest(aidUrl, body)
}

func SenHttpRequest(url string, body string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST",
		url,
		strings.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	return err
}

