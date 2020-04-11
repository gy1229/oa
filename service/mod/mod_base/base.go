package mod_base

type FormData struct {
	Key string `json:"key"`
	Value string `json:"value"`
	Postion string `json:"postion"`
	Options []string `json:"options"`
}

var TriggerGroup map[int64]Trigger
var ActionGroup map[int64]Action

func ActionInit() {

}

func TriggerInit() {

}
