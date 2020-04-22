package kafka

import (
	"testing"
)

func TestAAA(t *testing.T) {
	ConsumerInit()
	ConsumerStart("trigger_message")
}

func TestProductStart(t *testing.T) {
	ProductInit()
	ProductStart("trigger_message", "hello %d")
}
