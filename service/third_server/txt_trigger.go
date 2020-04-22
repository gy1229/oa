package third_server

import (
	"encoding/json"
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/database/automation"
	"github.com/gy1229/oa/database/stage"
	"github.com/gy1229/oa/gredis"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/sirupsen/logrus"
	"strconv"
)

const (
	TOContentChange  = "文本内容变更"
	TOPropertyChange = "文件属性变更"
)

const (
	TxtSelectFile           = "txt_select_file"
	TxtSelectFileTitle      = "选择文件"
	TxtTriggerFunction      = "txt_trigger_function"
	TxtTriggerFunctionTitle = "触发方式"
)

const (
	TxtRedisKey = "txt_redis"
)

const (
	NewTxtContent = "new_txt_content"
	OldTxtContent = "old_txt_content"
)

type TxtFileTrigger struct {
	mod_base.BaseTrigger
	Param map[string]string
}

type Option struct {
	Id    string
	Value string
}

type TxtRedisStruct struct {
	FileId           int64
	Change           int
	FlowDefinationId int64
}

func (b *TxtFileTrigger) GetTriggerName() string {
	return "TxtFileTrigger"
}

func (b *TxtFileTrigger) PreInitAction() {
	return
}

func (b *TxtFileTrigger) GetFrontStruct(userId int64) []*mod_base.FormData {
	return []*mod_base.FormData{
		{
			Title:    TxtSelectFileTitle,
			Key:      constant.ItemSingle,
			Value:    "",
			Position: "0",
			Options:  GenFlieList(userId),
		},
		{
			Title:    TxtTriggerFunctionTitle,
			Key:      constant.ItemSingle,
			Value:    "",
			Position: "1",
			Options:  []string{TOContentChange, TOPropertyChange},
		},
	}
}

func (b *TxtFileTrigger) SetRedisTrigger(fId int64, fd []*mod_base.FormData) error {
	var fileId int64
	var change int
	var err error
	for _, v := range fd {
		if v.Title == TxtSelectFileTitle {
			fileId, err = strconv.ParseInt(v.Value, 10, 64)
			if err != nil {
				logrus.Error("[SetRedisTrigger] file id cannt find, fId", fId)
			}
		} else {
			if v.Value == TOContentChange {
				change = 1
			} else {
				change = 2
			}
		}
	}
	err = gredis.LPush(TxtRedisKey, TxtRedisStruct{
		FileId:           fileId,
		Change:           change,
		FlowDefinationId: fId,
	})
	return err
}

func GenFlieList(userId int64) []Option {
	files, err := stage.DGetFileListByUserId(userId)
	if err != nil {
		return nil
	}
	op := make([]Option, 0)
	for _, v := range files {
		op = append(op, Option{
			Id:    strconv.FormatInt(v.Id, 10),
			Value: v.Name,
		})
	}
	return op
}

func TxtTriggerExec(userId int64, parma map[string]interface{}) {
	rs, err := gredis.LRange(TxtRedisKey)
	if err != nil {
		return
	}
	fileId := parma[mod_base.FileId].(int64)
	fileText, err := stage.DGetTextFileByFileId(fileId)
	if err != nil {
		return
	}
	for _, v := range rs {
		trs := &TxtRedisStruct{}
		err = json.Unmarshal([]byte(v), &trs)
		if err != nil {
			return
		}
		if trs.FileId == fileId {
			if fileText.Content != parma[NewTxtContent].(string) {
				TOContentChangeExec(fileText.Content, parma[NewTxtContent].(string), fileText.Name, trs.FlowDefinationId, userId)
			}
		}
	}
}

func TOContentChangeExec(old, new, name string, flowId, userId int64) {
	actionDef, err := automation.FindActionDefinationByFDefId2(flowId)
	if err != nil {
		return
	}
	action := mod_base.ActionGroup[actionDef.ActionId]
	param := make(map[string]interface{})
	param[mod_base.UserId] = userId

	// ---------- trigger 属性 --------
	param[NewTxtContent] = new
	param[OldTxtContent] = old
	param[TxtSelectFile] = name
	param[TxtTriggerFunction] = TOContentChange

	// ---------- action 属性 --------
	action.PreExecAction(actionDef.ActionId, param)
	action.ExecAction()
}

func TOPropertyChangeExec() {

}
