package third_server

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/gredis"
	"github.com/gy1229/oa/service/mod/mod_base"
	"github.com/sirupsen/logrus"
	"strconv"
)

const (
	TOContentChange = "文本内容变更"
	TOProperty = "文件属性变更"
)

const (
	txtSelectFile = "选择文件"
	txtTriggerFunction = "触发方式"
)

const (
	TxtRedisKey = "txt_redis"
)

type TxtFileTrigger struct {
	mod_base.BaseTrigger
	Msg map[string]string
}

type Option struct {
	Id string
	Value string
}

type TxtRedisStruct struct {
	FileId int64
	Change int
}

func(b *TxtFileTrigger) GetTriggerName() string {
	return "TxtFileTrigger"
}

func(b *TxtFileTrigger) PreInitAction()  {
	return
}

func(b *TxtFileTrigger) GetFrontStruct() []*mod_base.FormData {
	return []*mod_base.FormData{
		{
			Title: txtSelectFile,
			Key: constant.ItemSingle,
			Value: "",
			Position:"0",
			Options:[]Option{},
		},
		{
			Title:txtTriggerFunction,
			Key: constant.ItemSingle,
			Value:"",
			Position:"1",
			Options:[]string{TOContentChange, TOProperty},
		},
	}
}

func(b *TxtFileTrigger) SetRedisTrigger(fId int64, fd []*mod_base.FormData) error {
	var fileId int64
	var change int
	var err error
	for _, v := range fd {
		if v.Title == txtSelectFile {
			fileId, err = strconv.ParseInt(v.Value, 10, 64)
			if err != nil {
				logrus.Error("[SetRedisTrigger] file id cannt find, fId", fId)
			}
		}else {
			if v.Value == TOContentChange {
				change = 1
			}else {
				change = 2
			}
		}
	}
	err = gredis.LPush(TxtRedisKey, TxtRedisStruct{
		FileId: fileId,
		Change: change,
	})
	return err
}
