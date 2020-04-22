package mod_base

type Trigger interface {
	GetTriggerName() string
	PreInitAction()
	GetFrontStruct(int64) []*FormData
	StartTrigger() error
	GetTriggerId() int64
	GetTriggerImageId() int64
	GetTriggerType() string
	SetRedisTrigger(int64, []*FormData) error
}

type BaseTrigger struct{}

func (b *BaseTrigger) GetTriggerType() string {
	return "1"
}

func (b *BaseTrigger) GetTriggerName() string {
	return ""
}

func (b *BaseTrigger) PreInitAction() {
	return
}

func (b *BaseTrigger) GetFrontStruct(int64) []*FormData {
	return nil
}

func (b *BaseTrigger) StartTrigger() error {
	return nil
}

func (b *BaseTrigger) GetTriggerId() int64 {
	return 1
}
func (b *BaseTrigger) GetTriggerImageId() int64 {
	return 13
}

func (b *BaseTrigger) SetRedisTrigger(int64, []*FormData) error {
	return nil
}
