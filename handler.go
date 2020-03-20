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

func GGetFileList(c *gin.Context) {
	var req json_struct.GetFileListRequest
	req.UserId = c.Query("userId")
	req.RepositoryId = c.Query("repositoryId")
	resp, err := stage.GetFileList(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))

}

func GetFileList(c *gin.Context) {
	var req json_struct.GetFileListRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.GetFileList(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func GGetFileContent(c *gin.Context) {
	var req json_struct.GetFileContentRequest
	req.FileId = c.Query("fileId")
	req.UserId = c.Query("userId")
	resp, err := stage.GetFileContent(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func GetFileContent(c *gin.Context) {
	var req json_struct.GetFileContentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.GetFileContent(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UpdateTextContent(c *gin.Context) {
	var req json_struct.UpdateTextContentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.UpdateTextContent(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UpdateTableContent(c *gin.Context) {
	var req json_struct.UpdateTableContentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.UpdateTableContent(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func CreateNewFile(c *gin.Context) {
	var req json_struct.CreateNewFileRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.CreateNewFile(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func DelFile(c *gin.Context) {
	var req json_struct.DelFileRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage.DelFile(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}
