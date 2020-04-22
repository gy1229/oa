package mod_base

type FormData struct {
	Title string `json:"title"`
	Key string `json:"key"`
	Value string `json:"value"`
	Position string `json:"position"`
	Options interface{} `json:"options"`
}

var TriggerGroup = make(map[int64]Trigger, 0)
var ActionGroup = make(map[int64]Action, 0)

const (
	FileId = "file_id"
	UserId = "user_id"
	)


