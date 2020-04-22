package gredis

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisClient *redis.Pool

func Setup() error {
	RedisClient = &redis.Pool{
		MaxIdle:     viper.GetInt("redis.MaxIdle"),
		MaxActive:   viper.GetInt("redis.MaxActive"),
		IdleTimeout: time.Duration(viper.GetInt64("redis.IdleTimeout")),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", viper.GetString("redis.Host"))
			if err != nil {
				return nil, err
			}
			if viper.GetString("redis.Password") != "" {
				if _, err := c.Do("AUTH", viper.GetString("redis.Password")); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) (bool, error) {
	conn := RedisClient.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	reply, err := conn.Do("SET", key, value)
	conn.Do("EXPIRE", key, time)
	if err != nil {
		logrus.Error("[redis Set] reply ", reply)
		return false, err
	}
	return true, err
}

func Exists(key string) (bool, error) {
	conn := RedisClient.Get()
	defer conn.Close()

	reply, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		logrus.Error("[redis Exists] reply ", reply, key)
		return false, err
	}
	return reply, err
}

func Get(key string) ([]byte, error) {
	conn := RedisClient.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		logrus.Error("[redis Get] reply ", reply, key)
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	reply, err := conn.Do("DEL", key)
	if err != nil {
		logrus.Error("[redis Delete] reply ", reply, key)
		return false, err
	}
	return true, nil
}

func LikeDeletes(key string) error {
	conn := RedisClient.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func LRange(key string) ([]string, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	rStr := make([]string, 0)
	reply, err := redis.Values(conn.Do("lrange", "key", 0, 10))
	if err != nil {
		return nil, err
	}
	for _, v := range reply {
		rStr = append(rStr, string(v.([]byte)))
	}
	return rStr, err
}

func LPush(key string, data interface{}) error {
	conn := RedisClient.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("lpush", key, value)
	return nil
}
