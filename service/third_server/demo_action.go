package third_server

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/service/mod/mod_base"
)

type DemoAction struct {
	mod_base.BaseAction
}

func(b *DemoAction) GetActionName() string {
	return "DemoAction"
}

func(b *DemoAction) GetFrontStruct() []*mod_base.FormData {
	return []*mod_base.FormData{
		{
			Title:"单选",
			Key: constant.ItemSingle,
			Value: "",
			Position:"0",
			Options: []string{"A", "B", "C"},
		},
		{
				Title:"多选",
				Key:constant.ItemCheckbox,
				Value:"",
			Position:"1",
				Options:[]string{"A", "B", "C"},
		},
		{
					Title:"文本",
					Key:constant.ItemText,
					Value:"",
			Position:"2",
		},
		{
						Title:"时间选择器",
				Key:constant.ItemDateTime,
				Value:"",
			Position:"3",
		},
		{
							Title:"多行文本",
			Key:constant.ItemMutiText,
			Value:"",
			Position:"4",
		},
	}
}

func(b *DemoAction) ExecAction() error {
	return nil
}
