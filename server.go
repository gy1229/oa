package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/gredis"
	"github.com/gy1229/oa/kafka"
	"github.com/gy1229/oa/service/file_server"
	"github.com/gy1229/oa/util"
	"net/http"
	"strings"
)

func main() {
	util.InitViper()
	util.InitFileType()
	database.InitDB()
	util.InitID()
	gredis.Setup()
	kafka.ConsumerInit()
	kafka.ProductInit()
	file_server.FileServerInit()
	r := gin.Default()
	r.Use(Cors())

	// user
	r.GET("/ping", TestHanlder)
	r.POST("/register", RegisterUser)
	r.POST("/login", LoginUser)
	r.POST("/updateUserMessage", UpdateUserMessage)
	r.POST("/loadUserMessage", LoadUserMessage)
	r.POST("/certainAccount", CertainAccount)
	r.POST("/uploadAvatar", UploadAvatar)
	r.POST("/getAvatar", GetAvatoar)

	// stage
	r.POST("/uploadFile", UploadFile) // todo
	r.POST("/createRepository", CreateRepository)
	r.POST("/getRepositoryList", GetRepositoryList)
	r.POST("/updateRepository", UpdateRepository)
	r.POST("/delRepository", DelRepository)

	r.GET("/getFileList/:userId/:repositoryId", GGetFileList)

	// file
	r.POST("/getFileList", GetFileList)
	r.POST("/getFileContent", GetFileContent)
	r.POST("/updateTextContent", UpdateTextContent)
	r.POST("/updateTableContent", UpdateTableContent)
	r.POST("/createNewFile", CreateNewFile)
	r.POST("/delFile", DelFile)

	r.GET("/getFileDetail/:userId/:fileId", GGetFileContent)


	// automation
	r.POST("/getActionList", GetActionList)
	r.POST("/getActionDefination", GetActionDefination)
	r.POST("/getFlowDefinationDetail", GetFlowDefinationDetail)
	r.POST("/createFlowDefination", CreateFlowDefination)
	r.POST("/updateFlowDefination", UpdateFlowDefination)
	r.POST("/getFlowDefinationList", GetFlowDefinationList)
	r.POST("/deleteFlowDeination", DeleteFlowDeination)

	r.Run(":19999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
