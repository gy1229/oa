package third_server

import (
	"errors"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database/automation"
	"github.com/gy1229/oa/database/stage"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/sirupsen/logrus"
)

type TxtAction struct {
	mod_base.BaseAction
	Param map[string]interface{}
}

const (
	TxtActionSelectFile      = "txt_action_select_file"
	TxtActionSelectFileTitle = "选择文件"
	TxtActionAddContent      = "txt_action_add_content"
	TxtActionAddContentTitle = "新增内容"
)

func (b *TxtAction) PreExecAction(actionId int64, param map[string]interface{}) {
	b.Param = param
	formDatas, err := automation.FindFormDataByADefId(actionId)
	if err != nil {
		return
	}
	for _, v := range formDatas {
		b.Param[v.Key] = v.Value
	}

}

func (b *TxtAction) GetActionName() string {
	return "TxtAction"
}

func (b *TxtAction) GetFrontStruct(userId int64) []*mod_base.FormData {
	//third, _ := data_user.FindThirdAccountByUserId(userId)
	return []*mod_base.FormData{
		{
			Title:    TxtActionSelectFileTitle,
			Key:      constant.ItemSingle,
			Value:    "",
			Position: "0",
			Options:  GenFlieList(userId),
		},
		{
			Title:    TxtActionAddContentTitle,
			Key:      constant.ItemMutiText,
			Value:    "",
			Position: "1",
		},
	}
}

func (b *TxtAction) ExecAction() error {
	//userId, ok := b.Param[mod_base.UserId].(int64)
	//if ok {
	//	return errors.New("cant get UserId")
	//}
	fileId, ok := b.Param[TxtActionSelectFile].(int64)
	if ok {
		logrus.Error("[TxtAction] cant get TxtActionSelectFile")
		return errors.New("cant get TxtActionSelectFile")
	}
	body, ok := b.Param[TxtActionAddContent].(string)
	if ok {
		logrus.Error("[TxtAction] cant get TxtActionAddContent")
		return errors.New("[TxtAction] cant get TxtActionAddContent")
	}
	return addTxtContent(fileId, body)
}
func addTxtContent(fileId int64, body string) error {
	file, err := stage.DGetTextFileByFileId(fileId)
	if err != nil {
		return err
	}
	content := file.Content + "\n" + body
	return stage.DUpdateTextContent(fileId, content, file.Name)
}