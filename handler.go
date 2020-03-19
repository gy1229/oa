package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/service/stage"
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


// CreateRepository 创建仓库
func CreateRepository(c *gin.Context) {
	var req json_struct.CreateRepositoryRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.CreateRepository(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

// GetRepositoryList 获取用户私人仓库列表
func GetRepositoryList(c *gin.Context) {
	var req json_struct.GetRepositoryListRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.GetRepositoryList(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

// UpdateRepository 更新仓库信息
func UpdateRepository(c *gin.Context) {
	var req json_struct.UpdateRepositoryRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.UpdateRepository(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

// DelRepository 删除仓库
func DelRepository(c *gin.Context) {
	var req json_struct.DelRepositoryRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.DelRepository(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}