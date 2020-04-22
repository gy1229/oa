package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func GenDefaultResp(body string) gin.H {
	m := make(map[string]interface{}, 0)
	m["body"] = body
	return m
}

func GenDefaultFailResp(err string) gin.H {
	m := make(map[string]interface{}, 0)
	m["error"] = err
	return m
}

func TranformStruct2GinH(s interface{}) gin.H {
	b, _ := json.Marshal(s)
	m := make(map[string]interface{}, 0)
	_ = json.Unmarshal(b, &m)
	return m
}

func TranVarToContent(content string, param map[string]interface{}) string {
	for k, v := range param {
		if str, ok := v.(string); ok {
			content = strings.Replace(content, fmt.Sprintf("{{%s}}", k), str, -1 )
		}
	}
	return content
}
