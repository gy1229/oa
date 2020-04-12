package mod_base

type FormData struct {
	Title string `json:"title"`
	Key string `json:"key"`
	Value string `json:"value"`
	Postion string `json:"postion"`
	Options []string `json:"options"`
}

var TriggerGroup = make(map[int64]Trigger, 0)
var ActionGroup = make(map[int64]Action, 0)

