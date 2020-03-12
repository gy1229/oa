package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", TestHanlder)
	r.POST("/register", RegisterUser)
	r.POST("/login", LoginUser)
	r.POST("/updateUserMessage", UpdateUserMessage)
	r.POST("/loadUserMessage", LoadUserMessage)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}