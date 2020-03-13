package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/tool"
	"github.com/gy1229/oa/util"
)

func main() {
	tool.InitViper()
	database.InitDB()
	util.InitID()
	r := gin.Default()
	r.GET("/ping", TestHanlder)
	r.POST("/register", RegisterUser)
	r.POST("/login", LoginUser)
	r.POST("/updateUserMessage", UpdateUserMessage)
	r.POST("/loadUserMessage", LoadUserMessage)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}