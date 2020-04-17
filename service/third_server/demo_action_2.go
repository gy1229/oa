package third_server

import (
	"github.com/gy1229/oa/constant"
	"github.com/gy1229/oa/service/mod/mod_base"
)

type DemoAction2 struct {
	mod_base.BaseAction
}

func (b *DemoAction2) GetActionName() string {
	return "DemoAction_2"
}

func (b *DemoAction2) GetFrontStruct() []*mod_base.FormData {
	return []*mod_base.FormData{
		{
			Title:    "单选-2",
			Key:      constant.ItemSingle,
			Value:    "",
			Position: "0",
			Options:  []string{"A4", "B3", "C3"},
		},
		{
			Title:    "多选=2",
			Key:      constant.ItemCheckbox,
			Value:    "",
			Position: "1",
			Options:  []string{"A1", "B2", "C3"},
		},
	}
}

func (b *DemoAction2) ExecAction() error {
	return nil
}
