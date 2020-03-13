package test

import (
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/tool"
	"github.com/gy1229/oa/util"
)

func InitTestConfig1(){
	tool.InitViper1()
	util.InitID()
	database.InitDB()
}

func InitTestConfig2() {
	tool.InitViper2()
	util.InitID()
	database.InitDB()
}
