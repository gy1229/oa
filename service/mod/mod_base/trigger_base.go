package mod_base


type Trigger interface {
	GetTriggerName() string
	PreInitAction()
	GetFrontStruct() []*FormData
	StartTrigger() error
}

type BaseTrigger struct {}

func(b *BaseTrigger) GetTriggerName() string {
	return ""
}

func(b *BaseTrigger) PreInitAction()  {
	return
}

func(b *BaseTrigger) GetFrontStruct() []*FormData {
	return nil
}

func(b *BaseTrigger) StartTrigger() error {
	return nil
}

