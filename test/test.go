package test

import (
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/util"
)

func InitTestConfig1() {
	util.InitViper1()
	util.InitID()
	database.InitDB()
}

func InitTestConfig2() {
	util.InitViper2()
	util.InitID()
	database.InitDB()
}
