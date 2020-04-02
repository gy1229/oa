package base

type BehaviorInstance struct {
	Key string `json:"key"`
	Value string `json:"value"`
	Postion string `json:"postion"`
	Options []string `json:"options"`
}

type Action interface {
	GetActionName() string
	GetFrontStruct() []*BehaviorInstance
	ExecAction() error
}


type TriggerAction interface {
	PreInitAction()
}