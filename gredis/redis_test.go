package gredis

import (
	"fmt"
	"github.com/gy1229/oa/util"
	"github.com/spf13/viper"
	"testing"
)

func TestSet(t *testing.T) {
	util.InitViper1()
	Setup()
	pass := viper.GetString("redis.PassWord")
	fmt.Printf("%s", pass)
	st, err := Set("key", "value", 10000)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", st)
}

func TestGet(t *testing.T) {
	util.InitViper1()
	Setup()
	data, err := Get("key")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s", data)
}

func TestExists(t *testing.T) {
	util.InitViper1()
	Setup()
	st, err := Exists("key")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", st)
}

func TestDelete(t *testing.T) {
	util.InitViper1()
	Setup()
	st, err := Delete("key")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", st)
}