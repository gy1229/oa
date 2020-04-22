package third_server

import (
	"errors"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database/automation"
	data_user "github.com/gy1229/oa/database/user"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"strconv"
	"strings"
)

const (
	SelectAccount      = "select_account"
	SelectAccountTitle = "选择账号"
	EmailRecevier      = "email_recevier"
	EmailRecevierTitle = "收件人"
	EmailSubject       = "email_subject"
	EmailSubjectTitle  = "主题"
	EmailContent       = "email_content"
	EmailContentTitle  = "邮件内容"
)

type MailAction struct {
	mod_base.BaseAction
	Param map[string]interface{}
}

func (b *MailAction) PreExecAction(actionId int64, param map[string]interface{}) {
	b.Param = param
	formDatas, err := automation.FindFormDataByADefId(actionId)
	if err != nil {
		return
	}
	for _, v := range formDatas {
		b.Param[v.Key] = v.Value
	}

}

func (b *MailAction) GetActionName() string {
	return "MailAction"
}

func (b *MailAction) GetFrontStruct(userId int64) []*mod_base.FormData {
	third, _ := data_user.FindThirdAccountByUserId(userId)
	return []*mod_base.FormData{
		{
			Title:    SelectAccountTitle,
			Key:      constant.ItemSingle,
			Value:    "",
			Position: "0",
			Options: []Option{
				{
					Id:    strconv.FormatInt(third.Id, 10),
					Value: third.Account,
				},
			},
		},
		{
			Title:    EmailRecevierTitle,
			Key:      constant.ItemText,
			Value:    "",
			Position: "1",
		},
		{
			Title:    EmailSubjectTitle,
			Key:      constant.ItemText,
			Value:    "",
			Position: "2",
		},
		{
			Title:    EmailContentTitle,
			Key:      constant.ItemMutiText,
			Value:    "",
			Position: "3",
		},
	}
}

func (b *MailAction) ExecAction() error {
	userId, ok := b.Param[mod_base.UserId].(int64)
	if ok {
		return errors.New("cant get user_id")
	}
	third, err := data_user.FindThirdAccountByUserId(userId)
	if err != nil {
		return err
	}
	recevier, ok := b.Param[EmailRecevier].(string)
	if ok {
		return errors.New("cant get sender")
	}
	body, ok := b.Param[EmailContent].(string)
	if ok {
		return errors.New("cant get EmailContent")
	}
	subject, ok := b.Param[EmailSubject].(string)
	if ok {
		return errors.New("cant get EmailSubject")
	}
	receviers := strings.Split(recevier, ";")
	return SendEmail(third.Account, third.Password, subject, body, receviers)
}

func SendEmail(sender, pass, subject, body string, mailTo []string) error {
	mailConn := map[string]string{
		"user": sender,
		"pass": pass,
		"host": "smtp.126.com",
		"port": "25",
	}
	port := 25
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "Automation Flow"))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	if err != nil {
		logrus.Error("[SendEmail] DialAndSend err msg :", err.Error())
		return err
	}
	return err
}
