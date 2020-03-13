package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GenDefaultResp(body string) gin.H{
	m := make(map[string]interface{}, 0)
	m["body"] = body
	return m
}

func GenDefaultFailResp(err string) gin.H {
	m := make(map[string]interface{}, 0)
	m["error"] = err
	return m
}

func TranformStruct2GinH(s interface{}) gin.H{
	b, _ := json.Marshal(s)
	m := make(map[string]interface{}, 0)
	_ = json.Unmarshal(b, &m)
	return m
}