package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gy1229/oa/json_struct"
)

func GenDefaultResp(body string) gin.H{
	b := json_struct.BaseResponse{Body:body}
	c, _ := json.Marshal(b)
	m := make(map[string]interface{}, 0)
	json.Unmarshal(c, m)
	return m
}

func TranformStruct2GinH(s interface{}) {
	//m := make(map[string]interface{}, 0)

}