package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/service/file_server"
	"github.com/gy1229/oa/service/stage_server"
	"github.com/gy1229/oa/service/user_server"
	"github.com/gy1229/oa/util"
	"io/ioutil"
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
	resp, err := user_server.InsertUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func LoginUser(c *gin.Context) {
	var req json_struct.LoginUserRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.LoginUser(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UpdateUserMessage(c *gin.Context) {
	var req json_struct.UpdateUserRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.UpdateUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func LoadUserMessage(c *gin.Context) {
	var req json_struct.LoadUserMessageRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.LoadUserMessage(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func CertainAccount(c *gin.Context) {
	var req json_struct.CertainAccountRequest
	util.GenHandlerRequest(c, &req)
	resp, err := user_server.CertainAccount(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UploadFile(c *gin.Context) {
	formFile, _ := c.FormFile("file")
	userId := c.PostForm("user_id")
	repId := c.PostForm("repository_id")
	err := stage_server.UploadFile2Stage(formFile, userId, repId)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(util.GenDefaultResp("success")))
	//header, err := c.FormFile(constant.UploadFileKey)
}

func UploadAvatar(c *gin.Context) {
	formFile, _ := c.FormFile("file")
	var req json_struct.UploadAvatarRequest
	file, _ := formFile.Open()
	req.Name = formFile.Filename
	fileByte, _ := ioutil.ReadAll(file)
	req.FileContent = fileByte
	//filename := file.Filename
	resp, err := file_server.UploadAvatar(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))

}

func GetAvatoar(c *gin.Context) {
	var req json_struct.GetAvatarRequest
	util.GenHandlerRequest(c, &req)
	resp, err := file_server.GetAvatoar(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.Writer.Write(resp.ImageFile)
	//resp, err := user.GetAvator(&req)
	//if err != nil {
	//	c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
	//	return
	//}
	//size, _ := c.Writer.WriteString(resp.ImageFile)
	//if err != nil {
	//	return
	//}
	//logrus.Info("[GetAvatoar] WriteString size is ", size)
	//c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

// CreateRepository 创建仓库
func CreateRepository(c *gin.Context) {
	var req json_struct.CreateRepositoryRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage_server.CreateRepository(&req)
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
	resp, err := stage_server.GetRepositoryList(&req)
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
	resp, err := stage_server.UpdateRepository(&req)
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
	resp, err := stage_server.DelRepository(&req)
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
	resp, err := stage_server.GetFileList(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))

}

func GetFileList(c *gin.Context) {
	var req json_struct.GetFileListRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage_server.GetFileList(&req)
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
	resp, err := stage_server.GetFileContent(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func GetFileContent(c *gin.Context) {
	var req json_struct.GetFileContentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage_server.GetFileContent(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UpdateTextContent(c *gin.Context) {
	var req json_struct.UpdateTextContentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage_server.UpdateTextContent(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func UpdateTableContent(c *gin.Context) {
	var req json_struct.UpdateTableContentRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage_server.UpdateTableContent(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func CreateNewFile(c *gin.Context) {
	var req json_struct.CreateNewFileRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage_server.CreateNewFile(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func DelFile(c *gin.Context) {
	var req json_struct.DelFileRequest
	util.GenHandlerRequest(c, &req)
	resp, err := stage_server.DelFile(&req)
	if err != nil {
		c.JSON(http.StatusOK, util.GenDefaultFailResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.TranformStruct2GinH(resp))
}

func GetActionList(c *gin.Context) {

}

func GetActionDefination(c *gin.Context) {

}

func GetFlowDefinationDetail(c *gin.Context) {

}

func CreateFlowDefination(c *gin.Context) {

}

func UpdateFlowDefination(c *gin.Context) {

}

func GetFlowDefinationList(c *gin.Context) {

}