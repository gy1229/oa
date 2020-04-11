package mod_base

type Action interface {
	GetActionName() string
	GetFrontStruct() []*FormData
	ExecAction() error
}

type BaseAction struct {}

func(b *BaseAction) GetActionName() string {
	return ""
}

func(b *BaseAction) GetFrontStruct() []*FormData {
	return nil
}

func(b *BaseAction) ExecAction() error {
	return nil
}

