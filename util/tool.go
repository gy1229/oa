package util

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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

func TranHttpStruct2Database(a interface{}, b interface{}) {
	aByte, err := json.Marshal(a)
	if err != nil {
		logrus.Error("TranHttpStruct2Database err", err)
		return
	}
	err = json.Unmarshal(aByte, &b)
	if err != nil {
		logrus.Error("TranHttpStruct2Database err", err)
		return
	}
}
func Byte2String(bs []int8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func TranUnixString2TimeStamp(str string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, str, loc)
	return theTime.Unix()
}
