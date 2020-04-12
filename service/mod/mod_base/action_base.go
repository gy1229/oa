package mod_base

type Action interface {
	GetActionName() string
	GetFrontStruct() []*FormData
	ExecAction() error
	GetActionId() int64
	GetActionImageId() int64
	GetActionType() string
}

type BaseAction struct {}

func(b *BaseAction) GetActionType() string {
	return "2"
}

func(b *BaseAction) GetActionName() string {
	return ""
}

func(b *BaseAction) GetFrontStruct() []*FormData {
	return nil
}

func(b *BaseAction) ExecAction() error {
	return nil
}

func(b *BaseAction) GetActionId() int64 {
	return 0
}

func(b *BaseAction) GetActionImageId() int64 {
	return 123
}

