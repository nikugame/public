// Copyright 2016 Nikugame. All Rights Reserved

package pool

import (
	"os"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/nikgame/public/config"
	"github.com/nikgame/public/log"
)

//Client reids连接池客户端
var Client *redis.Pool

func init() {
	conf, _ := config.NewConfig("ini", "conf/settings.conf")
	path := strings.Join([]string{conf.DefaultString("Redis::host", "127.0.0.1"), ":", conf.DefaultString("Redis::port", "6379")}, "")
	if conf.DefaultBool("Redis::mode", false) {
		Client = &redis.Pool{
			MaxIdle:     conf.DefaultInt("Redis::idle", 5),
			MaxActive:   conf.DefaultInt("Redis::active", 500),
			IdleTimeout: time.Duration(conf.DefaultInt("Redis::timeout", 180)) * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", path)
				if err != nil {
					log.Log("Redis Connect ERROR: %s", err)
					panic(err)
				}
				if _, err := c.Do("AUTH", conf.DefaultString("Redis::password", "")); err != nil {
					c.Close()
					return nil, err
				}
				return c, nil
			},
		}
	} else {
		Client = &redis.Pool{
			MaxIdle:     conf.DefaultInt("Redis::idle", 5),
			MaxActive:   conf.DefaultInt("Redis::active", 500),
			IdleTimeout: time.Duration(conf.DefaultInt("Redis::timeout", 180)) * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", path)
				if err != nil {
					log.Log("Redis Connect ERROR: %s", err)
					os.Exit(1)
				}
				return c, nil
			},
		}
	}

}
