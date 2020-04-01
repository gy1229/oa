package util

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenHandlerRequest(c *gin.Context, req interface{}) {
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}