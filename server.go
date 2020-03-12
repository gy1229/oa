package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gy1229/oa/tool"
)

func main() {
	tool.InitViper()
	tool.InitDatabse()
	r := gin.Default()
	r.GET("/ping", TestHanlder)
	r.POST("/register", RegisterUser)
	r.POST("/login", LoginUser)
	r.POST("/updateUserMessage", UpdateUserMessage)
	r.POST("/loadUserMessage", LoadUserMessage)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}