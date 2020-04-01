package constant

const (
	ItemSingle   = "Single"   // 单选
	ItemCheckbox = "Checkbox" // 多选
	ItemText     = "Text"     // 单行文本
	ItemDateTime = "DateTime" // 时间
	ItemMutiText = "MutiText" // 多行文本
)

const (
	Trigger = 0
	Action = 1
)

const (
	TriggerEmail = 1
	TriggerTable = 3
	TriggerFile = 5
	TriggerRepository = 7
)

const (
	ActionEmail = 2
	ActionTable = 3
	ActionFile = 6
	ActionRepository = 8
)

const (
	FlowInstanceStart = 1
	FlowInstanceSuccess = 2
	FlowInstanceFail = 3
)