package user

import (
	"fmt"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/test"
	"testing"
)

func TestLoadUserMessage(t *testing.T) {
	test.InitTestConfig2()
	resp, _ := LoadUserMessage(&json_struct.LoadUserMessageRequest{UserId: "123"})
	fmt.Println(resp)
}

func TestCertainAccount(t *testing.T) {
	test.InitTestConfig2()
	resp, _ := CertainAccount(&json_struct.CertainAccountRequest{Account: "guyu"})
	fmt.Println(resp)
}

func TestInsertUserMessage(t *testing.T) {
	test.InitTestConfig2()
	InsertUserMessage(&json_struct.RegisterUserRequest{
		UserBase: &json_struct.UserBase{
			Account:  "12223",
			Password: "guyu1998",
		},
		UserName:"Helloa",
	})
}
func TestUploadAvator(t *testing.T) {
	test.InitTestConfig2()
	UploadAvator(&json_struct.UploadAvatarRequest{
		FileContent: "hello, img",
		FileSize:    123,
		Name:        "hello",
	})
}