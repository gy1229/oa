package util

import (
	"github.com/gitstliu/go-id-worker"
	"github.com/sirupsen/logrus"
)

var IdP *idworker.IdWorker

func InitID() {
	IdP = &idworker.IdWorker{}
	IdP.InitIdWorker(1000, 1)
}
func GenId() int64 {
	newId, err := IdP.NextId()
	if err != nil {
		newId = -1
		logrus.Error("GenId err")
	}
	return newId
}
