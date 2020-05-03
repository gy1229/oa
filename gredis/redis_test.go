package gredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
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

func TestLPush(t *testing.T) {
	util.InitViper1()
	Setup()
	pass := viper.GetString("redis.PassWord")
	fmt.Printf("%s", pass)
	conn := RedisClient.Get()
	defer conn.Close()
	//m := struct {
	//	Key string
	//	Value string
	//}{"ds","qw"}
	//value, _ := json.Marshal(m)
	reply, _ := redis.Values(conn.Do("lrange", "txt_redis", 0, 10))
	for _, v := range reply {
		fmt.Printf("value: \n%v\n", string(v.([]byte)))
	}
	fmt.Printf("%v", reply)
}
