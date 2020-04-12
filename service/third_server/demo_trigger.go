package third_server

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/service/mod/mod_base"
)

type DemoTrigger struct {
	mod_base.BaseTrigger
}

func(b *DemoTrigger) GetTriggerName() string {
	return "DemoTrigger"
}

func(b *DemoTrigger) PreInitAction()  {
	return
}

func(b *DemoTrigger) GetFrontStruct() []*mod_base.FormData {
	return []*mod_base.FormData{
		{
			Title:"单选",
			Key: constant.ItemSingle,
			Value: "",
			Postion:"0",
			Options: []string{"A", "B", "C"},
		},
		{
			Title:"多选",
			Key:constant.ItemCheckbox,
			Value:"",
			Postion:"1",
			Options:[]string{"A", "B", "C"},
		},
		{
			Title:"文本",
			Key:constant.ItemText,
			Value:"",
			Postion:"2",
		},
		{
			Title:"时间选择器",
			Key:constant.ItemDateTime,
			Value:"",
			Postion:"3",
		},
		{
			Title:"多行文本",
			Key:constant.ItemMutiText,
			Value:"",
			Postion:"4",
		},
	}
}

func(b *DemoTrigger) StartTrigger() error {
	return nil
}

