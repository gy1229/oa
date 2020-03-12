package util

import (
	"github.com/gy1229/oa/json_struct"
	"log"
	"testing"
)

func TestGenDefaultResp(t *testing.T) {
	GenDefaultResp("cc")
}

func TestTranformStruct2GinH(t *testing.T) {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	req := json_struct.RegisterUserRequest{
		UserBase: &json_struct.UserBase{
			Account:"123",
			Password:"1qwe",
		},
		UserName: "asd",
	}
	TranformStruct2GinH(req)
}