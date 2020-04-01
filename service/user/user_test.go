package user

import (
	"fmt"
	"github.com/gy1229/oa/database"
	"github.com/gy1229/oa/json_struct"
	"github.com/gy1229/oa/test"
	"github.com/sirupsen/logrus"
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
func TestHello(t *testing.T) {
	test.InitTestConfig2()
	a := database.FileTable{
		Id:   123,
		Name: "啊哈哈",
	}
	if err := database.DB.Create(a).Error; err != nil {
		logrus.Error("[InsertUserMessage] err msg", err.Error())
	}
}