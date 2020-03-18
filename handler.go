package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/service/user"
	"github.com/gy1229/oa/util"
	"net/http"
)

func TestHanlder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func RegisterUser(c *gin.Context) {
	var req json_struct.RegisterUserRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user.InsertUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func LoginUser(c *gin.Context) {
	var req json_struct.LoginUserRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user.LoginUser(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UpdateUserMessage(c *gin.Context) {
	var req json_struct.UpdateUserRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user.UpdateUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func LoadUserMessage(c *gin.Context) {
	var req json_struct.LoadUserMessageRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user.LoadUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func CertainAccount(c *gin.Context) {
	var req json_struct.CertainAccountRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user.CertainAccount(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UploadFile(c *gin.Context) {
	//header, err := c.FormFile(constant.UploadFileKey)
}
